package models

type ListCouponsQuery struct {
	PerPage    int
	PageNumber int
}

type ListCoupons struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	PerPage      uint `json:"per_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []Coupon `json:"elements,omitempty"`
}

type Coupon struct {
	Code           string    `json:"code,omitempty"`
	Category       string    `json:"category,omitempty"`
	Available      bool      `json:"available,omitempty"`
	Owner          *Owner    `json:"owner,omitempty"`
	Campaign       *Campaign `json:"campaign,omitempty"`
	ExpirationDate string    `json:"expiration_date,omitempty"`
}

type Owner struct {
	Id        int    `json:"id,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Firstname string `json:"firstname,omitempty"`
}

type Campaign struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ListCouponsCategory struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	PerPage      uint `json:"per_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []CouponCategory `json:"elements,omitempty"`
}

type CouponCategory struct {
	Category  string `json:"category,omitempty"`
	Total     uint   `json:"total,omitempty"`
	Available uint   `json:"available,omitempty"`
	Expired   uint   `json:"expired,omitempty"`
	Assigned  uint   `json:"assigned,omitempty"`
}
