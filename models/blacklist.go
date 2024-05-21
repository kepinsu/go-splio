package models

type BlackListCellphoneQuery struct {
	Source     string `json:"source,omitempty"`
	PerPage    int    `json:"per_page,omitempty"`
	PageNumber int    `json:"page_number,omitempty"`
	Term       string `json:"term,omitempty"`
}

type ListBlackListCellphone struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	PerPage      uint `json:"per_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []BlackListCellphoneElement `json:"elements,omitempty"`
}

type BlackListCellphoneElement struct {
	Number       int    `json:"number,omitempty"`
	CreationDate string `json:"creation_date,omitempty"`
	Source       string `json:"source,omitempty"`
}

type TouchPointBody struct {
	Data []string `json:"data,omitempty"`
}

type TouchPointResponse struct {
	Success  []string `json:"success,omitempty"`
	Failures []string `json:"failures,omitempty"`
}

type BlackListEmailQuery struct {
	Source     string `json:"source,omitempty"`
	PerPage    int    `json:"per_page,omitempty"`
	PageNumber int    `json:"page_number,omitempty"`
	Term       string `json:"term,omitempty"`
}

type ListBlackListEmail struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	PerPage      uint `json:"per_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []BlackListCellphoneElement `json:"elements,omitempty"`
}

type BlackListEmailElement struct {
	Number       int    `json:"number,omitempty"`
	CreationDate string `json:"creation_date,omitempty"`
	Source       string `json:"source,omitempty"`
}
