package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// Server is a HTTP server.
type Server struct {
	srv *http.Server
}

// NewServer new a HTTP server.
func NewServer() *Server {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "okay")
		},
	))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	return &Server{srv: srv}
}

// Start start the HTTP server.
func (s *Server) Start() error {
	log.Printf("[HTTP] Listening on: %s\n", s.srv.Addr)
	return s.srv.ListenAndServe()
}

// Shutdown shutdown the HTTP server.
func (s *Server) Shutdown(ctx context.Context) error {
	log.Printf("[HTTP] Shutdown on: %s\n", s.srv.Addr)
	return s.srv.Shutdown(ctx)
}
