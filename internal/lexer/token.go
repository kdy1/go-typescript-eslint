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
	IDENT    // identifier
	NUMBER   // 123, 0x1A, 3.14, 1_000
	STRING   // "abc", 'abc'
	TEMPLATE // template string part
	REGEXP   // /pattern/flags

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
	StringKeyword
	NumberKeyword
	SYMBOL
	UNDEFINED

	// Operators and punctuation
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	AND  // &
	OR   // |
	XOR  // ^
	BNOT // ~ (bitwise NOT)
	SHL  // <<
	SHR  // >>

	AddAssign // +=
	SubAssign // -=
	MulAssign // *=
	QuoAssign // /=
	RemAssign // %=

	AndAssign // &=
	OrAssign  // |=
	XorAssign // ^=
	ShlAssign // <<=
	ShrAssign // >>=

	LAND    // &&
	LOR     // ||
	INC     // ++
	DEC     // --
	NULLISH // ??

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !

	NEQ       // !=
	LEQ       // <=
	GEQ       // >=
	EqlStrict // ===
	NeqStrict // !==

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

	ARROW         // =>
	ELLIPSIS      // ...
	OPTIONAL      // ?.
	NullishAssign // ??=

	// JSX Tokens
	JSXText            // JSX text content
	JSXTagStart        // <
	JSXTagEnd          // >
	JSXSelfClosingEnd  // />
	JSXAttributeString // JSX attribute string value

	// Additional operators
	POWER             // **
	PowerAssign       // **=
	SHRUnsigned       // >>> (unsigned right shift)
	ShrUnsignedAssign // >>>=

	// Template literal tokens
	TemplateHead   // `...${
	TemplateMiddle // }...${
	TemplateTail   // }...`
	TemplateNoSub  // `...` (no substitution)
)

// Token represents a lexical token.
type Token struct {
	Literal string
	Type    TokenType
	Pos     int // byte offset of the token start
	End     int // byte offset of the token end
	Line    int // line number (1-based)
	Column  int // column number (0-based)
}

