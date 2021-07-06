package commercetools

//setLineItemTotalPrice
//Sets the price of a line item and changes the priceMode of the line item to ExternalPrice
//If the price mode of the line item is ExternalPrice and no externalPrice is given, the external
//price is unset and the priceMode is set to Platform.
func setLineItemTotalPrice(lineItemId string, price BaseMoney, totalPrice BaseMoney) interface{} {
	type CartActions struct {
		Action        string `json:"action"`
		LineItemId    string `json:"lineItemId"`
		ExternalPrice struct {
			Price      BaseMoney `json:"price"`
			TotalPrice BaseMoney `json:"totalPrice"`
		} `json:"externalTotalPrice"`
	}

	return CartActions{
		Action:     "setLineItemTotalPrice",
		LineItemId: lineItemId,
		ExternalPrice: struct {
			Price      BaseMoney `json:"price"`
			TotalPrice BaseMoney `json:"totalPrice"`
		}{Price: price, TotalPrice: totalPrice},
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

func setCustomShippingMethod(currencyCode string, centAmount int, taxID string) interface{} {
	type CartAction struct {
		Action             string `json:"action"`
		ShippingMethodName string `json:"shippingMethodName"`
		ShippingRate       struct {
			Price struct {
				CurrencyCode string `json:"currencyCode"`
				CentAmount   int    `json:"centAmount"`
			} `json:"price"`
		} `json:"shippingRate"`
		TaxCategory struct {
			ID     string `json:"id"`
			TypeId string `json:"typeId"`
		} `json:"taxCategory"`
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
		TaxCategory: struct {
			ID     string `json:"id"`
			TypeId string `json:"typeId"`
		}{ID: taxID, TypeId: "tax-category"},
	}
}
