package storage

import (
	"context"

	"smartHome/internal/entity"
)

type UserStorage interface {
	CreateUser(ctx context.Context, user user.User) (string, error)
	FindOne(ctx context.Context, id string) (user.User, error)
}
