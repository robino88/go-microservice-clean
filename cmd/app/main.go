package main

import (
	"fmt"
	"github.com/robino88/go-microservice-clean/app/router"
	"github.com/robino88/go-microservice-clean/app/server"
	"github.com/robino88/go-microservice-clean/config"
	LOG "github.com/robino88/go-microservice-clean/util/logger"
	"net/http"
)

//main function running the application
func main() {
	appConfig := config.AppConfig()

	logger := LOG.NewLogger(appConfig.Debug)

	server := server.NewServer(logger)

	appRouter := router.NewRouter(server)

	address := fmt.Sprintf(":%d", appConfig.Server.Port)

	logger.Info().Msgf("Starting server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConfig.Server.TimeoutRead,
		WriteTimeout: appConfig.Server.TimeoutWrite,
		IdleTimeout:  appConfig.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}

}
