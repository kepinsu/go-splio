package gosplio

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kepinsu/go-splio/internal"
	"github.com/kepinsu/go-splio/models"
	"github.com/kepinsu/go-splio/pkg/authentication"
	"github.com/kepinsu/go-splio/pkg/blacklist"
	"github.com/kepinsu/go-splio/pkg/coupon"
	"github.com/kepinsu/go-splio/pkg/export"
	"github.com/kepinsu/go-splio/pkg/fields"
	"github.com/kepinsu/go-splio/pkg/interactions"
	"github.com/kepinsu/go-splio/pkg/universe"
)

// Client is the entrypoint to all Splio channels.
type Client struct {
	apiKey     string
	httpClient *http.Client

	// All interfaces
	Authentication authentication.Authentication
	Universe       universe.Universe
	Fields         fields.Fields
	Exports        export.Export
	BlackList      blacklist.BlackList
	Coupon         coupon.Coupon
	Interaction    interactions.Interactions
}

// NewClient returns a client object using the provided apiKey.
// If a client is not provided using options, a default one is created.
func NewClient(ctx context.Context, apiKey string, options ...func(*Client)) (Client, error) {

	if ctx == nil { // Avoid to had any error
		ctx = context.TODO()
	}

	c := Client{apiKey: apiKey, httpClient: http.DefaultClient}
	for _, opt := range options {
		opt(&c)
	}

	internalHandler, err := internal.NewClient(c.httpClient, apiKey)
	if err != nil {
		return Client{}, err
	}
	c.Authentication = &authentication.Requester{H: internalHandler}

	// Try to do authentication request to get token
	token, resp, err := c.Authentication.Authentication(ctx, models.ApiKey{ApiKey: apiKey})
	if err != nil {
		return Client{}, err
	}
	if resp.Status != 0 { // if the API refuse the authentication
		return Client{}, fmt.Errorf("the splio api refuse the authentication %v", resp.Errors)
	}
	internalHandler.Bearer = token.Token

	c.Universe = &universe.Requester{H: internalHandler}
	c.Fields = &fields.Requester{H: internalHandler}
	c.Exports = &export.Requester{H: internalHandler}
	c.BlackList = &blacklist.Requester{H: internalHandler}
	c.Coupon = &coupon.Requester{H: internalHandler}
	c.Interaction = &interactions.Requester{H: internalHandler}
	return c, nil
}

func WithHTTPClient(httpClient *http.Client) func(*Client) {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
