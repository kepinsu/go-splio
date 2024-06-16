package authentication

import (
	"context"

	"github.com/kepinsu/go-splio/models"
)

func (r *Requester) Ping(
	ctx context.Context,
) (models.Ping, models.ErrorResponse, error) {
	var ping models.Ping
	resp, err := r.H.Get(ctx, sendping, nil, nil, &ping)
	return ping, resp, err
}
