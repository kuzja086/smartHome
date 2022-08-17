package apperror

import (
	"errors"
	"net/http"
)

type apphandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h apphandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *AppError
		err := h(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")

			if errors.As(err, &appErr) {
				if errors.Is(err, NotConfirmPass) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write(NotConfirmPass.Marshal())
					return
				} else if errors.Is(err, EmptyUsername) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write(EmptyUsername.Marshal())
					return
				} else if errors.Is(err, UserExists) {
					w.WriteHeader(http.StatusConflict)
					w.Write(UserExists.Marshal())
					return
				}

				err = err.(*AppError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(appErr.Marshal())
				return
			}
			w.WriteHeader(http.StatusTeapot)
			w.Write(systemError(err).Marshal())
		}
	}
}
