package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const RouteVersionOne = "v1"

type V1 struct {
	router *gin.Engine
	Path   string
}

func (s *V1) Router() *gin.Engine {
	return s.router
}

func (s *V1) RouteGroup() *gin.RouterGroup {
	group := s.router.Group(s.Path)
	group.GET("/", s.Get)
	return group
}

func (s *V1) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
