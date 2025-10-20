// Package parser implements the syntactic analysis phase of TypeScript parsing.
//
// The parser takes a stream of tokens from the lexer and constructs an
// Abstract Syntax Tree (AST) that represents the structure of the TypeScript
// program according to the language grammar.
//
// # Features
//
//   - Recursive descent parsing with operator precedence
//   - Full TypeScript syntax support
//   - Comprehensive error recovery
//   - Detailed error messages with position information
//   - Support for all ECMAScript features
//   - TypeScript-specific syntax (types, interfaces, generics, etc.)
//   - JSX/TSX support
//
// # Parser Architecture
//
// The parser uses a recursive descent approach with the following components:
//   - Token stream management
//   - Lookahead for parsing decisions
//   - Error recovery and reporting
//   - AST node construction
//   - Context tracking (strict mode, async, await, etc.)
//
// # Parsing Strategy
//
// The parser follows these principles:
//   - Top-down parsing from program to expressions
//   - Operator precedence parsing for expressions
//   - Error recovery at statement boundaries
//   - Permissive parsing with detailed diagnostics
//
// # Grammar
//
// The parser implements the TypeScript grammar, which extends ECMAScript:
//   - Statements (var, let, const, if, for, while, etc.)
//   - Expressions (literals, operators, function calls, etc.)
//   - Declarations (functions, classes, interfaces, types)
//   - Type annotations and type expressions
//   - Decorators and metadata
//
// This package is internal and should not be imported by external code.
// External code should use the public API in pkg/typescriptestree.
package parser
