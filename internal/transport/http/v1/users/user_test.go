package v1user

import (
	"testing"

	"github.com/kuzja086/smartHome/internal/apperror"
	httpdto "github.com/kuzja086/smartHome/internal/transport/http/v1/dto"
	"github.com/stretchr/testify/require"
)

func TestCreateUserValidate(t *testing.T) {
	h := UserHandler{}
	in := httpdto.CreateUserDTO{
		Username:       "testUser",
		Password:       "testPass",
		Email:          "test@test.ru",
		RepeatPassword: "testPass",
	}

	err := h.validateCreateRequest(in)
	require.NoError(t, err)
}

func TestCreateUserValidateError(t *testing.T) {
	cases := []struct {
		name   string
		in     httpdto.CreateUserDTO
		expErr *apperror.AppError
	}{
		{
			name: "bad repeat password",
			in: httpdto.CreateUserDTO{
				Password:       "testPass",
				RepeatPassword: "testPasserr",
			},
			expErr: apperror.NotConfirmPass,
		},
		{
			name:   "empty password",
			in:     httpdto.CreateUserDTO{},
			expErr: apperror.EmptyPassword,
		},
		{
			name: "empty username",
			in: httpdto.CreateUserDTO{
				Password:       "testPass",
				RepeatPassword: "testPass",
			},
			expErr: apperror.EmptyUsername,
		},
	}

	h := UserHandler{}
	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			err := h.validateCreateRequest(tCase.in)
			require.Error(t, err)
			require.EqualError(t, tCase.expErr, err.Error())
		})
	}
}

func TestAuthUserValidate(t *testing.T) {
	h := UserHandler{}
	in := httpdto.AuthDTO{
		Username: "testUser",
		Password: "testPass",
	}

	err := h.validateAuthRequest(in)
	require.NoError(t, err)
}

func TestAuthUserValidateError(t *testing.T) {
	cases := []struct {
		name   string
		in     httpdto.AuthDTO
		expErr *apperror.AppError
	}{
		{
			name:   "empty password",
			in:     httpdto.AuthDTO{},
			expErr: apperror.EmptyPassword,
		},
		{
			name: "empty username",
			in: httpdto.AuthDTO{
				Password: "testPass",
			},
			expErr: apperror.EmptyUsername,
		},
	}

	h := UserHandler{}
	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			err := h.validateAuthRequest(tCase.in)
			require.Error(t, err)
			require.EqualError(t, tCase.expErr, err.Error())
		})
	}
}
