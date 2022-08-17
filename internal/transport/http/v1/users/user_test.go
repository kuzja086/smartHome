package v1user

import (
	"testing"

	"github.com/kuzja086/smartHome/internal/apperror"
	httpdto "github.com/kuzja086/smartHome/internal/transport/http/v1/dto"
	"github.com/kuzja086/smartHome/pkg/logging"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	logger := logging.GetLogger()
	h := UserHandler{}
	in := httpdto.CreateUserDTO{
		Username:       "testUser",
		Password:       "testPass",
		Email:          "test@test.ru",
		RepeatPassword: "testPass",
	}

	err := h.validateRequest(in, logger)
	require.NoError(t, err)
}

func TestValidateError(t *testing.T) {
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
	logger := logging.GetLogger()
	h := UserHandler{}
	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			err := h.validateRequest(tCase.in, logger)
			require.Error(t, err)
			require.EqualError(t, tCase.expErr, err.Error())
		})
	}
}
