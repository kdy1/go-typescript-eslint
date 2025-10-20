package lexer

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

// scanIdentifier scans an identifier or keyword.
//
//nolint:gocyclo,cyclop // Identifier scanning requires checking many character classes
func (s *Scanner) scanIdentifier() Token {
	start := s.pos
	ch := s.char()

	// Check if this is a valid identifier start
	if !isIdentifierStart(ch) {
		return s.createToken(ILLEGAL, string(ch))
	}

	// Fast path for ASCII-only identifiers
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_' || ch == '$' {
		hasUnicode := false
		for {
			s.next()
			ch = s.char()
			if ch >= utf8.RuneSelf {
				// Contains Unicode, switch to slow path
				hasUnicode = true
				break
			}
			//nolint:staticcheck // De Morgan's law would make this harder to read
			if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') ||
				(ch >= '0' && ch <= '9') || ch == '_' || ch == '$') {
				break
			}
		}

		// If we found Unicode characters, continue with Unicode scanning
		if hasUnicode {
			for isIdentifierPart(s.char()) {
				s.nextRune()
			}
		}
	} else {
		// Unicode identifier - use full rune scanning
		s.nextRune()
		for isIdentifierPart(s.char()) {
			s.nextRune()
		}
	}

	literal := s.source[start:s.pos]
	tokenType := lookupKeyword(literal)

	return s.createToken(tokenType, literal)
}

// scanNumber scans a numeric literal (decimal, hex, binary, octal, float, bigint).
//
//nolint:gocognit,gocyclo,cyclop,dupl,funlen // Number scanning handles many formats with similar patterns
func (s *Scanner) scanNumber() Token {
	start := s.pos
	ch := s.char()

	// Hexadecimal: 0x1234, 0xABCD
	if ch == '0' && (s.peek(1) == 'x' || s.peek(1) == 'X') {
		s.next() // consume '0'
		s.next() // consume 'x' or 'X'

		if !isHexDigit(s.char()) {
			return s.createToken(ILLEGAL, s.source[start:s.pos])
		}

		for isHexDigit(s.char()) || s.char() == '_' {
			s.next()
		}

		// Check for BigInt suffix
		if s.char() == 'n' {
			s.next()
		}

		return s.createToken(NUMBER, s.source[start:s.pos])
	}

	// Binary: 0b1010
	if ch == '0' && (s.peek(1) == 'b' || s.peek(1) == 'B') {
		s.next() // consume '0'
		s.next() // consume 'b' or 'B'

		if !isBinaryDigit(s.char()) {
			return s.createToken(ILLEGAL, s.source[start:s.pos])
		}

		for isBinaryDigit(s.char()) || s.char() == '_' {
			s.next()
		}

		// Check for BigInt suffix
		if s.char() == 'n' {
			s.next()
		}

		return s.createToken(NUMBER, s.source[start:s.pos])
	}

	// Octal: 0o777
	if ch == '0' && (s.peek(1) == 'o' || s.peek(1) == 'O') {
		s.next() // consume '0'
		s.next() // consume 'o' or 'O'

		if !isOctalDigit(s.char()) {
			return s.createToken(ILLEGAL, s.source[start:s.pos])
		}

		for isOctalDigit(s.char()) || s.char() == '_' {
			s.next()
		}

		// Check for BigInt suffix
		if s.char() == 'n' {
			s.next()
		}

		return s.createToken(NUMBER, s.source[start:s.pos])
	}

	// Decimal number (including floats)
	for isDigit(s.char()) || s.char() == '_' {
		s.next()
	}

	// Fractional part
	if s.char() == '.' && isDigit(s.peek(1)) {
		s.next() // consume '.'
		for isDigit(s.char()) || s.char() == '_' {
			s.next()
		}
	}

	// Exponent part
	if s.char() == 'e' || s.char() == 'E' {
		s.next()
		if s.char() == '+' || s.char() == '-' {
			s.next()
		}
		if !isDigit(s.char()) {
			return s.createToken(ILLEGAL, s.source[start:s.pos])
		}
		for isDigit(s.char()) || s.char() == '_' {
			s.next()
		}
	}

	// Check for BigInt suffix
	if s.char() == 'n' {
		s.next()
	}

	return s.createToken(NUMBER, s.source[start:s.pos])
}

// scanString scans a string literal (single or double quoted).
func (s *Scanner) scanString(quote rune) Token {
	start := s.pos
	s.next() // consume opening quote

	var sb strings.Builder

	for {
		ch := s.char()

		if ch == -1 {
			// Unterminated string
			return s.createToken(ILLEGAL, s.source[start:s.pos])
		}

		if ch == quote {
			s.next() // consume closing quote
			break
		}

		//nolint:staticcheck // Switch would not improve readability here
		if ch == '\\' {
			// Escape sequence
			s.next()
			escaped := s.scanEscapeSequence()
			sb.WriteString(escaped)
		} else if ch == '\n' || ch == '\r' {
			// Unescaped newline in string
			return s.createToken(ILLEGAL, s.source[start:s.pos])
		} else {
			sb.WriteRune(ch)
			s.next()
		}
	}

	// Return the processed string value (without quotes)
	return s.createToken(STRING, sb.String())
}

