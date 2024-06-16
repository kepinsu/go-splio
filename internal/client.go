package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"

	"github.com/kepinsu/go-splio/models"
)

const (
	baseURL = "https://api.splio.com"
)

type HTTPRequest interface {
	Delete(ctx context.Context, resourcePath string, body interface{},
		queryParams map[string]string, response interface{}) (models.ErrorResponse, error)
	Get(ctx context.Context, resourcePath string, body interface{},
		queryParams map[string]string, response interface{}) (models.ErrorResponse, error)
	Patch(ctx context.Context, resourcePath string, body interface{},
		queryParams map[string]string, response interface{}) (models.ErrorResponse, error)
	Post(ctx context.Context, resourcePath string, body interface{},
		queryParams map[string]string, response interface{}) (models.ErrorResponse, error)
}

type Client struct {
	// API Key for authentication
	APIKey string
	// Bearer from Authentication endpoint
	Bearer     string
	baseURL    string
	HTTPClient *http.Client
}

func NewClient(h *http.Client, APIkey string) (*Client, error) {
	if APIkey == "" {
		return nil, fmt.Errorf("missing api key")
	}
	if h == nil {
		h = http.DefaultClient
	}
	// Overwrite the http.Transport with the authentication
	h.Transport = &transport{Base: h.Transport, APIKey: APIkey}
	c := &Client{
		HTTPClient: h,
		APIKey:     APIkey,
		baseURL:    baseURL,
	}
	return c, nil

}

func (h *Client) Post(
	ctx context.Context,
	resourcePath string,
	body interface{},
	queryParams map[string]string,
	response interface{},
) (models.ErrorResponse, error) {

	var b io.Reader
	if body != nil {
		rawByte, err := json.Marshal(body)
		if err != nil {
			return models.ErrorResponse{}, err
		}
		b = bytes.NewBuffer(rawByte)
	}

	req, err := h.newRequest(ctx, http.MethodPost, resourcePath, b, queryParams)
	if err != nil {
		return models.ErrorResponse{}, err
	}
	return h.do(req, response)
}

func (h *Client) Get(
	ctx context.Context,
	resourcePath string,
	body interface{},
	queryParams map[string]string,
	response interface{},
) (models.ErrorResponse, error) {

	req, err := h.newRequest(ctx, http.MethodGet, resourcePath, nil, queryParams)
	if err != nil {
		return models.ErrorResponse{}, err
	}
	return h.do(req, response)
}

func (h *Client) Delete(
	ctx context.Context,
	resourcePath string,
	body interface{},
	queryParams map[string]string,
	response interface{},
) (models.ErrorResponse, error) {

	req, err := h.newRequest(ctx, http.MethodDelete, resourcePath, nil, queryParams)
	if err != nil {
		return models.ErrorResponse{}, err
	}
	return h.do(req, response)
}
func (h *Client) Patch(
	ctx context.Context,
	resourcePath string,
	body interface{},
	queryParams map[string]string,
	response interface{},
) (models.ErrorResponse, error) {

	var b io.Reader
	if body != nil {
		rawByte, err := json.Marshal(body)
		if err != nil {
			return models.ErrorResponse{}, err
		}
		b = bytes.NewBuffer(rawByte)
	}

	req, err := h.newRequest(ctx, http.MethodPatch, resourcePath, b, queryParams)
	if err != nil {
		return models.ErrorResponse{}, err
	}
	return h.do(req, response)
}

// newRequest verifies, constructs and authorize an HTTP request.
func (h *Client) newRequest(
	ctx context.Context,
	method string,
	resourcePath string,
	body io.Reader,
	queryParams map[string]string,
) (*http.Request, error) {

	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", h.baseURL, resourcePath), body)
	if err != nil {
		return nil, err
	}
	req.Header = h.generateCommonHeaders()
	req.URL.RawQuery = generateQueryParams(queryParams)
	return req, nil
}

func (h *Client) do(r *http.Request, response interface{}) (models.ErrorResponse, error) {

	client := h.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	res, err := client.Do(r)
	if err != nil {
		return models.ErrorResponse{}, err
	}

	// decode the response
	payload, err := io.ReadAll(r.Response.Body)
	if err != nil {
		return models.ErrorResponse{}, err
	}
	defer r.Response.Body.Close()

	switch {
	case res.StatusCode < 200 || res.StatusCode >= 400:
		var response models.ErrorResponse
		if decodeErr := json.Unmarshal(payload, &response); decodeErr != nil {
			return models.ErrorResponse{}, fmt.Errorf("error %w decoding the response for an HTTP error code: %d",
				err, res.StatusCode)
		}
	case res.StatusCode == http.StatusNoContent:
		return models.ErrorResponse{}, nil
	case res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden:
		return models.ErrorResponse{}, fmt.Errorf("status code %d and text %s", res.StatusCode, payload)
	case res.StatusCode >= 200 || res.StatusCode >= 300:
		if decodeErr := json.Unmarshal(payload, &response); decodeErr != nil {
			return models.ErrorResponse{}, fmt.Errorf("error %w decoding the response for an HTTP error code: %d",
				err, res.StatusCode)
		}
	}
	return models.ErrorResponse{}, nil

}

func (h *Client) generateCommonHeaders() http.Header {
	header := http.Header{}
	header.Add("Accept", "application/json")
	if h.Bearer != "" {
		header.Add("Authorization", fmt.Sprintf("Bearer %s", h.Bearer))
	}
	header.Add("Content-Type", "application/json")
	header.Add("User-Agent", "splio/go-sdk/v1"+" go/"+runtime.Version())
	return header
}

func generateQueryParams(params map[string]string) string {
	q := url.Values{}
	for name, value := range params {
		if value != "" {
			q.Add(name, value)
		}
	}
	return q.Encode()
}
