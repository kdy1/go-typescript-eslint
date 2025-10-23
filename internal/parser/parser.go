package parser

import (
	"fmt"

	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/lexer"
)

// Parser represents a TypeScript parser that implements recursive descent parsing.
// It consumes tokens from a scanner and produces an AST that conforms to the
// ESTree specification with TypeScript extensions.
type Parser struct {
	scanner *lexer.Scanner
	current lexer.Token
	peek    lexer.Token

	// Parser state
	errors      []ParseError
	inFunction  bool
	inLoop      bool
	inSwitch    bool
	inAsync     bool
	inGenerator bool
	inClass     bool
	allowYield  bool
	allowAwait  bool
	strictMode  bool

	// Module state
	sourceType string // "script" or "module"

	// JSX state
	jsxEnabled bool

	// All tokens and comments for the AST
	allTokens   []lexer.Token
	allComments []ast.Comment
}

// ParseError represents a parsing error.
type ParseError struct {
	Message string
	Line    int
	Column  int
	Pos     int
}

// Error implements the error interface for ParseError.
func (e ParseError) Error() string {
	return fmt.Sprintf("%s at line %d, column %d", e.Message, e.Line, e.Column)
}

// New creates a new parser for the given source code.
func New(source string) *Parser {
	scanner := lexer.NewScanner(source)
	scanner.SetSkipComments(false) // We want to capture comments

	p := &Parser{
		scanner:    scanner,
		errors:     []ParseError{},
		sourceType: "module", // Default to module
		jsxEnabled: false,
	}

	// Prime the parser with two tokens
	p.nextToken()
	p.nextToken()

	return p
}

// SetSourceType sets the source type ("script" or "module").
func (p *Parser) SetSourceType(sourceType string) {
	p.sourceType = sourceType
}

// SetJSXEnabled enables or disables JSX parsing.
func (p *Parser) SetJSXEnabled(enabled bool) {
	p.jsxEnabled = enabled
	p.scanner.SetJSXMode(enabled)
}

// SetStrictMode enables or disables strict mode parsing.
func (p *Parser) SetStrictMode(strict bool) {
	p.strictMode = strict
}

// nextToken advances to the next token in the stream.
func (p *Parser) nextToken() {
	p.current = p.peek
	p.peek = p.scanner.Scan()

	// Store all tokens for the AST
	if p.current.Type != lexer.COMMENT {
		p.allTokens = append(p.allTokens, p.current)
	} else {
		// Store comments separately
		p.allComments = append(p.allComments, ast.Comment{
			Type:  "Line", // Will be refined based on comment content
			Value: p.current.Literal,
			Range: &ast.Range{p.current.Pos, p.current.End},
		})
	}
}

// expect checks if the current token is of the expected type and advances.
// Returns an error if the token doesn't match.
func (p *Parser) expect(typ lexer.TokenType) error {
	if p.current.Type != typ {
		return p.errorAtCurrent(fmt.Sprintf("expected %v, got %v", typ, p.current.Type))
	}
	p.nextToken()
	return nil
}

// match checks if the current token matches any of the given types.
func (p *Parser) match(types ...lexer.TokenType) bool {
	for _, typ := range types {
		if p.current.Type == typ {
			return true
		}
	}
	return false
}

// consume advances if the current token matches the expected type.
func (p *Parser) consume(typ lexer.TokenType) bool {
	if p.current.Type == typ {
		p.nextToken()
		return true
	}
	return false
}

// errorAtCurrent creates an error at the current token position.
func (p *Parser) errorAtCurrent(message string) error {
	err := ParseError{
		Message: message,
		Line:    p.current.Line,
		Column:  p.current.Column,
		Pos:     p.current.Pos,
	}
	p.errors = append(p.errors, err)
	return err
}

// errorAtToken creates an error at a specific token position.
func (p *Parser) errorAtToken(token lexer.Token, message string) error {
	err := ParseError{
		Message: message,
		Line:    token.Line,
		Column:  token.Column,
		Pos:     token.Pos,
	}
	p.errors = append(p.errors, err)
	return err
}

// isAtEnd checks if we've reached the end of the token stream.
func (p *Parser) isAtEnd() bool {
	return p.current.Type == lexer.EOF
}

// Parse parses the source code and returns a Program node (the root of the AST).
//
//nolint:ireturn // This returns an interface by design as it's the base node type for the AST
func (p *Parser) Parse() (ast.Node, error) {
	program := &ast.Program{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeProgram.String(),
		},
		SourceType: p.sourceType,
		Body:       []ast.Statement{},
	}

	// Parse all top-level statements
	for !p.isAtEnd() {
		stmt, err := p.parseStatementListItem()
		if err != nil {
			// Try to recover by synchronizing to the next statement
			p.synchronize()
			continue
		}
		program.Body = append(program.Body, stmt)
	}

	// Attach all comments and tokens
	program.Comments = p.allComments

	// Convert lexer tokens to AST tokens
	for _, tok := range p.allTokens {
		program.Tokens = append(program.Tokens, ast.Token{
			Type:  tok.Type.String(),
			Value: tok.Literal,
			Range: &ast.Range{tok.Pos, tok.End},
		})
	}

	if len(p.errors) > 0 {
		return program, p.errors[0]
	}

	return program, nil
}

// synchronize attempts to recover from a parse error by advancing to the next statement.
func (p *Parser) synchronize() {
	p.nextToken()

	for !p.isAtEnd() {
		// If we just passed a semicolon, we're likely at a statement boundary
		if p.current.Type == lexer.SEMICOLON {
			p.nextToken()
			return
		}

		// Check if we're at a statement boundary
		switch p.current.Type {
		case lexer.CLASS, lexer.FUNCTION, lexer.VAR, lexer.LET, lexer.CONST,
			lexer.FOR, lexer.IF, lexer.WHILE, lexer.DO, lexer.SWITCH,
			lexer.RETURN, lexer.TRY, lexer.THROW, lexer.BREAK, lexer.CONTINUE,
			lexer.IMPORT, lexer.EXPORT, lexer.INTERFACE, lexer.TYPE, lexer.ENUM:
			return
		}

		p.nextToken()
	}
}

// Errors returns the list of parsing errors.
func (p *Parser) Errors() []ParseError {
	return p.errors
}

// ErrNotImplemented is returned when parsing is not yet fully implemented.
var ErrNotImplemented = fmt.Errorf("parsing not yet implemented")
