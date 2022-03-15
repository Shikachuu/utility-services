package pkg

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/log"
	"net/http"
)

type Server struct {
	router *chi.Mux
	logger log.Logger
}

func NewServer(router *chi.Mux, logger log.Logger) *Server {
	s := &Server{router: router, logger: logger}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
