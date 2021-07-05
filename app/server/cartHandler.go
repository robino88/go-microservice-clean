package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/robino88/go-microservice-clean/util/commercetools"
	"github.com/robino88/go-microservice-clean/util/mock"
	"net/http"
	"strings"
)

func (s *Server) HandleCartApplyCustomer(w http.ResponseWriter, r *http.Request) {
	s.log.Debug().Msg("HandleCartApplyCustomer called")
	ctx := context.TODO()

	// We parse the request to a workable struct
	request, err := parseRequest(ctx, r)
	if err != nil {
		return //todo: implement proper error
	}

	// we retrieve the customerID from he cart
	customerId := request.Resource.Cart.CustomerId

	// We can use that customerID to retrieve the customerKey From the Customer
	customerKey, err := getCustomerKey(ctx, customerId, s.ct)
	if err != nil {
		return //todo: implement proper error
	}

	// We create a UpdateAction to return back as a response
	actions := commercetools.CreateUpdateActionForCustomerKeyAppend(customerKey)
	response := commercetools.NewUpdateResponse(actions)

	// We create the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

	s.log.Debug().Msg("HandleCartApplyCustomer finished")
}

func (s *Server) HandleCartUpdateLineItems(w http.ResponseWriter, r *http.Request) {
	s.log.Debug().Msg("HandleCartUpdateLineItems called")
	ctx := context.TODO()

	// We parse the request to a workable struct
	request, err := parseRequest(ctx, r)
	if err != nil {
		return //todo: implement proper error
	}
	if request != nil &&
		request.Resource != nil &&
		request.Resource.Cart.LineItems != nil {
		w.WriteHeader(http.StatusOK)
		return
	}
	// We retrieve the lineItems from the cart
	lineItems := request.Resource.Cart.LineItems
	currencyCode := request.Resource.Cart.TotalPrice.CurrencyCode

	sapIds := getSapIDs(lineItems)

	// Do call to service
	//todo: implement real database
	prices := mock.FakePriceGenerator(sapIds)

	// We create a UpdateAction to return back as a response
	actions := commercetools.CreateUpdateActionForLineItemPrices(lineItems, prices, currencyCode)
	response := commercetools.NewUpdateResponse(actions)

	// We create the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

	s.log.Debug().Msg("HandleCartUpdateLineItems finished")
}

func (s *Server) HandleCartUpdateLSurCharges(w http.ResponseWriter, r *http.Request) {
	s.log.Debug().Msg("HandleCartUpdateLSurCharges called")
	ctx := context.TODO()

	// We parse the request to a workable struct
	request, err := parseRequest(ctx, r)
	if err != nil {
		return //todo: implement proper error
	}
	if request != nil &&
		request.Resource != nil &&
		request.Resource.Cart.CustomLineItems != nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	items := request.Resource.Cart.CustomLineItems
	currencyCode := request.Resource.Cart.TotalPrice.CurrencyCode
	surchargeCodes := getSurchargeCodes(items)
	prices := mock.FakePriceGenerator(surchargeCodes)

	// We create a UpdateAction to return back as a response
	actions := commercetools.CreateUpdateActionForSurCharges(items, prices, currencyCode)
	response := commercetools.NewUpdateResponse(actions)

	// We create the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

	s.log.Debug().Msg("HandleCartUpdateLSurCharges finished")
}

func (s *Server) HandleCartUpdateShippingCost(w http.ResponseWriter, r *http.Request) {
	s.log.Debug().Msg("HandleCartUpdateShippingCost called")
	ctx := context.TODO()

	// We parse the request to a workable struct
	request, err := parseRequest(ctx, r)
	if err != nil {
		return //todo: implement proper error
	}
	if request != nil &&
		request.Resource != nil &&
		request.Resource.Cart.ShippingAddress != nil &&
		request.Resource.Cart.ShippingAddress.PostalCode != "" {
		w.WriteHeader(http.StatusOK)
		return
	}

	postalCode := request.Resource.Cart.ShippingAddress.PostalCode

	actions := commercetools.CreateUpdateActionShippingCost(postalCode)
	response := commercetools.NewUpdateResponse(actions)

	// We create the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

	s.log.Debug().Msg("HandleCartUpdateShippingCost finished")
}

func (s *Server) HandleCartExtension(writer http.ResponseWriter, req *http.Request) {
	//we always want to send back the data as json
	writer.Header().Set("Content-Type", "application/json")
	s.printRequest(req.Body)
	writer.WriteHeader(http.StatusOK)
}

func getCustomerKey(ctx context.Context, id string, ct *commercetools.Client) (string, error) {
	customer, response, err := ct.Customer.Get(ctx, id)
	if err != nil {
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("CT returned code %v please check logs : %v ", response.StatusCode, response.Body))
	}

	return customer.Key, nil
}

func getSapIDs(items []*commercetools.LineItem) string {
	var sapIds string
	for _, item := range items {
		sapId := ""
		for _, attribute := range item.Variant.Attributes {
			if attribute.Name == "sap-number" {
				sapId = fmt.Sprintf("%v", attribute.Value)
			}
		}
		sapIds += sapId + ","
	}

	return strings.TrimSuffix(sapIds, ",")
}

func getSurchargeCodes(items []*commercetools.CustomLineItem) string {
	var codes string
	for _, item := range items {
		codes += item.Slug + ","
	}
	return strings.TrimSuffix(codes, ",")
}
