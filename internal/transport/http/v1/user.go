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
	var d httpdto.CreateUserDTO
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// validate?

	// MAPPING dto.CreateBookDTO --> book_usecase.CreateBookDTO
	CreateUserDTO := entity.CreateUserDTO{
		Username: d.Username,
		Email:    d.Email,
		Password: d.Password}

	id, err := h.userService.CreateUser(r.Context(), CreateUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
