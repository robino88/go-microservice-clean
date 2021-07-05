package mock

import "strings"

/// below stuff is just here to keep on working on the implementation
func FakePriceGenerator(sapIDs string) []*PriceResp {
	var prices []*PriceResp
	for _, sapId := range strings.Split(sapIDs, ",") {
		prices = append(prices, newPriceResp(sapId, 100000000))
	}
	return prices
}

type PriceResp struct {
	SapID string
	Price int64
}

func newPriceResp(sapID string, price int64) *PriceResp {
	return &PriceResp{SapID: sapID, Price: price}
}
