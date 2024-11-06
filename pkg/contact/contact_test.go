package contact

import (
	"context"
	"reflect"
	"testing"

	"github.com/kepinsu/go-splio/internal"
	"github.com/kepinsu/go-splio/models"
)

func TestRequester_ListContact(t *testing.T) {
	type fields struct {
		H internal.HTTPRequest
	}
	type args struct {
		ctx     context.Context
		queries models.ListSearchContactsQuery
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.ListSearchContacts
		want1   models.ErrorResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Requester{
				H: tt.fields.H,
			}
			got, got1, err := r.ListContact(tt.args.ctx, tt.args.queries)
			if (err != nil) != tt.wantErr {
				t.Errorf("Requester.ListContact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Requester.ListContact() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Requester.ListContact() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
