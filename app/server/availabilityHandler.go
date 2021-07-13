package server

import (
	"github.com/robino88/go-microservice-clean/util/mock"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"
)

func (s *Server) HandleProductAvailability(w http.ResponseWriter, r *http.Request) {
	s.log.Debug().Msg("HandleProductAvailability called")

	plantID, ok := r.URL.Query()["plantID"]
	if !ok || len(plantID[0]) < 1 {
		s.log.Debug().Msg("Url Param 'plantID' is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	matIDs, ok := r.URL.Query()["matIDs"]
	if !ok || len(matIDs[0]) < 1 {
		s.log.Debug().Msg("Url Param 'matIDs' is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	deliveryDateQuery, ok := r.URL.Query()["deliveryDate"]
	deliveryDate := time.Now()
	if deliveryDateQuery != nil {
		selectedDate, err := time.Parse("2006-01-02T15:04:05Z-07:00", deliveryDateQuery[0])
		if err != nil {
			s.log.Error().Err(err).Msg("Unable to parse the deliveryDate")
		} else {
			deliveryDate = selectedDate
		}
	}

	fakeResp := mock.FakeStockGenerator(plantID[0], matIDs[0], deliveryDate)

	response, err := protojson.Marshal(fakeResp)
	if err != nil {
		s.log.Warn().Err(err).Msg("Unable to marshall response")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
