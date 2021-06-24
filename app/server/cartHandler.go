package server

import (
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

func (server *Server) HandleCartExtension(resp http.ResponseWriter, req *http.Request) {
	// serialize the data
	data, err := ioutil.ReadAll(req.Body)
	log.Debug().Msgf("Request body: %v", string(data))
	if err != nil {
		server.logger.Warn().Err(err).Msg("Unable to serialize data")
	}

	resp.WriteHeader(http.StatusOK)
}
