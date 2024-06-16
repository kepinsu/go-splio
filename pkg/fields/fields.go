package fields

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kepinsu/go-splio/internal"
	"github.com/kepinsu/go-splio/models"
)

const (
	sendCustomField         = "/data/fields"
	deleteCustomField       = "/data/fields/%s/%s"
	getContactFields        = "/data/fields/contacts"
	getOrdersFields         = "/data/fields/orders"
	getOrderedProductFields = "/data/fields/order-lines"
	getProductFields        = "/data/fields/products"
	getRewardsFields        = "/data/fields/rewards"
	getMembersFields        = "/data/fields/members"
	getStoreFields          = "/data/fields/stores"
)

type Fields interface {
	CreateCustomField(
		ctx context.Context,
		body models.CustomFieldBody,
	) (models.CustomFieldResponse, models.ErrorResponse, error)

	DeleteCustomField(
		ctx context.Context,
		id string,
		scope string,
	) (models.ErrorResponse, error)

	GetContactFields(
		ctx context.Context,
	) (models.ContactsFields, models.ErrorResponse, error)

	ListOrdersFields(
		ctx context.Context,
		id int,
		queries models.ListOrdersFieldsQuery,
	) (models.ListOrdersFields, models.ErrorResponse, error)

	ListOrderedFields(
		ctx context.Context,
		id int,
		queries models.ListOrderedProductFieldsQuery,
	) (models.ListOrderedProductFields, models.ErrorResponse, error)

	ListProductsFields(
		ctx context.Context,
		id int,
		queries models.ListProductFieldsQuery,
	) (models.ListProductFields, models.ErrorResponse, error)

	ListRewardsFields(
		ctx context.Context,
		id int,
		queries models.ListRewardsFieldsQuery,
	) (models.ListRewardsFields, models.ErrorResponse, error)

	ListLoyaltyMembersFields(
		ctx context.Context,
		id int,
		queries models.ListStoreQuery,
	) (models.ListStoreFields, models.ErrorResponse, error)
}

type Requester struct {
	H internal.HTTPRequest
}

func (r *Requester) CreateCustomField(
	ctx context.Context,
	body models.CustomFieldBody,
) (models.CustomFieldResponse, models.ErrorResponse, error) {
	if err := body.Validate(); err != nil {
		return models.CustomFieldResponse{}, models.ErrorResponse{}, err
	}
	var response models.CustomFieldResponse
	resp, err := r.H.Post(ctx, sendCustomField, body, nil, &response)
	return response, resp, err
}

func (r *Requester) DeleteCustomField(
	ctx context.Context,
	id string,
	scope string,
) (models.ErrorResponse, error) {
	return r.H.Delete(ctx, fmt.Sprintf(deleteCustomField, id, scope), nil, nil, nil)
}

func (r *Requester) GetContactFields(
	ctx context.Context,
) (models.ContactsFields, models.ErrorResponse, error) {
	var response models.ContactsFields
	resp, err := r.H.Get(ctx, getContactFields, nil, nil, &response)
	return response, resp, err
}

func (r *Requester) ListOrdersFields(
	ctx context.Context,
	id int,
	queries models.ListOrdersFieldsQuery,
) (models.ListOrdersFields, models.ErrorResponse, error) {
	// Prepare query
	query := make(map[string]string)
	if queries.PageNumber > 0 {
		query["page_number"] = strconv.Itoa(queries.PageNumber)
	}
	if queries.PerPage > 0 {
		query["per_page"] = strconv.Itoa(queries.PerPage)
	}
	var response models.ListOrdersFields
	resp, err := r.H.Get(ctx, getOrdersFields, nil, query, &response)
	return response, resp, err
}

func (r *Requester) ListOrderedFields(
	ctx context.Context,
	id int,
	queries models.ListOrderedProductFieldsQuery,
) (models.ListOrderedProductFields, models.ErrorResponse, error) {
	// Prepare query
	query := make(map[string]string)
	if queries.PageNumber > 0 {
		query["page_number"] = strconv.Itoa(queries.PageNumber)
	}
	if queries.PerPage > 0 {
		query["per_page"] = strconv.Itoa(queries.PerPage)
	}
	var response models.ListOrderedProductFields
	resp, err := r.H.Get(ctx, getOrderedProductFields, nil, query, &response)
	return response, resp, err
}

func (r *Requester) ListProductsFields(
	ctx context.Context,
	id int,
	queries models.ListProductFieldsQuery,
) (models.ListProductFields, models.ErrorResponse, error) {
	// Prepare query
	query := make(map[string]string)
	if queries.PageNumber > 0 {
		query["page_number"] = strconv.Itoa(queries.PageNumber)
	}
	if queries.PerPage > 0 {
		query["per_page"] = strconv.Itoa(queries.PerPage)
	}
	var response models.ListProductFields
	resp, err := r.H.Get(ctx, getProductFields, nil, query, &response)
	return response, resp, err
}

func (r *Requester) ListRewardsFields(
	ctx context.Context,
	id int,
	queries models.ListRewardsFieldsQuery,
) (models.ListRewardsFields, models.ErrorResponse, error) {
	// Prepare query
	query := make(map[string]string)
	if queries.PageNumber > 0 {
		query["page_number"] = strconv.Itoa(queries.PageNumber)
	}
	if queries.PerPage > 0 {
		query["per_page"] = strconv.Itoa(queries.PerPage)
	}
	var response models.ListRewardsFields
	resp, err := r.H.Get(ctx, getProductFields, nil, query, &response)
	return response, resp, err
}

func (r *Requester) ListLoyaltyMembersFields(
	ctx context.Context,
	id int,
	queries models.ListStoreQuery,
) (models.ListStoreFields, models.ErrorResponse, error) {
	// Prepare query
	query := make(map[string]string)
	if queries.PageNumber > 0 {
		query["page_number"] = strconv.Itoa(queries.PageNumber)
	}
	if queries.PerPage > 0 {
		query["per_page"] = strconv.Itoa(queries.PerPage)
	}
	var response models.ListStoreFields
	resp, err := r.H.Get(ctx, getStoreFields, nil, query, &response)
	return response, resp, err
}
