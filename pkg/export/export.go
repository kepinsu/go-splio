package export

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kepinsu/go-splio/internal"
	"github.com/kepinsu/go-splio/models"
)

// Constant for all paths
const (
	sendUniversesConsumption = "/reports/smtprelay/%s/recipients"
)

// Export provides methods to interact with the Splio Export API.
// Export API docs: https://dev-scp.splio.com/reference/get_reports-smtprelay-action-id-recipients
type Export interface {
	ListRecipients(
		ctx context.Context,
		actionID string,
		queries models.ListExportQuery,
	) (models.ListRecipientForwardCampagn, models.ErrorResponse, error)
}

type Requester struct {
	H internal.HTTPRequest
}

func (r *Requester) ListRecipients(
	ctx context.Context,
	actionID string,
	queries models.ListExportQuery,
) (models.ListRecipientForwardCampagn, models.ErrorResponse, error) {
	// Prepare query
	query := make(map[string]string)
	if queries.PageNumber > 0 {
		query["page_number"] = strconv.Itoa(queries.PageNumber)
	}
	if queries.PerPage > 0 {
		query["per_page"] = strconv.Itoa(queries.PerPage)
	}
	var list models.ListRecipientForwardCampagn
	resp, err := r.H.Get(ctx, fmt.Sprintf(sendUniversesConsumption, actionID), nil, query, &list)
	return list, resp, err
}