// scanEscapeSequence scans an escape sequence and returns the escaped character.
//
//nolint:gocyclo,cyclop,funlen // Escape sequence handling requires many cases
func (s *Scanner) scanEscapeSequence() string {
	ch := s.char()
	s.next()

	switch ch {
	case 'b':
		return "\b"
	case 't':
		return "\t"
	case 'n':
		return "\n"
	case 'v':
		return "\v"
	case 'f':
		return "\f"
	case 'r':
		return "\r"
	case '"':
		return "\""
	case '\'':
		return "'"
	case '\\':
		return "\\"
	case '0':
		// Null character (only if not followed by a digit)
		if !isDigit(s.char()) {
			return "\x00"
		}
		// Otherwise, it's an octal escape
		return s.scanOctalEscape(ch)
	case '1', '2', '3', '4', '5', '6', '7':
		return s.scanOctalEscape(ch)
	case 'x':
		// Hex escape: \xAB
		//nolint:mnd // 2 hex digits for \xNN escape sequence
		return s.scanHexEscape(2)
	case 'u':
		// Unicode escape: \uABCD or \u{10FFFF}
		if s.char() == '{' {
			s.next() // consume '{'
			hex := s.scanHexDigits()
			if s.char() != '}' {
				return string(utf8.RuneError)
			}
			s.next() // consume '}'
			val, err := strconv.ParseInt(hex, 16, 32)
			if err != nil || val > 0x10FFFF {
				return string(utf8.RuneError)
			}
			return string(rune(val))
		}
		//nolint:mnd // 4 hex digits for \uNNNN escape sequence
		return s.scanHexEscape(4)
	case '\r', '\n':
		// Line continuation
		if ch == '\r' && s.char() == '\n' {
			s.next()
		}
		return ""
	default:
		return string(ch)
	}
}

// scanOctalEscape scans an octal escape sequence (\0-\377).
func (s *Scanner) scanOctalEscape(first rune) string {
	octal := string(first)

	// Octal escapes can be 1-3 digits
	for i := 0; i < 2 && isOctalDigit(s.char()); i++ {
		octal += string(s.char())
		s.next()
	}

	val, err := strconv.ParseInt(octal, 8, 32)
	if err != nil {
		return string(utf8.RuneError)
	}

	return string(rune(val))
}

// scanHexEscape scans a hex escape sequence (\xNN or \uNNNN).
func (s *Scanner) scanHexEscape(digits int) string {
	hex := ""
	for i := 0; i < digits && isHexDigit(s.char()); i++ {
		hex += string(s.char())
		s.next()
	}

	if len(hex) != digits {
		return string(utf8.RuneError)
	}

	val, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return string(utf8.RuneError)
	}

	return string(rune(val))
}

// scanHexDigits scans a sequence of hex digits (for Unicode escapes).
func (s *Scanner) scanHexDigits() string {
	start := s.pos
	for isHexDigit(s.char()) {
		s.next()
	}
	return s.source[start:s.pos]
}

// scanTemplate scans a template literal.
func (s *Scanner) scanTemplate() Token {
	start := s.pos
	s.next() // consume opening backtick

	var sb strings.Builder
	hasSubstitution := false

	for {
		ch := s.char()

		if ch == -1 {
			// Unterminated template
			return s.createToken(ILLEGAL, s.source[start:s.pos])
		}

		if ch == '`' {
			s.next() // consume closing backtick
			break
		}

		if ch == '$' && s.peek(1) == '{' {
			// Template substitution
			s.next() // consume '$'
			s.next() // consume '{'

			// Return template head or middle token
			tokenType := TemplateHead
			if start != s.offset {
				tokenType = TemplateMiddle
			}

			return s.createToken(tokenType, sb.String())
		}

		if ch == '\\' {
			// Escape sequence
			s.next()
			escaped := s.scanEscapeSequence()
			sb.WriteString(escaped)
		} else {
			sb.WriteRune(ch)
			s.next()
		}
	}

	// Template with no substitution
	if !hasSubstitution {
		return s.createToken(TemplateNoSub, sb.String())
	}

	return s.createToken(TemplateTail, sb.String())
}

// scanLineComment scans a single-line comment (//...).
func (s *Scanner) scanLineComment() Token {
	start := s.pos
	s.next() // consume first '/'
	s.next() // consume second '/'

	for s.char() != '\n' && s.char() != '\r' && s.char() != -1 {
		s.next()
	}

	return s.createToken(COMMENT, s.source[start:s.pos])
}

// scanBlockComment scans a multi-line comment (/* ... */).
func (s *Scanner) scanBlockComment() Token {
	start := s.pos
	s.next() // consume '/'
	s.next() // consume '*'

	for {
		ch := s.char()

		if ch == -1 {
			// Unterminated comment
			return s.createToken(ILLEGAL, s.source[start:s.pos])
		}

		if ch == '*' && s.peek(1) == '/' {
			s.next() // consume '*'
			s.next() // consume '/'
			break
		}

		s.next()
	}

	return s.createToken(COMMENT, s.source[start:s.pos])
}

// scanRegExp scans a regular expression literal.
//
//nolint:gocyclo,cyclop,unused // RegExp scanning requires handling many special cases, will be used in future
func (s *Scanner) scanRegExp() Token {
	start := s.pos
	s.next() // consume opening '/'

	inCharClass := false

	for {
		ch := s.char()

		if ch == -1 || ch == '\n' || ch == '\r' {
			// Unterminated regex
			return s.createToken(ILLEGAL, s.source[start:s.pos])
		}

		//nolint:staticcheck // Switch would not improve readability for character class handling
		if ch == '\\' {
			// Escaped character
			s.next()
			s.next()
			continue
		} else if ch == '[' {
			inCharClass = true
		} else if ch == ']' {
			inCharClass = false
		}

		if ch == '/' && !inCharClass {
			s.next() // consume closing '/'
			break
		}

		s.next()
	}

	// Scan regex flags (g, i, m, s, u, y)
	//nolint:revive // Loop pattern is clearer than inverted logic
	for {
		ch := s.char()
		if ch == 'g' || ch == 'i' || ch == 'm' || ch == 's' || ch == 'u' || ch == 'y' {
			s.next()
		} else {
			break
		}
	}

	return s.createToken(REGEXP, s.source[start:s.pos])
}
