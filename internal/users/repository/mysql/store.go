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
	tx := r.db.Begin()

	if err := tx.Table(usermodel.User{}.TableName()).Create(user).Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, usermodel.ErrCannotCreateUser.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (*usermodel.User, error) {
	var user usermodel.User

	if err := r.db.Table(usermodel.User{}.TableName()).Where("id = ?", id).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(err, usermodel.ErrUserNotFound.Error())
		}
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*usermodel.User, error) {
	var user usermodel.User

	if err := r.db.Table(usermodel.User{}.TableName()).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(err, usermodel.ErrUserNotFound.Error())
		}
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*usermodel.User, error) {
	var user usermodel.User
	if err := r.db.Table(usermodel.User{}.TableName()).Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(err, usermodel.ErrUserNotFound.Error())
		}
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *usermodel.User) error {
	tx := r.db.Begin()

	if err := tx.Table(usermodel.User{}.TableName()).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, usermodel.ErrCannotUpdateUser.Error())
	}
	if err := tx.Commit().Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	tx := r.db.Begin()

	if err := tx.Table(usermodel.User{}.TableName()).Where("id = ?", id).Delete(&usermodel.User{}).Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, usermodel.ErrCannotDeleteUser.Error())
	}
	if err := tx.Commit().Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *userRepository) List(ctx context.Context, limit, offset int) ([]*usermodel.User, error) {
	var users []*usermodel.User

	if err := r.db.Table(usermodel.User{}.TableName()).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return users, nil
}

func (r *userRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.Table(usermodel.User{}.TableName()).Select("id").Count(&count).Error; err != nil {
		return 0, errors.WithStack(err)
	}
	return count, nil
}
