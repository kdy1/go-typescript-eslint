package lexer

// Lexer represents a lexical analyzer for TypeScript source code.
type Lexer struct {
	tokens []Token
	source string
	pos    int
	line   int
	column int
}

// New creates a new Lexer for the given source code.
func New(source string) *Lexer {
	return &Lexer{
		source: source,
		pos:    0,
		line:   1,
		column: 0,
		tokens: []Token{},
	}
}

// NextToken returns the next token from the source code.
func (l *Lexer) NextToken() Token {
	// TODO: Implement tokenization logic
	return Token{
		Type:    EOF,
		Literal: "",
		Pos:     l.pos,
		End:     l.pos,
		Line:    l.line,
		Column:  l.column,
	}
}

// Tokenize tokenizes the entire source code and returns all tokens.
func (l *Lexer) Tokenize() []Token {
	// TODO: Implement full tokenization
	return []Token{}
}
