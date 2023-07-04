.PHONY: build test lint clean

clean:
	@go mod tidy
	@buf generate api/grpc

build: clean
	@go build -o ports-service ./cmd/service
	@go build -o ports-ingestor ./cmd/ingestor

test: clean
	@go test ./... -count=1

lint: clean
	@golangci-lint run
	@buf lint api/grpc
