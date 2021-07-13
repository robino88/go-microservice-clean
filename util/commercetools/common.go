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

func NewUpdateResponse(updateActions []interface{}) []byte {
	response, _ := json.Marshal(Response{
		Actions: updateActions,
	})
	return response
}

func CreateUpdateActionForCustomerKeyAppend(customTypeID string, customerKey string) []interface{} {
	var updateActions []interface{}
	updateActions = append(updateActions, setCustomType(customTypeID))
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

func CreateUpdateActionShippingCost(CurrencyCode string, shippingCost int, taxID string) []interface{} {
	var updateActions []interface{}
	updateActions = append(updateActions, setCustomShippingMethod(CurrencyCode, shippingCost, taxID))
	return updateActions
}

func RequestCartCustomTypeID(ctx context.Context, typeKey string, ct *Client) (string, error) {
	customType, response, err := ct.CustomTypes.GetByKey(ctx, typeKey)
	if err != nil {
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("CT returned code %v please check logs : %v ", response.StatusCode, response.Body))
	}

	return customType.ID, nil
}

func RequestTaxID(ctx context.Context, key string, ct *Client) (string, error) {
	tax, response, err := ct.Taxes.GetByKey(ctx, key)
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("CT returned code %v please check logs : %v ", response.StatusCode, response.Body))
	}

	return tax.ID, nil
}

func RequestCustomerExternalID(ctx context.Context, id string, ct *Client) (string, error) {
	customer, response, err := ct.Customer.Get(ctx, id)
	if err != nil {
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("CT returned code %v please check logs : %v ", response.StatusCode, response.Body))
	}

	return customer.ExternalId, nil
}

func GetSapIDs(items []*LineItem) string {
	var sapIds string
	for _, item := range items {
		sapId := ""
		for _, attribute := range item.Variant.Attributes {
			if attribute.Name == "sap-number" && attribute.Value != "none" {
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

func getLineItemId(items []*LineItem, sapID string) (string, int64) {
	for _, lineItem := range items {
		for _, attribute := range lineItem.Variant.Attributes {
			if attribute.Name == "sap-number" && attribute.Value == sapID {
				return lineItem.Id, lineItem.Quantity
			}
		}
	}
	return "", 0
}

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

type BaseMoney struct {
	Type           string `json:"type"`
	CurrencyCode   string `json:"currencyCode"`
	CentAmount     int64  `json:"centAmount"`
	FractionDigits int    `json:"fractionDigits,omitempty"`
	PreciseAmount  int    `json:"preciseAmount,omitempty"`
}

type Address struct {
	FirstName             string `json:"firstName"`
	LastName              string `json:"lastName"`
	StreetName            string `json:"streetName"`
	StreetNumber          string `json:"streetNumber"`
	AdditionalStreetInfo  string `json:"additionalStreetInfo"`
	PostalCode            string `json:"postalCode"`
	City                  string `json:"city"`
	Country               string `json:"country"`
	Region                string `json:"region"`
	Email                 string `json:"email"`
	Phone                 string `json:"phone"`
	Company               string `json:"company"`
	AdditionalAddressInfo string `json:"additionalAddressInfo"`
}
