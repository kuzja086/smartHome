package service

import (
	"context"
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
