package commercetools

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/robino88/go-microservice-clean/config"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type service struct {
	client *Client
}

// A Client manages communication with Commercetools
type Client struct {
	client  *http.Client
	baseURL *url.URL
	config  config.CommercetoolsConfig

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	// Services used for talking to different parts of the Commercetools API.
	Project *ProjectService
	Carts   *CartService
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(code string, message string) []byte {
	err, _ := json.Marshal(ErrorResponse{
		Code:    code,
		Message: message,
	})
	return err
}

// NewClient returns a new Commercetools API client. To use API methods which require
// authentication, provide a valid private or personal token.
func NewClient(ctx context.Context, config config.CommercetoolsConfig) *Client {
	c := &Client{}
	authConfig := clientcredentials.Config{
		ClientID:     config.ClientId,
		ClientSecret: config.ClientSecret,
		TokenURL:     strings.Join([]string{config.OauthUrl, "oauth", "token"}, "/"),
		Scopes:       config.Scopes,
	}

	client := authConfig.Client(ctx)

	c.client = client

	c.baseURL = config.ApiUrl
	c.baseURL.Path += "/" + config.Project + "/"
	c.config = config

	c.common.client = c
	c.Project = (*ProjectService)(&c.common)
	c.Carts = (*CartService)(&c.common)

	return c
}

func (c *Client) newRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "go-api-extension")
	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.bareDo(ctx, req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

func (c *Client) bareDo(ctx context.Context, req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	return resp, err
}
