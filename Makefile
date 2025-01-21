all: lint test

lint:
	@echo "Running linter..."
	golangci-lint run ./...

test:
	@echo "Running tests..."
	go test ./...

