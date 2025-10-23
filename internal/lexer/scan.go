package lexer

// Scan returns the next token from the source code.
// It is the main entry point for the scanner.
//
//nolint:gocognit,gocyclo,cyclop,funlen // Scanner functions are inherently complex
func (s *Scanner) Scan() Token {
	// Skip whitespace and comments (if configured)
	for {
		s.skipWhitespace()
		s.fullOffset = s.pos
		s.offset = s.pos

		// Check for comments
		//nolint:nestif // Comment handling requires nested conditions
		if s.char() == '/' {
			if s.peek(1) == '/' {
				token := s.scanLineComment()
				if !s.skipComments {
					s.current = token
					return token
				}
				continue
			}
			if s.peek(1) == '*' {
				token := s.scanBlockComment()
				if !s.skipComments {
					s.current = token
					return token
				}
				continue
			}
		}
		break
	}

	s.offset = s.pos
	s.tokenLine = s.line
	s.tokenColumn = s.column
	ch := s.char()

	// EOF
	if ch == -1 {
		s.current = s.createToken(EOF, "")
		return s.current
	}

	// Identifiers and keywords
	if isIdentifierStart(ch) {
		token := s.scanIdentifier()
		s.current = token
		return token
	}

	// Numbers
	if isDigit(ch) {
		token := s.scanNumber()
		s.current = token
		return token
	}

	// Operators and punctuation
	switch ch {
	// String literals
	case '"', '\'':
		token := s.scanString(ch)
		s.current = token
		return token

	// Template literals
	case '`':
		token := s.scanTemplate()
		s.current = token
		return token

	// Punctuation
	case '(':
		s.next()
		s.current = s.createToken(LPAREN, "(")
		return s.current

	case ')':
		s.next()
		s.current = s.createToken(RPAREN, ")")
		return s.current

	case '{':
		s.next()
		s.current = s.createToken(LBRACE, "{")
		return s.current

	case '}':
		s.next()
		s.current = s.createToken(RBRACE, "}")
		return s.current

	case '[':
		s.next()
		s.current = s.createToken(LBRACK, "[")
		return s.current

	case ']':
		s.next()
		s.current = s.createToken(RBRACK, "]")
		return s.current

	case ';':
		s.next()
		s.current = s.createToken(SEMICOLON, ";")
		return s.current

	case ',':
		s.next()
		s.current = s.createToken(COMMA, ",")
		return s.current

	case ':':
		s.next()
		s.current = s.createToken(COLON, ":")
		return s.current

	case '~':
		s.next()
		s.current = s.createToken(BNOT, "~")
		return s.current

	// Operators with multiple characters
	case '.':
		s.next()
		// Check for ... (ellipsis)
		if s.char() == '.' && s.peek(1) == '.' {
			s.next()
			s.next()
			s.current = s.createToken(ELLIPSIS, "...")
			return s.current
		}
		// Check for number starting with . (e.g., .123)
		if isDigit(s.char()) {
			s.pos = s.offset // Reset position
			token := s.scanNumber()
			s.current = token
			return token
		}
		s.current = s.createToken(PERIOD, ".")
		return s.current

	case '?':
		s.next()
		// Check for ??= (nullish coalescing assignment)
		if s.char() == '?' && s.peek(1) == '=' {
			s.next()
			s.next()
			s.current = s.createToken(NullishAssign, "??=")
			return s.current
		}
		// Check for ?? (nullish coalescing)
		if s.char() == '?' {
			s.next()
			s.current = s.createToken(NULLISH, "??")
			return s.current
		}
		// Check for ?. (optional chaining)
		if s.char() == '.' && !isDigit(s.peek(1)) {
			s.next()
			s.current = s.createToken(OPTIONAL, "?.")
			return s.current
		}
		s.current = s.createToken(QUESTION, "?")
		return s.current

	case '+':
		s.next()
		// Check for += (addition assignment)
		if s.char() == '=' {
			s.next()
			s.current = s.createToken(AddAssign, "+=")
			return s.current
		}
		// Check for ++ (increment)
		if s.char() == '+' {
			s.next()
			s.current = s.createToken(INC, "++")
			return s.current
		}
		s.current = s.createToken(ADD, "+")
		return s.current

	case '-':
		s.next()
		// Check for -= (subtraction assignment)
		if s.char() == '=' {
			s.next()
			s.current = s.createToken(SubAssign, "-=")
			return s.current
		}
		// Check for -- (decrement)
		if s.char() == '-' {
			s.next()
			s.current = s.createToken(DEC, "--")
			return s.current
		}
		s.current = s.createToken(SUB, "-")
		return s.current

	case '*':
		s.next()
		// Check for **= (exponentiation assignment)
		if s.char() == '*' && s.peek(1) == '=' {
			s.next()
			s.next()
			s.current = s.createToken(PowerAssign, "**=")
			return s.current
		}
		// Check for ** (exponentiation)
		if s.char() == '*' {
			s.next()
			s.current = s.createToken(POWER, "**")
			return s.current
		}
		// Check for *= (multiplication assignment)
		if s.char() == '=' {
			s.next()
			s.current = s.createToken(MulAssign, "*=")
			return s.current
		}
		s.current = s.createToken(MUL, "*")
		return s.current

	case '/':
		s.next()
		// Check for /= (division assignment)
		if s.char() == '=' {
			s.next()
			s.current = s.createToken(QuoAssign, "/=")
			return s.current
		}
		s.current = s.createToken(QUO, "/")
		return s.current

	case '%':
		s.next()
		// Check for %= (modulo assignment)
		if s.char() == '=' {
			s.next()
			s.current = s.createToken(RemAssign, "%=")
			return s.current
		}
		s.current = s.createToken(REM, "%")
		return s.current

	case '&':
		s.next()
		// Check for &&= (logical AND assignment)
		if s.char() == '&' && s.peek(1) == '=' {
			s.next()
			s.next()
			s.current = s.createToken(AndAssign, "&&=")
			return s.current
		}
		// Check for && (logical AND)
		if s.char() == '&' {
			s.next()
			s.current = s.createToken(LAND, "&&")
			return s.current
		}
		// Check for &= (bitwise AND assignment)
		if s.char() == '=' {
			s.next()
			s.current = s.createToken(AndAssign, "&=")
			return s.current
		}
		s.current = s.createToken(AND, "&")
		return s.current

	case '|':
		s.next()
		// Check for ||= (logical OR assignment)
		if s.char() == '|' && s.peek(1) == '=' {
			s.next()
			s.next()
			s.current = s.createToken(OrAssign, "||=")
			return s.current
		}
		// Check for || (logical OR)
		if s.char() == '|' {
			s.next()
			s.current = s.createToken(LOR, "||")
			return s.current
		}
		// Check for |= (bitwise OR assignment)
		if s.char() == '=' {
			s.next()
			s.current = s.createToken(OrAssign, "|=")
			return s.current
		}
		s.current = s.createToken(OR, "|")
		return s.current

	case '^':
		s.next()
		// Check for ^= (bitwise XOR assignment)
		if s.char() == '=' {
			s.next()
			s.current = s.createToken(XorAssign, "^=")
			return s.current
		}
		s.current = s.createToken(XOR, "^")
		return s.current

	case '<':
		s.next()
		// Check for <<<= (unsigned left shift assignment) - doesn't exist in JS/TS
		// Check for <<= (left shift assignment)
		if s.char() == '<' && s.peek(1) == '=' {
			s.next()
			s.next()
			s.current = s.createToken(ShlAssign, "<<=")
			return s.current
		}
		// Check for << (left shift)
		if s.char() == '<' {
			s.next()
			s.current = s.createToken(SHL, "<<")
			return s.current
		}
		// Check for <= (less than or equal)
		if s.char() == '=' {
			s.next()
			s.current = s.createToken(LEQ, "<=")
			return s.current
		}
		s.current = s.createToken(LSS, "<")
		return s.current

	case '>':
		s.next()
		// Check for >>>= (unsigned right shift assignment)
		//nolint:mnd // 2 is the lookahead offset for checking third character
		if s.char() == '>' && s.peek(1) == '>' && s.peek(2) == '=' {
			s.next()
			s.next()
			s.next()
			s.current = s.createToken(ShrUnsignedAssign, ">>>=")
			return s.current
		}
		// Check for >>> (unsigned right shift)
		if s.char() == '>' && s.peek(1) == '>' {
			s.next()
			s.next()
			s.current = s.createToken(SHRUnsigned, ">>>")
			return s.current
		}
		// Check for >>= (right shift assignment)
		if s.char() == '>' && s.peek(1) == '=' {
			s.next()
			s.next()
			s.current = s.createToken(ShrAssign, ">>=")
			return s.current
		}
		// Check for >> (right shift)
		if s.char() == '>' {
			s.next()
			s.current = s.createToken(SHR, ">>")
			return s.current
		}
		// Check for >= (greater than or equal)
		if s.char() == '=' {
			s.next()
			s.current = s.createToken(GEQ, ">=")
			return s.current
		}
		s.current = s.createToken(GTR, ">")
		return s.current

	case '=':
		s.next()
		// Check for => (arrow function)
		if s.char() == '>' {
			s.next()
			s.current = s.createToken(ARROW, "=>")
			return s.current
		}
		// Check for === (strict equality)
		if s.char() == '=' && s.peek(1) == '=' {
			s.next()
			s.next()
			s.current = s.createToken(EqlStrict, "===")
			return s.current
		}
		// Check for == (equality)
		if s.char() == '=' {
			s.next()
			s.current = s.createToken(EQL, "==")
			return s.current
		}
		s.current = s.createToken(ASSIGN, "=")
		return s.current

	case '!':
		s.next()
		// Check for !== (strict inequality)
		if s.char() == '=' && s.peek(1) == '=' {
			s.next()
			s.next()
			s.current = s.createToken(NeqStrict, "!==")
			return s.current
		}
		// Check for != (inequality)
		if s.char() == '=' {
			s.next()
			s.current = s.createToken(NEQ, "!=")
			return s.current
		}
		s.current = s.createToken(NOT, "!")
		return s.current

	default:
		// Unknown character
		s.next()
		s.current = s.createToken(ILLEGAL, string(ch))
		return s.current
	}
}

// Current returns the most recently scanned token.
func (s *Scanner) Current() Token {
	return s.current
}

// Reset resets the scanner to the beginning of the source.
func (s *Scanner) Reset() {
	s.pos = 0
	s.offset = 0
	s.fullOffset = 0
	s.line = 1
	s.column = 0
}
