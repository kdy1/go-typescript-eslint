.PHONY: help build test lint fmt imports vet coverage clean install-tools

# Default target
help:
	@echo "Available targets:"
	@echo "  build         - Build the project"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  lint          - Run golangci-lint"
	@echo "  fmt           - Format code with gofmt"
	@echo "  imports       - Fix imports with goimports"
	@echo "  vet           - Run go vet"
	@echo "  coverage      - Generate and open coverage report"
	@echo "  clean         - Remove build artifacts"
	@echo "  install-tools - Install development tools"
	@echo "  ci            - Run all CI checks locally"

# Build the project
build:
	@echo "Building..."
	go build -v ./...

# Run tests
test:
	@echo "Running tests..."
	go test -v -race ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -func=coverage.out

# Run golangci-lint
lint:
	@echo "Running golangci-lint..."
	golangci-lint run

# Format code with gofmt
fmt:
	@echo "Formatting code..."
	gofmt -s -w .

# Fix imports with goimports
imports:
	@echo "Fixing imports..."
	goimports -w .

# Run go vet
vet:
	@echo "Running go vet..."
	go vet ./...

# Generate and open coverage report
coverage: test-coverage
	@echo "Generating HTML coverage report..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f coverage.out coverage.html
	go clean -cache -testcache

# Install development tools
install-tools:
	@echo "Installing development tools..."
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run all CI checks locally
ci: fmt imports vet lint test-coverage
	@echo "All CI checks passed!"
