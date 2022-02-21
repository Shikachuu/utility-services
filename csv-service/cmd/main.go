package main

import (
	"github.com/Shikachuu/utility-services/csv-service/pkg"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("encountered an error: %s", err)
	}
}

func run() error {
	router := mux.NewRouter()
	srv := pkg.NewServer(router)

	httpSrv := &http.Server{
		Handler:      srv,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return httpSrv.ListenAndServe()
}
