package api

import (
	"net/http"
)

// Opt configures a RestService during New.
type Opt func(*RestService)

type Config struct {
	Port int
}

func WithPort(port int) Opt {
	return func(s *RestService) {
		s.config.Port = port
	}
}

// WithHandler sets the handler the server serves. Taking an http.Handler
// rather than a *gin.Engine keeps the web framework out of this package's API.
func WithHandler(h http.Handler) Opt {
	return func(s *RestService) {
		s.handler = h
	}
}
