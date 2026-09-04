package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CoduranceExperiments/go-test-skeleton/api/routes"
)

// TestV1 is the worked example for HTTP handler tests in this repo: build a
// router, drive it with httptest, assert on the recorded response.
func TestV1(t *testing.T) {
	t.Parallel()

	router, err := routes.NewRouter(routes.V1{})
	if err != nil {
		t.Fatalf("NewRouter() error = %v", err)
	}

	tests := []struct {
		name     string
		method   string
		path     string
		wantCode int
		wantBody string
	}{
		{
			name:     "status returns ok",
			method:   http.MethodGet,
			path:     "/v1",
			wantCode: http.StatusOK,
			wantBody: `{"status":"ok"}`,
		},
		{
			name:     "unknown path under v1 is not found",
			method:   http.MethodGet,
			path:     "/v1/does-not-exist",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "post is not allowed",
			method:   http.MethodPost,
			path:     "/v1",
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, httptest.NewRequest(tt.method, tt.path, nil))

			if rec.Code != tt.wantCode {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantCode)
			}
			if tt.wantBody != "" && rec.Body.String() != tt.wantBody {
				t.Errorf("body = %q, want %q", rec.Body.String(), tt.wantBody)
			}
		})
	}
}
