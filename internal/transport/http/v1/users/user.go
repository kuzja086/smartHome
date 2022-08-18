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
		apperror.NewAppError("Incorrect body", "", "", err)
	}

	h.logger.Debug("Validate DTO")
	err := h.validateRequest(d, h.logger)
	if err != nil {
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
	return apperror.NewAppError("", "", "", nil)
}

func (h *UserHandler) validateRequest(req httpdto.CreateUserDTO, l *logging.Logger) error {
	l.Debug("check password and confirm password")
	if req.Password != req.RepeatPassword {
		l.Info("reapeat pass wrong")
		return apperror.NotConfirmPass
	}
	if req.Password == "" {
		l.Info("empty password")
		return apperror.EmptyPassword
	}

	if req.Username == "" {
		l.Info("empty username")
		return apperror.EmptyUsername
	}
	return nil
}
