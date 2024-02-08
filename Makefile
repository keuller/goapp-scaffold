.PHONY: deps build

deps:
	@go install github.com/cosmtrek/air@latest
	@go get github.com/go-chi/chi/v5
	
dev:
	@go mod tidy
	@air
	
build:
	@go build ldflags="-s -w" -o app ./cmd/main.go
