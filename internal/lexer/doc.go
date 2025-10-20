// Package lexer implements the lexical analysis (tokenization) phase
// of TypeScript parsing.
//
// The lexer reads TypeScript source code character by character and
// produces a stream of tokens that represent the syntactic elements
// of the language (keywords, identifiers, operators, literals, etc.).
//
// # Features
//
//   - Unicode support for identifiers
//   - Template literal tokenization
//   - JSX/TSX syntax support
//   - TypeScript-specific tokens (type annotations, decorators, etc.)
//   - Accurate position tracking for error reporting
//   - Comment preservation (both single-line and multi-line)
//
// # Usage
//
// The lexer is used internally by the parser and should not be used
// directly by external code. The public API is in pkg/typescriptestree.
//
// # Token Types
//
// The lexer recognizes all standard ECMAScript tokens plus TypeScript
// extensions:
//   - Keywords (const, let, var, function, class, etc.)
//   - TypeScript keywords (type, interface, enum, namespace, etc.)
//   - Operators (+, -, *, /, &&, ||, etc.)
//   - TypeScript operators (as, is, satisfies, etc.)
//   - Literals (numbers, strings, booleans, null, undefined)
//   - Template literals with interpolations
//   - Regular expressions
//   - Comments
//
// This package is internal and should not be imported by external code.
package lexer
