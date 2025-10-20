// Package ast defines the Abstract Syntax Tree node types used by the
// TypeScript ESTree parser.
//
// This package contains the internal representation of AST nodes that
// conform to the ESTree specification with TypeScript extensions.
//
// The node types are designed to be:
//   - Compatible with ESTree specification
//   - Extended with TypeScript-specific nodes
//   - Efficient for parsing and traversal
//   - Serializable to JSON for interoperability
//
// Node hierarchy follows the ESTree specification:
//   - All nodes implement the Node interface
//   - Expression nodes implement the Expression interface
//   - Statement nodes implement the Statement interface
//   - Pattern nodes implement the Pattern interface
//
// This package is internal and should not be imported by external code.
// External code should use the public API in pkg/typescriptestree.
package ast
