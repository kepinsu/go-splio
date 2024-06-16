package interactions

import (
	"context"

	"github.com/kepinsu/go-splio/internal"
	"github.com/kepinsu/go-splio/models"
)

type Interactions interface {
	NewInteraction(
		ctx context.Context,
		interaction models.Interaction,
	) (models.Interaction, models.ErrorResponse, error)

	NewBulkInteraction(
		ctx context.Context,
		interactions models.Interactions,
	) (models.InteractionsResponse, models.ErrorResponse, error)
}

const (
	sendInteraction     = "/events"
	sendBulkInteraction = "/events/bulk"
)

type Requester struct {
	H internal.HTTPRequest
}

func (r *Requester) NewInteraction(
	ctx context.Context,
	interaction models.Interaction,
) (models.Interaction, models.ErrorResponse, error) {

	// sanity check
	if err := interaction.Validate(); err != nil {
		return models.Interaction{}, models.ErrorResponse{}, err
	}
	var response models.Interaction
	resp, err := r.H.Post(ctx, sendInteraction, interaction, nil, &response)
	return response, resp, err
}

func (r *Requester) NewBulkInteraction(
	ctx context.Context,
	interactions models.Interactions,
) (models.InteractionsResponse, models.ErrorResponse, error) {
	var response models.InteractionsResponse
	resp, err := r.H.Post(ctx, sendInteraction, interactions, nil, &response)
	return response, resp, err
}
