LDFLAGS := "-s -w"
export GO111MODULE=on
export CGO_ENABLED=0
export GOOS=linux

default:
	rm -rf bin/csv-service*
	go mod vendor
	GOARCH=amd64 go build -mod=vendor -o bin/csv-service-amd64 -a -installsuffix cgo -ldflags $(LDFLAGS)
	GOARCH=arm64 go build -mod=vendor -o bin/csv-service-arm64 -a -installsuffix cgo -ldflags $(LDFLAGS)