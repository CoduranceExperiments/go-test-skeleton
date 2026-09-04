package routes_test

import (
	"testing"

	"github.com/CoduranceExperiments/go-test-skeleton/api/routes"
	"github.com/gin-gonic/gin"
)

// stubVersion lets the error paths of NewRouter be exercised without a real
// API version.
type stubVersion struct{ prefix string }

func (s stubVersion) Prefix() string         { return s.prefix }
func (s stubVersion) Register(_ gin.IRouter) {}

func TestNewRouterErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		versions []routes.Version
		wantErr  bool
	}{
		{
			name:     "no versions",
			versions: nil,
			wantErr:  true,
		},
		{
			name:     "empty prefix",
			versions: []routes.Version{stubVersion{prefix: ""}},
			wantErr:  true,
		},
		{
			name: "duplicate prefix",
			versions: []routes.Version{
				stubVersion{prefix: "v1"},
				stubVersion{prefix: "v1"},
			},
			wantErr: true,
		},
		{
			name: "distinct prefixes",
			versions: []routes.Version{
				stubVersion{prefix: "v1"},
				stubVersion{prefix: "v2"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := routes.NewRouter(tt.versions...)
			if gotErr := err != nil; gotErr != tt.wantErr {
				t.Errorf("NewRouter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestAllVersionsMount guards against a version being added to All with a
// prefix that collides with an existing one.
func TestAllVersionsMount(t *testing.T) {
	t.Parallel()

	if _, err := routes.NewRouter(routes.All()...); err != nil {
		t.Fatalf("NewRouter(All()...) error = %v", err)
	}
}
