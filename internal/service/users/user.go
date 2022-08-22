package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/kuzja086/smartHome/internal/apperror"
	entity "github.com/kuzja086/smartHome/internal/entity/users"
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
	fmt.Println(user.PasswordHash)
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

	u.logger.Debug("create user")

	id, err = u.storage.CreateUser(ctx, user)

	if err != nil {
		return id, apperror.NewAppError("error create user", "", apperror.ErrorCreateUser, err)
	}

	return id, nil
}

func (u *UserService) Auth(ctx context.Context, dto entity.AuthDTO) (string, error) {
	u.logger.Debug("Check user exists")
	user, err := u.storage.FindByUsername(ctx, dto.Username)
	if err != nil {
		u.logger.Debug(err.Error())
		return "", apperror.AuthFaild
	}

	errCheck := entity.CheckPassword(user.PasswordHash, dto.Password)
	if errCheck != nil {
		u.logger.Debug(errCheck.Error())
		return "", apperror.AuthFaild
	}

	return user.ID, nil
}
