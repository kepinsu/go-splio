package models

import "fmt"

type Order struct {
	ExternalID     int         `json:"external_id,omitempty"`
	CreatedAt      string      `json:"created_at,omitempty"`
	OrderedAt      string      `json:"ordered_at,omitempty"`
	ShippingAmount float32     `json:"shipping_amount,omitempty"`
	DiscountAmount float32     `json:"discount_amount,omitempty"`
	TotalPrice     float32     `json:"total_price,omitempty"`
	TaxAmount      float32     `json:"tax_amount,omitempty"`
	Currency       string      `json:"currency,omitempty"`
	SalesPerson    string      `json:"sales_person,omitempty"`
	Store          *OrderStore `json:"store,omitempty"`
	CustomFields   []struct {
		ID    int    `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"custom_fields,omitempty"`
	Contact OrderContact `json:"contact,omitempty"`
}

type OrderContact struct {
	ID        uint   `json:"id,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Email     string `json:"email,omitempty"`
}

type CreateOrder struct {
	ExternalID      string  `json:"external_id,omitempty"`
	ContactID       string  `json:"contact_id,omitempty"`
	CardCode        string  `json:"card_code,omitempty"`
	StoreExternalID string  `json:"store_external_id,omitempty"`
	OrderedAt       string  `json:"ordered_at,omitempty"`
	ShippingAmount  float32 `json:"shipping_amount,omitempty"`
	DiscountAmount  float32 `json:"discount_amount,omitempty"`
	TotalPrice      float32 `json:"total_price,omitempty"`
	TaxAmount       float32 `json:"tax_amount,omitempty"`
	Completed       bool    `json:"completed,omitempty"`
	Currency        string  `json:"currency,omitempty"`
	SalesPerson     string  `json:"sales_person,omitempty"`
	CustomFields    []struct {
		ID    int    `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"custom_fields,omitempty"`
}

func (c *CreateOrder) Validate() error {

	if c.ExternalID == "" {
		return fmt.Errorf("missing external ID")
	}
	if c.ContactID == "" {
		return fmt.Errorf("missing contacgt ID")
	}
	if c.StoreExternalID == "" {
		return fmt.Errorf("missing store external ID")
	}
	return nil
}
