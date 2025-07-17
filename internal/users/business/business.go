package userbiz

import (
	"blog-tech/common"
	userdto "blog-tech/internal/users/dto"
	usermodel "blog-tech/internal/users/model"
	usermysql "blog-tech/internal/users/repository/mysql"
	"context"
	"fmt"
	"strings"
)

type UserBusiness interface {
	Register(ctx context.Context, req *usermodel.CreateUserRequest) (*usermodel.User, string, error)
	Login(ctx context.Context, req *userdto.LoginRequest) (*usermodel.User, string, string, error)
	GetProfile(ctx context.Context, userID int) (*usermodel.User, error)
	UpdateProfile(ctx context.Context, userID int, req *usermodel.UpdateUserRequest) (*usermodel.User, error)
	ChangePassword(ctx context.Context, userID int, req *usermodel.ChangePasswordRequest) error
	ListUsers(ctx context.Context, limit, offset int) ([]*usermodel.User, int64, error)
	DeactivateUser(ctx context.Context, userID int) error
	RefreshToken(ctx context.Context, req *userdto.RefreshTokenRequest) (string, string, error)
}

type userBusiness struct {
	userRepo   usermysql.UserRepository
	JwtManager *common.JwtManager
}

func NewUserBusiness(userRepo usermysql.UserRepository, jwtManager *common.JwtManager) *userBusiness {
	return &userBusiness{
		userRepo:   userRepo,
		JwtManager: jwtManager,
	}
}

func (b *userBusiness) Register(ctx context.Context, req *usermodel.CreateUserRequest) (*usermodel.User, string, error) {
	if err := req.Validate(); err != nil {
		return nil, "", err
	}

	if _, err := b.userRepo.GetUserByEmail(ctx, req.Email); err == nil {
		return nil, "", usermodel.ErrEmailExists
	}

	if _, err := b.userRepo.GetUserByUsername(ctx, req.Username); err == nil {
		return nil, "", usermodel.ErrUsernameExists
	}

	hashPassword, err := common.HashPassword(req.Password)
	if err != nil {
		return nil, "", err
	}

	user := &usermodel.User{
		Username:     req.Username,
		Email:        strings.ToLower(req.Email),
		PasswordHash: hashPassword,
		FullName:     req.FullName,
		IsActive:     true,
	}

	if err := b.userRepo.Create(ctx, user); err != nil {
		return nil, "", fmt.Errorf("failed to create user: %w", err)
	}

	token, err := b.JwtManager.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, "", usermodel.ErrTokenGeneration
	}

	return user, token, nil
}

func (b *userBusiness) Login(ctx context.Context, req *userdto.LoginRequest) (*usermodel.User, string, string, error) {
	if req.Email == "" || req.Password == "" {
		return nil, "", "", usermodel.ErrRequiredField
	}

	user, err := b.userRepo.GetUserByEmail(ctx, strings.ToLower(req.Email))
	if err != nil {
		if err == usermodel.ErrUserNotFound {
			return nil, "", "", usermodel.ErrInvalidCredentials
		}
		return nil, "", "", err
	}

	if !user.IsActive {
		return nil, "", "", usermodel.ErrUserInactive
	}

	if !common.CheckPassword(req.Password, user.PasswordHash) {
		return nil, "", "", usermodel.ErrInvalidCredentials
	}

	accessToken, refreshToken, err := b.JwtManager.GenerateTokens(user.ID, user.Email)
	if err != nil {
		return nil, "", "", usermodel.ErrTokenGeneration
	}

	return user, accessToken, refreshToken, nil
}

func (b *userBusiness) GetProfile(ctx context.Context, userID int) (*usermodel.User, error) {
	user, err := b.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (b *userBusiness) UpdateProfile(ctx context.Context, userID int, req *usermodel.UpdateUserRequest) (*usermodel.User, error) {
	// Get existing user
	user, err := b.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Save updated user
	if err := b.userRepo.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

func (b *userBusiness) ChangePassword(ctx context.Context, userID int, req *usermodel.ChangePasswordRequest) error {
	user, err := b.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	if !common.CheckPassword(req.OldPassword, user.PasswordHash) {
		return usermodel.ErrInvalidPassword
	}
	if len(req.NewPassword) < 6 {
		return usermodel.ErrPasswordTooShort
	}

	hashedPassword, err := common.HashPassword(req.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to hash new password: %w", err)
	}

	user.PasswordHash = hashedPassword
	if err := b.userRepo.Update(ctx, user); err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}

func (b *userBusiness) ListUsers(ctx context.Context, limit, offset int) ([]*usermodel.User, int64, error) {
	users, err := b.userRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := b.userRepo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (b *userBusiness) DeactivateUser(ctx context.Context, userID int) error {
	user, err := b.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	user.IsActive = false
	return b.userRepo.Update(ctx, user)
}

func (b *userBusiness) RefreshToken(ctx context.Context, req *userdto.RefreshTokenRequest) (string, string, error) {
	claims, err := b.JwtManager.ValidateRefreshToken(req.RefreshToken)

	if err != nil {
		return "", "", usermodel.ErrTokenInvalid
	}

	user, err := b.userRepo.GetUserByID(ctx, claims.UserID)
	if err != nil {
		return "", "", usermodel.ErrUserNotFound
	}

	if !user.IsActive {
		return "", "", usermodel.ErrUserInactive
	}

	accessToken, err := b.JwtManager.GenerateToken(claims.UserID, claims.Email)
	if err != nil {
		return "", "", usermodel.ErrTokenGeneration
	}

	return accessToken, "", nil
}
