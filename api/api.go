package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	defaultPort       = 3000
	readHeaderTimeout = 10 * time.Second
)

// New builds a RestService from opts. A handler is required; everything else
// has a default.
func New(opts ...Opt) (*RestService, error) {
	s := &RestService{
		config: &Config{
			Port: defaultPort,
		},
	}
	for _, o := range opts {
		o(s)
	}
	if s.handler == nil {
		return nil, errors.New("api: no handler configured, pass WithHandler")
	}

	s.server = &http.Server{
		Addr:              fmt.Sprintf(":%d", s.config.Port),
		Handler:           s.handler,
		ReadHeaderTimeout: readHeaderTimeout,
	}
	return s, nil
}
