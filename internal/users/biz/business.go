package userbiz

import (
	"blog-tech/common"
	usermodel "blog-tech/internal/users/model"
	usermysql "blog-tech/internal/users/repository/mysql"
	"fmt"
	"strings"
)

type UserBusiness interface {
	Register(req *usermodel.CreateUserRequest) (*usermodel.User, string, error)
	Login(req *usermodel.LoginRequest) (*usermodel.User, string, error)
	GetProfile(userID int) (*usermodel.User, error)
	UpdateProfile(userID int, req *usermodel.UpdateUserRequest) (*usermodel.User, error)
	ChangePassword(userID int, req *usermodel.ChangePasswordRequest) error
	ListUsers(limit, offset int) ([]*usermodel.User, int64, error)
	DeactivateUser(userID int) error
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

func (b *userBusiness) Register(req *usermodel.CreateUserRequest) (*usermodel.User, string, error) {
	if err := req.Validate(); err != nil {
		return nil, "", err
	}

	if _, err := b.userRepo.GetUserByEmail(req.Email); err == nil {
		return nil, "", usermodel.ErrEmailExists
	}

	if _, err := b.userRepo.GetUserByUsername(req.Username); err == nil {
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

	if err := b.userRepo.Create(user); err != nil {
		return nil, "", fmt.Errorf("failed to create user: %w", err)
	}

	token, err := b.JwtManager.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, "", usermodel.ErrTokenGeneration
	}

	return user, token, nil
}

func (b *userBusiness) Login(req *usermodel.LoginRequest) (*usermodel.User, string, error) {
	if req.Email == "" || req.Password == "" {
		return nil, "", usermodel.ErrRequiredField
	}

	user, err := b.userRepo.GetUserByEmail(strings.ToLower(req.Email))
	if err != nil {
		if err == usermodel.ErrUserNotFound {
			return nil, "", usermodel.ErrInvalidCredentials
		}
		return nil, "", err
	}

	if !user.IsActive {
		return nil, "", usermodel.ErrUserInactive
	}

	if !common.CheckPassword(req.Password, user.PasswordHash) {
		return nil, "", usermodel.ErrInvalidCredentials
	}

	token, err := b.JwtManager.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, "", usermodel.ErrTokenGeneration
	}

	return user, token, nil
}

func (b *userBusiness) GetProfile(userID int) (*usermodel.User, error) {
	user, err := b.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (b *userBusiness) UpdateProfile(userID int, req *usermodel.UpdateUserRequest) (*usermodel.User, error) {
	// Get existing user
	user, err := b.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.FullName != "" {
		user.FullName = req.FullName
	}
	if req.Bio != "" {
		user.Bio = req.Bio
	}
	if req.AvatarURL != "" {
		user.AvatarURL = req.AvatarURL
	}

	// Save updated user
	if err := b.userRepo.Update(user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

func (b *userBusiness) ChangePassword(userID int, req *usermodel.ChangePasswordRequest) error {
	user, err := b.userRepo.GetUserByID(userID)
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
	if err := b.userRepo.Update(user); err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}

func (b *userBusiness) ListUsers(limit, offset int) ([]*usermodel.User, int64, error) {
	users, err := b.userRepo.List(limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := b.userRepo.Count()
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (b *userBusiness) DeactivateUser(userID int) error {
	user, err := b.userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}

	user.IsActive = false
	return b.userRepo.Update(user)
}
