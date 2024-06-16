package models

type ErrorResponse struct {
	Status int     `json:"status,omitempty"`
	Errors []Error `json:"errors,omitempty"`
}

type Error struct {
	ErrorKey         string `json:"error_key,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}
