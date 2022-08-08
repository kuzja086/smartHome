package auth

import (
	"net/http"
	"smartHome/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	logger *logging.Logger
}

func NewUserHandler(logger *logging.Logger) *UserHandler {
	return &UserHandler{
		logger: logger,
	}
}

func (h *UserHandler) Register(router *httprouter.Router) {
	router.POST("/auth/signup", h.SignUp)
	router.POST("/auth/signin", h.SignIn)
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
