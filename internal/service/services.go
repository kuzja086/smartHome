package service

import (
	"context"

	entityUsers "github.com/kuzja086/smartHome/internal/entity/users"
)

type User interface {
	CreateUser(ctx context.Context, user entityUsers.CreateUserDTO) (string, error)
	FindByUsername(ctx context.Context, username string) (entityUsers.User, error)
	// FindOne(ctx context.Context, id string) (entity.User, error)
}
