package router

import (
	"github.com/go-chi/chi"
	"github.com/robino88/go-microservice-clean/app/handler"
	"github.com/robino88/go-microservice-clean/app/server"
)

func NewRouter(server *server.Server) *chi.Mux {
	logger := server.Logger()
	router := chi.NewRouter()

	router.Method("GET", "/ping", handler.NewHandler(server.HandlePingGET, logger))
	router.Method("POST", "/ping", handler.NewHandler(server.HandlePingPOST, logger))

	return router
}
