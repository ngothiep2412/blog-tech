package usermodel

import "errors"

var (
	// User errors
	ErrUserNotFound    = errors.New("user not found")
	ErrUserExists      = errors.New("user already exists")
	ErrEmailExists     = errors.New("email already exists")
	ErrUsernameExists  = errors.New("username already exists")
	ErrInvalidPassword = errors.New("invalid password")
	ErrUserInactive    = errors.New("user account is inactive")

	// Validation errors
	ErrInvalidEmail     = errors.New("invalid email format")
	ErrInvalidUsername  = errors.New("invalid username format")
	ErrPasswordTooShort = errors.New("password must be at least 6 characters")
	ErrRequiredField    = errors.New("required field is missing")

	// Auth errors
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrTokenGeneration    = errors.New("failed to generate token")
	ErrTokenInvalid       = errors.New("invalid token")
	ErrUnauthorized       = errors.New("unauthorized access")

	ErrCannotCreateUser = errors.New("cannot create user")
	ErrCannotUpdateUser = errors.New("cannot update user")
	ErrCannotDeleteUser = errors.New("cannot delete user")
)
