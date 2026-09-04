package api_test

import (
	"net/http"
	"testing"

	"github.com/CoduranceExperiments/go-test-skeleton/api"
)

func TestNew(t *testing.T) {
	t.Parallel()

	noop := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})

	tests := []struct {
		name    string
		opts    []api.Opt
		wantErr bool
	}{
		{
			name:    "handler is required",
			opts:    []api.Opt{api.WithPort(8080)},
			wantErr: true,
		},
		{
			name:    "handler is enough",
			opts:    []api.Opt{api.WithHandler(noop)},
			wantErr: false,
		},
		{
			name:    "port and handler",
			opts:    []api.Opt{api.WithPort(8080), api.WithHandler(noop)},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv, err := api.New(tt.opts...)
			if gotErr := err != nil; gotErr != tt.wantErr {
				t.Fatalf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && srv == nil {
				t.Error("New() returned a nil service and no error")
			}
		})
	}
}
