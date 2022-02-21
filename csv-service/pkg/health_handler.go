package pkg

import (
	"encoding/json"
	"net/http"
)

func (s *Server) newHealthHandler() http.HandlerFunc {
	type response struct {
		http bool
	}
	return func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(response{true})
	}
}
