# go-typescript-eslint

[![CI](https://github.com/kdy1/go-typescript-eslint/actions/workflows/ci.yml/badge.svg)](https://github.com/kdy1/go-typescript-eslint/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kdy1/go-typescript-eslint)](https://goreportcard.com/report/github.com/kdy1/go-typescript-eslint)
[![License](https://img.shields.io/github/license/kdy1/go-typescript-eslint)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/kdy1/go-typescript-eslint.svg)](https://pkg.go.dev/github.com/kdy1/go-typescript-eslint)

A Go port of [TypeScript ESTree](https://github.com/typescript-eslint/typescript-eslint/tree/main/packages/typescript-estree), which converts TypeScript source code into an ESTree-compatible Abstract Syntax Tree (AST).

## Overview

This project provides a pure Go implementation of the TypeScript ESTree parser, enabling Go-based tools to parse and analyze TypeScript code. The AST produced is compatible with the ESTree specification, which is widely used by JavaScript/TypeScript tooling ecosystems.

## Features

- **Full TypeScript Support**: Parse all TypeScript syntax including types, interfaces, generics, and decorators
- **ESTree Compatible**: Output conforms to ESTree specification for interoperability
- **Type-Aware Parsing**: Optional TypeScript compiler integration for type information
- **High Performance**: Native Go implementation optimized for speed
- **Cross-Platform**: Works on all platforms supported by Go
- **JSX/TSX Support**: Parse React components with TypeScript
- **Comprehensive Error Reporting**: Detailed parse errors with location information

## Installation

### As a Library

```bash
go get github.com/kdy1/go-typescript-eslint
```

### As a CLI Tool

```bash
go install github.com/kdy1/go-typescript-eslint/cmd/go-typescript-eslint@latest
```

## Usage

### Library Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"
)

func main() {
    source := `const x: number = 42;`

    options := typescriptestree.ParseOptions{
        ECMAVersion: 2023,
        SourceType:  "module",
        Loc:         true,
        Range:       true,
    }

    ast, err := typescriptestree.Parse(source, options)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Parsed AST: %+v\n", ast)
}
```

### CLI Usage

```bash
# Parse a TypeScript file
go-typescript-eslint file.ts

# Parse with location and range information
go-typescript-eslint -loc -range file.ts

# Pretty print the AST
go-typescript-eslint -format pretty file.ts

# Include tokens and comments
go-typescript-eslint -tokens -comments file.ts
```

## Project Structure

This project follows Go best practices for module layout:

```
.
├── cmd/
│   └── go-typescript-eslint/    # CLI tool
│       ├── doc.go
│       └── main.go
├── pkg/
│   └── typescriptestree/        # Public API
│       ├── doc.go
│       └── parse.go
├── internal/                    # Internal packages (not exported)
│   ├── ast/                     # AST node definitions
│   │   ├── doc.go
│   │   └── node.go
│   ├── lexer/                   # Tokenization
│   │   ├── doc.go
│   │   └── token.go
│   ├── parser/                  # Parser implementation
│   │   ├── doc.go
│   │   └── parser.go
│   └── types/                   # Type system representation
│       ├── doc.go
│       └── types.go
├── examples/                    # Usage examples
│   ├── README.md
│   └── doc.go
├── .github/
│   └── workflows/
│       └── ci.yml               # CI/CD pipeline
├── .golangci.yml                # Linter configuration
├── Makefile                     # Development tasks
├── go.mod                       # Go module definition
├── README.md                    # This file
├── CONTRIBUTING.md              # Contribution guidelines
└── LICENSE                      # License file
```

### Package Organization

- **`pkg/typescriptestree`**: Public API for parsing TypeScript code
- **`internal/lexer`**: Tokenization and lexical analysis
- **`internal/parser`**: Syntactic analysis and AST construction
- **`internal/ast`**: AST node type definitions
- **`internal/types`**: TypeScript type system representation
- **`cmd/go-typescript-eslint`**: Command-line tool
- **`examples/`**: Example code and usage patterns

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

# Build the CLI tool
make build

# Run tests
make test

# Run all CI checks locally
make ci
```

### Available Make Targets

```bash
make help           # Show all available targets
make build          # Build the CLI tool
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

## Examples

See the [examples/](examples/) directory for complete usage examples:

- Basic parsing
- Type-aware parsing
- Custom AST traversal
- Error handling
- JSX/TSX parsing

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed contribution guidelines.

Quick checklist:

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

## Architecture

### Parser Pipeline

```
TypeScript Source
       ↓
   [Lexer]           → Tokenization
       ↓
   [Parser]          → AST Construction
       ↓
   [AST]             → ESTree-compatible tree
       ↓
[Type Checker]       → Optional type information
       ↓
  Output (JSON)
```

### Internal Packages

- **lexer**: Converts source code into tokens
  - Handles all TypeScript syntax including JSX
  - Tracks position information for error reporting
  - Preserves comments and whitespace when requested

- **parser**: Constructs AST from tokens
  - Recursive descent parser with operator precedence
  - Full TypeScript grammar support
  - Error recovery and detailed diagnostics

- **ast**: Defines node types
  - ESTree-compatible node definitions
  - TypeScript-specific extensions
  - JSON serialization support

- **types**: Type system representation
  - TypeScript type definitions
  - Type checking and inference (future)
  - Compatibility checking (future)

## Compatibility

This implementation aims for compatibility with:

- [@typescript-eslint/typescript-estree](https://github.com/typescript-eslint/typescript-eslint/tree/main/packages/typescript-estree)
- [ESTree specification](https://github.com/estree/estree)
- TypeScript 5.x syntax

## License

See [LICENSE](LICENSE) file for details.

## Resources

### TypeScript ESTree

- [TypeScript ESLint](https://typescript-eslint.io/)
- [TypeScript ESTree Package](https://github.com/typescript-eslint/typescript-eslint/tree/main/packages/typescript-estree)
- [ESTree Specification](https://github.com/estree/estree)

### Go Development

- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Modules](https://go.dev/doc/modules)
- [Go Testing](https://go.dev/doc/tutorial/add-a-test)

### CI/CD

- [GitHub Actions for Go](https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go)
- [golangci-lint Documentation](https://golangci-lint.run/)
- [Go CI Best Practices](https://medium.com/@tedious/go-linting-best-practices-for-ci-cd-with-github-actions-aa6d96e0c509)

## Acknowledgments

This project is a Go port of the excellent TypeScript ESTree project by the TypeScript ESLint team.
