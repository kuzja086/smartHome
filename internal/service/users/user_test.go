package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kuzja086/smartHome/internal/apperror"
	entity "github.com/kuzja086/smartHome/internal/entity/users"
	mock_storage "github.com/kuzja086/smartHome/internal/storage/mocks"
	"github.com/kuzja086/smartHome/pkg/logging"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_storage.NewMockUserStorage(ctl)
	ctx := context.Background()
	logger := logging.GetLogger()

	mockReq := entity.CreateUserDTO{
		Username:       "testUser",
		Password:       "testPass",
		Email:          "test@test.ru",
		RepeatPassword: "testPass",
	}

	id := "62f94cdc51e47edc761ab15b"

	repo.EXPECT().FindByUsername(ctx, mockReq.Username).Return(entity.User{}, apperror.UserNotFound).Times(1)
	repo.EXPECT().CreateUser(ctx, gomock.Any()).Return(id, nil).Times(1)

	UseCase := NewUserService(logger, repo)
	idExp, err := UseCase.CreateUser(ctx, mockReq)
	require.NoError(t, err)
	require.Equal(t, id, idExp)
}

func TestCreateUserErr(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_storage.NewMockUserStorage(ctl)
	ctx := context.Background()
	logger := logging.GetLogger()

	mockReq := entity.CreateUserDTO{
		Username:       "testUser",
		Password:       "testPass",
		Email:          "test@test.ru",
		RepeatPassword: "testPass",
	}

	errDB := errors.New("error with db")

	repo.EXPECT().FindByUsername(ctx, mockReq.Username).Return(entity.User{}, nil).Times(1)
	repo.EXPECT().FindByUsername(ctx, mockReq.Username).Return(entity.User{}, errDB).Times(1)

	cases := []struct {
		name   string
		in     entity.CreateUserDTO
		expErr *apperror.AppError
	}{
		{
			name:   "exists",
			in:     mockReq,
			expErr: apperror.UserExists,
		},
		{
			name:   "error db",
			in:     mockReq,
			expErr: apperror.NewAppError("error with db", "", apperror.InternalError, errDB),
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			UseCase := NewUserService(logger, repo)
			_, err := UseCase.CreateUser(ctx, tCase.in)
			require.Equal(t, err, tCase.expErr)
		})
	}
}

func TestAuth(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_storage.NewMockUserStorage(ctl)
	ctx := context.Background()
	logger := logging.GetLogger()

	mockReq := entity.AuthDTO{
		Username: "testUser",
		Password: "testPass",
	}

	id := "62f94cdc51e47edc761ab15b"

	mockResp := entity.User{
		ID:           id,
		Username:     "testUser",
		Email:        "test@test.ru",
		PasswordHash: "$2a$04$lmQ..8n/.jxq4uNVcrfUaO9A/b4q.xaS5OLwTmGU/55GB92b/0X22",
	}

	repo.EXPECT().FindByUsername(ctx, mockReq.Username).Return(mockResp, nil).Times(1)

	UseCase := NewUserService(logger, repo)
	idExp, err := UseCase.Auth(ctx, mockReq)
	require.NoError(t, err)
	require.Equal(t, id, idExp)
}

func TestErrAuth(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_storage.NewMockUserStorage(ctl)
	ctx := context.Background()
	logger := logging.GetLogger()

	mockReq := entity.AuthDTO{
		Username: "testUser",
		Password: "testPass1",
	}

	id := "62f94cdc51e47edc761ab15b"
	mockResp := entity.User{
		ID:           id,
		Username:     "testUser",
		Email:        "test@test.ru",
		PasswordHash: "$2a$04$lmQ..8n/.jxq4uNVcrfUaO9A/b4q.xaS5OLwTmGU/55GB92b/0X22",
	}

	repo.EXPECT().FindByUsername(ctx, mockReq.Username).Return(mockResp, nil).Times(1)

	repo.EXPECT().FindByUsername(ctx, mockReq.Username).Return(entity.User{}, apperror.UserNotFound).Times(1)

	cases := []struct {
		name   string
		in     entity.AuthDTO
		expErr *apperror.AppError
	}{
		{
			name:   "not exists",
			in:     mockReq,
			expErr: apperror.AuthFaild,
		},
		{
			name:   "compare pass",
			in:     mockReq,
			expErr: apperror.AuthFaild,
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			UseCase := NewUserService(logger, repo)
			_, err := UseCase.Auth(ctx, mockReq)
			require.ErrorIs(t, err, tCase.expErr)
		})
	}
}
