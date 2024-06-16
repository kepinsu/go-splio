package internal

import (
	"context"
	"io"
	"net/http"
	"testing"
)

func TestClient_newRequest(t *testing.T) {
	type fields struct {
		Bearer     string
		BaseURL    string
		HTTPClient *http.Client
	}
	type args struct {
		ctx          context.Context
		method       string
		resourcePath string
		body         io.Reader
		queryParams  map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "http.method unknown",
			wantErr: true,
		},
		{
			name: "missing context.context",
			args: args{
				method: http.MethodPatch,
			},
			wantErr: true,
		},
		{
			name: "empty request",
			args: args{
				ctx:    context.TODO(),
				method: http.MethodPatch,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Client{
				Bearer:     tt.fields.Bearer,
				baseURL:    tt.fields.BaseURL,
				HTTPClient: tt.fields.HTTPClient,
			}
			_, err := h.newRequest(tt.args.ctx, tt.args.method, tt.args.resourcePath, tt.args.body, tt.args.queryParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.newRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
