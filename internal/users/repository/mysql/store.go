package usermysql

import (
	usermodel "blog-tech/internal/users/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *usermodel.User) error
	GetUserByID(id int) (*usermodel.User, error)
	GetUserByEmail(email string) (*usermodel.User, error)
	GetUserByUsername(username string) (*usermodel.User, error)
	Update(user *usermodel.User) error
	Delete(id int) error
	List(limit, offset int) ([]*usermodel.User, error)
	Count() (int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *usermodel.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *userRepository) GetUserByID(id int) (*usermodel.User, error) {
	var user usermodel.User

	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, usermodel.ErrUserNotFound
		}
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*usermodel.User, error) {
	var user usermodel.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, usermodel.ErrUserNotFound
		}
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsername(username string) (*usermodel.User, error) {
	var user usermodel.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, usermodel.ErrUserNotFound
		}
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (r *userRepository) Update(user *usermodel.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *userRepository) Delete(id int) error {
	if err := r.db.Delete(&usermodel.User{}, id).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *userRepository) List(limit, offset int) ([]*usermodel.User, error) {
	var users []*usermodel.User
	if err := r.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return users, nil
}

func (r *userRepository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&usermodel.User{}).Count(&count).Error; err != nil {
		return 0, errors.WithStack(err)
	}
	return count, nil
}
