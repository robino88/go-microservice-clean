package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/robino88/go-microservice-clean/util/commercetools"
	"io/ioutil"
	"net/http"
	"strings"
)

func (server *Server) HandleCartExtension(writer http.ResponseWriter, req *http.Request) {
	//we always want to send back the data as json
	writer.Header().Set("Content-Type", "application/json")

	//this peace will log the req and put it back on the body for debugging purposes.
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		server.logger.Error().Err(err).Msg("")
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	server.logger.Debug().Msgf("Request body: %v", string(buf))
	reader := ioutil.NopCloser(bytes.NewBuffer(buf))
	req.Body = reader

	// Serialize the data and return the error is something goes wrong
	resp, err := SerializeResponse(req)
	if err != nil {
		server.logger.Error().Err(err).Msg("")
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(commercetools.NewErrorResponse("InvalidInput",
			"The cart received from commercetools wasn't valid and could not be serialized"))
		return
	}

	cart := resp.Resource.Cart
	server.logger.Debug().
		Msgf("Received Cart: %v (version %v)", cart.ID, cart.Version)

	//only work on carts with active states
	if cart.CartState != "Active" {
		server.logger.Info().
			Msgf("Skipping cart %v because its state is %v", cart.ID, cart.CartState)
		writer.WriteHeader(http.StatusOK)
		return
	}

	// Extracting id's from cart, if no id's are found return error.
	sapIds := extractSapNumbersFromCart(cart)
	if len(sapIds) < 1 {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(commercetools.NewErrorResponse("InvalidInput",
			"The cart did not contain any valid sap-ids please check the cart if if it contained valid products ("+cart.ID+")"))
		return
	}
	server.logger.Debug().
		Msgf("Getting prices for customer %v (%v)", cart.CustomerId, sapIds)

	// Do call to service
	calculatedPrices := fakePriceGenerator(sapIds)
	server.logger.Debug().
		Msgf("Retrieving the custom prices for customer %v (%v)", cart.CustomerId, calculatedPrices[0].price)

	// create updateActions and update request to commercetools and execute
	updateActions := createPriceUpdatesForCart(cart, calculatedPrices)
	updateCart := commercetools.UpdateCart{
		Version: cart.Version,
		Actions: updateActions,
	}

	_, ctResp, err := server.commercetools.Carts.Update(context.TODO(), cart.ID, updateCart)
	server.logger.Debug().Msgf("commercetools update send with the status: %v", ctResp.StatusCode)
	if err != nil {
		server.logger.Error().Err(err).Msg("")
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(commercetools.NewErrorResponse("InvalidOperation",
			"There was an error while updating the cart, Check the logs of the API extension. CartID: "+cart.ID))
		return
	}

	if ctResp.StatusCode == 409 {
		get, _, _ := server.commercetools.Carts.Get(context.TODO(), cart.ID)
		server.logger.Debug().Msgf("got a version conflict trying it again with version %v", get.Version)
		updateCart.Version = get.Version
		_, ctResp, err = server.commercetools.Carts.Update(context.TODO(), cart.ID, updateCart)
		server.logger.Debug().Msgf("commercetools update send with the status: %v", ctResp.StatusCode)
		if err != nil {
			server.logger.Error().Err(err).Msg("")
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write(commercetools.NewErrorResponse("InvalidOperation",
				"There was an error while updating the cart, Check the logs of the API extension. CartID: "+cart.ID))
			return
		}

	}

	jsonReq, _ := json.Marshal(updateCart)
	server.logger.Debug().Msgf("Request body: %v", string(jsonReq))

	writer.WriteHeader(http.StatusOK)
}

//SerializeResponse Just takes the request and
func SerializeResponse(req *http.Request) (*commercetools.UpdateResponse, error) {
	updateResponse := &commercetools.UpdateResponse{}
	if err := json.NewDecoder(req.Body).Decode(updateResponse); err != nil {
		return nil, err
	}
	return updateResponse, nil
}

func createPriceUpdatesForCart(cart *commercetools.Cart, prices []*priceResp) []interface{} {
	var updateActions []interface{}
	for _, price := range prices {
		id := getLineItemId(cart, price.sapID)
		updateActions = append(updateActions,
			commercetools.CartActions{}.SetLineItemPrice(id, commercetools.BaseMoney{
				Type:           "centPrecision",
				CurrencyCode:   cart.TotalPrice.CurrencyCode, //ugly hack but it works
				CentAmount:     price.price,
				FractionDigits: 2,
			}))
	}
	return updateActions
}

func extractSapNumbersFromCart(cart *commercetools.Cart) string {
	var sapIds string
	for _, item := range cart.LineItems {
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

func getLineItemId(cart *commercetools.Cart, sapID string) string {
	for _, lineItem := range cart.LineItems {
		for _, attribute := range lineItem.Variant.Attributes {
			if attribute.Name == "sap-number" && attribute.Value == sapID {
				return lineItem.Id
			}
		}
	}
	return ""
}

/// below stuff is just here to keep on working on the implementation
func fakePriceGenerator(sapIDs string) []*priceResp {
	var prices []*priceResp
	for _, sapId := range strings.Split(sapIDs, ",") {
		prices = append(prices, newPriceResp(sapId, 100000000))
	}
	return prices
}

type priceResp struct {
	sapID string
	price int64
}

func newPriceResp(sapID string, price int64) *priceResp {
	return &priceResp{sapID: sapID, price: price}
}
