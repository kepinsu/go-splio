package contact

import (
	"context"
	"strconv"

	"github.com/kepinsu/go-splio/internal"
	"github.com/kepinsu/go-splio/models"
)

type Requester struct {
	H internal.HTTPRequest
}

// List of path for any splio call
const (
	getContact = "/data/contacts"
)

func (r *Requester) ListContact(
	ctx context.Context,
	queries models.ListSearchContactsQuery,
) (models.ListSearchContacts, models.ErrorResponse, error) {
	// Prepare query
	query := make(map[string]string)
	if queries.PerPage > 0 {
		query["per_page"] = strconv.Itoa(queries.PerPage)
	}
	if queries.PageNumber != "" {
		query["page.number"] = queries.PageNumber
	}

	var response models.ListSearchContacts
	resp, err := r.H.Get(ctx, getContact, queries, nil, &response)
	return response, resp, err
}
