package auth

import (
	"net/http"
	"smartHome/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

type AuthHandler struct {
	logger *logging.Logger
}

func NewAuthHandler(logger *logging.Logger) *AuthHandler {
	return &AuthHandler{
		logger: logger,
	}
}

func (h *AuthHandler) Register(router *httprouter.Router) {
	router.POST("/auth/signup", h.SignUp)
	router.POST("/auth/signin", h.SignIn)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
