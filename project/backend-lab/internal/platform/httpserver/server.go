package httpserver

import (
	"context"
	"net/http"
	"time"
)

type Server struct{ http *http.Server }

func New(address string, handler http.Handler) *Server {
	return &Server{http: &http.Server{Addr: address, Handler: handler, ReadHeaderTimeout: 5 * time.Second}}
}
func (server *Server) ListenAndServe() error              { return server.http.ListenAndServe() }
func (server *Server) Shutdown(ctx context.Context) error { return server.http.Shutdown(ctx) }
