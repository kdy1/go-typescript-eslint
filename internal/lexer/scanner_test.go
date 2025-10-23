package lexer

import (
	"testing"
)

//nolint:funlen // Table-driven test with many test cases
func TestScannerBasicTokens(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []TokenType
	}{
		{
			name:     "punctuation",
			input:    "( ) { } [ ] ; , : .",
			expected: []TokenType{LPAREN, RPAREN, LBRACE, RBRACE, LBRACK, RBRACK, SEMICOLON, COMMA, COLON, PERIOD, EOF},
		},
		{
			name:     "arithmetic operators",
			input:    "+ - * / %",
			expected: []TokenType{ADD, SUB, MUL, QUO, REM, EOF},
		},
		{
			name:     "comparison operators",
			input:    "< > <= >= == != === !==",
			expected: []TokenType{LSS, GTR, LEQ, GEQ, EQL, NEQ, EqlStrict, NeqStrict, EOF},
		},
		{
			name:     "logical operators",
			input:    "! && ||",
			expected: []TokenType{NOT, LAND, LOR, EOF},
		},
		{
			name:     "bitwise operators",
			input:    "& | ^ ~ << >> >>>",
			expected: []TokenType{AND, OR, XOR, BNOT, SHL, SHR, SHRUnsigned, EOF},
		},
		{
			name:  "assignment operators",
			input: "= += -= *= /= %= &= |= ^= <<= >>= >>>=",
			expected: []TokenType{
				ASSIGN, AddAssign, SubAssign, MulAssign, QuoAssign, RemAssign,
				AndAssign, OrAssign, XorAssign, ShlAssign, ShrAssign, ShrUnsignedAssign, EOF,
			},
		},
		{
			name:     "special operators",
			input:    "++ -- ... => ? ?. ?? ??=",
			expected: []TokenType{INC, DEC, ELLIPSIS, ARROW, QUESTION, OPTIONAL, NULLISH, NullishAssign, EOF},
		},
		{
			name:     "exponentiation",
			input:    "** **=",
			expected: []TokenType{POWER, PowerAssign, EOF},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(tt.input)
			for i, expected := range tt.expected {
				token := scanner.Scan()
				if token.Type != expected {
					t.Errorf("token %d: expected %v, got %v (literal: %q)", i, expected, token.Type, token.Literal)
				}
			}
		})
	}
}

//nolint:funlen // Table-driven test with many test cases
func TestScannerKeywords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected TokenType
	}{
		// JavaScript keywords
		{"break", "break", BREAK},
		{"case", "case", CASE},
		{"catch", "catch", CATCH},
		{"class", "class", CLASS},
		{"const", "const", CONST},
		{"continue", "continue", CONTINUE},
		{"debugger", "debugger", DEBUGGER},
		{"default", "default", DEFAULT},
		{"delete", "delete", DELETE},
		{"do", "do", DO},
		{"else", "else", ELSE},
		{"enum", "enum", ENUM},
		{"export", "export", EXPORT},
		{"extends", "extends", EXTENDS},
		{"false", "false", FALSE},
		{"finally", "finally", FINALLY},
		{"for", "for", FOR},
		{"function", "function", FUNCTION},
		{"if", "if", IF},
		{"import", "import", IMPORT},
		{"in", "in", IN},
		{"instanceof", "instanceof", INSTANCEOF},
		{"new", "new", NEW},
		{"null", "null", NULL},
		{"return", "return", RETURN},
		{"super", "super", SUPER},
		{"switch", "switch", SWITCH},
		{"this", "this", THIS},
		{"throw", "throw", THROW},
		{"true", "true", TRUE},
		{"try", "try", TRY},
		{"typeof", "typeof", TYPEOF},
		{"var", "var", VAR},
		{"void", "void", VOID},
		{"while", "while", WHILE},
		{"with", "with", WITH},
		{"yield", "yield", YIELD},

		// TypeScript keywords
		{"as", "as", AS},
		{"async", "async", ASYNC},
		{"await", "await", AWAIT},
		{"declare", "declare", DECLARE},
		{"interface", "interface", INTERFACE},
		{"let", "let", LET},
		{"module", "module", MODULE},
		{"namespace", "namespace", NAMESPACE},
		{"of", "of", OF},
		{"private", "private", PRIVATE},
		{"protected", "protected", PROTECTED},
		{"public", "public", PUBLIC},
		{"readonly", "readonly", READONLY},
		{"static", "static", STATIC},
		{"type", "type", TYPE},
		{"from", "from", FROM},
		{"satisfies", "satisfies", SATISFIES},
		{"implements", "implements", IMPLEMENTS},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(tt.input)
			token := scanner.Scan()
			if token.Type != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, token.Type)
			}
			if token.Literal != tt.input {
				t.Errorf("expected literal %q, got %q", tt.input, token.Literal)
			}
		})
	}
}

