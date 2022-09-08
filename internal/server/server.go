package server

import (
	"context"
	"github.com/vadimpk/go-gym-manager-api/internal/config"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:           ":" + cfg.Server.Port,
			Handler:        handler,
			ReadTimeout:    cfg.Server.ReadTimeout,
			WriteTimeout:   cfg.Server.WriteTimeout,
			MaxHeaderBytes: cfg.Server.MaxHeaderMegabytes << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
