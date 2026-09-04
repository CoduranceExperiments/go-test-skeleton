package routes

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type ServiceAPI interface {
	V1() ServiceAPI
	// TODO: Demonstrate how you would create a V2 API
}

type ServiceRoutes interface {
	Router() *gin.Engine
	RouteGroup() *gin.RouterGroup
	Get(ctx *gin.Context)
}

func New(version string) (ServiceRoutes, error) {
	// TODO: Demonstrate how you would create a V2 API
	switch version {
	case RouteVersionOne:
		v1 := &V1{Path: RouteVersionOne, router: gin.Default()}
		v1.RouteGroup()
		return v1, nil
	default:
		return nil, errors.New("version not supported")
	}
}
