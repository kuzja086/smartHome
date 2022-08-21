package v1user_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kuzja086/smartHome/internal/apperror"
	entity "github.com/kuzja086/smartHome/internal/entity/users"
	uservice "github.com/kuzja086/smartHome/internal/service/users"
	mock_storage "github.com/kuzja086/smartHome/internal/storage/mocks"
	v1user "github.com/kuzja086/smartHome/internal/transport/http/v1/users"
	"github.com/kuzja086/smartHome/pkg/logging"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	logger := logging.GetLogger()
	ctx := context.Background()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_storage.NewMockUserStorage(ctl)
	mockReq := entity.CreateUserDTO{
		Username:       "testUser",
		Password:       "testPass",
		Email:          "test@test.ru",
		RepeatPassword: "testPass",
	}
	idexp := "62f94cdc51e47edc761ab15b"

	repo.EXPECT().FindByUsername(ctx, mockReq.Username).Return(entity.User{}, apperror.UserNotFound).Times(1)
	repo.EXPECT().CreateUser(ctx, gomock.Any()).Return(idexp, nil).Times(1)

	userService := uservice.NewUserService(logger, repo)
	h := v1user.NewUserHandler(logger, userService)
	rec := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodPost,
		"/auth/signup",
		bytes.NewBuffer([]byte(
			[]byte(`
			{
				"Username": "testUser",
				"Password": "testPass",
				"Email": "test@test.ru",
				"RepeatPassword": "testPass"
			}
			`),
		)),
	)

	expBody := fmt.Sprintf("{\"id\":\"%s\"}\n", idexp)

	h.SignUp(rec, req)

	res := rec.Result()

	require.Equal(t, 200, res.StatusCode)

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)

	require.Equal(t, expBody, string(data))
}

func TestAuthUser(t *testing.T) {
	logger := logging.GetLogger()
	ctx := context.Background()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_storage.NewMockUserStorage(ctl)
	mockReq := entity.AuthDTO{
		Username: "testUser",
		Password: "testPass",
	}
	id := "62f94cdc51e47edc761ab15b"
	mockResp := entity.User{
		ID:           id,
		Username:     "testUser",
		PasswordHash: "$2a$04$lmQ..8n/.jxq4uNVcrfUaO9A/b4q.xaS5OLwTmGU/55GB92b/0X22",
	}

	repo.EXPECT().FindByUsername(ctx, mockReq.Username).Return(mockResp, nil).Times(1)

	userService := uservice.NewUserService(logger, repo)
	h := v1user.NewUserHandler(logger, userService)
	rec := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodPost,
		"/auth/signin",
		bytes.NewBuffer([]byte(
			[]byte(`
			{
				"username": "testUser",
				"password": "testPass"
			}
			`),
		)),
	)

	h.SignIn(rec, req)

	res := rec.Result()
	fmt.Println(res)
	require.Equal(t, 204, res.StatusCode)
	require.Equal(t, rec.Result().Header.Get("Iddd"), id)
}
