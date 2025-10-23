# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Comprehensive documentation including README, ARCHITECTURE, MIGRATION, PERFORMANCE guides
- Example code demonstrating various use cases
- Complete API documentation with godoc comments

## [0.1.0] - 2024-01-XX

### Added
- Initial public release
- Complete TypeScript 5.x syntax support
- `Parse()` function for basic TypeScript parsing
- `ParseAndGenerateServices()` function for type-aware parsing
- Full ESTree-compatible AST output
- Builder pattern for configuration with `ParseOptionsBuilder` and `ParseAndGenerateServicesOptionsBuilder`
- AST node type constants (`AST_NODE_TYPES`) with 177+ node types
- Token type constants (`AST_TOKEN_TYPES`) with 90+ token types
- JSX/TSX parsing support with automatic detection
- Comment and token collection
- Location and range information tracking
- TypeScript program management with caching
- tsconfig.json parsing and resolution
- Parser services with bidirectional node mappings
- Program cache management (`ClearProgramCache()`)
- Comprehensive test suite with unit, integration, and example tests
- Benchmark tests for performance measurement
- CLI tool for command-line parsing
- Complete internal packages:
  - `internal/lexer`: Scanner and tokenizer
  - `internal/parser`: Recursive descent parser
  - `internal/ast`: AST node definitions and utilities
  - `internal/converter`: ESTree conversion
  - `internal/program`: Program management and caching
  - `internal/tstype`: Type system representation (foundation)

### Features
- Zero external dependencies for parsing
- Pure Go implementation (no CGo)
- Linear time complexity (O(n))
- Memory-efficient with minimal allocations
- Thread-safe program caching
- Error recovery and partial AST generation
- Support for all TypeScript 5.x features:
  - Interfaces, type aliases, and generics
  - Enums and namespaces
  - Decorators and metadata
  - Type assertions and satisfies expressions
  - All TypeScript keywords and operators
  - JSX/TSX elements and expressions

### Compatibility
- Compatible with @typescript-eslint/typescript-estree version 8.x API
- ESTree specification compliant
- TypeScript 5.x syntax support
- Go 1.21+ required

## Release Process

### Versioning Strategy

This project follows [Semantic Versioning](https://semver.org/):

- **MAJOR** version for incompatible API changes
- **MINOR** version for new functionality in a backwards-compatible manner
- **PATCH** version for backwards-compatible bug fixes

### Pre-1.0 Development

During 0.x releases:
- **0.x.0** (minor): New features, may include breaking changes
- **0.x.y** (patch): Bug fixes and non-breaking improvements

### Release Checklist

Before releasing a new version:

- [ ] Update CHANGELOG.md with changes
- [ ] Update version in documentation
- [ ] Ensure all tests pass (`make test`)
- [ ] Ensure linters pass (`make lint`)
- [ ] Update benchmarks if performance changed
- [ ] Tag release: `git tag -a v0.x.y -m "Release v0.x.y"`
- [ ] Push tag: `git push origin v0.x.y`
- [ ] Create GitHub release with changelog
- [ ] Verify go.mod versioning

## Categories

Changes are grouped by category:

### Added
New features and capabilities.

### Changed
Changes to existing functionality.

### Deprecated
Features that will be removed in future versions.

### Removed
Features that have been removed.

### Fixed
Bug fixes.

### Security
Security-related changes and fixes.

### Performance
Performance improvements and optimizations.

## Migration Guides

For major version changes, see [MIGRATION.md](MIGRATION.md) for detailed upgrade instructions.

## Support Policy

### Supported Versions

| Version | Supported | Go Version | TypeScript Version |
|---------|-----------|------------|-------------------|
| 0.1.x   | ✅ Yes     | 1.21+      | 5.x               |

### Support Timeline

- **Current version**: Full support with bug fixes and features
- **Previous minor**: Bug fixes only for 6 months after next minor release
- **Older versions**: Community support only

## Reporting Issues

If you encounter issues with a specific version:

1. Check if the issue is already fixed in the latest version
2. Search [existing issues](https://github.com/kdy1/go-typescript-eslint/issues)
3. If not found, open a new issue with:
   - Version number
   - Go version (`go version`)
   - Minimal reproduction case
   - Expected vs actual behavior

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for how to contribute changes.

## Links

- [GitHub Repository](https://github.com/kdy1/go-typescript-eslint)
- [Issue Tracker](https://github.com/kdy1/go-typescript-eslint/issues)
- [Discussions](https://github.com/kdy1/go-typescript-eslint/discussions)
- [Documentation](https://pkg.go.dev/github.com/kdy1/go-typescript-eslint)

---

[Unreleased]: https://github.com/kdy1/go-typescript-eslint/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/kdy1/go-typescript-eslint/releases/tag/v0.1.0