func TestScannerIdentifiers(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"simple", "foo"},
		{"camelCase", "fooBar"},
		{"underscore", "_private"},
		{"dollar", "$jQuery"},
		{"mixed", "foo_bar$123"},
		{"unicode", "naïve"},
		{"emoji", "test😀"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(tt.input)
			token := scanner.Scan()
			if token.Type != IDENT {
				t.Errorf("expected IDENT, got %v", token.Type)
			}
			if token.Literal != tt.input {
				t.Errorf("expected %q, got %q", tt.input, token.Literal)
			}
		})
	}
}

func TestScannerNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"integer", "123", "123"},
		{"float", "123.456", "123.456"},
		{"exponent", "1e10", "1e10"},
		{"exponent_positive", "1e+10", "1e+10"},
		{"exponent_negative", "1e-10", "1e-10"},
		{"hex", "0x1A2B", "0x1A2B"},
		{"hex_upper", "0XABCD", "0XABCD"},
		{"binary", "0b1010", "0b1010"},
		{"octal", "0o777", "0o777"},
		{"separator", "1_000_000", "1_000_000"},
		{"bigint", "123n", "123n"},
		{"hex_bigint", "0xABn", "0xABn"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(tt.input)
			token := scanner.Scan()
			if token.Type != NUMBER {
				t.Errorf("expected NUMBER, got %v", token.Type)
			}
			if token.Literal != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, token.Literal)
			}
		})
	}
}

func TestScannerStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"double_quotes", `"hello"`, "hello"},
		{"single_quotes", `'hello'`, "hello"},
		{"empty", `""`, ""},
		{"escape_newline", `"hello\nworld"`, "hello\nworld"},
		{"escape_tab", `"hello\tworld"`, "hello\tworld"},
		{"escape_quote", `"hello\"world"`, `hello"world`},
		{"escape_backslash", `"hello\\world"`, `hello\world`},
		{"hex_escape", `"\x41"`, "A"},
		{"unicode_escape", `"\u0041"`, "A"},
		{"unicode_extended", `"\u{1F600}"`, "😀"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(tt.input)
			token := scanner.Scan()
			if token.Type != STRING {
				t.Errorf("expected STRING, got %v", token.Type)
			}
			if token.Literal != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, token.Literal)
			}
		})
	}
}

func TestScannerTemplates(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected TokenType
	}{
		{"simple", "`hello`", TemplateNoSub},
		{"empty", "``", TemplateNoSub},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(tt.input)
			token := scanner.Scan()
			if token.Type != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, token.Type)
			}
		})
	}
}

func TestScannerComments(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected TokenType
	}{
		{"line_comment", "// comment", COMMENT},
		{"block_comment", "/* comment */", COMMENT},
		{"multiline_block", "/* line1\nline2 */", COMMENT},
		{"jsdoc", "/** JSDoc */", COMMENT},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(tt.input)
			scanner.SetSkipComments(false)
			token := scanner.Scan()
			if token.Type != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, token.Type)
			}
		})
	}
}

func TestScannerSkipComments(t *testing.T) {
	input := "// comment\nfoo"
	scanner := NewScanner(input)
	scanner.SetSkipComments(true)

	token := scanner.Scan()
	if token.Type != IDENT {
		t.Errorf("expected IDENT, got %v", token.Type)
	}
	if token.Literal != "foo" {
		t.Errorf("expected 'foo', got %q", token.Literal)
	}
}

func TestScannerPositionTracking(t *testing.T) {
	input := "foo\nbar"
	scanner := NewScanner(input)

	token1 := scanner.Scan()
	if token1.Line != 1 {
		t.Errorf("expected line 1, got %d", token1.Line)
	}

	token2 := scanner.Scan()
	if token2.Line != 2 {
		t.Errorf("expected line 2, got %d", token2.Line)
	}
}

