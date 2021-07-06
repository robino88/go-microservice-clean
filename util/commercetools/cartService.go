package commercetools

import (
	"context"
	"net/http"
)

//CartService is the interface to the /carts/ endpoint at commercetools
type CartService service

func (s *CartService) Update(ctx context.Context, id string, body UpdateCart) (*Cart, *http.Response, error) {
	req, err := s.client.newRequest(http.MethodPost, "carts/"+id, body)
	if err != nil {
		return nil, nil, err
	}

	var cart *Cart

	resp, err := s.client.do(ctx, req, &cart)
	if err != nil {
		return nil, resp, err
	}

	return cart, resp, nil
}

func (s *CartService) Get(ctx context.Context, id string) (*Cart, *http.Response, error) {
	req, err := s.client.newRequest(http.MethodGet, "carts/"+id, nil)

	if err != nil {
		return nil, nil, err
	}

	var cart *Cart
	resp, err := s.client.do(ctx, req, &cart)

	return cart, resp, err
}

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
	Quantity  int64    `json:"quantity"`
	Variant   *Variant `json:"variant"`
}

type CustomLineItem struct {
	Id         string    `json:"id"`
	Money      BaseMoney `json:"money"`
	TotalPrice BaseMoney `json:"totalPrice"`
	Slug       string    `json:"slug"`
	Quantity   int64     `json:"quantity"`
}
