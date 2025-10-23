// Package ast defines the Abstract Syntax Tree (AST) node types used by the parser.
//
// This package provides comprehensive data structures for representing
// TypeScript and JavaScript code as a structured tree. It follows the ESTree
// specification with TypeScript extensions, providing full compatibility with
// the TypeScript ESLint parser.
//
// # Node Type System
//
// The package defines 177+ node types organized into the following categories:
//
//   - Program & Core: Root AST node and fundamental types
//   - Identifiers & Literals: Variable names, values, and constants
//   - Expressions: Nodes that produce values (36+ types)
//   - Statements: Nodes that perform actions (24+ types)
//   - Declarations: Nodes that introduce bindings (18+ types)
//   - Patterns: Destructuring patterns (4 types)
//   - JSX: React JSX/TSX support (16 types)
//   - Decorators: Stage 3 decorator support
//   - TypeScript Types: Type system nodes (74+ types)
//
// # Base Types
//
// All node types embed BaseNode and implement the Node interface:
//
//   - Node: Base interface with Type(), Pos(), End() methods
//   - BaseNode: Common fields (NodeType, Loc, Range)
//   - Expression: Interface for expression nodes
//   - Statement: Interface for statement nodes
//   - Pattern: Interface for pattern nodes (destructuring)
//   - Declaration: Interface for declaration nodes
//   - TSNode: Interface for TypeScript-specific type nodes
//
// # Node Type Enumeration
//
// The NodeType enumeration (using iota) provides constants for all 177 node types,
// from NodeTypeProgram to NodeTypeTSClassImplements. Use these constants to identify
// node types at runtime.
//
// # JSON Serialization
//
// All node types include JSON tags for ESTree-compatible serialization:
//
//   - "type": Node type name (e.g., "Identifier", "BinaryExpression")
//   - "loc": Source location (line/column information)
//   - "range": Character range [start, end)
//   - Node-specific fields as per ESTree specification
//
// # TypeScript Extensions
//
// TypeScript-specific nodes (prefixed with TS*) extend the base ESTree
// specification to support TypeScript's type system:
//
//   - Type keywords: TSAnyKeyword, TSStringKeyword, etc.
//   - Type expressions: TSArrayType, TSUnionType, TSIntersectionType, etc.
//   - Type declarations: TSInterfaceDeclaration, TSTypeAliasDeclaration, etc.
//   - Type annotations: TSTypeAnnotation for variables and parameters
//   - Type assertions: TSAsExpression, TSTypeAssertion, etc.
//   - Import/Export: TSImportEqualsDeclaration, TSExportAssignment, etc.
//
// # Usage Example
//
//	// Create an identifier node
//	ident := &ast.Identifier{
//	    BaseNode: ast.BaseNode{
//	        NodeType: "Identifier",
//	        Start:    0,
//	        EndPos:   3,
//	    },
//	    Name: "foo",
//	}
//
//	// Create a binary expression
//	expr := &ast.BinaryExpression{
//	    BaseNode: ast.BaseNode{
//	        NodeType: "BinaryExpression",
//	    },
//	    Operator: "+",
//	    Left:     leftExpr,
//	    Right:    rightExpr,
//	}
//
// # References
//
//   - ESTree specification: https://github.com/estree/estree
//   - TypeScript ESTree: https://typescript-eslint.io/packages/typescript-estree/ast-spec/
//   - AST_NODE_TYPES: https://typescript-eslint.io/packages/typescript-estree/
//
// This package is internal and should not be imported by external code.
// External code should use the public API in pkg/typescriptestree.
package ast
