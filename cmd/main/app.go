package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/kuzja086/smartHome/internal/config"
	service "github.com/kuzja086/smartHome/internal/service/users"
	mongodbStorage "github.com/kuzja086/smartHome/internal/storage/db/mongodb"
	v1user "github.com/kuzja086/smartHome/internal/transport/http/v1/users"
	"github.com/kuzja086/smartHome/pkg/client/mongodb"
	"github.com/kuzja086/smartHome/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create logger")
	router := httprouter.New()

	cfg := config.GetConfig()

	logger.Info("create mongoDB client")
	mongoDBClient, err := mongodb.NewClient(context.Background(), cfg.MongoDB.HostMDB, cfg.MongoDB.PortMDB, cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Database, cfg.MongoDB.AuthDB)
	if err != nil {
		panic(err)
	}

	logger.Info("storage init")
	storage := mongodbStorage.NewUserStorage(mongoDBClient, cfg.MongoDB.Collection, logger)

	logger.Info("register ping")
	router.GET("/ping", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) { w.Write([]byte("it's work!")) })

	logger.Info("register servisec")
	userService := service.NewUserService(logger, storage)

	logger.Info("register handlers")
	authHandler := v1user.NewUserHandler(logger, userService)
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
