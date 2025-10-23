package typescriptestree

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/lexer"
)

// AST_NODE_TYPES provides the string values for every single AST node's type property.
// This is equivalent to the AST_NODE_TYPES enum in @typescript-eslint/typescript-estree.
//
// Example usage:
//
//	if node.Type() == typescriptestree.AST_NODE_TYPES.Identifier {
//		// Handle identifier node
//	}
var AST_NODE_TYPES = struct {
	// ==================== Program & Core ====================
	Program string

	// ==================== Identifiers & Literals ====================
	Identifier        string
	PrivateIdentifier string
	Literal           string

	// ==================== Expressions ====================
	ThisExpression             string
	Super                      string
	ArrayExpression            string
	ObjectExpression           string
	Property                   string
	FunctionExpression         string
	ArrowFunctionExpression    string
	ClassExpression            string
	UnaryExpression            string
	UpdateExpression           string
	BinaryExpression           string
	LogicalExpression          string
	AssignmentExpression       string
	ConditionalExpression      string
	SequenceExpression         string
	MemberExpression           string
	CallExpression             string
	NewExpression              string
	MetaProperty               string
	TemplateLiteral            string
	TaggedTemplateExpression   string
	TemplateElement            string
	YieldExpression            string
	AwaitExpression            string
	ChainExpression            string
	ImportExpression           string
	SpreadElement              string

	// ==================== Statements ====================
	BlockStatement       string
	ExpressionStatement  string
	EmptyStatement       string
	DebuggerStatement    string
	ReturnStatement      string
	BreakStatement       string
	ContinueStatement    string
	LabeledStatement     string
	IfStatement          string
	SwitchStatement      string
	SwitchCase           string
	WhileStatement       string
	DoWhileStatement     string
	ForStatement         string
	ForInStatement       string
	ForOfStatement       string
	ThrowStatement       string
	TryStatement         string
	CatchClause          string
	WithStatement        string

	// ==================== Declarations ====================
	VariableDeclaration      string
	VariableDeclarator       string
	FunctionDeclaration      string
	ClassDeclaration         string
	ClassBody                string
	MethodDefinition         string
	PropertyDefinition       string
	AccessorProperty         string
	StaticBlock              string
	ImportDeclaration        string
	ImportSpecifier          string
	ImportDefaultSpecifier   string
	ImportNamespaceSpecifier string
	ImportAttribute          string
	ExportNamedDeclaration   string
	ExportDefaultDeclaration string
	ExportAllDeclaration     string
	ExportSpecifier          string

	// ==================== Patterns (Destructuring) ====================
	ArrayPattern      string
	ObjectPattern     string
	RestElement       string
	AssignmentPattern string

	// ==================== JSX (React) ====================
	JSXElement              string
	JSXFragment             string
	JSXOpeningElement       string
	JSXClosingElement       string
	JSXOpeningFragment      string
	JSXClosingFragment      string
	JSXAttribute            string
	JSXSpreadAttribute      string
	JSXIdentifier           string
	JSXNamespacedName       string
	JSXMemberExpression     string
	JSXExpressionContainer  string
	JSXEmptyExpression      string
	JSXText                 string
	JSXSpreadChild          string

	// ==================== Decorators ====================
	Decorator string

	// ==================== TypeScript Type Keywords ====================
	TSAnyKeyword       string
	TSBigIntKeyword    string
	TSBooleanKeyword   string
	TSIntrinsicKeyword string
	TSNeverKeyword     string
	TSNullKeyword      string
	TSNumberKeyword    string
	TSObjectKeyword    string
	TSStringKeyword    string
	TSSymbolKeyword    string
	TSUndefinedKeyword string
	TSUnknownKeyword   string
	TSVoidKeyword      string

	// ==================== TypeScript Type Expressions ====================
	TSArrayType             string
	TSTupleType             string
	TSUnionType             string
	TSIntersectionType      string
	TSConditionalType       string
	TSInferType             string
	TSTypeReference         string
	TSTypeQuery             string
	TSTypeLiteral           string
	TSFunctionType          string
	TSConstructorType       string
	TSMappedType            string
	TSLiteralType           string
	TSIndexedAccessType     string
	TSOptionalType          string
	TSRestType              string
	TSThisType              string
	TSTypeOperator          string
	TSTemplateLiteralType   string

	// ==================== TypeScript Type Declarations ====================
	TSTypeAnnotation        string
	TSTypeAliasDeclaration  string
	TSInterfaceDeclaration  string
	TSInterfaceBody         string
	TSInterfaceHeritage     string
	TSEnumDeclaration       string
	TSEnumBody              string
	TSEnumMember            string
	TSModuleDeclaration     string
	TSModuleBlock           string

	// ==================== TypeScript Type Components ====================
	TSTypeParameter                 string
	TSTypeParameterDeclaration      string
	TSTypeParameterInstantiation    string
	TSCallSignatureDeclaration      string
	TSConstructSignatureDeclaration string
	TSPropertySignature             string
	TSMethodSignature               string
	TSIndexSignature                string
	TSNamedTupleMember              string

	// ==================== TypeScript Type Assertions & Expressions ====================
	TSAsExpression            string
	TSTypeAssertion           string
	TSNonNullExpression       string
	TSSatisfiesExpression     string
	TSInstantiationExpression string

	// ==================== TypeScript Type Predicates ====================
	TSTypePredicate string

	// ==================== TypeScript Modifier Keywords ====================
	TSAbstractKeyword  string
	TSAsyncKeyword     string
	TSDeclareKeyword   string
	TSExportKeyword    string
	TSPrivateKeyword   string
	TSProtectedKeyword string
	TSPublicKeyword    string
	TSReadonlyKeyword  string
	TSStaticKeyword    string

	// ==================== TypeScript Abstract Members ====================
	TSAbstractAccessorProperty   string
	TSAbstractMethodDefinition   string
	TSAbstractPropertyDefinition string

	// ==================== TypeScript Import/Export ====================
	TSImportEqualsDeclaration    string
	TSImportType                 string
	TSExternalModuleReference    string
	TSExportAssignment           string
	TSNamespaceExportDeclaration string

	// ==================== TypeScript Other ====================
	TSQualifiedName                string
	TSParameterProperty            string
	TSDeclareFunction              string
	TSEmptyBodyFunctionExpression  string
	TSClassImplements              string
}{
	// Initialize all node type strings
	Program:                         ast.NodeTypeProgram.String(),
	Identifier:                      ast.NodeTypeIdentifier.String(),
	PrivateIdentifier:               ast.NodeTypePrivateIdentifier.String(),
	Literal:                         ast.NodeTypeLiteral.String(),
	ThisExpression:                  ast.NodeTypeThisExpression.String(),
	Super:                           ast.NodeTypeSuper.String(),
	ArrayExpression:                 ast.NodeTypeArrayExpression.String(),
	ObjectExpression:                ast.NodeTypeObjectExpression.String(),
	Property:                        ast.NodeTypeProperty.String(),
	FunctionExpression:              ast.NodeTypeFunctionExpression.String(),
	ArrowFunctionExpression:         ast.NodeTypeArrowFunctionExpression.String(),
	ClassExpression:                 ast.NodeTypeClassExpression.String(),
	UnaryExpression:                 ast.NodeTypeUnaryExpression.String(),
	UpdateExpression:                ast.NodeTypeUpdateExpression.String(),
	BinaryExpression:                ast.NodeTypeBinaryExpression.String(),
	LogicalExpression:               ast.NodeTypeLogicalExpression.String(),
	AssignmentExpression:            ast.NodeTypeAssignmentExpression.String(),
	ConditionalExpression:           ast.NodeTypeConditionalExpression.String(),
	SequenceExpression:              ast.NodeTypeSequenceExpression.String(),
	MemberExpression:                ast.NodeTypeMemberExpression.String(),
	CallExpression:                  ast.NodeTypeCallExpression.String(),
	NewExpression:                   ast.NodeTypeNewExpression.String(),
	MetaProperty:                    ast.NodeTypeMetaProperty.String(),
	TemplateLiteral:                 ast.NodeTypeTemplateLiteral.String(),
	TaggedTemplateExpression:        ast.NodeTypeTaggedTemplateExpression.String(),
	TemplateElement:                 ast.NodeTypeTemplateElement.String(),
	YieldExpression:                 ast.NodeTypeYieldExpression.String(),
	AwaitExpression:                 ast.NodeTypeAwaitExpression.String(),
	ChainExpression:                 ast.NodeTypeChainExpression.String(),
	ImportExpression:                ast.NodeTypeImportExpression.String(),
	SpreadElement:                   ast.NodeTypeSpreadElement.String(),
	BlockStatement:                  ast.NodeTypeBlockStatement.String(),
	ExpressionStatement:             ast.NodeTypeExpressionStatement.String(),
	EmptyStatement:                  ast.NodeTypeEmptyStatement.String(),
	DebuggerStatement:               ast.NodeTypeDebuggerStatement.String(),
	ReturnStatement:                 ast.NodeTypeReturnStatement.String(),
	BreakStatement:                  ast.NodeTypeBreakStatement.String(),
	ContinueStatement:               ast.NodeTypeContinueStatement.String(),
	LabeledStatement:                ast.NodeTypeLabeledStatement.String(),
	IfStatement:                     ast.NodeTypeIfStatement.String(),
	SwitchStatement:                 ast.NodeTypeSwitchStatement.String(),
	SwitchCase:                      ast.NodeTypeSwitchCase.String(),
	WhileStatement:                  ast.NodeTypeWhileStatement.String(),
	DoWhileStatement:                ast.NodeTypeDoWhileStatement.String(),
	ForStatement:                    ast.NodeTypeForStatement.String(),
	ForInStatement:                  ast.NodeTypeForInStatement.String(),
	ForOfStatement:                  ast.NodeTypeForOfStatement.String(),
	ThrowStatement:                  ast.NodeTypeThrowStatement.String(),
	TryStatement:                    ast.NodeTypeTryStatement.String(),
	CatchClause:                     ast.NodeTypeCatchClause.String(),
	WithStatement:                   ast.NodeTypeWithStatement.String(),
	VariableDeclaration:             ast.NodeTypeVariableDeclaration.String(),
	VariableDeclarator:              ast.NodeTypeVariableDeclarator.String(),
	FunctionDeclaration:             ast.NodeTypeFunctionDeclaration.String(),
	ClassDeclaration:                ast.NodeTypeClassDeclaration.String(),
	ClassBody:                       ast.NodeTypeClassBody.String(),
	MethodDefinition:                ast.NodeTypeMethodDefinition.String(),
	PropertyDefinition:              ast.NodeTypePropertyDefinition.String(),
	AccessorProperty:                ast.NodeTypeAccessorProperty.String(),
	StaticBlock:                     ast.NodeTypeStaticBlock.String(),
	ImportDeclaration:               ast.NodeTypeImportDeclaration.String(),
	ImportSpecifier:                 ast.NodeTypeImportSpecifier.String(),
	ImportDefaultSpecifier:          ast.NodeTypeImportDefaultSpecifier.String(),
	ImportNamespaceSpecifier:        ast.NodeTypeImportNamespaceSpecifier.String(),
	ImportAttribute:                 ast.NodeTypeImportAttribute.String(),
	ExportNamedDeclaration:          ast.NodeTypeExportNamedDeclaration.String(),
	ExportDefaultDeclaration:        ast.NodeTypeExportDefaultDeclaration.String(),
	ExportAllDeclaration:            ast.NodeTypeExportAllDeclaration.String(),
	ExportSpecifier:                 ast.NodeTypeExportSpecifier.String(),
	ArrayPattern:                    ast.NodeTypeArrayPattern.String(),
	ObjectPattern:                   ast.NodeTypeObjectPattern.String(),
	RestElement:                     ast.NodeTypeRestElement.String(),
	AssignmentPattern:               ast.NodeTypeAssignmentPattern.String(),
	JSXElement:                      ast.NodeTypeJSXElement.String(),
	JSXFragment:                     ast.NodeTypeJSXFragment.String(),
	JSXOpeningElement:               ast.NodeTypeJSXOpeningElement.String(),
	JSXClosingElement:               ast.NodeTypeJSXClosingElement.String(),
	JSXOpeningFragment:              ast.NodeTypeJSXOpeningFragment.String(),
	JSXClosingFragment:              ast.NodeTypeJSXClosingFragment.String(),
	JSXAttribute:                    ast.NodeTypeJSXAttribute.String(),
	JSXSpreadAttribute:              ast.NodeTypeJSXSpreadAttribute.String(),
	JSXIdentifier:                   ast.NodeTypeJSXIdentifier.String(),
	JSXNamespacedName:               ast.NodeTypeJSXNamespacedName.String(),
	JSXMemberExpression:             ast.NodeTypeJSXMemberExpression.String(),
	JSXExpressionContainer:          ast.NodeTypeJSXExpressionContainer.String(),
	JSXEmptyExpression:              ast.NodeTypeJSXEmptyExpression.String(),
	JSXText:                         ast.NodeTypeJSXText.String(),
	JSXSpreadChild:                  ast.NodeTypeJSXSpreadChild.String(),
	Decorator:                       ast.NodeTypeDecorator.String(),
	TSAnyKeyword:                    ast.NodeTypeTSAnyKeyword.String(),
	TSBigIntKeyword:                 ast.NodeTypeTSBigIntKeyword.String(),
	TSBooleanKeyword:                ast.NodeTypeTSBooleanKeyword.String(),
	TSIntrinsicKeyword:              ast.NodeTypeTSIntrinsicKeyword.String(),
	TSNeverKeyword:                  ast.NodeTypeTSNeverKeyword.String(),
	TSNullKeyword:                   ast.NodeTypeTSNullKeyword.String(),
	TSNumberKeyword:                 ast.NodeTypeTSNumberKeyword.String(),
	TSObjectKeyword:                 ast.NodeTypeTSObjectKeyword.String(),
	TSStringKeyword:                 ast.NodeTypeTSStringKeyword.String(),
	TSSymbolKeyword:                 ast.NodeTypeTSSymbolKeyword.String(),
	TSUndefinedKeyword:              ast.NodeTypeTSUndefinedKeyword.String(),
	TSUnknownKeyword:                ast.NodeTypeTSUnknownKeyword.String(),
	TSVoidKeyword:                   ast.NodeTypeTSVoidKeyword.String(),
	TSArrayType:                     ast.NodeTypeTSArrayType.String(),
	TSTupleType:                     ast.NodeTypeTSTupleType.String(),
	TSUnionType:                     ast.NodeTypeTSUnionType.String(),
	TSIntersectionType:              ast.NodeTypeTSIntersectionType.String(),
	TSConditionalType:               ast.NodeTypeTSConditionalType.String(),
	TSInferType:                     ast.NodeTypeTSInferType.String(),
	TSTypeReference:                 ast.NodeTypeTSTypeReference.String(),
	TSTypeQuery:                     ast.NodeTypeTSTypeQuery.String(),
	TSTypeLiteral:                   ast.NodeTypeTSTypeLiteral.String(),
	TSFunctionType:                  ast.NodeTypeTSFunctionType.String(),
	TSConstructorType:               ast.NodeTypeTSConstructorType.String(),
	TSMappedType:                    ast.NodeTypeTSMappedType.String(),
	TSLiteralType:                   ast.NodeTypeTSLiteralType.String(),
	TSIndexedAccessType:             ast.NodeTypeTSIndexedAccessType.String(),
	TSOptionalType:                  ast.NodeTypeTSOptionalType.String(),
	TSRestType:                      ast.NodeTypeTSRestType.String(),
	TSThisType:                      ast.NodeTypeTSThisType.String(),
	TSTypeOperator:                  ast.NodeTypeTSTypeOperator.String(),
	TSTemplateLiteralType:           ast.NodeTypeTSTemplateLiteralType.String(),
	TSTypeAnnotation:                ast.NodeTypeTSTypeAnnotation.String(),
	TSTypeAliasDeclaration:          ast.NodeTypeTSTypeAliasDeclaration.String(),
	TSInterfaceDeclaration:          ast.NodeTypeTSInterfaceDeclaration.String(),
	TSInterfaceBody:                 ast.NodeTypeTSInterfaceBody.String(),
	TSInterfaceHeritage:             ast.NodeTypeTSInterfaceHeritage.String(),
	TSEnumDeclaration:               ast.NodeTypeTSEnumDeclaration.String(),
	TSEnumBody:                      ast.NodeTypeTSEnumBody.String(),
	TSEnumMember:                    ast.NodeTypeTSEnumMember.String(),
	TSModuleDeclaration:             ast.NodeTypeTSModuleDeclaration.String(),
	TSModuleBlock:                   ast.NodeTypeTSModuleBlock.String(),
	TSTypeParameter:                 ast.NodeTypeTSTypeParameter.String(),
	TSTypeParameterDeclaration:      ast.NodeTypeTSTypeParameterDeclaration.String(),
	TSTypeParameterInstantiation:    ast.NodeTypeTSTypeParameterInstantiation.String(),
	TSCallSignatureDeclaration:      ast.NodeTypeTSCallSignatureDeclaration.String(),
	TSConstructSignatureDeclaration: ast.NodeTypeTSConstructSignatureDeclaration.String(),
	TSPropertySignature:             ast.NodeTypeTSPropertySignature.String(),
	TSMethodSignature:               ast.NodeTypeTSMethodSignature.String(),
	TSIndexSignature:                ast.NodeTypeTSIndexSignature.String(),
	TSNamedTupleMember:              ast.NodeTypeTSNamedTupleMember.String(),
	TSAsExpression:                  ast.NodeTypeTSAsExpression.String(),
	TSTypeAssertion:                 ast.NodeTypeTSTypeAssertion.String(),
	TSNonNullExpression:             ast.NodeTypeTSNonNullExpression.String(),
	TSSatisfiesExpression:           ast.NodeTypeTSSatisfiesExpression.String(),
	TSInstantiationExpression:       ast.NodeTypeTSInstantiationExpression.String(),
	TSTypePredicate:                 ast.NodeTypeTSTypePredicate.String(),
	TSAbstractKeyword:               ast.NodeTypeTSAbstractKeyword.String(),
	TSAsyncKeyword:                  ast.NodeTypeTSAsyncKeyword.String(),
	TSDeclareKeyword:                ast.NodeTypeTSDeclareKeyword.String(),
	TSExportKeyword:                 ast.NodeTypeTSExportKeyword.String(),
	TSPrivateKeyword:                ast.NodeTypeTSPrivateKeyword.String(),
	TSProtectedKeyword:              ast.NodeTypeTSProtectedKeyword.String(),
	TSPublicKeyword:                 ast.NodeTypeTSPublicKeyword.String(),
	TSReadonlyKeyword:               ast.NodeTypeTSReadonlyKeyword.String(),
	TSStaticKeyword:                 ast.NodeTypeTSStaticKeyword.String(),
	TSAbstractAccessorProperty:      ast.NodeTypeTSAbstractAccessorProperty.String(),
	TSAbstractMethodDefinition:      ast.NodeTypeTSAbstractMethodDefinition.String(),
	TSAbstractPropertyDefinition:    ast.NodeTypeTSAbstractPropertyDefinition.String(),
	TSImportEqualsDeclaration:       ast.NodeTypeTSImportEqualsDeclaration.String(),
	TSImportType:                    ast.NodeTypeTSImportType.String(),
	TSExternalModuleReference:       ast.NodeTypeTSExternalModuleReference.String(),
	TSExportAssignment:              ast.NodeTypeTSExportAssignment.String(),
	TSNamespaceExportDeclaration:    ast.NodeTypeTSNamespaceExportDeclaration.String(),
	TSQualifiedName:                 ast.NodeTypeTSQualifiedName.String(),
	TSParameterProperty:             ast.NodeTypeTSParameterProperty.String(),
	TSDeclareFunction:               ast.NodeTypeTSDeclareFunction.String(),
	TSEmptyBodyFunctionExpression:   ast.NodeTypeTSEmptyBodyFunctionExpression.String(),
	TSClassImplements:               ast.NodeTypeTSClassImplements.String(),
}

