package models

import "fmt"

type Interaction struct {
	Channel      string `json:"channel,omitempty"`
	Event        string `json:"event,omitempty"`
	ExtId        string `json:"ext_id,omitempty"`
	ExtProvider  string `json:"ext_provider,omitempty"`
	IndividualID string `json:"individual_id,omitempty"`
	SourceDate   string `json:"source_date,omitempty"`
	SourceID     string `json:"source_id,omitempty"`
	ValueDate    string `json:"value_date,omitempty"`
}

func (i *Interaction) Validate() error {

	if i.Event == "" {
		return fmt.Errorf("missing event field")
	}
	if i.IndividualID == "" {
		return fmt.Errorf("missing IndividualID field")
	}
	if i.ValueDate == "" {
		return fmt.Errorf("missing ValueDate field")
	}

	return nil
}

type Interactions struct {
	Events []Interaction `json:"events,omitempty"`
}

type InteractionsResponse struct {
	Errors int `json:"errors,omitempty"`
	Items  []struct {
		Code        int    `json:"code,omitempty"`
		Description string `json:"description,omitempty"`
		ExtID       string `json:"ext_id,omitempty"`
		ExtProvider string `json:"ext_provider,omitempty"`
	} `json:"items,omitempty"`
}
