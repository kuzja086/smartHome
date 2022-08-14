package storage

import (
	"context"

	"smartHome/internal/entity"
)

type UserStorage interface {
	CreateUser(ctx context.Context, user entity.User) (string, error)
	FindByUsername(ctx context.Context, username string) (entity.User, error)
}
