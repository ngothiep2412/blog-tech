package usermysql

import (
	usermodel "blog-tech/internal/users/model"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *usermodel.User) error
	GetUserByID(ctx context.Context, id int) (*usermodel.User, error)
	GetUserByEmail(ctx context.Context, email string) (*usermodel.User, error)
	GetUserByUsername(ctx context.Context, username string) (*usermodel.User, error)
	Update(ctx context.Context, user *usermodel.User) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, limit, offset int) ([]*usermodel.User, error)
	Count(ctx context.Context) (int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *usermodel.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return errors.Wrap(err, usermodel.ErrCannotCreateUser.Error())
	}
	return nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (*usermodel.User, error) {
	var user usermodel.User

	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(err, usermodel.ErrUserNotFound.Error())
		}
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*usermodel.User, error) {
	var user usermodel.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(err, usermodel.ErrUserNotFound.Error())
		}
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*usermodel.User, error) {
	var user usermodel.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(err, usermodel.ErrUserNotFound.Error())
		}
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *usermodel.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return errors.Wrap(err, usermodel.ErrCannotUpdateUser.Error())
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.Delete(&usermodel.User{}, id).Error; err != nil {
		return errors.Wrap(err, usermodel.ErrCannotDeleteUser.Error())
	}
	return nil
}

func (r *userRepository) List(ctx context.Context, limit, offset int) ([]*usermodel.User, error) {
	var users []*usermodel.User
	if err := r.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return users, nil
}

func (r *userRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.Model(&usermodel.User{}).Count(&count).Error; err != nil {
		return 0, errors.WithStack(err)
	}
	return count, nil
}
