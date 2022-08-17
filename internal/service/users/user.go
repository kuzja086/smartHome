package service

import (
	"context"
	"errors"

	"github.com/kuzja086/smartHome/internal/apperror"
	"github.com/kuzja086/smartHome/internal/entity"
	"github.com/kuzja086/smartHome/internal/storage"
	"github.com/kuzja086/smartHome/pkg/logging"
)

type UserService struct {
	storage storage.UserStorage
	logger  *logging.Logger
}

func NewUserService(logger *logging.Logger, storage storage.UserStorage) *UserService {
	return &UserService{
		logger:  logger,
		storage: storage,
	}
}

func (u *UserService) CreateUser(ctx context.Context, dto entity.CreateUserDTO) (id string, err error) {
	user, err := entity.NewUser(dto)
	if err != nil {
		u.logger.Errorf("failed to create user due to error %v", err)
		return id, err
	}

	u.logger.Debug("Check user exists")
	_, err = u.storage.FindByUsername(ctx, user.Username)
	if err == nil {
		return id, apperror.UserExists
	} else if !errors.Is(err, apperror.UserNotFound) {
		return id, apperror.NewAppError("error with db", "", apperror.InternalError, err)
	}

	u.logger.Debug("generate password hash")

	id, err = u.storage.CreateUser(ctx, user)

	if err != nil {
		return id, apperror.NewAppError("error create user", "", apperror.ErrorCreateUser, err)
	}

	return id, nil
}

func (u *UserService) FindByUsername(ctx context.Context, username string) (user entity.User, err error) {
	// TODO
	return user, nil
}
