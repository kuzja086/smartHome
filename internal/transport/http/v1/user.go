package v1user

import (
	"encoding/json"
	"net/http"
	"smartHome/internal/apperror"
	"smartHome/internal/entity"
	"smartHome/internal/service"
	httpdto "smartHome/internal/transport/http/v1/dto"
	"smartHome/pkg/logging"

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
		apperror.NewAppError("Incorrect body", "", "", err)
	}

	h.logger.Debug("Validate DTO")
	err := validateRequest(d, h.logger)
	if err != nil {
		return err
	}

	h.logger.Debug("Map DTO")
	CreateUserDTO := entity.CreateUserDTO{
		Username:       d.Username,
		Email:          d.Email,
		Password:       d.Password,
		RepeatPassword: d.RepeatPassword}

	id, err := h.userService.CreateUser(r.Context(), CreateUserDTO)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
	return nil
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) error {
	return apperror.NewAppError("", "", "", nil)
}

func validateRequest(req httpdto.CreateUserDTO, l *logging.Logger) error {
	l.Debug("check password and confirm password")
	if req.Password != req.RepeatPassword {
		l.Info("reapeat pass wrong")
		return apperror.NotConfirmPass
	}

	if req.Username == "" {
		l.Info("Empty username")
		return apperror.EmptyUsername
	}
	return nil
}
