package api

import (
	"github.com/gin-gonic/gin"
)

type Opt func(r *RestService)

type Config struct {
	Port int
}

func WithPort(port int) Opt {
	return func(s *RestService) {
		s.config.Port = port
	}
}

func WithRouter(r *gin.Engine) Opt {
	return func(s *RestService) {
		s.router = r
	}
}
