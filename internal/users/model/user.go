package usermodel

import (
	"blog-tech/common"
	articlemodel "blog-tech/internal/articles/model"
)

type User struct {
	common.SqlModel `json:",inline"`
	Username        string `json:"username" gorm:"column:username;unique"`
	Email           string `json:"email" gorm:"column:email;unique"`
	Password        string `json:"password" gorm:"column:password"`
	FullName        string `json:"full_name" gorm:"column:full_name"`
	Bio             string `json:"bio" gorm:"column:bio"`
	AvatarURL       string `json:"avatar_url" gorm:"column:avatar_url"`
	IsActive        bool   `json:"is_active" gorm:"column:is_active"`

	Articles []articlemodel.Article `json:"articles" gorm:"preload:false"`
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type UserUpdate struct {
	FullName  string `json:"full_name"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
}
