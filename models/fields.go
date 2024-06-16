package models

import "fmt"

type Scope string

const (
	ScopeContacts     = "contacts"
	ScopeStore        = "store"
	ScopeProduct      = "product"
	ScopeOrderProduct = "order_product"
	ScopeOrder        = "order"
	ScopeRewards      = "rewards"
	ScopeLoyalty      = "loyalty"
)

type CustomType string

const (
	CustomTypeDate   = "date"
	CustomTypeText   = "text"
	CustomTypeInt    = "int"
	CustomTypeDouble = "double"
)

type CustomFieldBody struct {
	Scope      Scope      `json:"scope,omitempty"`
	Name       string     `json:"name,omitempty"`
	CustomType CustomType `json:"type,omitempty"`
}

func (c *CustomFieldBody) Validate() error {
	switch c.Scope {
	case ScopeContacts, ScopeStore, ScopeProduct, ScopeOrderProduct,
		ScopeOrder, ScopeRewards, ScopeLoyalty:
	default:
		return fmt.Errorf("scope field value unknown %s", c.Scope)
	}
	if c.Name == "" {
		return fmt.Errorf("missing name field")
	}
	switch c.CustomType {
	case CustomTypeDate, CustomTypeText, CustomTypeInt, CustomTypeDouble:
	default:
		return fmt.Errorf("type field value unknown %s", c.Scope)
	}
	return nil
}

type CustomFieldResponse struct {
	ID       int        `json:"id,omitempty"`
	Name     string     `json:"name,omitempty"`
	DataType CustomType `json:"data_type,omitempty"`
}

type ContactsFields struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	PerPage      uint `json:"per_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []FieldElement `json:"elements,omitempty"`
}
type ListOrdersFieldsQuery struct {
	PageNumber int
	PerPage    int
}

type ListOrdersFields struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	PerPage      uint `json:"per_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []FieldElement `json:"elements,omitempty"`
}

type ListOrderedProductFieldsQuery struct {
	PageNumber int
	PerPage    int
}
type ListOrderedProductFields struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	PerPage      uint `json:"per_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []FieldElement `json:"elements,omitempty"`
}

type ListProductFieldsQuery struct {
	PageNumber int
	PerPage    int
}
type ListProductFields struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	PerPage      uint `json:"per_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []FieldElement `json:"elements,omitempty"`
}

type ListRewardsFieldsQuery struct {
	PageNumber int
	PerPage    int
}
type ListRewardsFields struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	PerPage      uint `json:"per_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []FieldElement `json:"elements,omitempty"`
}
type ListLoyaltyMembersQuery struct {
	PageNumber int
	PerPage    int
}
type ListLoyaltyMembersFields struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	PerPage      uint `json:"per_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []FieldElement `json:"elements,omitempty"`
}
type ListStoreQuery struct {
	PageNumber int
	PerPage    int
}
type ListStoreFields struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	PerPage      uint `json:"per_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []FieldElement `json:"elements,omitempty"`
}

type FieldElement struct {
	Intid          *int   `json:"intid,omitempty"`
	Id             int    `json:"id,omitempty"`
	IsSystem       bool   `json:"is_system,omitempty"`
	Filling        int    `json:"filling,omitempty"`
	UniqueValues   int    `json:"unique_values,omitempty"`
	Name           string `json:"name,omitempty"`
	TranslationKey string `json:"translation_key,omitempty"`
	Label          string `json:"label,omitempty"`
	FieldKey       string `json:"field_key,omitempty"`
}

type ListFieldsQuery struct {
	PageNumber int
	PerPage    int
}
