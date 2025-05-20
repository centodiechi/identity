package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GetPasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func ComparePasswordHash(password []byte, pass string) error {
	return bcrypt.CompareHashAndPassword(password, []byte(pass))
}
