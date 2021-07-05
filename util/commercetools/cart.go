package commercetools

//https://docs.commercetools.com/http-api-projects-carts

type Cart struct {
	Type                  string            `json:"type"`
	ID                    string            `json:"id"`
	Version               int               `json:"version"`
	CustomerId            string            `json:"customerId,omitempty"`
	CustomerEmail         string            `json:"customerEmail,omitempty"`
	CartState             string            `json:"cartState,omitempty"`
	LineItems             []*LineItem       `json:"lineItems"`
	CustomLineItems       []*CustomLineItem `json:"customLineItems"`
	TotalPrice            *BaseMoney        `json:"totalPrice"`
	ShippingAddress       *Address          `json:"shippingAddress,omitempty"`
	BillingAddress        *Address          `json:"billingAddress,omitempty"`
	Country               string            `json:"country,omitempty"`
	ItemShippingAddresses []*Address        `json:"itemShippingAddresses,omitempty"`
}

type UpdateCart struct {
	Version int           `json:"version"`
	Actions []interface{} `json:"actions"`
}

type Variant struct {
	Id         int          `json:"id"`
	Attributes []*Attribute `json:"attributes"`
}

type Attribute struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type LineItem struct {
	Id        string   `json:"id"`
	ProductId string   `json:"productId"`
	Quantity  uint     `json:"quantity"`
	Variant   *Variant `json:"variant"`
}

type CustomLineItem struct {
	Id         string    `json:"id"`
	Money      BaseMoney `json:"money"`
	TotalPrice BaseMoney `json:"totalPrice"`
	Slug       string    `json:"slug"`
	Quantity   uint      `json:"quantity"`
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