// String returns a string representation of the token type.
func (t TokenType) String() string {
	//nolint:exhaustive,gocyclo // Complete token type mapping
	switch t {
	// Special tokens
	case EOF:
		return "EOF"
	case ILLEGAL:
		return "ILLEGAL"
	case COMMENT:
		return "COMMENT"

	// Literals
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

	// Keywords
	case BREAK:
		return "BREAK"
	case CASE:
		return "CASE"
	case CATCH:
		return "CATCH"
	case CLASS:
		return "CLASS"
	case CONST:
		return "CONST"
	case CONTINUE:
		return "CONTINUE"
	case DEBUGGER:
		return "DEBUGGER"
	case DEFAULT:
		return "DEFAULT"
	case DELETE:
		return "DELETE"
	case DO:
		return "DO"
	case ELSE:
		return "ELSE"
	case ENUM:
		return "ENUM"
	case EXPORT:
		return "EXPORT"
	case EXTENDS:
		return "EXTENDS"
	case FALSE:
		return "FALSE"
	case FINALLY:
		return "FINALLY"
	case FOR:
		return "FOR"
	case FUNCTION:
		return "FUNCTION"
	case IF:
		return "IF"
	case IMPORT:
		return "IMPORT"
	case IN:
		return "IN"
	case INSTANCEOF:
		return "INSTANCEOF"
	case NEW:
		return "NEW"
	case NULL:
		return "NULL"
	case RETURN:
		return "RETURN"
	case SUPER:
		return "SUPER"
	case SWITCH:
		return "SWITCH"
	case THIS:
		return "THIS"
	case THROW:
		return "THROW"
	case TRUE:
		return "TRUE"
	case TRY:
		return "TRY"
	case TYPEOF:
		return "TYPEOF"
	case VAR:
		return "VAR"
	case VOID:
		return "VOID"
	case WHILE:
		return "WHILE"
	case WITH:
		return "WITH"
	case YIELD:
		return "YIELD"

	// TypeScript keywords
	case AS:
		return "AS"
	case ASYNC:
		return "ASYNC"
	case AWAIT:
		return "AWAIT"
	case DECLARE:
		return "DECLARE"
	case INTERFACE:
		return "INTERFACE"
	case LET:
		return "LET"
	case MODULE:
		return "MODULE"
	case NAMESPACE:
		return "NAMESPACE"
	case OF:
		return "OF"
	case PACKAGE:
		return "PACKAGE"
	case PRIVATE:
		return "PRIVATE"
	case PROTECTED:
		return "PROTECTED"
	case PUBLIC:
		return "PUBLIC"
	case READONLY:
		return "READONLY"
	case REQUIRE:
		return "REQUIRE"
	case STATIC:
		return "STATIC"
	case TYPE:
		return "TYPE"
	case FROM:
		return "FROM"
	case SATISFIES:
		return "SATISFIES"
	case IMPLEMENTS:
		return "IMPLEMENTS"
	case ANY:
		return "ANY"
	case BOOLEAN:
		return "BOOLEAN"
	case CONSTRUCTOR:
		return "CONSTRUCTOR"
	case GET:
		return "GET"
	case SET:
		return "SET"
	case NEVER:
		return "NEVER"
	case UNKNOWN:
		return "UNKNOWN"
	case StringKeyword:
		return "StringKeyword"
	case NumberKeyword:
		return "NumberKeyword"
	case SYMBOL:
		return "SYMBOL"
	case UNDEFINED:
		return "UNDEFINED"

	// Operators and punctuation
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case MUL:
		return "MUL"
	case QUO:
		return "QUO"
	case REM:
		return "REM"
	case AND:
		return "AND"
	case OR:
		return "OR"
	case XOR:
		return "XOR"
	case BNOT:
		return "BNOT"
	case SHL:
		return "SHL"
	case SHR:
		return "SHR"
	case AddAssign:
		return "AddAssign"
	case SubAssign:
		return "SubAssign"
	case MulAssign:
		return "MulAssign"
	case QuoAssign:
		return "QuoAssign"
	case RemAssign:
		return "RemAssign"
	case AndAssign:
		return "AndAssign"
	case OrAssign:
		return "OrAssign"
	case XorAssign:
		return "XorAssign"
	case ShlAssign:
		return "ShlAssign"
	case ShrAssign:
		return "ShrAssign"
	case LAND:
		return "LAND"
	case LOR:
		return "LOR"
	case INC:
		return "INC"
	case DEC:
		return "DEC"
	case NULLISH:
		return "NULLISH"
	case EQL:
		return "EQL"
	case LSS:
		return "LSS"
	case GTR:
		return "GTR"
	case ASSIGN:
		return "ASSIGN"
	case NOT:
		return "NOT"
	case NEQ:
		return "NEQ"
	case LEQ:
		return "LEQ"
	case GEQ:
		return "GEQ"
	case EqlStrict:
		return "EqlStrict"
	case NeqStrict:
		return "NeqStrict"
	case LPAREN:
		return "LPAREN"
	case LBRACK:
		return "LBRACK"
	case LBRACE:
		return "LBRACE"
	case COMMA:
		return "COMMA"
	case PERIOD:
		return "PERIOD"
	case RPAREN:
		return "RPAREN"
	case RBRACK:
		return "RBRACK"
	case RBRACE:
		return "RBRACE"
	case SEMICOLON:
		return "SEMICOLON"
	case COLON:
		return "COLON"
	case QUESTION:
		return "QUESTION"
	case ARROW:
		return "ARROW"
	case ELLIPSIS:
		return "ELLIPSIS"
	case OPTIONAL:
		return "OPTIONAL"
	case NullishAssign:
		return "NullishAssign"

	// JSX Tokens
	case JSXText:
		return "JSXText"
	case JSXTagStart:
		return "JSXTagStart"
	case JSXTagEnd:
		return "JSXTagEnd"
	case JSXSelfClosingEnd:
		return "JSXSelfClosingEnd"
	case JSXAttributeString:
		return "JSXAttributeString"

	// Additional operators
	case POWER:
		return "POWER"
	case PowerAssign:
		return "PowerAssign"
	case SHRUnsigned:
		return "SHRUnsigned"
	case ShrUnsignedAssign:
		return "ShrUnsignedAssign"

	// Template literal tokens
	case TemplateHead:
		return "TemplateHead"
	case TemplateMiddle:
		return "TemplateMiddle"
	case TemplateTail:
		return "TemplateTail"
	case TemplateNoSub:
		return "TemplateNoSub"

	default:
		return "INVALID_TOKEN"
	}
}
