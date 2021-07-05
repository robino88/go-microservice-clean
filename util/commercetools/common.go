package commercetools

import (
	"encoding/json"
	"github.com/robino88/go-microservice-clean/util/mock"
)

type Response struct {
	Actions []interface{} `json:"actions"`
}

type Request struct {
	Action   string `json:"action"`
	Resource *struct {
		TypeID string `json:"typeId"`
		ID     string `json:"id"`
		Cart   *Cart  `json:"obj"`
	} `json:"resource"`
}

func NewErrorResponse(code string, message string) []byte {
	type Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	err, _ := json.Marshal(Error{
		Code:    code,
		Message: message,
	})
	return err
}

func NewUpdateResponse(updateActions []interface{}) []byte {
	response, _ := json.Marshal(Response{
		Actions: updateActions,
	})
	return response
}

func CreateUpdateActionForCustomerKeyAppend(customerKey string) []interface{} {
	var updateActions []interface{}
	updateActions = append(updateActions, setCustomType("6e9f44ed-542f-4792-a22f-ffecb6392044"))
	updateActions = append(updateActions, setCustomField("customer-sap-id", customerKey))
	return updateActions
}

func CreateUpdateActionForLineItemPrices(items []*LineItem, prices []*mock.PriceResp, code string) []interface{} {
	var updateActions []interface{}
	if items != nil {
		return nil
	}

	for _, price := range prices {
		id := getLineItemId(items, price.SapID)
		updateActions = append(updateActions,
			setLineItemPrice(id, BaseMoney{
				Type:           "centPrecision",
				CurrencyCode:   code,
				CentAmount:     price.Price,
				FractionDigits: 2,
			}))
	}
	return updateActions
}

func CreateUpdateActionForSurCharges(items []*CustomLineItem, prices []*mock.PriceResp, code string) []interface{} {
	var updateActions []interface{}
	if items == nil {
		return nil
	}

	for _, lineItem := range items {
		if lineItem.Slug == "crane" {
			//updateActions = append(updateActions,)
		}
	}

	return updateActions
}

func CreateUpdateActionShippingCost(postalCode string) []interface{} {
	var updateActions []interface{}

	return updateActions
}
