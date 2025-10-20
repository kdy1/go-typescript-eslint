# go-typescript-eslint

[![CI](https://github.com/kdy1/go-typescript-eslint/actions/workflows/ci.yml/badge.svg)](https://github.com/kdy1/go-typescript-eslint/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kdy1/go-typescript-eslint)](https://goreportcard.com/report/github.com/kdy1/go-typescript-eslint)
[![License](https://img.shields.io/github/license/kdy1/go-typescript-eslint)](LICENSE)

A TypeScript ESLint implementation in Go.

## Features

- TypeScript parser implementation
- ESLint compatibility
- High performance
- Cross-platform support

## Installation

```bash
go install github.com/kdy1/go-typescript-eslint@latest
```

## Development

### Prerequisites

- Go 1.21 or higher
- golangci-lint (for linting)
- goimports (for import management)

### Quick Start

```bash
# Clone the repository
git clone https://github.com/kdy1/go-typescript-eslint.git
cd go-typescript-eslint

# Install development tools
make install-tools

# Run tests
make test

# Run all CI checks locally
make ci
```

### Available Make Targets

```bash
make help           # Show all available targets
make build          # Build the project
make test           # Run tests
make test-coverage  # Run tests with coverage report
make lint           # Run golangci-lint
make fmt            # Format code with gofmt
make imports        # Fix imports with goimports
make vet            # Run go vet
make coverage       # Generate and open coverage report
make clean          # Remove build artifacts
make install-tools  # Install development tools
make ci             # Run all CI checks locally
```

## CI/CD Pipeline

This project uses GitHub Actions for continuous integration and deployment. The CI pipeline includes:

### Workflow Jobs

1. **Lint** - Code quality checks using golangci-lint v2.2.0
   - Runs 60+ linters including gosec, govet, staticcheck
   - Configured via `.golangci.yml`
   - Shows only new issues on PRs
   - Uses caching for faster execution

2. **Format Check** - Code formatting verification
   - `gofmt -s` for standard formatting
   - `goimports` for import organization
   - Fails if code is not properly formatted

3. **Test Matrix** - Cross-version testing
   - Tests on Go 1.21, 1.22, and 1.23
   - Race detection enabled
   - Coverage reporting on Go 1.23
   - Parallel execution for speed

4. **Security Scan** - Security vulnerability detection
   - Gosec static security analyzer
   - SARIF report generation
   - GitHub Security tab integration

5. **CI Success** - Gateway check
   - Ensures all jobs pass
   - Required for PR merges

### Workflow Features

- **Concurrency Control**: Cancels in-progress runs for the same branch
- **Smart Caching**: Caches Go modules and build artifacts
- **Coverage Reports**: Uploaded as artifacts, available for 30 days
- **GitHub Summary**: Coverage summary in workflow summary page
- **Matrix Testing**: Ensures compatibility across Go versions

### Linters Enabled

The project uses a comprehensive set of linters including:

- **Error Handling**: errcheck, errorlint, nilerr
- **Security**: gosec (60+ security rules)
- **Performance**: prealloc, gocritic (performance checks)
- **Style**: gofmt, gofumpt, goimports, revive, stylecheck
- **Complexity**: gocyclo, gocognit, cyclop, funlen
- **Best Practices**: govet, staticcheck, unused, ineffassign
- **Code Quality**: dupl, goconst, misspell, unconvert

See `.golangci.yml` for complete configuration.

## Testing

```bash
# Run tests
go test ./...

# Run tests with race detection
go test -race ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Or use make
make test-coverage
make coverage  # Opens HTML report
```

## Code Quality

### Running Linters Locally

```bash
# Run all linters
make lint

# Or directly
golangci-lint run

# Run specific linter
golangci-lint run --disable-all --enable=gosec
```

### Formatting Code

```bash
# Format code
make fmt

# Fix imports
make imports

# Or run both as part of CI checks
make ci
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run CI checks locally (`make ci`)
5. Commit your changes (`git commit -m 'Add some amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

### Code Review Checklist

- [ ] Tests pass locally (`make test`)
- [ ] Linters pass (`make lint`)
- [ ] Code is formatted (`make fmt imports`)
- [ ] Coverage is maintained or improved
- [ ] Documentation is updated
- [ ] Commit messages are clear

## License

See [LICENSE](LICENSE) file for details.

## Resources

- [GitHub Actions for Go](https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go)
- [golangci-lint Documentation](https://golangci-lint.run/)
- [Go Testing](https://go.dev/doc/tutorial/add-a-test)
- [Effective Go](https://go.dev/doc/effective_go)
