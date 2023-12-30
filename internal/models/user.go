package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"-" validate:"required,email"`
	Password string `json:"-" validate:"required"`
	Role     string `json:"role"`
}

func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {

		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePasswords(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
