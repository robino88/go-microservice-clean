package server

import "github.com/robino88/go-microservice-clean/util/logger"

type Server struct {
	logger *logger.Logger
}

func NewServer(logger *logger.Logger) *Server {
	return &Server{logger: logger}
}

func (server *Server) Logger() *logger.Logger {
	return server.logger
}
