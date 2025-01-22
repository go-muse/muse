all: vet lint test

vet:
	@echo "Running vet..."
	go vet ./...

lint:
	@echo "Running linter..."
	golangci-lint run ./...

test:
	@echo "Running tests..."
	go test ./...

coverage:
	@echo "Generating test coverage report..."
	mkdir -p coverage
	go test -coverprofile=coverage/coverage.out ./...
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html
