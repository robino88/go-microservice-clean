package commercetools

import (
	"context"
	"net/http"
)

type CustomerService struct {
	client *Client
}

func (s *CustomerService) Get(ctx context.Context, id string) (*Customer, *http.Response, error) {
	req, err := s.client.newRequest(http.MethodGet, "customers/"+id, nil)
	if err != nil {
		return nil, nil, err
	}

	var customer *Customer
	resp, err := s.client.do(ctx, req, &customer)

	return customer, resp, err
}
