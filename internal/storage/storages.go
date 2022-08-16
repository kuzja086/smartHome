package storage

import (
	"context"

	"github.com/kuzja086/smartHome/internal/entity"
)

type UserStorage interface {
	CreateUser(ctx context.Context, user entity.User) (string, error)
	FindByUsername(ctx context.Context, username string) (entity.User, error)
}
