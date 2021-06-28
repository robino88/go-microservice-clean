package commercetools

import (
	"context"
	"net/http"
)

type CartService struct {
	client *Client
}

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
