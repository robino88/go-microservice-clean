package server

import (
	"github.com/robino88/go-microservice-clean/util/commercetools"
	"github.com/robino88/go-microservice-clean/util/logger"
)

type Server struct {
	logger        *logger.Logger
	commercetools *commercetools.Client
}

func NewServer(logger *logger.Logger, commercetools *commercetools.Client) *Server {
	return &Server{logger: logger, commercetools: commercetools}
}

func (server *Server) Logger() *logger.Logger {
	return server.logger
}

func (server *Server) Commercetools() *commercetools.Client {
	return server.commercetools
}
