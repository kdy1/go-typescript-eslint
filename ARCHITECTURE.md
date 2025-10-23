# Architecture Documentation

This document provides a detailed overview of the go-typescript-eslint architecture, design decisions, and implementation details.

## Table of Contents

- [Overview](#overview)
- [Design Principles](#design-principles)
- [Parser Pipeline](#parser-pipeline)
- [Internal Packages](#internal-packages)
- [Public API](#public-api)
- [Data Flow](#data-flow)
- [Key Design Decisions](#key-design-decisions)
- [Performance Considerations](#performance-considerations)
- [Future Enhancements](#future-enhancements)

## Overview

go-typescript-eslint is a pure Go port of [@typescript-eslint/typescript-estree](https://github.com/typescript-eslint/typescript-eslint/tree/main/packages/typescript-estree), which converts TypeScript source code into an ESTree-compatible Abstract Syntax Tree (AST).

The architecture follows a multi-stage pipeline design where each stage is responsible for a specific transformation:

```
Source Code → Tokens → TypeScript AST → ESTree AST → Result
```

### Key Components

1. **Lexer** (`internal/lexer`): Tokenizes source code
2. **Parser** (`internal/parser`): Builds TypeScript AST from tokens
3. **Converter** (`internal/converter`): Transforms TypeScript AST to ESTree format
4. **Program Manager** (`internal/program`): Manages TypeScript programs and type information
5. **Public API** (`pkg/typescriptestree`): Provides ergonomic interface to consumers

## Design Principles

### 1. Separation of Concerns

Each package has a single, well-defined responsibility:

- **Lexer**: Only tokenization, no parsing logic
- **Parser**: Only AST construction, no ESTree conversion
- **Converter**: Only AST transformation, no parsing logic
- **Program**: Only TypeScript program management, no parsing

This separation makes the codebase easier to understand, test, and maintain.

### 2. Internal Implementation, Public API

All implementation details are in `internal/` packages, which are not importable by external code. The `pkg/typescriptestree` package provides the only public API, ensuring:

- Clean API surface
- Freedom to refactor internals without breaking changes
- Clear boundary between implementation and interface

### 3. Immutable AST Nodes

AST nodes are designed to be immutable after creation. This:

- Prevents accidental mutations
- Enables safe concurrent access
- Makes code easier to reason about

### 4. Builder Pattern for Configuration

Options are constructed using the builder pattern:

```go
opts := typescriptestree.NewBuilder().
    WithSourceType(typescriptestree.SourceTypeModule).
    WithLoc(true).
    MustBuild()
```

Benefits:

- Fluent, readable API
- Type-safe configuration
- Clear validation errors
- Discoverability (IDE autocomplete)

### 5. Zero Dependencies

The project has no external dependencies for parsing. This:

- Simplifies deployment
- Reduces security surface
- Improves build times
- Ensures long-term maintainability

## Parser Pipeline

### Stage 1: Lexical Analysis (Tokenization)

**Package**: `internal/lexer`

The lexer scans the source code character by character and produces a stream of tokens.

```
Source: "const x = 42;"
↓
Tokens: [
  { Type: "const", Value: "const", Start: 0, End: 5 },
  { Type: "identifier", Value: "x", Start: 6, End: 7 },
  { Type: "=", Value: "=", Start: 8, End: 9 },
  { Type: "number", Value: "42", Start: 10, End: 12 },
  { Type: ";", Value: ";", Start: 12, End: 13 }
]
```

**Key Features**:

- **Streaming**: Tokens are generated on-demand, not all at once
- **Position Tracking**: Every token knows its exact location in source
- **Comment Preservation**: Comments are tracked separately
- **JSX Support**: Handles JSX syntax with context-aware tokenization
- **Error Recovery**: Continues tokenizing even after errors

**Implementation Details**:

- Scanner maintains current position, line, and column
- Lookahead for multi-character tokens (`>>`, `===`, `?.`)
- Context switching for JSX vs. TypeScript mode
- Unicode handling for identifiers and strings

### Stage 2: Syntactic Analysis (Parsing)

**Package**: `internal/parser`

The parser consumes tokens and builds a TypeScript AST using recursive descent parsing.

```
Tokens: [const, x, =, 42, ;]
↓
TypeScript AST:
Program {
  body: [
    VariableDeclaration {
      kind: "const",
      declarations: [
        VariableDeclarator {
          id: Identifier { name: "x" },
          init: NumericLiteral { value: 42 }
        }
      ]
    }
  ]
}
```

**Key Features**:

- **Recursive Descent**: Top-down parsing with operator precedence
- **Full TypeScript Grammar**: All TypeScript 5.x syntax
- **Error Recovery**: Produces partial AST even with syntax errors
- **Context Awareness**: Handles ambiguous syntax (JSX, type assertions)
- **Source Type Support**: Script vs. module parsing modes

**Implementation Details**:

- Parser maintains token position and lookahead
- Precedence climbing for expression parsing
- Context stack for nested scopes
- Error node creation for invalid syntax

### Stage 3: AST Transformation (Conversion)

**Package**: `internal/converter`

The converter transforms the TypeScript AST into an ESTree-compatible format.

```
TypeScript AST: VariableDeclaration
↓
ESTree AST: VariableDeclaration {
  type: "VariableDeclaration",
  kind: "const",
  declarations: [...],
  loc: { start: {...}, end: {...} },
  range: [0, 13]
}
```

**Key Features**:

- **ESTree Compatibility**: Matches ESTree specification exactly
- **Node Mapping**: Bidirectional mapping between TypeScript and ESTree nodes
- **TypeScript Extensions**: Adds TS-specific nodes (TSInterfaceDeclaration, etc.)
- **Metadata Preservation**: Maintains location, range, comments, tokens
- **Optional Mappings**: Can skip node maps for performance

**Implementation Details**:

- Visitor pattern for node transformation
- Separate methods for each node type
- Position and location conversion
- Comment and token attachment

### Stage 4: Type Information (Program Services)

**Package**: `internal/program`

The program manager creates and caches TypeScript programs for type-aware parsing.

```
tsconfig.json → TypeScript Program → Type Checker
                       ↓
                 CompilerOptions
                 SourceFiles
                 Type Information
```

**Key Features**:

- **tsconfig.json Parsing**: Reads and resolves TypeScript configuration
- **Program Caching**: Caches programs to avoid repeated parsing
- **Compiler Options**: Provides access to TypeScript compiler settings
- **Project Management**: Handles multiple projects and configurations
- **Cache Lifecycle**: Configurable cache expiry

**Implementation Details**:

- LRU cache for programs
- Config file resolution (extends, include, exclude)
- Compiler option normalization
- Thread-safe program access

## Internal Packages

### `internal/lexer` - Lexical Analysis

**Files**:
- `scanner.go`: Character-level scanning
- `lexer.go`: Token production interface
- `token.go`: Token definitions and types
- `scan.go`: Main scanning logic
- `scan_methods.go`: Helper methods for scanning

**Key Types**:
- `Token`: Represents a single token
- `Scanner`: Stateful scanner for source code
- `TokenType`: Enum of all token types

**Responsibilities**:
- Convert source code to tokens
- Track position information
- Handle Unicode correctly
- Support JSX tokenization

### `internal/parser` - Syntactic Analysis

**Files**:
- `parser.go`: Main parser implementation
- `expressions.go`: Expression parsing
- `statements.go`: Statement parsing
- `declarations.go`: Declaration parsing
- `typescript.go`: TypeScript-specific syntax
- `jsx.go`: JSX/TSX parsing
- `patterns.go`: Destructuring patterns
- `functions.go`: Function declarations and expressions

**Key Types**:
- `Parser`: Stateful parser
- Various AST node types

**Responsibilities**:
- Parse tokens into TypeScript AST
- Handle operator precedence
- Manage parsing context
- Produce error diagnostics

### `internal/ast` - AST Definitions

**Files**:
- `node.go`: Base node types and interfaces
- `node_types.go`: Node type enums
- `types.go`: AST node definitions
- `typescript.go`: TypeScript-specific nodes
- `jsx.go`: JSX node definitions
- `tokens.go`: Token-related types
- `comments.go`: Comment types
- `traverse.go`: AST traversal utilities
- `visitor_keys.go`: Node visitor key mappings
- `guards.go`: Type guard functions
- `utils.go`: AST utility functions

**Key Types**:
- `Node`: Base interface for all AST nodes
- `Program`: Root node
- `Expression`, `Statement`, `Declaration`: Node categories
- `Visitor`: Interface for AST traversal

**Responsibilities**:
- Define AST node structure
- Provide node type checking
- Enable AST traversal
- Support AST analysis

### `internal/converter` - ESTree Conversion

**Files**:
- `converter.go`: Main converter implementation
- `expressions.go`: Expression conversion
- `statements.go`: Statement conversion
- `declarations.go`: Declaration conversion
- `typescript.go`: TypeScript node conversion
- `patterns.go`: Pattern conversion
- `helpers.go`: Conversion utilities

**Key Types**:
- `Converter`: Stateful converter
- `NodeMaps`: Bidirectional node mappings
- `Options`: Conversion options

**Responsibilities**:
- Convert TypeScript AST to ESTree format
- Maintain node mappings
- Add ESTree metadata
- Handle TypeScript extensions

### `internal/program` - Program Management

**Files**:
- `program.go`: Program creation and management
- `tsconfig.go`: tsconfig.json parsing
- `cache.go`: Program caching
- `tsconfig_test.go`: Config parsing tests
- `cache_test.go`: Cache tests
- `program_test.go`: Program tests

**Key Types**:
- `Program`: TypeScript program instance
- `CompilerOptions`: Compiler settings
- `ProgramCache`: Cached programs
- `ProgramOptions`: Program creation options

**Responsibilities**:
- Parse tsconfig.json files
- Create TypeScript programs
- Cache programs for reuse
- Provide type information access

### `internal/tstype` - Type System (Planned)

**Files**:
- `types.go`: Type definitions
- `doc.go`: Package documentation

**Key Types**:
- Type representations (planned)
- Type checker interface (planned)

**Responsibilities**:
- Represent TypeScript types
- Provide type checking
- Enable type-aware analysis

## Public API

### `pkg/typescriptestree` - Public API

**Files**:
- `doc.go`: Package documentation
- `parse.go`: Main parsing functions
- `options.go`: Configuration types and builders
- `services.go`: Parser services for type information
- `constants.go`: AST node and token type constants
- `cache.go`: Cache management utilities
- `parse_test.go`: Parse function tests
- `examples_test.go`: Example code
- `services_test.go`: Services tests
- `options_examples_test.go`: Options examples

**Exported Functions**:
- `Parse()`: Basic parsing without type information
- `ParseAndGenerateServices()`: Type-aware parsing
- `ClearProgramCache()`: Cache management
- `NewBuilder()`: Parse options builder
- `NewServicesBuilder()`: Services options builder

**Exported Types**:
- `Result`: Parse result with AST and optional services
- `ParseOptions`: Basic parsing configuration
- `ParseAndGenerateServicesOptions`: Type-aware parsing configuration
- `Services`: Type information and node mappings
- `SourceType`: Script vs. module enum
- `JSDocParsingMode`: JSDoc parsing configuration

**Exported Constants**:
- `AST_NODE_TYPES`: All ESTree node type constants (177+)
- `AST_TOKEN_TYPES`: All token type constants (90+)

## Data Flow

### Basic Parsing Flow

```
User Code:
  Parse(source, opts)
       ↓
  [Create Parser]
       ↓
  [Tokenize Source] (lexer.Scanner)
       ↓
  [Parse to TypeScript AST] (parser.Parser)
       ↓
  [Convert to ESTree] (converter.Converter)
       ↓
  [Apply Options] (filter comments/tokens/loc/range)
       ↓
  Return Result{AST}
```

### Type-Aware Parsing Flow

```
User Code:
  ParseAndGenerateServices(source, opts)
       ↓
  [Load/Create TypeScript Program] (program.Program)
       ↓
  [Parse tsconfig.json if needed]
       ↓
  [Check Program Cache]
       ↓
  [Create Parser]
       ↓
  [Tokenize Source] (lexer.Scanner)
       ↓
  [Parse to TypeScript AST] (parser.Parser)
       ↓
  [Convert to ESTree with Node Maps] (converter.Converter)
       ↓
  [Create Parser Services]
       ↓
  [Add Node Mappings]
       ↓
  [Apply Options]
       ↓
  Return Result{AST, Services}
```

## Key Design Decisions

### 1. Recursive Descent Parser

**Decision**: Use recursive descent parsing instead of parser generators.

**Rationale**:
- Full control over error recovery
- Easy to understand and debug
- No external tools required
- Excellent performance
- Natural mapping to TypeScript grammar

**Trade-offs**:
- More code to write manually
- Requires careful precedence handling
- Limited left-recursion support

### 2. Internal Packages

**Decision**: Place all implementation in `internal/` packages.

**Rationale**:
- Prevents accidental API exposure
- Freedom to refactor without breaking changes
- Clear public API boundary
- Go best practice

**Trade-offs**:
- Cannot extend parser from external code
- Must go through public API

### 3. Builder Pattern for Options

**Decision**: Use fluent builder pattern for configuration.

**Rationale**:
- More readable than large option structs
- Type-safe configuration
- Clear validation points
- Better IDE support

**Trade-offs**:
- More code than simple structs
- Cannot use struct literals

### 4. Bidirectional Node Mappings

**Decision**: Maintain both ESTree→TypeScript and TypeScript→ESTree mappings.

**Rationale**:
- Required for ESLint rule implementation
- Enables type-aware linting
- Matches typescript-estree behavior

**Trade-offs**:
- Extra memory overhead
- Maintenance burden
- Can be disabled for performance

### 5. Program Caching

**Decision**: Cache TypeScript programs globally by default.

**Rationale**:
- Huge performance improvement for repeated parsing
- Matches typescript-estree behavior
- Configurable cache lifetime

**Trade-offs**:
- Memory usage for cached programs
- Need to clear cache in tests
- Thread-safety requirements

### 6. Two-Function API

**Decision**: Separate `Parse()` and `ParseAndGenerateServices()` functions.

**Rationale**:
- Clear distinction between fast parsing and type-aware parsing
- Matches typescript-estree API
- Users can choose appropriate trade-off

**Trade-offs**:
- API duplication
- More code to maintain

## Performance Considerations

### Memory Efficiency

1. **Streaming Tokens**: Lexer generates tokens on-demand, not all at once
2. **Minimal Allocations**: Reuse buffers where possible
3. **Optional Features**: Comments, tokens, loc, range can be disabled
4. **Node Map Control**: Can skip node mappings if not needed

### Time Complexity

1. **Linear Parsing**: O(n) where n is source code size
2. **Program Caching**: O(1) cache lookup after initial parse
3. **No Backtracking**: Recursive descent without excessive lookahead

### Optimization Strategies

1. **String Interning**: Common identifiers could be interned
2. **Arena Allocation**: Could use arena for AST nodes
3. **Parallel Parsing**: Multiple files could be parsed concurrently
4. **Incremental Parsing**: Could support incremental updates (planned)

## Future Enhancements

### Short-term

1. **Project Service**: TypeScript project service integration
2. **Enhanced Error Recovery**: Better partial AST for invalid code
3. **Performance Profiling**: Identify and optimize hot paths
4. **Benchmark Suite**: Comprehensive performance benchmarks

### Medium-term

1. **Type Checker Integration**: Full TypeScript type checking
2. **Incremental Parsing**: Only reparse changed regions
3. **Parallel Parsing**: Parse multiple files concurrently
4. **Source Maps**: Support for source map generation

### Long-term

1. **Language Server**: LSP implementation for editors
2. **Transformer API**: Programmatic AST transformation
3. **Custom Type System**: Go-native type representation
4. **WASM Support**: Compile to WebAssembly

## Testing Strategy

### Unit Tests

Each internal package has comprehensive unit tests:

- Lexer: Token generation for all syntax
- Parser: AST construction for all nodes
- Converter: ESTree conversion correctness
- Program: Config parsing and caching

### Integration Tests

Public API tests in `pkg/typescriptestree`:

- End-to-end parsing scenarios
- Options configuration
- Error handling
- Example code

### Test Organization

- Table-driven tests for comprehensive coverage
- Example tests for documentation
- Benchmark tests for performance
- Golden file tests for AST output (planned)

## Error Handling

### Error Recovery

The parser attempts to recover from errors and produce a partial AST:

1. **Skip to Sync Points**: Skip tokens until statement/declaration boundary
2. **Insert Missing Tokens**: Synthesize missing `;`, `)`, `}`, etc.
3. **Error Nodes**: Create error nodes for invalid syntax
4. **Continue Parsing**: Keep parsing after errors

### Error Reporting

Errors include:

- Error message
- Source location (line, column)
- Surrounding context
- Suggestion for fix (when possible)

## Debugging

### Debug Options

- `DebugLevel`: Enable detailed logging for specific modules
- `ErrorOnUnknownASTType`: Catch unsupported syntax early

### Logging

Structured logging via `LoggerFn` option:

```go
opts := NewBuilder().
    WithLoggerFn(func(msg string) {
        log.Println("[parser]", msg)
    }).
    Build()
```

## Compatibility

### TypeScript Compatibility

- Supports TypeScript 5.x syntax
- New syntax added as TypeScript evolves
- Deprecation warnings for old syntax

### ESTree Compatibility

- Matches ESTree specification
- TypeScript extensions clearly marked
- Compatible with ESLint and other tools

### Go Compatibility

- Requires Go 1.21+
- Uses standard library only
- No CGo dependencies

## Contributing to Architecture

When proposing architectural changes:

1. Open an issue describing the problem
2. Discuss design alternatives
3. Consider backward compatibility
4. Update this document
5. Implement with tests
6. Update examples

See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.

---

For questions or clarifications about the architecture, please open a [GitHub Discussion](https://github.com/kdy1/go-typescript-eslint/discussions).
