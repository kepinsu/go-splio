package models

import "testing"

func TestCustomFieldBody_Validate(t *testing.T) {
	type fields struct {
		Scope      Scope
		Name       string
		CustomType CustomType
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CustomFieldBody{
				Scope:      tt.fields.Scope,
				Name:       tt.fields.Name,
				CustomType: tt.fields.CustomType,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("CustomFieldBody.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
