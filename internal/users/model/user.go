package usermodel

import (
	"blog-tech/common"
	articlemodel "blog-tech/internal/articles/model"
)

type User struct {
	common.SqlModel `json:",inline"`
	Username        string `json:"username" gorm:"column:username" db:"username"`
	Email           string `json:"email" gorm:"column:email" db:"email"`
	PasswordHash    string `json:"-" gorm:"column:password_hash" db:"password_hash"`
	FullName        string `json:"full_name" gorm:"column:full_name" db:"full_name"`
	Bio             string `json:"bio" gorm:"column:bio" db:"bio"`
	AvatarURL       string `json:"avatar_url" gorm:"column:avatar_url" db:"avatar_url"`
	IsActive        bool   `json:"is_active" gorm:"column:is_active" db:"is_active"`

	Articles []articlemodel.Article `json:"articles,omitempty" gorm:"preload:false"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	FullName string `json:"full_name" validate:"required,min=1,max=150"`
}

type UpdateUserRequest struct {
	FullName  *string `json:"full_name" validate:"omitempty,min=1,max=150"`
	Bio       *string `json:"bio" validate:"omitempty,max=1000"`
	AvatarURL *string `json:"avatar_url" validate:"omitempty,url"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}
