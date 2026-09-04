package api

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestService struct {
	server *http.Server
	config *Config
	router *gin.Engine
}

func (r *RestService) Serve(ctx context.Context) error {
	log.Printf("Listening on 0.0.0.0:%d", r.config.Port)
	// TODO: Implement graceful shutdown...
	return r.server.ListenAndServe()

}

func (r *RestService) Stop(ctx context.Context) error {
	return r.server.Shutdown(context.Background())
}
