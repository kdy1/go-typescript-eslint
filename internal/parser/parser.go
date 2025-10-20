package parser

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/lexer"
)

// Parser represents a TypeScript parser.
type Parser struct {
	lexer   *lexer.Lexer
	tokens  []lexer.Token
	current int
	errors  []ParseError
}

// ParseError represents a parsing error.
type ParseError struct {
	Message string
	Pos     int
	Line    int
	Column  int
}

// New creates a new parser for the given source code.
func New(source string) *Parser {
	// TODO: Initialize lexer and tokenize
	return &Parser{
		tokens:  []lexer.Token{},
		current: 0,
		errors:  []ParseError{},
	}
}

// Parse parses the source code and returns the AST.
func (p *Parser) Parse() (ast.Node, error) {
	// TODO: Implement parsing logic
	return nil, nil
}

// Errors returns the list of parsing errors.
func (p *Parser) Errors() []ParseError {
	return p.errors
}

// peek returns the current token without consuming it.
func (p *Parser) peek() lexer.Token {
	if p.current >= len(p.tokens) {
		return lexer.Token{Type: lexer.EOF}
	}
	return p.tokens[p.current]
}

// next consumes and returns the current token.
func (p *Parser) next() lexer.Token {
	token := p.peek()
	if p.current < len(p.tokens) {
		p.current++
	}
	return token
}

// expect checks if the current token matches the expected type.
func (p *Parser) expect(tokenType lexer.TokenType) bool {
	return p.peek().Type == tokenType
}
