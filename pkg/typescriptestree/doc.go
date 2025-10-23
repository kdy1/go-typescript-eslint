// Package typescriptestree is a Go port of TypeScript ESTree, which converts
// TypeScript source code into an ESTree-compatible Abstract Syntax Tree (AST).
//
// This package is the main public API for parsing TypeScript code. It provides
// functions to parse TypeScript source code into an AST that is compatible with
// the ESTree specification, which is widely used by JavaScript/TypeScript tooling.
//
// # Main Functions
//
// The package exports two main parsing functions:
//
//   - Parse: Converts TypeScript source code into an ESTree-compatible AST
//   - ParseAndGenerateServices: Parses code and generates TypeScript program services
//     for type-aware linting and analysis
//
// # Constants
//
// The package also exports constants for AST node and token types:
//
//   - AST_NODE_TYPES: String values for every AST node's type property
//   - AST_TOKEN_TYPES: String values for every AST token's type property
//
// # Example Usage
//
// Basic parsing:
//
//	import "github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"
//
//	source := "const x: number = 42;"
//	opts := typescriptestree.NewBuilder().
//		WithSourceType(typescriptestree.SourceTypeModule).
//		WithLoc(true).
//		WithRange(true).
//		MustBuild()
//
//	result, err := typescriptestree.Parse(source, opts)
//	if err != nil {
//		// Handle error
//	}
//	// Use result.AST...
//
// Type-aware parsing:
//
//	opts := typescriptestree.NewServicesBuilder().
//		WithProject("./tsconfig.json").
//		WithTSConfigRootDir(".").
//		Build()
//
//	result, err := typescriptestree.ParseAndGenerateServices(source, opts)
//	if err != nil {
//		// Handle error
//	}
//	// Use result.AST and result.Services...
//
// Using node type constants:
//
//	if node.Type() == typescriptestree.AST_NODE_TYPES.Identifier {
//		// Handle identifier node
//	}
//
// # Architecture
//
// The package integrates several internal components:
//
//   - internal/lexer: Tokenizes TypeScript source code
//   - internal/parser: Parses tokens into an AST using recursive descent parsing
//   - internal/converter: Transforms TypeScript AST into ESTree format
//   - internal/ast: Defines AST node types and utilities
//   - internal/program: Manages TypeScript programs and tsconfig.json parsing
//
// # Compatibility
//
// This implementation aims for full compatibility with @typescript-eslint/typescript-estree,
// the reference TypeScript ESTree implementation used by the TypeScript ESLint project.
// It matches the API surface and behavior of typescript-estree version 8.x.
//
// For more information, see:
//   - https://typescript-eslint.io/packages/typescript-estree/
//   - https://github.com/typescript-eslint/typescript-eslint
package typescriptestree
