package auth

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
}

func (h *Handler) Register(router *httprouter.Router) {
	router.POST("/auth/signup", h.SignUp)
	router.POST("/auth/signin", h.SignIn)
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
