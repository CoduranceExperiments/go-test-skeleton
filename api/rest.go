package api

import (
	"context"
	"log/slog"
	"net/http"
)

type RestService struct {
	server  *http.Server
	config  *Config
	handler http.Handler
}

// Serve starts the server and blocks until it stops.
//
// TODO(candidate): Serve currently ignores ctx and returns an error even on a
// clean shutdown. Make it shut down gracefully when ctx is cancelled, and have
// cmd/serve.go cancel ctx on SIGINT/SIGTERM. In-flight requests should be
// allowed to finish, under a bounded timeout.
func (s *RestService) Serve(ctx context.Context) error {
	slog.Info("server listening", "addr", s.server.Addr)
	return s.server.ListenAndServe()
}

// Stop gracefully shuts the server down, waiting for in-flight requests until
// ctx is done.
func (s *RestService) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
