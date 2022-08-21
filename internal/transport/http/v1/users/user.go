package v1user

import (
	"encoding/json"
	"net/http"

	"github.com/kuzja086/smartHome/internal/apperror"
	usersEntity "github.com/kuzja086/smartHome/internal/entity/users"
	"github.com/kuzja086/smartHome/internal/service"
	httpdto "github.com/kuzja086/smartHome/internal/transport/http/v1/dto"
	"github.com/kuzja086/smartHome/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	logger      *logging.Logger
	userService service.User
}

func NewUserHandler(logger *logging.Logger, us service.User) *UserHandler {
	return &UserHandler{
		logger:      logger,
		userService: us,
	}
}

func (h *UserHandler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, "/auth/signup", apperror.Middleware(h.SignUp))
	router.HandlerFunc(http.MethodPost, "/auth/signin", apperror.Middleware(h.SignIn))
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) error {
	h.logger.Info("Sign Up")
	var d httpdto.CreateUserDTO
	defer r.Body.Close()

	h.logger.Debug("Decode body")
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		h.logger.Debug([]byte(err.Error()))
		return apperror.NewAppError("Incorrect body", "", "", err)
	}

	h.logger.Debug("Validate DTO")
	err := h.validateCreateRequest(d)
	if err != nil {
		h.logger.Info(err.Error())
		return err
	}

	h.logger.Debug("Map DTO")
	CreateUserDTO := usersEntity.CreateUserDTO{
		Username:       d.Username,
		Email:          d.Email,
		Password:       d.Password,
		RepeatPassword: d.RepeatPassword}

	id, err := h.userService.CreateUser(r.Context(), CreateUserDTO)
	if err != nil {
		return err
	}

	res := usersEntity.CreateUserResp{ID: id}
	json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
	return nil
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) error {
	h.logger.Info("Sign In")
	var d httpdto.AuthDTO
	defer r.Body.Close()

	h.logger.Debug("Decode body")
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		h.logger.Debug(err.Error())
		return apperror.NewAppError("Incorrect body", "", "", err)
	}

	h.logger.Debug("Validate AuthDTO")
	err := h.validateAuthRequest(d)
	if err != nil {
		h.logger.Info(err.Error())
		return err
	}

	h.logger.Debug("Map DTO")
	AuthUserDTO := usersEntity.AuthDTO{
		Username: d.Username,
		Password: d.Password,
	}

	id, autherr := h.userService.Auth(r.Context(), AuthUserDTO)
	if autherr != nil {
		return autherr
	}

	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("iddd", id)
	h.logger.Info(w.Header())
	return nil
}

func (h *UserHandler) validateCreateRequest(req httpdto.CreateUserDTO) error {
	if req.Password != req.RepeatPassword {
		return apperror.NotConfirmPass
	}

	return checkFillUserPassword(req.Username, req.Password)
}

func (h *UserHandler) validateAuthRequest(req httpdto.AuthDTO) error {
	return checkFillUserPassword(req.Username, req.Password)
}

func checkFillUserPassword(username, password string) error {
	if password == "" {
		return apperror.EmptyPassword
	}

	if username == "" {
		return apperror.EmptyUsername
	}
	return nil
}
