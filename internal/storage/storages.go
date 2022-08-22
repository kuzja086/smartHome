package storage

import (
	"context"

	usersEntity "github.com/kuzja086/smartHome/internal/entity/users"
)

type UserStorage interface {
	CreateUser(ctx context.Context, user usersEntity.User) (string, error)
	FindByUsername(ctx context.Context, username string) (usersEntity.User, error)
}