func TestScannerComplexExpression(t *testing.T) {
	input := `const x: number = 42;`
	expected := []TokenType{CONST, IDENT, COLON, NumberKeyword, ASSIGN, NUMBER, SEMICOLON, EOF}

	scanner := NewScanner(input)
	for i, expectedType := range expected {
		token := scanner.Scan()
		if token.Type != expectedType {
			t.Errorf("token %d: expected %v, got %v (literal: %q)", i, expectedType, token.Type, token.Literal)
		}
	}
}

func TestScannerFunctionDeclaration(t *testing.T) {
	input := `function add(a, b) { return a + b; }`
	expected := []TokenType{
		FUNCTION, IDENT, LPAREN, IDENT, COMMA, IDENT, RPAREN,
		LBRACE, RETURN, IDENT, ADD, IDENT, SEMICOLON, RBRACE, EOF,
	}

	scanner := NewScanner(input)
	for i, expectedType := range expected {
		token := scanner.Scan()
		if token.Type != expectedType {
			t.Errorf("token %d: expected %v, got %v (literal: %q)", i, expectedType, token.Type, token.Literal)
		}
	}
}

func TestScannerArrowFunction(t *testing.T) {
	input := `const add = (a, b) => a + b`
	expected := []TokenType{
		CONST, IDENT, ASSIGN, LPAREN, IDENT, COMMA, IDENT, RPAREN,
		ARROW, IDENT, ADD, IDENT, EOF,
	}

	scanner := NewScanner(input)
	for i, expectedType := range expected {
		token := scanner.Scan()
		if token.Type != expectedType {
			t.Errorf("token %d: expected %v, got %v (literal: %q)", i, expectedType, token.Type, token.Literal)
		}
	}
}

func TestScannerClassDeclaration(t *testing.T) {
	input := `class Foo extends Bar {}`
	expected := []TokenType{CLASS, IDENT, EXTENDS, IDENT, LBRACE, RBRACE, EOF}

	scanner := NewScanner(input)
	for i, expectedType := range expected {
		token := scanner.Scan()
		if token.Type != expectedType {
			t.Errorf("token %d: expected %v, got %v (literal: %q)", i, expectedType, token.Type, token.Literal)
		}
	}
}

func TestScannerOptionalChaining(t *testing.T) {
	input := `obj?.prop`
	expected := []TokenType{IDENT, OPTIONAL, IDENT, EOF}

	scanner := NewScanner(input)
	for i, expectedType := range expected {
		token := scanner.Scan()
		if token.Type != expectedType {
			t.Errorf("token %d: expected %v, got %v (literal: %q)", i, expectedType, token.Type, token.Literal)
		}
	}
}

func TestScannerNullishCoalescing(t *testing.T) {
	input := `x ?? y`
	expected := []TokenType{IDENT, NULLISH, IDENT, EOF}

	scanner := NewScanner(input)
	for i, expectedType := range expected {
		token := scanner.Scan()
		if token.Type != expectedType {
			t.Errorf("token %d: expected %v, got %v (literal: %q)", i, expectedType, token.Type, token.Literal)
		}
	}
}

func TestScannerDestructuring(t *testing.T) {
	input := `const { a, b, ...rest } = obj`
	expected := []TokenType{
		CONST, LBRACE, IDENT, COMMA, IDENT, COMMA, ELLIPSIS, IDENT, RBRACE, ASSIGN, IDENT, EOF,
	}

	scanner := NewScanner(input)
	for i, expectedType := range expected {
		token := scanner.Scan()
		if token.Type != expectedType {
			t.Errorf("token %d: expected %v, got %v (literal: %q)", i, expectedType, token.Type, token.Literal)
		}
	}
}

func TestLexerTokenize(t *testing.T) {
	input := `const x = 42;`
	lexer := New(input)
	tokens := lexer.Tokenize()

	if len(tokens) == 0 {
		t.Error("expected tokens, got none")
	}

	// Check that last token is EOF
	lastToken := tokens[len(tokens)-1]
	if lastToken.Type != EOF {
		t.Errorf("expected last token to be EOF, got %v", lastToken.Type)
	}
}

func TestLexerNextToken(t *testing.T) {
	input := `const x = 42;`
	lexer := New(input)

	token1 := lexer.NextToken()
	if token1.Type != CONST {
		t.Errorf("expected CONST, got %v", token1.Type)
	}

	token2 := lexer.NextToken()
	if token2.Type != IDENT {
		t.Errorf("expected IDENT, got %v", token2.Type)
	}
}
