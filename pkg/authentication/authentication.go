package authentication

import (
	"context"

	"github.com/kepinsu/go-splio/internal"
	"github.com/kepinsu/go-splio/models"
)

// Authentication provides methods to interact with the Splio Authentication API.
// Authentication API docs: https://dev-scp.splio.com/reference/post_authenticate
type Authentication interface {
	Authentication(
		ctx context.Context,
		APIKey models.ApiKey,
	) (models.Token, models.ErrorResponse, error)

	Ping(
		ctx context.Context,
	) (models.Ping, models.ErrorResponse, error)
}

type Requester struct {
	H internal.HTTPRequest
}

// Constant for all paths
const (
	sendAuthentication = "/authentication"
	sendping           = "/ping"
)

func (r *Requester) Authentication(
	ctx context.Context,
	APIKey models.ApiKey,
) (models.Token, models.ErrorResponse, error) {
	var token models.Token
	resp, err := r.H.Post(ctx, sendAuthentication, APIKey, nil, &token)
	return token, resp, err
}
