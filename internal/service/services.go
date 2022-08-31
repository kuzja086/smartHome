package service

import (
	"context"

	entityUsers "github.com/kuzja086/smartHome/internal/entity/users"
)

type User interface {
	CreateUser(ctx context.Context, user entityUsers.CreateUserDTO) (string, error)
	Auth(ctx context.Context, user entityUsers.AuthDTO) (string, error)
}
