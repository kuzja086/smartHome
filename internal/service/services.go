package service

import (
	"context"

	"smartHome/internal/entity"
)

type User interface {
	CreateUser(ctx context.Context, user entity.CreateUserDTO) (string, error)
	// FindOne(ctx context.Context, id string) (entity.User, error)
}
