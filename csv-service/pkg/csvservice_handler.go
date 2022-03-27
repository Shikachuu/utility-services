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
		Error string `json:"error"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var strResponse []string
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
			return
		}

		strCtx := r.PostFormValue("query")
		f, fi, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
			return
		}
		defer f.Close()

		csvReader := csv.NewReader(f)
		csvReader.Comma = ';'
		csvReader.Comment = '#'
		csvReader.TrimLeadingSpace = true
		csvReader.LazyQuotes = true

		recs, err := csvReader.ReadAll()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
			return
		}

		for _, rec := range recs {
			templateStr := strCtx
			for i, col := range rec {
				templateStr = strings.ReplaceAll(templateStr, "$"+strconv.Itoa(i+1), col)
			}
			strResponse = append(strResponse, templateStr)
		}

		w.Header().Set("Content-Type", "text/plain")
		fileSize := strconv.Itoa(int(fi.Size)) + "byte"
		_ = s.logger.Log("level", "info", "query_length", strconv.Itoa(len(strCtx))+"char", "filesize", fileSize)

		_, _ = w.Write([]byte(strings.Join(strResponse, "\n")))
	}
}
