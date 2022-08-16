package service

import (
	"context"

	"github.com/kuzja086/smartHome/internal/entity"
)

type User interface {
	CreateUser(ctx context.Context, user entity.CreateUserDTO) (string, error)
	FindByUsername(ctx context.Context, username string) (entity.User, error)
	// FindOne(ctx context.Context, id string) (entity.User, error)
}
