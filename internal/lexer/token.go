package lexer

// TokenType represents the type of a lexical token.
type TokenType int

// Token types
const (
	// Special tokens
	EOF TokenType = iota
	ILLEGAL
	COMMENT

	// Literals
	IDENT  // identifier
	NUMBER // 123, 0x1A, 3.14, 1_000
	STRING // "abc", 'abc'
	TEMPLATE // template string part
	REGEXP // /pattern/flags

	// Keywords
	BREAK
	CASE
	CATCH
	CLASS
	CONST
	CONTINUE
	DEBUGGER
	DEFAULT
	DELETE
	DO
	ELSE
	ENUM
	EXPORT
	EXTENDS
	FALSE
	FINALLY
	FOR
	FUNCTION
	IF
	IMPORT
	IN
	INSTANCEOF
	NEW
	NULL
	RETURN
	SUPER
	SWITCH
	THIS
	THROW
	TRUE
	TRY
	TYPEOF
	VAR
	VOID
	WHILE
	WITH
	YIELD

	// TypeScript keywords
	AS
	ASYNC
	AWAIT
	DECLARE
	INTERFACE
	LET
	MODULE
	NAMESPACE
	OF
	PACKAGE
	PRIVATE
	PROTECTED
	PUBLIC
	READONLY
	REQUIRE
	STATIC
	TYPE
	FROM
	SATISFIES
	IMPLEMENTS
	ANY
	BOOLEAN
	CONSTRUCTOR
	GET
	SET
	NEVER
	UNKNOWN
	STRING_KEYWORD
	NUMBER_KEYWORD
	SYMBOL
	UNDEFINED

	// Operators and punctuation
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	AND // &
	OR  // |
	XOR // ^
	SHL // <<
	SHR // >>

	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	QUO_ASSIGN // /=
	REM_ASSIGN // %=

	AND_ASSIGN // &=
	OR_ASSIGN  // |=
	XOR_ASSIGN // ^=
	SHL_ASSIGN // <<=
	SHR_ASSIGN // >>=

	LAND  // &&
	LOR   // ||
	INC   // ++
	DEC   // --
	NULLISH // ??

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !

	NEQ      // !=
	LEQ      // <=
	GEQ      // >=
	EQL_STRICT // ===
	NEQ_STRICT // !==

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :
	QUESTION  // ?

	ARROW        // =>
	ELLIPSIS     // ...
	OPTIONAL     // ?.
	NULLISH_ASSIGN // ??=
)

// Token represents a lexical token.
type Token struct {
	Type    TokenType
	Literal string
	Pos     int // byte offset of the token start
	End     int // byte offset of the token end
	Line    int // line number (1-based)
	Column  int // column number (0-based)
}

// String returns a string representation of the token type.
func (t TokenType) String() string {
	switch t {
	case EOF:
		return "EOF"
	case ILLEGAL:
		return "ILLEGAL"
	case COMMENT:
		return "COMMENT"
	case IDENT:
		return "IDENT"
	case NUMBER:
		return "NUMBER"
	case STRING:
		return "STRING"
	case TEMPLATE:
		return "TEMPLATE"
	case REGEXP:
		return "REGEXP"
	default:
		return "UNKNOWN"
	}
}
