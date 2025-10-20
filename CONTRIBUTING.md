# Contributing to go-typescript-eslint

Thank you for your interest in contributing to go-typescript-eslint! This document provides guidelines and instructions for contributing.

## Code of Conduct

Please be respectful and constructive in all interactions with the community.

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/go-typescript-eslint.git
   cd go-typescript-eslint
   ```
3. Add the upstream repository:
   ```bash
   git remote add upstream https://github.com/kdy1/go-typescript-eslint.git
   ```

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Git
- Make (optional but recommended)

### Install Development Tools

```bash
make install-tools
```

This will install:
- golangci-lint - for linting
- goimports - for import management

## Making Changes

### 1. Create a Branch

```bash
git checkout -b feature/your-feature-name
```

Use descriptive branch names:
- `feature/` for new features
- `fix/` for bug fixes
- `docs/` for documentation changes
- `refactor/` for code refactoring
- `test/` for test improvements

### 2. Make Your Changes

- Write clean, idiomatic Go code
- Follow the existing code style
- Add tests for new functionality
- Update documentation as needed

### 3. Test Your Changes

Run the full CI test suite locally:

```bash
# Run all checks
make ci

# Or run individual checks
make test           # Run tests
make lint           # Run linters
make fmt            # Format code
make imports        # Fix imports
make test-coverage  # Check coverage
```

### 4. Commit Your Changes

Write clear, descriptive commit messages:

```bash
git add .
git commit -m "feat: add new parser feature"
```

Commit message format:
- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation changes
- `test:` for test changes
- `refactor:` for code refactoring
- `perf:` for performance improvements
- `chore:` for maintenance tasks

### 5. Push and Create Pull Request

```bash
git push origin feature/your-feature-name
```

Then create a Pull Request on GitHub with:
- Clear title describing the change
- Detailed description of what and why
- Reference any related issues

## Code Review Process

1. All submissions require review before merging
2. CI checks must pass (lint, test, format, security)
3. Maintainers may request changes
4. Once approved, a maintainer will merge your PR

## Testing Guidelines

### Writing Tests

- Place tests in `*_test.go` files
- Use table-driven tests where appropriate
- Aim for high code coverage (target: 80%+)
- Test edge cases and error conditions

Example test structure:

```go
func TestFeature(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    interface{}
        wantErr bool
    }{
        {
            name:    "valid input",
            input:   "test",
            want:    expectedResult,
            wantErr: false,
        },
        // More test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Feature(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with race detection
go test -race ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific package tests
go test ./parser/...

# Run specific test
go test -run TestFeature ./...
```

## Linting and Code Quality

### Running Linters

```bash
# Run all configured linters
make lint

# Or use golangci-lint directly
golangci-lint run

# Run specific linter
golangci-lint run --disable-all --enable=gosec
```

### Code Formatting

```bash
# Format code
gofmt -s -w .

# Or use make
make fmt

# Fix imports
goimports -w .

# Or use make
make imports
```

### Handling Linter Issues

If you need to disable a linter for a specific line:

```go
//nolint:lintername // reason for disabling
problematicCode()
```

Only do this when absolutely necessary and always provide a reason.

## Documentation

- Update README.md for user-facing changes
- Add godoc comments to exported functions, types, and packages
- Update CONTRIBUTING.md if changing development workflow
- Add inline comments for complex logic

### Writing Good Godoc

```go
// Package parser provides TypeScript parsing functionality.
package parser

// Parser represents a TypeScript ESLint parser that can parse
// TypeScript source code into an abstract syntax tree.
type Parser struct {
    Options map[string]interface{}
}

// New creates and initializes a new Parser with the given options.
// If options is nil, default options will be used.
func New(options map[string]interface{}) *Parser {
    return &Parser{
        Options: options,
    }
}
```

## Pull Request Checklist

Before submitting your PR, ensure:

- [ ] Code follows Go best practices and project style
- [ ] All tests pass (`make test`)
- [ ] Linters pass (`make lint`)
- [ ] Code is formatted (`make fmt imports`)
- [ ] New code has tests with good coverage
- [ ] Documentation is updated
- [ ] Commit messages are clear and descriptive
- [ ] PR description explains the changes
- [ ] No merge conflicts with main branch

## CI/CD Pipeline

All pull requests automatically run through our CI pipeline:

1. **Lint**: Runs 60+ linters via golangci-lint
2. **Format Check**: Verifies code formatting
3. **Test Matrix**: Tests on Go 1.21, 1.22, and 1.23
4. **Security Scan**: Runs gosec security analysis
5. **Coverage**: Generates coverage reports

PRs must pass all checks before merging.

## Getting Help

- Check existing issues and PRs
- Read the documentation in README.md
- Ask questions in PR comments
- Open an issue for bugs or feature requests

## License

By contributing, you agree that your contributions will be licensed under the same license as the project (see LICENSE file).

## Recognition

Contributors are recognized in:
- Git commit history
- GitHub contributors page
- Release notes (for significant contributions)

Thank you for contributing to go-typescript-eslint!
