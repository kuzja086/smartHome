package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"smartHome/internal/config"
	"smartHome/internal/transport/http/v1/auth"
	"smartHome/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create logger")
	router := httprouter.New()

	cfg := config.GetConfig()

	router.GET("/ping", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) { w.Write([]byte("it's work!")) })

	logger.Info("register handlers")
	authHandler := auth.NewAuthHandler(logger)
	authHandler.Register(router)

	logger.Info("create listener")
	logger.Debugf("create listener %s:%d", cfg.App.Host, cfg.App.Port)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port))
	if err != nil {
		panic(err.Error())
	}

	logger.Info("start server")
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
