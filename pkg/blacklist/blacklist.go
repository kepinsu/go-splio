package blacklist

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kepinsu/go-splio/internal"
	"github.com/kepinsu/go-splio/models"
)

const (
	getBlackListedCellPhones = "/data/blacklists/cellphones"
	addTouchpoint            = "/data/blacklists/%s"
	deleteTouchpoint         = "/data/blacklists/%s/%s"
	getBlackListedEmails     = "/data/blacklists/emails"
)

type BlackList interface {
	GetBlackListedCellPhones(
		ctx context.Context,
		queries models.BlackListCellphoneQuery,
	) (models.ListBlackListCellphone, models.ErrorResponse, error)

	AddTouchPoints(
		ctx context.Context,
		channel string,
		body models.TouchPointBody,
	) (models.TouchPointResponse, models.ErrorResponse, error)

	DeleteTouchPoints(
		ctx context.Context,
		channel string,
		source string,
		body models.TouchPointBody,
	) (models.TouchPointResponse, models.ErrorResponse, error)

	GetBlackListedEmail(
		ctx context.Context,
		queries models.BlackListEmailQuery,
	) (models.ListBlackListEmail, models.ErrorResponse, error)
}

type Requester struct {
	H internal.HTTPRequest
}

func (r *Requester) GetBlackListedCellPhones(
	ctx context.Context,
	queries models.BlackListCellphoneQuery,
) (models.ListBlackListCellphone, models.ErrorResponse, error) {
	// Prepare query
	query := make(map[string]string)
	if queries.PageNumber > 0 {
		query["page_number"] = strconv.Itoa(queries.PageNumber)
	}
	if queries.PerPage > 0 {
		query["per_page"] = strconv.Itoa(queries.PerPage)
	}
	if queries.Source != "" {
		query["source"] = queries.Source
	}
	if queries.Term != "" {
		query["term"] = queries.Term
	}
	var response models.ListBlackListCellphone
	resp, err := r.H.Get(ctx, getBlackListedCellPhones, queries, nil, &response)
	return response, resp, err
}

func (r *Requester) AddTouchPoints(
	ctx context.Context,
	channel string,
	body models.TouchPointBody,
) (models.TouchPointResponse, models.ErrorResponse, error) {
	var response models.TouchPointResponse
	resp, err := r.H.Post(ctx, fmt.Sprintf(addTouchpoint, channel), body, nil, &response)
	return response, resp, err
}

func (r *Requester) DeleteTouchPoints(
	ctx context.Context,
	channel string,
	source string,
	body models.TouchPointBody,
) (models.TouchPointResponse, models.ErrorResponse, error) {
	var response models.TouchPointResponse
	resp, err := r.H.Delete(ctx, fmt.Sprintf(deleteTouchpoint, channel, source), body, nil, &response)
	return response, resp, err
}

func (r *Requester) GetBlackListedEmail(
	ctx context.Context,
	queries models.BlackListEmailQuery,
) (models.ListBlackListEmail, models.ErrorResponse, error) {
	// Prepare query
	query := make(map[string]string)
	if queries.PageNumber > 0 {
		query["page_number"] = strconv.Itoa(queries.PageNumber)
	}
	if queries.PerPage > 0 {
		query["per_page"] = strconv.Itoa(queries.PerPage)
	}
	if queries.Source != "" {
		query["source"] = queries.Source
	}
	if queries.Term != "" {
		query["term"] = queries.Term
	}
	var response models.ListBlackListEmail
	resp, err := r.H.Get(ctx, getBlackListedEmails, nil, query, &response)
	return response, resp, err
}
