package service

import (
	"context"
	"errors"
	"fmt"
	"smartHome/internal/entity"
	"smartHome/internal/storage"
	"smartHome/pkg/logging"
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
	u.logger.Debug("check password and repeat password")
	if dto.Password != dto.RepeatPassword {
		return id, errors.New("reapeat pass")
	}

	u.logger.Debug("generate password hash")
	user, err := entity.NewUser(dto)

	if err != nil {
		u.logger.Errorf("failed to create user due to error %v", err)
		return
	}

	//TODO поиск по логину и паролю:

	id, err = u.storage.CreateUser(ctx, user)

	if err != nil {
		// if errors.Is(err, apperror.ErrNotFound) {
		// 	return id, err
		// }
		return id, fmt.Errorf("failed to create user. error: %w", err)
	}

	return id, nil
}
