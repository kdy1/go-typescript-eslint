package lexer

// Lexer represents a lexical analyzer for TypeScript source code.
// It wraps the Scanner to provide a higher-level API for tokenization.
type Lexer struct {
	scanner *Scanner
	tokens  []Token
}

// New creates a new Lexer for the given source code.
func New(source string) *Lexer {
	return &Lexer{
		scanner: NewScanner(source),
		tokens:  []Token{},
	}
}

// NextToken returns the next token from the source code.
func (l *Lexer) NextToken() Token {
	return l.scanner.Scan()
}

// Tokenize tokenizes the entire source code and returns all tokens.
func (l *Lexer) Tokenize() []Token {
	if len(l.tokens) > 0 {
		return l.tokens
	}

	l.scanner.Reset()
	tokens := []Token{}

	for {
		token := l.scanner.Scan()
		tokens = append(tokens, token)

		if token.Type == EOF {
			break
		}
	}

	l.tokens = tokens
	return tokens
}
