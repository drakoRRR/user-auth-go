package utils

import (
	"errors"
	"github.com/drakoRRR/user-auth-go/internal/models"
	"regexp"
)

func ValidateUserData(user *models.User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if !IsValidEmail(user.Email) {
		return errors.New("invalid email format")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func IsValidEmail(email string) bool {
	const emailRegex = `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
