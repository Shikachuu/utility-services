GIT_COMMIT := $(shell git rev-parse --short HEAD)
PKG_NAME := $(shell cat go.mod | grep module | sed -e "s/^module\s//")
LDFLAGS := "-s -w -X $(PKG_NAME)/internal.GitCommit=$(GIT_COMMIT)"
export GO111MODULE=on
export CGO_ENABLED=0
export GOOS=linux
export GIT_DISCOVERY_ACROSS_FILESYSTEM=1

build:
	go build -mod=vendor -o bin/csv-service -ldflags "-X $(PKG_NAME)/internal.GitCommit=$(GIT_COMMIT)"

image:
	docker build . -t "ghcr.io/shikachuu/csv-service:$(GIT_COMMIT)"
	docker tag "ghcr.io/shikachuu/csv-service:$(GIT_COMMIT)" "ghcr.io/shikachuu/csv-service:latest"

dist:
	rm -rf bin/csv-service*
	go mod vendor
	GOARCH=amd64 go build -mod=vendor -o bin/csv-service-amd64 -a -installsuffix cgo -ldflags $(LDFLAGS)
	GOARCH=arm64 go build -mod=vendor -o bin/csv-service-arm64 -a -installsuffix cgo -ldflags $(LDFLAGS)