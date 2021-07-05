package server

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/robino88/go-microservice-clean/util/commercetools"
	"github.com/robino88/go-microservice-clean/util/logger"
	"io"
	"io/ioutil"
	"net/http"
)

type Server struct {
	log *logger.Logger
	ct  *commercetools.Client
}

func NewServer(logger *logger.Logger, commercetools *commercetools.Client) *Server {
	return &Server{log: logger, ct: commercetools}
}

func (s *Server) Logger() *logger.Logger {
	return s.log
}

func (s *Server) Commercetools() *commercetools.Client {
	return s.ct
}

func (s *Server) printRequest(body io.ReadCloser) error {
	if !s.Logger().Debug().Enabled() {
		return nil
	}

	buf, err := ioutil.ReadAll(body)
	if err != nil {
		s.log.Error().Err(err).Msg("")
		return err
	}

	s.log.Debug().Msgf("Request body: %v", string(buf))
	reader := ioutil.NopCloser(bytes.NewBuffer(buf))
	body = reader
	return nil
}

func parseRequest(ctx context.Context, r *http.Request) (*commercetools.Request, error) {
	request := &commercetools.Request{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}
	return request, nil
}
