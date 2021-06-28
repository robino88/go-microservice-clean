package main

import (
	"context"
	"fmt"
	"github.com/robino88/go-microservice-clean/app/router"
	"github.com/robino88/go-microservice-clean/app/server"
	"github.com/robino88/go-microservice-clean/config"
	"github.com/robino88/go-microservice-clean/util/commercetools"
	"github.com/robino88/go-microservice-clean/util/logger"
	"net/http"
)

//main function running the application
func main() {
	ctx := context.Background()

	appConfig := config.AppConfig()

	log := logger.NewLogger(appConfig.Debug)
	ct := commercetools.NewClient(ctx, appConfig.Commercetools)
	srv := server.NewServer(log, ct)
	appRouter := router.NewRouter(srv)

	address := fmt.Sprintf(":%d", appConfig.Server.Port)

	project, _, err := ct.Project.Get(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("EXIT")
		return
	}
	log.Info().Msgf("Connected to commercetools instance : %s (%s)", project.Name, project.Key)

	log.Info().Msgf("Starting server %v", address)
	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConfig.Server.TimeoutRead,
		WriteTimeout: appConfig.Server.TimeoutWrite,
		IdleTimeout:  appConfig.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Server startup failed")
	}

}
