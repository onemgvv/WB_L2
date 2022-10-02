package server

import (
	"context"
	"github.com/onemgvv/WB_L2/develop/dev11/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			Handler:        handler,
			ReadTimeout:    cfg.HTTP.Timeouts.Read,
			WriteTimeout:   cfg.HTTP.Timeouts.Write,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
