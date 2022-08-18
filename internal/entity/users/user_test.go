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
