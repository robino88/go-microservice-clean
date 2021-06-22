package main

import (
	"fmt"
	"github.com/robino88/go-microservice-clean/config"
	"log"
	"net/http"
)

//main function running the application
func main() {
	appConfig := config.AppConfig()

	mux := http.NewServeMux()

	//adding the Greet function to the /ping endpoint
	mux.HandleFunc("/ping", Greet)

	address := fmt.Sprintf(":%d", appConfig.Server.Port)

	log.Printf("Starting server at %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  appConfig.Server.TimeoutRead,
		WriteTimeout: appConfig.Server.TimeoutWrite,
		IdleTimeout:  appConfig.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}

}

//Greet function will write pong back
func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")

}
