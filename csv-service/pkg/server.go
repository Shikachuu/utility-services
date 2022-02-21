package pkg

import (
	"github.com/go-kit/log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	logger log.Logger
}

func NewServer(router *mux.Router, logger log.Logger) *Server {
	s := &Server{router: router, logger: logger}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
