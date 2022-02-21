package pkg

import (
	"encoding/json"
	"net/http"
)

func (s *Server) newHealthHandler() http.HandlerFunc {
	type response struct {
		Http bool `json:"http"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response{true})
	}
}