// AST_TOKEN_TYPES provides the string values for every single AST token's type property.
// This is equivalent to the AST_TOKEN_TYPES enum in @typescript-eslint/typescript-estree.
//
// Example usage:
//
//	if token.Type == typescriptestree.AST_TOKEN_TYPES.Identifier {
//		// Handle identifier token
//	}
var AST_TOKEN_TYPES = struct {
	// Special tokens
	EOF     string
	Illegal string
	Comment string

	// Literals
	Identifier string
	Number     string
	String     string
	Template   string
	RegExp     string

	// Keywords
	Break      string
	Case       string
	Catch      string
	Class      string
	Const      string
	Continue   string
	Debugger   string
	Default    string
	Delete     string
	Do         string
	Else       string
	Enum       string
	Export     string
	Extends    string
	False      string
	Finally    string
	For        string
	Function   string
	If         string
	Import     string
	In         string
	Instanceof string
	New        string
	Null       string
	Return     string
	Super      string
	Switch     string
	This       string
	Throw      string
	True       string
	Try        string
	Typeof     string
	Var        string
	Void       string
	While      string
	With       string
	Yield      string

	// TypeScript keywords
	As         string
	Async      string
	Await      string
	Declare    string
	Interface  string
	Let        string
	Module     string
	Namespace  string
	Of         string
	Package    string
	Private    string
	Protected  string
	Public     string
	Readonly   string
	Require    string
	Static     string
	Type       string
	From       string
	Satisfies  string
	Implements string
	Any        string
	Boolean    string
	Never      string
	Unknown    string
	Symbol     string
	Undefined  string

	// Operators and punctuation
	Add           string // +
	Sub           string // -
	Mul           string // *
	Quo           string // /
	Rem           string // %
	And           string // &
	Or            string // |
	Xor           string // ^
	BitwiseNot    string // ~
	ShiftLeft     string // <<
	ShiftRight    string // >>
	AddAssign     string // +=
	SubAssign     string // -=
	MulAssign     string // *=
	QuoAssign     string // /=
	RemAssign     string // %=
	AndAssign     string // &=
	OrAssign      string // |=
	XorAssign     string // ^=
	ShlAssign     string // <<=
	ShrAssign     string // >>=
	LogicalAnd    string // &&
	LogicalOr     string // ||
	Increment     string // ++
	Decrement     string // --
	Nullish       string // ??
	Equal         string // ==
	Less          string // <
	Greater       string // >
	Assign        string // =
	Not           string // !
	NotEqual      string // !=
	LessEqual     string // <=
	GreaterEqual  string // >=
	StrictEqual   string // ===
	StrictNotEqual string // !==
	LeftParen     string // (
	LeftBracket   string // [
	LeftBrace     string // {
	Comma         string // ,
	Period        string // .
	RightParen    string // )
	RightBracket  string // ]
	RightBrace    string // }
	Semicolon     string // ;
	Colon         string // :
	Question      string // ?
	Arrow         string // =>
	Ellipsis      string // ...
}{
	// Initialize all token type strings
	EOF:            lexer.EOF.String(),
	Illegal:        lexer.ILLEGAL.String(),
	Comment:        lexer.COMMENT.String(),
	Identifier:     lexer.IDENT.String(),
	Number:         lexer.NUMBER.String(),
	String:         lexer.STRING.String(),
	Template:       lexer.TEMPLATE.String(),
	RegExp:         lexer.REGEXP.String(),
	Break:          lexer.BREAK.String(),
	Case:           lexer.CASE.String(),
	Catch:          lexer.CATCH.String(),
	Class:          lexer.CLASS.String(),
	Const:          lexer.CONST.String(),
	Continue:       lexer.CONTINUE.String(),
	Debugger:       lexer.DEBUGGER.String(),
	Default:        lexer.DEFAULT.String(),
	Delete:         lexer.DELETE.String(),
	Do:             lexer.DO.String(),
	Else:           lexer.ELSE.String(),
	Enum:           lexer.ENUM.String(),
	Export:         lexer.EXPORT.String(),
	Extends:        lexer.EXTENDS.String(),
	False:          lexer.FALSE.String(),
	Finally:        lexer.FINALLY.String(),
	For:            lexer.FOR.String(),
	Function:       lexer.FUNCTION.String(),
	If:             lexer.IF.String(),
	Import:         lexer.IMPORT.String(),
	In:             lexer.IN.String(),
	Instanceof:     lexer.INSTANCEOF.String(),
	New:            lexer.NEW.String(),
	Null:           lexer.NULL.String(),
	Return:         lexer.RETURN.String(),
	Super:          lexer.SUPER.String(),
	Switch:         lexer.SWITCH.String(),
	This:           lexer.THIS.String(),
	Throw:          lexer.THROW.String(),
	True:           lexer.TRUE.String(),
	Try:            lexer.TRY.String(),
	Typeof:         lexer.TYPEOF.String(),
	Var:            lexer.VAR.String(),
	Void:           lexer.VOID.String(),
	While:          lexer.WHILE.String(),
	With:           lexer.WITH.String(),
	Yield:          lexer.YIELD.String(),
	As:             lexer.AS.String(),
	Async:          lexer.ASYNC.String(),
	Await:          lexer.AWAIT.String(),
	Declare:        lexer.DECLARE.String(),
	Interface:      lexer.INTERFACE.String(),
	Let:            lexer.LET.String(),
	Module:         lexer.MODULE.String(),
	Namespace:      lexer.NAMESPACE.String(),
	Of:             lexer.OF.String(),
	Package:        lexer.PACKAGE.String(),
	Private:        lexer.PRIVATE.String(),
	Protected:      lexer.PROTECTED.String(),
	Public:         lexer.PUBLIC.String(),
	Readonly:       lexer.READONLY.String(),
	Require:        lexer.REQUIRE.String(),
	Static:         lexer.STATIC.String(),
	Type:           lexer.TYPE.String(),
	From:           lexer.FROM.String(),
	Satisfies:      lexer.SATISFIES.String(),
	Implements:     lexer.IMPLEMENTS.String(),
	Any:            lexer.ANY.String(),
	Boolean:        lexer.BOOLEAN.String(),
	Never:          lexer.NEVER.String(),
	Unknown:        lexer.UNKNOWN.String(),
	Symbol:         lexer.SYMBOL.String(),
	Undefined:      lexer.UNDEFINED.String(),
	Add:            lexer.ADD.String(),
	Sub:            lexer.SUB.String(),
	Mul:            lexer.MUL.String(),
	Quo:            lexer.QUO.String(),
	Rem:            lexer.REM.String(),
	And:            lexer.AND.String(),
	Or:             lexer.OR.String(),
	Xor:            lexer.XOR.String(),
	BitwiseNot:     lexer.BNOT.String(),
	ShiftLeft:      lexer.SHL.String(),
	ShiftRight:     lexer.SHR.String(),
	AddAssign:      lexer.AddAssign.String(),
	SubAssign:      lexer.SubAssign.String(),
	MulAssign:      lexer.MulAssign.String(),
	QuoAssign:      lexer.QuoAssign.String(),
	RemAssign:      lexer.RemAssign.String(),
	AndAssign:      lexer.AndAssign.String(),
	OrAssign:       lexer.OrAssign.String(),
	XorAssign:      lexer.XorAssign.String(),
	ShlAssign:      lexer.ShlAssign.String(),
	ShrAssign:      lexer.ShrAssign.String(),
	LogicalAnd:     lexer.LAND.String(),
	LogicalOr:      lexer.LOR.String(),
	Increment:      lexer.INC.String(),
	Decrement:      lexer.DEC.String(),
	Nullish:        lexer.NULLISH.String(),
	Equal:          lexer.EQL.String(),
	Less:           lexer.LSS.String(),
	Greater:        lexer.GTR.String(),
	Assign:         lexer.ASSIGN.String(),
	Not:            lexer.NOT.String(),
	NotEqual:       lexer.NEQ.String(),
	LessEqual:      lexer.LEQ.String(),
	GreaterEqual:   lexer.GEQ.String(),
	StrictEqual:    lexer.EqlStrict.String(),
	StrictNotEqual: lexer.NeqStrict.String(),
	LeftParen:      lexer.LPAREN.String(),
	LeftBracket:    lexer.LBRACK.String(),
	LeftBrace:      lexer.LBRACE.String(),
	Comma:          lexer.COMMA.String(),
	Period:         lexer.PERIOD.String(),
	RightParen:     lexer.RPAREN.String(),
	RightBracket:   lexer.RBRACK.String(),
	RightBrace:     lexer.RBRACE.String(),
	Semicolon:      lexer.SEMICOLON.String(),
	Colon:          lexer.COLON.String(),
	Question:       lexer.QUESTION.String(),
	Arrow:          lexer.ARROW.String(),
	Ellipsis:       lexer.ELLIPSIS.String(),
}
