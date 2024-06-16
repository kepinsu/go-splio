package coupon

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kepinsu/go-splio/internal"
	"github.com/kepinsu/go-splio/models"
)

// Coupon provides methods to interact with the Splio Coupon API.
// Coupon API docs: https://dev-scp.splio.com/reference/get_data-coupons-categories
type Coupon interface {
	ListCoupons(
		ctx context.Context,
		id int,
		queries models.ListCouponsQuery,
	) (models.ListCoupons, models.ErrorResponse, error)

	ListCouponCategories(
		ctx context.Context,
		queries models.ListCouponsQuery,
	) (models.ListCouponsCategory, models.ErrorResponse, error)
}

type Requester struct {
	H internal.HTTPRequest
}

const (
	sendGetCoupons         = "/data/coupons/categories/%d"
	sendGetCouponsCategory = "/data/coupons/categories"
)

func (r *Requester) ListCoupons(
	ctx context.Context,
	id int,
	queries models.ListCouponsQuery,
) (models.ListCoupons, models.ErrorResponse, error) {
	// Prepare query
	query := make(map[string]string)
	if queries.PageNumber > 0 {
		query["page_number"] = strconv.Itoa(queries.PageNumber)
	}
	if queries.PerPage > 0 {
		query["per_page"] = strconv.Itoa(queries.PerPage)
	}
	var response models.ListCoupons
	resp, err := r.H.Get(ctx, fmt.Sprintf(sendGetCoupons, id), nil, query, &response)
	return response, resp, err
}

func (r *Requester) ListCouponCategories(
	ctx context.Context,
	queries models.ListCouponsQuery,
) (models.ListCouponsCategory, models.ErrorResponse, error) {
	// Prepare query
	query := make(map[string]string)
	if queries.PageNumber > 0 {
		query["page_number"] = strconv.Itoa(queries.PageNumber)
	}
	if queries.PerPage > 0 {
		query["per_page"] = strconv.Itoa(queries.PerPage)
	}
	var response models.ListCouponsCategory
	resp, err := r.H.Get(ctx, sendGetCouponsCategory, nil, query, &response)
	return response, resp, err
}
