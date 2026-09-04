package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// V1 is version 1 of the service API.
type V1 struct{}

func (V1) Prefix() string { return "v1" }

func (v V1) Register(r gin.IRouter) {
	// Register the version root as "" rather than "/" so that GET /v1 answers
	// directly instead of redirecting to GET /v1/.
	r.GET("", v.Status)
}

// Status reports that the service is up.
func (V1) Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
