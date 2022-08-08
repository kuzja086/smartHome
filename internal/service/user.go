package service

import (
	"context"
	"os/user"
	"smartHome/pkg/logging"
)

type UserService struct {
	storage UserStorage
	logger  *logging.Logger
}

type UserStorage interface {
	CreateUser(ctx context.Context, user user.User) (string, error)
	FindOne(ctx context.Context, id string) (user.User, error)
}
