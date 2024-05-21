package models

import "time"

type ListSRecipientForwardCampagn struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	PerPage      uint `json:"per_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []RecipientForwardCampagnElement `json:"elements,omitempty"`
}

type RecipientForwardCampagnElement struct {
	ID            *uint      `json:"id,omitempty"`
	ExtID         *uint      `json:"extid,omitempty"`
	Recipient     string     `json:"rcpt,omitempty"`
	Subject       string     `json:"subject,omitempty"`
	Status        string     `json:"status,omitempty"`
	Smtpmsg       string     `json:"smtpmsg,omitempty"`
	Openings      *uint      `json:"openings,omitempty"`
	LastOpenDate  *time.Time `json:"last_open_date,omitempty"`
	Clicks        *uint      `json:"clicks,omitempty"`
	LastClickData *time.Time `json:"last_click_data,omitempty"`
	Unsub         bool       `json:"unsub,omitempty"`
	SpamCompliant bool       `json:"spam_compliant,omitempty"`
	SalesCount    *uint      `json:"sales_count,omitempty"`
	SalesAmount   *uint      `json:"sales_amount,omitempty"`
}

type ListExportQuery struct {
	PageNumber int
	PerPage    int
}
