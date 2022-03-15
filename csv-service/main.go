package main

import (
	"net/http"
	"os"
	"time"

	"github.com/Shikachuu/utility-services/csv-service/internal"
	"github.com/Shikachuu/utility-services/csv-service/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/log"
)

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "commit", internal.GitCommit)

	_ = logger.Log("level", "info", "msg", "service started")

	if err := run(logger); err != nil {
		_ = logger.Log("level", "error", "encountered an error", err)
		os.Exit(1)
	}
}

func run(logger log.Logger) error {
	router := chi.NewRouter()
	srv := pkg.NewServer(router, logger)

	httpSrv := &http.Server{
		Handler:      srv,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return httpSrv.ListenAndServe()
}
