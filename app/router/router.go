package router

import (
	"github.com/go-chi/chi"
	"github.com/robino88/go-microservice-clean/app/server"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.MethodFunc("GET", "/ping", server.HandlePing)

	return r
}
