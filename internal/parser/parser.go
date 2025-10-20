package parser

import (
	"fmt"

	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/lexer"
)

// Parser represents a TypeScript parser.
type Parser struct {
	tokens  []lexer.Token
	errors  []ParseError
	current int
}

// ParseError represents a parsing error.
type ParseError struct {
	Message string
	Line    int
	Column  int
	Pos     int
}

// New creates a new parser for the given source code.
func New(_ string) *Parser {
	// TODO: Initialize lexer and tokenize
	return &Parser{
		tokens:  []lexer.Token{},
		current: 0,
		errors:  []ParseError{},
	}
}

// Parse parses the source code and returns the AST.
//
//nolint:ireturn // This returns an interface by design as it's the base node type for the AST
func (p *Parser) Parse() (ast.Node, error) {
	// TODO: Implement parsing logic
	return nil, ErrNotImplemented
}

// ErrNotImplemented is returned when parsing is not yet fully implemented.
var ErrNotImplemented = fmt.Errorf("parsing not yet implemented")

// Errors returns the list of parsing errors.
func (p *Parser) Errors() []ParseError {
	return p.errors
}
