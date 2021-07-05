package mock

import (
	"strconv"
	"strings"
)

/// below stuff is just here to keep on working on the implementation
func FakePriceGenerator(sapIDs string) []*PriceResp {
	var prices []*PriceResp
	for _, sapId := range strings.Split(sapIDs, ",") {
		prices = append(prices, newPriceResp(sapId, 100000000))
	}
	return prices
}

func FakeShippingCostCalculator(postalCode string) int {
	atoi, err := strconv.Atoi(postalCode)
	if err != nil {
		return 1
	}
	return atoi
}

type PriceResp struct {
	SapID string
	Price int64
}

func newPriceResp(sapID string, price int64) *PriceResp {
	return &PriceResp{SapID: sapID, Price: price}
}
