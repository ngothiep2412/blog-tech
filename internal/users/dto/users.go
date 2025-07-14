package dto

import usermodel "blog-tech/internal/users/model"

type RegisterUserResponse struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

type LoginUserResponse struct {
	User  *usermodel.User `json:"user"`
	Token string          `json:"token"`
}
