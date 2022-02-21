package pkg

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	router *mux.Router
}

func NewServer(router *mux.Router) *Server {
	return &Server{router: router}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
