package typescriptestree

import (
	"errors"
)

// ErrNotImplemented is returned when a feature is not yet implemented.
var ErrNotImplemented = errors.New("feature not yet implemented")

// ParseOptions configures the parser behavior.
type ParseOptions struct {
	// SourceType specifies the source type: "script" or "module".
	SourceType string

	// FilePath is the path to the file being parsed (for error messages).
	FilePath string

	// ECMAVersion specifies the ECMAScript version to parse.
	// Defaults to the latest supported version.
	ECMAVersion int

	// Loc indicates whether to include location information in the AST.
	Loc bool

	// Range indicates whether to include range information in the AST.
	Range bool

	// Comment indicates whether to include comments in the AST.
	Comment bool

	// Tokens indicates whether to include tokens in the AST.
	Tokens bool
}

// AST represents the Abstract Syntax Tree produced by parsing.
// This is a placeholder and will be expanded with proper node types.
type AST struct {
	Loc      *Location `json:"loc,omitempty"`
	Range    *[2]int   `json:"range,omitempty"`
	Body     []ASTNode `json:"body"`
	Comments []Comment `json:"comments,omitempty"`
	Tokens   []Token   `json:"tokens,omitempty"`
	Type     string    `json:"type"`
}

// ASTNode represents a node in the Abstract Syntax Tree.
// This is a placeholder interface that will be expanded.
type ASTNode interface {
	Node()
}

// Comment represents a comment in the source code.
type Comment struct {
	Value string    `json:"value"`
	Loc   *Location `json:"loc,omitempty"`
	Range *[2]int   `json:"range,omitempty"`
	Type  string    `json:"type"`
}

// Token represents a token in the source code.
type Token struct {
	Value string    `json:"value"`
	Loc   *Location `json:"loc,omitempty"`
	Range *[2]int   `json:"range,omitempty"`
	Type  string    `json:"type"`
}

// Location represents the location of a node in the source code.
type Location struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

// Position represents a position in the source code.
type Position struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

// Parse parses TypeScript source code into an AST.
// This is the main entry point for parsing TypeScript code.
func Parse(_ string, _ ParseOptions) (*AST, error) {
	// TODO: Implement full TypeScript parsing
	// This will use the internal lexer and parser packages
	return nil, ErrNotImplemented
}

// ParseAndGenerateServices parses TypeScript source code and generates
// TypeScript program services for type-aware linting.
func ParseAndGenerateServices(_ string, _ ParseOptions) (*AST, error) {
	// TODO: Implement with TypeScript services support
	return nil, ErrNotImplemented
}
