// Package typescriptestree is a Go port of TypeScript ESTree, which converts
// TypeScript source code into an ESTree-compatible Abstract Syntax Tree (AST).
//
// This package is the main public API for parsing TypeScript code. It provides
// functions to parse TypeScript source code into an AST that is compatible with
// the ESTree specification, which is widely used by JavaScript/TypeScript tooling.
//
// # Example Usage
//
//	import "github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"
//
//	// Parse TypeScript code
//	source := "const x: number = 42;"
//	options := typescriptestree.ParseOptions{
//		ECMAVersion: 2023,
//		SourceType:  "module",
//	}
//	ast, err := typescriptestree.Parse(source, options)
//	if err != nil {
//		// Handle error
//	}
//	// Use ast...
//
// # Architecture
//
// The package uses several internal components:
//   - internal/lexer: Tokenizes TypeScript source code
//   - internal/parser: Parses tokens into an AST
//   - internal/ast: Defines AST node types
//   - internal/types: Type system representation
//
// # Compatibility
//
// This implementation aims for compatibility with @typescript-eslint/typescript-estree,
// the reference TypeScript ESTree implementation used by the TypeScript ESLint project.
package typescriptestree
