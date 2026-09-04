package api

import (
	"context"
	"fmt"
	"net/http"
)

type Server interface {
	Serve(ctx context.Context) error
	Stop(ctx context.Context) error
}

func New(opt ...Opt) *RestService {
	r := &RestService{
		config: &Config{
			Port: 3000,
		},
	}
	for _, o := range opt {
		o(r)
	}
	r.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", r.config.Port),
		Handler: r.router.Handler(),
	}
	return r
}
