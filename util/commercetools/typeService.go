package commercetools

import (
	"context"
	"net/http"
)

type TypeService service

func (s *TypeService) GetByKey(ctx context.Context, key string) (*CustomType, *http.Response, error) {
	req, err := s.client.newRequest(http.MethodGet, "types/"+key, nil)

	if err != nil {
		return nil, nil, err
	}

	var customType *CustomType
	resp, err := s.client.do(ctx, req, &customType)

	return customType, resp, err
}

type CustomType struct {
	Id      string `json:"id"`
	Version int    `json:"version"`
	//we dont need the rest... yet
}