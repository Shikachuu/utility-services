Version := $(shell git describe --tags --dirty)
GitCommit := $(shell git rev-parse HEAD)
PKG_NAME := $(shell cat go.mod | grep module | sed -e "s/^module\s//")
LDFLAGS := "-s -w -X $(PKG_NAME)/internal.Version=$(Version) -X $(PKG_NAME)/internal.GitCommit=$(GitCommit)"
export GO111MODULE=on

build:
	CGO_ENABLED=0 GOOS=linux go build cmd/main.go -mod=vendor -o bin/csv-service

dist:
	rm -rf bin/csv-service*
	go mod vendor
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/main.go -mod=vendor -o bin/csv-service-amd64 -a -installsuffix cgo -ldflags $(LDFLAGS)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build cmd/main.go -mod=vendor -o bin/csv-service-arm64 -a -installsuffix cgo -ldflags $(LDFLAGS)