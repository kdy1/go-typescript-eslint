// Package types provides TypeScript type system representation and
// type checking functionality.
//
// This package implements the type system used by TypeScript, including:
//   - Type definitions (primitive types, object types, union types, etc.)
//   - Type inference
//   - Type checking rules
//   - Type compatibility checking
//   - Generic type instantiation
//
// # Type Hierarchy
//
// The TypeScript type system includes:
//   - Primitive types: string, number, boolean, null, undefined, symbol, bigint
//   - Object types: interfaces, classes, type literals
//   - Union and intersection types
//   - Tuple types
//   - Array types
//   - Function types
//   - Generic types
//   - Literal types
//   - Conditional types
//   - Mapped types
//   - Template literal types
//
// # Type Checking
//
// The package provides functionality for:
//   - Type assignment compatibility
//   - Structural type checking
//   - Generic type constraint checking
//   - Type narrowing
//   - Control flow analysis
//
// This package is internal and should not be imported by external code.
// Type information is exposed through the public API in pkg/typescriptestree
// when type-aware parsing is enabled.
package tstype
