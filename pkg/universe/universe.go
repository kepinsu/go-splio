package universe

import (
	"context"

	"github.com/kepinsu/go-splio/internal"
	"github.com/kepinsu/go-splio/models"
)

// Constant for all paths
const (
	sendUniversesConsumption = "/universes/consumption"
	sendUniversesUniqueKey   = "/universes/unique-key"
)

// Universe provides methods to interact with the Splio Universe API.
// Universe API docs: https://dev-scp.splio.com/reference/get_universes-consumption
type Universe interface {
	UniverseConsumption(
		ctx context.Context,
	) (models.UniverseConsumption, models.ErrorResponse, error)

	UniverseUniqueKey(
		ctx context.Context,
	) (models.UniverseUniqueKey, models.ErrorResponse, error)
}

type Requester struct {
	H internal.HTTPRequest
}

func (r *Requester) UniverseConsumption(
	ctx context.Context,
) (models.UniverseConsumption, models.ErrorResponse, error) {
	var universe models.UniverseConsumption
	resp, err := r.H.Get(ctx, sendUniversesConsumption, nil, nil, &universe)
	return universe, resp, err
}

func (r *Requester) UniverseUniqueKey(
	ctx context.Context,
) (models.UniverseUniqueKey, models.ErrorResponse, error) {
	var universe models.UniverseUniqueKey
	resp, err := r.H.Get(ctx, sendUniversesUniqueKey, nil, nil, &universe)
	return universe, resp, err
}
