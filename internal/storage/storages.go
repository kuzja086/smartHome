package storage

import (
	"context"

	"smartHome/internal/entity"
)

type UserStorage interface {
	CreateUser(ctx context.Context, user entity.User) (string, error)
	FindOne(ctx context.Context, id string) (entity.User, error)
}
