package server

import (
	"context"
	"encoding/json"
	"github.com/robino88/go-microservice-clean/util/commercetools"
	"github.com/robino88/go-microservice-clean/util/mock"
	"net/http"
	"time"
)

//HandleCartApplyCustomer This handle is called upon when the cart is created
// We expect the cart to have a customerID,
// We retrieve the 'eternalID from the customer resource and we retrieve the ID of the type with the key `cart-key`
func (s *Server) HandleCartApplyCustomer(w http.ResponseWriter, r *http.Request) {
	s.log.Debug().Msg("HandleCartApplyCustomer called")
	//s.printRequest(r.Body)
	ctx := context.TODO()

	s.log.Info().Msg("testing the timeout so now we will sleep for 30 seconds....")
	time.Sleep(time.Second * 30)
	s.log.Info().Msg("Yawn.... is the time already over well lets get back to work then. ")

	// We parse the request to a workable struct
	request, err := parseRequest(ctx, r)
	if err != nil {
		s.log.Error().Err(err).Msg("HandleCartApplyCustomer: got a bad request the data was incomplete")
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	if request == nil ||
		request.Resource == nil ||
		request.Resource.Cart == nil {
		s.log.Info().Msg("HandleCartApplyCustomer: Got a bad request the data was incomplete")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if request.Resource.Cart.CustomerId == "" {
		s.log.Info().Msg("HandleCartApplyCustomer: Got a cart without customerID")
		w.WriteHeader(http.StatusOK)
		return
	}

	// We retrieve the customerID from he cart
	// We can use that customerID to retrieve the customerKey From the Customer
	customerId := request.Resource.Cart.CustomerId
	customerKey, err := commercetools.RequestCustomerExternalID(ctx, customerId, s.ct)
	if err != nil {
		s.log.Error().Err(err).Msg("HandleCartApplyCustomer: Couldn't get the customer key")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	customTypeID, err := commercetools.RequestCartCustomTypeID(ctx, "cart-type", s.ct)
	if err != nil {
		s.log.Error().Err(err).Msg("HandleCartApplyCustomer: Couldn't retrieve the custom type that is" +
			" needed to map sap ID to.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// We create a UpdateAction to return back as a response
	actions := commercetools.CreateUpdateActionForCustomerKeyAppend(customTypeID, customerKey)
	response := commercetools.NewUpdateResponse(actions)

	// We create the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	s.log.Info().Msgf("response: %v", string(response))
	w.Write(response)

	s.log.Debug().Msg("HandleCartApplyCustomer finished")
}

//HandleCartUpdateLineItems This handle is called upon when the cart is updated
// once the cart has LineItems we extract the sapIDs from the variant and pass them along, together with
// the sapCustomerID to the database, this will response with the prices for that lineitem.
// these prices are added to the lineItem as externalPrice and externalTotalPrice (price * Quantity)
func (s *Server) HandleCartUpdateLineItems(w http.ResponseWriter, r *http.Request) {
	s.log.Debug().Msg("HandleCartUpdateLineItems called")
	ctx := context.TODO()

	// We parse the request to a workable struct
	request, err := parseRequest(ctx, r)
	if err != nil {
		s.log.Error().Err(err).Msg("HandleCartUpdateLineItems: Got some broken request from commercetools")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if request == nil ||
		request.Resource == nil ||
		request.Resource.Cart == nil ||
		request.Resource.Cart.LineItems == nil {
		s.log.Info().Msg("HandleCartUpdateLineItems: Missing necessary data on request skipping the call")
		w.WriteHeader(http.StatusOK)
		return
	}

	// We retrieve the lineItems from the cart
	// When a cart is created it by default always has a totalPrice with a currencyCode so we can use that for our next
	// requests.
	// We also extract the sapIds from the lineitems in a comma seperated string value
	lineItems := request.Resource.Cart.LineItems
	currencyCode := request.Resource.Cart.TotalPrice.CurrencyCode
	sapIds := commercetools.GetSapIDs(lineItems)
	s.log.Info().Msgf("retrieved sapID's : %v from lineItems", sapIds)

	// Do call to service
	//todo: implement real database
	prices := mock.FakePriceGenerator(sapIds)
	marshal, err := json.Marshal(prices)
	if err != nil {
		s.log.Error().Err(err).Msg("HandleCartUpdateLineItems: Got some broken request from commercetools")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	s.log.Info().Msgf("HandleCartUpdateLineItems: Got the prices from the Backend: %v", string(marshal))

	// We create a UpdateAction to return back as a response
	actions := commercetools.CreateUpdateActionForLineItemPrices(lineItems, prices, currencyCode)
	response := commercetools.NewUpdateResponse(actions)

	// We create the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	s.log.Info().Msgf("response: %v", string(response))
	w.Write(response)
	s.log.Debug().Msg("HandleCartUpdateLineItems finished")
}

//HandleCartUpdateSurCharges This handle is called upon when the cart is updated
//This is a method that will look at the surcharges.
func (s *Server) HandleCartUpdateSurCharges(w http.ResponseWriter, r *http.Request) {
	s.log.Debug().Msg("HandleCartUpdateLSurCharges called")
	ctx := context.TODO()

	// We parse the request to a workable struct
	request, err := parseRequest(ctx, r)
	if err != nil {
		s.log.Error().Err(err).Msg("HandleCartUpdateLSurCharges: Got some broken request from commercetools")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if request == nil ||
		request.Resource == nil ||
		request.Resource.Cart == nil ||
		request.Resource.Cart.CustomLineItems == nil {
		s.log.Info().Msg("HandleCartUpdateLSurCharges: Missing necessary data on request skipping the call")
		w.WriteHeader(http.StatusOK)
		return
	}

	items := request.Resource.Cart.CustomLineItems
	currencyCode := request.Resource.Cart.TotalPrice.CurrencyCode
	surchargeCodes := commercetools.GetSurchargeCodes(items)
	s.log.Info().Msgf("Retrieved the surCharges from the cart:  %v", surchargeCodes)
	prices := mock.FakePriceGenerator(surchargeCodes)

	// We create a UpdateAction to return back as a response
	actions := commercetools.CreateUpdateActionForSurCharges(items, prices, currencyCode)
	response := commercetools.NewUpdateResponse(actions)

	// We create the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	s.log.Info().Msgf("response: %v", string(response))
	w.Write(response)

	s.log.Debug().Msg("HandleCartUpdateLSurCharges finished")
}

//HandleCartUpdateShippingCost This handle is called upon when the cart is updated
// We expect the cart to have a shipping adres (postalcode to be precise) otherwise it will do nothing,
// With the postal code we check the database for the shipping rate and calcluate the shpiing cost.
// We do need to do a call towards commercetools to retrieve the tax category that applies
func (s *Server) HandleCartUpdateShippingCost(w http.ResponseWriter, r *http.Request) {
	s.log.Debug().Msg("HandleCartUpdateShippingCost called")
	ctx := context.TODO()

	// We parse the request to a workable struct
	request, err := parseRequest(ctx, r)
	if err != nil {
		s.log.Error().Err(err).Msg("HandleCartUpdateShippingCost: Got some broken request from commercetools")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if request == nil ||
		request.Resource == nil ||
		request.Resource.Cart == nil ||
		request.Resource.Cart.ShippingAddress == nil ||
		request.Resource.Cart.ShippingAddress.PostalCode == "" {
		s.log.Info().Msg("HandleCartUpdateShippingCost: Missing necessary data on request skipping the call")
		w.WriteHeader(http.StatusOK)
		return
	}

	taxID, err := commercetools.RequestTaxID(ctx, "standard", s.ct)
	if err != nil {
		s.log.Error().Err(err).Msg("HandleCartApplyCustomer: Couldn't get the Tax ID of the standard Tax")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postalCode := request.Resource.Cart.ShippingAddress.PostalCode
	currencyCode := request.Resource.Cart.TotalPrice.CurrencyCode
	shippingCost := mock.FakeShippingCostCalculator(postalCode)

	actions := commercetools.CreateUpdateActionShippingCost(currencyCode, shippingCost, taxID)
	response := commercetools.NewUpdateResponse(actions)

	// We create the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	s.log.Info().Msgf("response: %v", string(response))
	w.Write(response)

	s.log.Debug().Msg("HandleCartUpdateShippingCost finished")
}
