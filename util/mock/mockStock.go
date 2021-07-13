package mock

import (
	"github.com/robino88/go-microservice-clean/spec"
	"math/rand"
	"strings"
	"time"
)

/// below stuff is just here to keep on working on the implementation
func FakeStockGenerator(plantID string, matIDs string, deliveryDate time.Time) *spec.Availabilities {
	today := time.Now()
	//sleep for 8 seconds to mock sap.

	if today.Before(deliveryDate) {
		time.Sleep(8 * time.Second)
	}

	availabilities := spec.Availabilities{
		Availability: nil,
	}

	for _, matID := range strings.Split(matIDs, ",") {
		quantity := rand.Intn(100)
		availability := "red"
		if quantity > 10 {
			availability = "amber"
		}
		if quantity > 60 {
			availability = "green"

		}
		availabilities.Availability = append(availabilities.Availability, newAvailabilityResp(int32(quantity), matID, availability))
	}
	return &availabilities
}

func newAvailabilityResp(q int32, m string, a string) *spec.Availability {
	return &spec.Availability{Quantity: q, MatID: m, Availability: a}
}
