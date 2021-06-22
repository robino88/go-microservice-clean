package main

import (
	"fmt"
	"github.com/robino88/go-microservice-clean/app/router"
	"github.com/robino88/go-microservice-clean/config"
	"log"
	"net/http"
)

//main function running the application
func main() {
	appConfig := config.AppConfig()

	appRouter := router.NewRouter()

	address := fmt.Sprintf(":%d", appConfig.Server.Port)

	log.Printf("Starting server at %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConfig.Server.TimeoutRead,
		WriteTimeout: appConfig.Server.TimeoutWrite,
		IdleTimeout:  appConfig.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}

}
