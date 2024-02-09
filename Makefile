.PHONY: deps build

deps:
	@go install github.com/cosmtrek/air@latest
	@go get github.com/go-chi/chi/v5
	@go get github.com/mattn/go-sqlite3
	@go get github.com/jmoiron/sqlx
	@go get github.com/matoous/go-nanoid/v2

dev:
	@go mod tidy
	@air
	
build:
	@go build ldflags="-s -w" -o app ./cmd/main.go
