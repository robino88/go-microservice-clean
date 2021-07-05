package commercetools

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/robino88/go-microservice-clean/util/mock"
	"net/http"
	"strings"
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
	if items == nil {
		return nil
	}

	for _, price := range prices {
		id, quantity := getLineItemId(items, price.SapID)

		updateActions = append(updateActions,
			setLineItemTotalPrice(id, BaseMoney{
				Type:           "centPrecision",
				CurrencyCode:   code,
				CentAmount:     price.Price,
				FractionDigits: 2,
			}, BaseMoney{
				Type:           "centPrecision",
				CurrencyCode:   code,
				CentAmount:     price.Price * quantity,
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

func CreateUpdateActionShippingCost(CurrencyCode string, shippingCost int) []interface{} {
	var updateActions []interface{}
	updateActions = append(updateActions, setCustomShippingMethod(CurrencyCode, shippingCost))
	return updateActions
}

func GetCartCustomTypeID(ctx context.Context, typeKey string, ct *Client) (string, error) {
	customType, response, err := ct.CustomTypes.GetByKey(ctx, typeKey)
	if err != nil {
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("CT returned code %v please check logs : %v ", response.StatusCode, response.Body))
	}

	return customType.Id, nil
}

func GetSapIDs(items []*LineItem) string {
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

func GetSurchargeCodes(items []*CustomLineItem) string {
	var codes string
	for _, item := range items {
		codes += item.Slug + ","
	}
	return strings.TrimSuffix(codes, ",")
}
