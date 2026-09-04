package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Version is a mountable version of the service API.
//
// To add a version, implement this interface in its own file and add it to
// All. Nothing else needs to change.
type Version interface {
	// Prefix is the path the version mounts under, without a leading slash,
	// e.g. "v1".
	Prefix() string
	// Register attaches the version's handlers to r, which is already scoped
	// to Prefix.
	Register(r gin.IRouter)
}

// All returns every version the server mounts.
func All() []Version {
	return []Version{
		V1{},
	}
}

// NewRouter returns an engine with each version mounted under its own prefix.
func NewRouter(versions ...Version) (*gin.Engine, error) {
	if len(versions) == 0 {
		return nil, fmt.Errorf("routes: no versions to mount")
	}

	engine := gin.Default()
	seen := make(map[string]struct{}, len(versions))
	for _, v := range versions {
		prefix := v.Prefix()
		if prefix == "" {
			return nil, fmt.Errorf("routes: version %T has an empty prefix", v)
		}
		if _, dup := seen[prefix]; dup {
			return nil, fmt.Errorf("routes: duplicate prefix %q", prefix)
		}
		seen[prefix] = struct{}{}

		v.Register(engine.Group(prefix))
	}
	return engine, nil
}
