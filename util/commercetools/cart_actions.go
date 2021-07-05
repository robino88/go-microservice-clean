package commercetools

//SetLineItemPrice
//Sets the price of a line item and changes the priceMode of the line item to ExternalPrice
//If the price mode of the line item is ExternalPrice and no externalPrice is given, the external
//price is unset and the priceMode is set to Platform.
func setLineItemPrice(lineItemId string, price BaseMoney) interface{} {
	type CartActions struct {
		Action        string    `json:"action"`
		LineItemId    string    `json:"lineItemId"`
		ExternalPrice BaseMoney `json:"externalPrice"`
	}

	return CartActions{
		Action:        "setLineItemPrice",
		LineItemId:    lineItemId,
		ExternalPrice: price,
	}
}

func setCustomField(name string, value string) interface{} {
	type CartActions struct {
		Action string `json:"action"`
		Name   string `json:"name"`
		Value  string `json:"value"`
	}

	return CartActions{
		Action: "setCustomField",
		Name:   name,
		Value:  value,
	}
}

func setCustomType(id string) interface{} {
	type CartActions struct {
		Action string `json:"action"`
		Type   struct {
			Id     string `json:"id"`
			TypeID string `json:"typeId"`
		} `json:"type"`
	}

	return CartActions{
		Action: "setCustomType",
		Type: struct {
			Id     string `json:"id"`
			TypeID string `json:"typeId"`
		}{Id: id,
			TypeID: "type"},
	}
}

func changeCustomLineItemMoney(name string, price BaseMoney) interface{} {
	type CartActions struct {
		Action           string    `json:"action"`
		CustomLineItemId string    `json:"customLineItemId"`
		ExternalPrice    BaseMoney `json:"externalPrice"`
	}

	return CartActions{
		Action:           "changeCustomLineItemMoney",
		CustomLineItemId: name,
		ExternalPrice:    price,
	}
}

func setCustomShippingMethod(currencyCode string, centAmount int) interface{} {
	type CartAction struct {
		Action             string `json:"action"`
		ShippingMethodName string `json:"shippingMethodName"`
		ShippingRate       struct {
			Price struct {
				CurrencyCode string `json:"currencyCode"`
				CentAmount   int    `json:"centAmount"`
			} `json:"price"`
		} `json:"shippingRate"`
	}

	return CartAction{
		Action:             "setCustomShippingMethod",
		ShippingMethodName: "external-calculated-shipping",
		ShippingRate: struct {
			Price struct {
				CurrencyCode string `json:"currencyCode"`
				CentAmount   int    `json:"centAmount"`
			} `json:"price"`
		}{Price: struct {
			CurrencyCode string `json:"currencyCode"`
			CentAmount   int    `json:"centAmount"`
		}{CurrencyCode: currencyCode, CentAmount: centAmount}},
	}
}

func getLineItemId(items []*LineItem, sapID string) string {
	for _, lineItem := range items {
		for _, attribute := range lineItem.Variant.Attributes {
			if attribute.Name == "sap-number" && attribute.Value == sapID {
				return lineItem.Id
			}
		}
	}
	return ""
}
