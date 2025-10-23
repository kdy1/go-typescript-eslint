// Package converter transforms TypeScript AST nodes into ESTree-compatible TSESTree format.
//
// This package implements the core AST conversion logic that bridges between
// TypeScript's internal AST representation and the ESTree specification with
// TypeScript extensions (TSESTree format used by typescript-eslint).
//
// # Overview
//
// The converter performs several key transformations:
//
//  1. Node Conversion: Transforms TypeScript AST nodes to ESTree equivalents
//  2. Token Conversion: Converts TypeScript tokens to ESTree token format
//  3. Comment Attachment: Attaches comments to appropriate AST nodes
//  4. Node Mapping: Creates bidirectional mappings for ParserServices
//
// # Architecture
//
// The Converter struct uses a visitor pattern to recursively traverse and
// transform the TypeScript AST:
//
//	converter := NewConverter(sourceCode, options)
//	estreeAST := converter.ConvertProgram(typescriptAST)
//	nodeMaps := converter.GetNodeMaps()
//
// # Node Mappings
//
// The converter maintains bidirectional maps between TypeScript and ESTree nodes,
// which are essential for type-aware linting rules that need to correlate ESTree
// nodes back to TypeScript's type checker:
//
//   - ESTreeNodeToTSNodeMap: Maps ESTree nodes to original TypeScript nodes
//   - TSNodeToESTreeNodeMap: Maps TypeScript nodes to converted ESTree nodes
//
// # Reference Implementation
//
// This package follows the design of @typescript-eslint/typescript-estree's
// ast-converter module:
// https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/typescript-estree/src/ast-converter.ts
//
// # ESTree Compatibility
//
// The output conforms to:
//   - ESTree specification: https://github.com/estree/estree
//   - TypeScript-ESTree extensions: https://typescript-eslint.io/packages/typescript-estree/ast-spec/
package converter
