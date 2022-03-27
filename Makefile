GIT_COMMIT := $(shell git rev-parse --short HEAD)

images:
	docker build csv-service -t "ghcr.io/shikachuu/csv-service:$(GIT_COMMIT)"
	docker tag "ghcr.io/shikachuu/csv-service:$(GIT_COMMIT)" "ghcr.io/shikachuu/csv-service:latest"