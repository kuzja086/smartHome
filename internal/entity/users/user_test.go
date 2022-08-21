package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUserDTO(t *testing.T) {
	in := CreateUserDTO{
		Username:       "testUser",
		Password:       "testPass",
		Email:          "test@test.ru",
		RepeatPassword: "testPass",
	}

	_, err := NewUser(in)
	require.NoError(t, err)
}

func TestCheckPass(t *testing.T) {
	hash := "$2a$04$lmQ..8n/.jxq4uNVcrfUaO9A/b4q.xaS5OLwTmGU/55GB92b/0X22"
	password := "testPass"
	err := CheckPassword(hash, password)

	require.NoError(t, err)
}

func TestErrCheckPass(t *testing.T) {
	hash := "$2a$04$lmQ..8n/.jxq4uNVcrfUaO9A/b4q.xaS5OLwTmGU/55GB92b/0X22"
	password := "testPass1"
	err := CheckPassword(hash, password)

	require.Error(t, err)
}
