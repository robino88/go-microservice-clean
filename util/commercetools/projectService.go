package commercetools

import (
	"context"
	"net/http"
)

type ProjectService service

func (s *ProjectService) Get(ctx context.Context) (*Project, *http.Response, error) {
	req, err := s.client.newRequest(http.MethodGet, "", nil)

	if err != nil {
		return nil, nil, err
	}

	var project *Project
	resp, err := s.client.do(ctx, req, &project)

	return project, resp, err
}

type Project struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}
