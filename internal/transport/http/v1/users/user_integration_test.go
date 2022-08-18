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
