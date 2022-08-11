package v1user

import (
	"encoding/json"
	"net/http"
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
	router.POST("/auth/signup", h.SignUp)
	router.POST("/auth/signin", h.SignIn)
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("Sign Up")
	var d httpdto.CreateUserDTO
	defer r.Body.Close()

	h.logger.Debug("Decode body")
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		h.logger.Debug([]byte(err.Error()))
		return
	}

	h.logger.Debug("Validate DTO")
	// validate

	h.logger.Debug("Map DTO")
	CreateUserDTO := entity.CreateUserDTO{
		Username:       d.Username,
		Email:          d.Email,
		Password:       d.Password,
		RepeatPassword: d.RepeatPassword}

	id, err := h.userService.CreateUser(r.Context(), CreateUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		h.logger.Debug([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
