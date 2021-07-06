package commercetools

import (
	"context"
	"net/http"
	"time"
)

//TaxService is the interface to the /tax-categories/ endpoint at commercetools
type TaxService service

func (s *TaxService) GetByKey(ctx context.Context, key string) (*TaxCategory, *http.Response, error) {
	req, err := s.client.newRequest(http.MethodGet, "tax-categories/key="+key, nil)
	if err != nil {
		return nil, nil, err
	}

	var tx *TaxCategory
	resp, err := s.client.do(ctx, req, &tx)
	if err != nil {
		return nil, resp, err
	}

	return tx, resp, nil
}

type TaxCategory struct {
	ID             string    `json:"id"`
	Key            string    `json:"key"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
}
