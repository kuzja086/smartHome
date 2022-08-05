package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"smartHome/internal/transport/http/v1/auth"
	"smartHome/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create logger")
	router := httprouter.New()
	// TODO Вынести отдельно
	router.GET("/ping", Ping)

	logger.Info("register handlers")
	authHandler := auth.NewAuthHandler(logger)
	authHandler.Register(router)

	logger.Info("create listener")
	listener, err := net.Listen("tcp", "127.0.0.1:2607")
	if err != nil {
		panic(err.Error())
	}

	logger.Info("start server")
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
