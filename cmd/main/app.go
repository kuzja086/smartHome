package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"smartHome/internal/delivery/http/v1/auth"
)

func main() {
	router := httprouter.New()
	// TODO Вынести отдельно
	router.GET("/ping", Ping)

	authHandler := auth.Handler{}
	authHandler.Register(router)

	listener, err := net.Listen("tcp", "127.0.0.1:2607")
	if err != nil {
		panic(err.Error())
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server.Serve(listener)
}

func Ping(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprint("it's work!")))
}
