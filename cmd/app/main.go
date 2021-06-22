package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//main function running the application
func main() {
	mux := http.NewServeMux()

	//adding the Greet function to the /ping endpoint
	mux.HandleFunc("/ping", Greet)

	log.Println("Starting server at :8080")

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")

	}

}

//Greet function will write pong back
func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")

}
