package userdto

import usermodel "blog-tech/internal/users/model"

type RegisterResponse struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User  *usermodel.User `json:"user"`
	Token string          `json:"token"`
}
