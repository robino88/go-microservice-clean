package commercetools

type CartActions struct{}

//SetLineItemPrice
//Sets the price of a line item and changes the priceMode of the line item to ExternalPrice
//If the price mode of the line item is ExternalPrice and no externalPrice is given, the external
//price is unset and the priceMode is set to Platform.
func (t CartActions) SetLineItemPrice(lineItemId string, price BaseMoney) interface{} {
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
