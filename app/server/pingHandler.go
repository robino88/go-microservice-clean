package server

import (
	"github.com/robino88/go-microservice-clean/spec"
	"google.golang.org/protobuf/encoding/protojson"
	"io/ioutil"
	"net/http"
)

func (server *Server) HandlePingGET(resp http.ResponseWriter, _ *http.Request) {
	result := &spec.PingResponse{Message: "pong"}

	response, err := protojson.Marshal(result)
	if err != nil {
		server.logger.Warn().Err(err).Msg("Unable to marshall response")
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(response)
}

func (server *Server) HandlePingPOST(resp http.ResponseWriter, req *http.Request) {
	// serialize the data
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		server.logger.Warn().Err(err).Msg("Unable to serialize data")
	}

	// map it to the proto Spec
	request := &spec.PingRequest{}
	err = protojson.Unmarshal(data, request)
	if err != nil {
		server.logger.Warn().Err(err).Msg("Unable to serialize data")
	}

	// create response and marshall it to json
	result := &spec.PingResponse{Message: request.Name}
	response, err := protojson.Marshal(result)
	if err != nil {
		server.logger.Warn().Err(err).Msg("Unable to marshall response")
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(response)
}
