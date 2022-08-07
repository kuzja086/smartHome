package service

import (
	"smartHome/internal/storage"
	"smartHome/pkg/logging"
)

type UserService struct {
	storage storage.UserStorage
	logger  *logging.Logger
}
