package lexer

import (
	"unicode"
	"unicode/utf8"
)

// Scanner represents a stateful lexical scanner for TypeScript source code.
// It uses rune-based scanning for proper Unicode support and maintains detailed
// position tracking for error reporting and source mapping.
type Scanner struct {
	source  string // Source code to scan
	current Token  // Current token state

	// Position tracking (grouped for better memory alignment)
	pos        int // Current byte position
	offset     int // Start of current token
	fullOffset int // Start including whitespace/comments
	length     int // Length of source in bytes

	// Line/column tracking (1-based line, 0-based column)
	line        int
	column      int
	tokenLine   int // Line number at token start
	tokenColumn int // Column number at token start

	// Configuration
	skipComments bool // Whether to skip comments
	jsxMode      bool // Whether JSX scanning is enabled
}

// NewScanner creates a new Scanner for the given source code.
func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		length: len(source),
		pos:    0,
		offset: 0,
		line:   1,
		column: 0,
	}
}

// SetSkipComments configures whether the scanner should skip comment tokens.
func (s *Scanner) SetSkipComments(skip bool) {
	s.skipComments = skip
}

// SetJSXMode enables or disables JSX-aware scanning.
func (s *Scanner) SetJSXMode(enabled bool) {
	s.jsxMode = enabled
}

// char returns the current character without advancing the position.
// Returns -1 if at EOF.
func (s *Scanner) char() rune {
	if s.pos >= s.length {
		return -1
	}
	return rune(s.source[s.pos])
}

// peek returns the character at the given offset from current position.
// Returns -1 if offset is out of bounds.
func (s *Scanner) peek(offset int) rune {
	p := s.pos + offset
	if p >= s.length {
		return -1
	}
	return rune(s.source[p])
}

// next advances the position by one character and returns it.
// Returns -1 if at EOF.
//
//nolint:unparam // Return value is used in some cases
func (s *Scanner) next() rune {
	if s.pos >= s.length {
		return -1
	}

	ch := rune(s.source[s.pos])
	s.pos++

	// Track line/column
	//nolint:gocritic // ifElseChain: if-else is clearer than switch for line tracking
	if ch == '\n' {
		s.line++
		s.column = 0
	} else if ch == '\r' {
		s.line++
		s.column = 0
		// Handle \r\n as single line break
		if s.char() == '\n' {
			s.pos++
		}
	} else {
		s.column++
	}

	return ch
}

// nextRune advances the position by one full UTF-8 rune and returns it.
// Returns -1 if at EOF.
//
//nolint:unparam // Return value is used in some cases
func (s *Scanner) nextRune() rune {
	if s.pos >= s.length {
		return -1
	}

	ch, size := utf8.DecodeRuneInString(s.source[s.pos:])
	if ch == utf8.RuneError && size == 1 {
		// Invalid UTF-8
		s.pos++
		s.column++
		return ch
	}

	s.pos += size

	// Track line/column
	//nolint:gocritic // ifElseChain: if-else is clearer than switch for line tracking
	if ch == '\n' {
		s.line++
		s.column = 0
	} else if ch == '\r' {
		s.line++
		s.column = 0
		// Handle \r\n as single line break
		if s.char() == '\n' {
			s.pos++
		}
	} else {
		s.column++
	}

	return ch
}

// skipWhitespace advances the scanner position past any whitespace characters.
func (s *Scanner) skipWhitespace() {
	for {
		ch := s.char()
		if ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n' {
			s.next()
		} else {
			break
		}
	}
}

// isLetter checks if a rune is a letter (including Unicode).
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_' || ch == '$' ||
		(ch >= utf8.RuneSelf && unicode.IsLetter(ch))
}

// isDigit checks if a rune is a decimal digit.
func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

// isHexDigit checks if a rune is a hexadecimal digit.
func isHexDigit(ch rune) bool {
	return isDigit(ch) || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F')
}

// isBinaryDigit checks if a rune is a binary digit (0 or 1).
func isBinaryDigit(ch rune) bool {
	return ch == '0' || ch == '1'
}

// isOctalDigit checks if a rune is an octal digit.
func isOctalDigit(ch rune) bool {
	return ch >= '0' && ch <= '7'
}

// isIdentifierStart checks if a rune can start an identifier.
func isIdentifierStart(ch rune) bool {
	return isLetter(ch)
}

// isIdentifierPart checks if a rune can be part of an identifier.
func isIdentifierPart(ch rune) bool {
	return isLetter(ch) || isDigit(ch) || (ch >= utf8.RuneSelf && unicode.IsDigit(ch))
}

// createToken creates a token with the current position information.
func (s *Scanner) createToken(typ TokenType, literal string) Token {
	return Token{
		Type:    typ,
		Literal: literal,
		Pos:     s.offset,
		End:     s.pos,
		Line:    s.tokenLine,
		Column:  s.tokenColumn,
	}
}

// Keywords maps string literals to their corresponding token types.
var keywords = map[string]TokenType{
	"break":       BREAK,
	"case":        CASE,
	"catch":       CATCH,
	"class":       CLASS,
	"const":       CONST,
	"continue":    CONTINUE,
	"debugger":    DEBUGGER,
	"default":     DEFAULT,
	"delete":      DELETE,
	"do":          DO,
	"else":        ELSE,
	"enum":        ENUM,
	"export":      EXPORT,
	"extends":     EXTENDS,
	"false":       FALSE,
	"finally":     FINALLY,
	"for":         FOR,
	"function":    FUNCTION,
	"if":          IF,
	"import":      IMPORT,
	"in":          IN,
	"instanceof":  INSTANCEOF,
	"new":         NEW,
	"null":        NULL,
	"return":      RETURN,
	"super":       SUPER,
	"switch":      SWITCH,
	"this":        THIS,
	"throw":       THROW,
	"true":        TRUE,
	"try":         TRY,
	"typeof":      TYPEOF,
	"var":         VAR,
	"void":        VOID,
	"while":       WHILE,
	"with":        WITH,
	"yield":       YIELD,
	"as":          AS,
	"async":       ASYNC,
	"await":       AWAIT,
	"declare":     DECLARE,
	"interface":   INTERFACE,
	"let":         LET,
	"module":      MODULE,
	"namespace":   NAMESPACE,
	"of":          OF,
	"package":     PACKAGE,
	"private":     PRIVATE,
	"protected":   PROTECTED,
	"public":      PUBLIC,
	"readonly":    READONLY,
	"require":     REQUIRE,
	"static":      STATIC,
	"type":        TYPE,
	"from":        FROM,
	"satisfies":   SATISFIES,
	"implements":  IMPLEMENTS,
	"any":         ANY,
	"boolean":     BOOLEAN,
	"constructor": CONSTRUCTOR,
	"get":         GET,
	"set":         SET,
	"never":       NEVER,
	"unknown":     UNKNOWN,
	"string":      StringKeyword,
	"number":      NumberKeyword,
	"symbol":      SYMBOL,
	"undefined":   UNDEFINED,
}

// lookupKeyword checks if an identifier is a keyword and returns the appropriate token type.
func lookupKeyword(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
