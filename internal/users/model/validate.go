package usermodel

import (
	"fmt"
	"regexp"
)

func (req *CreateUserRequest) Validate() error {
	if req.Username == "" {
		return ErrRequiredField
	}
	if len(req.Username) < 3 || len(req.Username) > 50 {
		return ErrInvalidUsername
	}
	if !isValidUsername(req.Username) {
		return ErrInvalidUsername
	}

	if req.Email == "" {
		return ErrRequiredField
	}
	if !isValidEmail(req.Email) {
		return ErrInvalidEmail
	}

	if req.Password == "" {
		return ErrRequiredField
	}
	if len(req.Password) < 6 {
		return ErrPasswordTooShort
	}

	if req.FullName == "" {
		return ErrRequiredField
	}
	if len(req.FullName) > 150 {
		return fmt.Errorf("full name must be less than 150 characters")
	}

	return nil
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func isValidUsername(username string) bool {
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return usernameRegex.MatchString(username)
}
