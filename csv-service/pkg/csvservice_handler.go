package pkg

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) newCSVServiceHandler() http.HandlerFunc {
	type errorResponse struct {
		Error error `json:"error"`
	}

	var strResponse []string

	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(errorResponse{Error: err})
			return
		}

		strCtx := r.PostFormValue("query")

		f, _, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(errorResponse{Error: err})
			return
		}

		csvReader := csv.NewReader(f)
		csvReader.Comma = ';'
		csvReader.Comment = '#'
		csvReader.TrimLeadingSpace = true
		csvReader.LazyQuotes = true

		recs, err := csvReader.ReadAll()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(errorResponse{Error: err})
			return
		}

		for _, rec := range recs {
			for i, col := range rec {
				strResponse = append(strResponse, strings.ReplaceAll(strCtx, "$"+strconv.Itoa(i), col))
			}
		}

		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte(strings.Join(strResponse, "\n")))
	}
}
