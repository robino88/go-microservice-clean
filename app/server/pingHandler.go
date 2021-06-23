package server

import (
	"net/http"
)

func (server *Server) HandlePing(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Length", "4")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Pong"))
}
