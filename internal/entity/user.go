package entity

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username,omitempty"`
	Email        string `json:"email" bson:"email,omitempty"`
	PasswordHash string `json:"-" bson:"password"`
}

type CreateUserDTO struct {
	Username       string
	Email          string
	Password       string
	RepeatPassword string
}

func NewUser(dto CreateUserDTO) (u User, err error) {
	pwd, err := generatePasswordHash(dto.Password)
	if err != nil {
		return u, fmt.Errorf("%v", err)
	}

	return User{
		Email:        dto.Email,
		Username:     dto.Username,
		PasswordHash: pwd,
	}, nil
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password due to error %w", err)
	}
	return string(hash), nil
}
