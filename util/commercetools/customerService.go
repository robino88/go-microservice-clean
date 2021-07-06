package commercetools

import (
	"context"
	"net/http"
)

//CustomerService is the interface to the /customers/ endpoint at commercetools
type CustomerService service

func (s *CustomerService) Get(ctx context.Context, id string) (*Customer, *http.Response, error) {
	req, err := s.client.newRequest(http.MethodGet, "customers/"+id, nil)
	if err != nil {
		return nil, nil, err
	}

	var customer *Customer
	resp, err := s.client.do(ctx, req, &customer)

	return customer, resp, err
}

type Customer struct {
	ID             string     `json:"id"`
	CustomerNumber string     `json:"customerNumber"`
	Key            string     `json:"key"`
	ExternalId     string     `json:"externalId"`
	Version        int        `json:"version"`
	CompanyName    string     `json:"companyName"`
	Email          string     `json:"email"`
	FirstName      string     `json:"firstName"`
	LastName       string     `json:"lastName"`
	Password       string     `json:"password"`
	Addresses      []*Address `json:"addresses,omitempty"`
	CustomerGroup  string     `json:"customerGroup"`
	Custom         string     `json:"custom"`
	Locale         string     `json:"locale"`
}
