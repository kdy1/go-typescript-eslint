package ast

// This file defines all ESTree and TypeScript AST node types.
// Based on: https://github.com/estree/estree and https://typescript-eslint.io/packages/typescript-estree/ast-spec/

// ==================== Core Types ====================

// Range represents the character range of a node in the source [start, end).
type Range [2]int

// TSNode represents a TypeScript-specific node type.
// This is an extension interface for TypeScript-only nodes.
type TSNode interface {
	Node
	// TSTypeNode marks this as a TypeScript type node.
	TSTypeNode()
}

// ==================== Program ====================

// Program represents the root node of an AST.
// It contains the entire program source.
//
//nolint:govet // Field order optimized for JSON output readability
type Program struct {
	BaseNode
	SourceType   string      `json:"sourceType"`   // "script" or "module"
	Body         []Statement `json:"body"`         // Top-level statements
	Comments     []Comment   `json:"comments,omitempty"`
	Tokens       []Token     `json:"tokens,omitempty"`
	Decorators   []Decorator `json:"decorators,omitempty"`   // Decorators (experimental)
	TSConfigPath *string     `json:"tsConfigPath,omitempty"` // TypeScript config path (TS extension)
}

// ==================== Identifiers ====================

// Identifier represents an identifier (variable name, function name, etc.).
//
//nolint:govet // Field order optimized for JSON output readability
type Identifier struct {
	BaseNode
	Name           string            `json:"name"`
	Decorators     []Decorator       `json:"decorators,omitempty"`
	Optional       bool              `json:"optional,omitempty"`
	TypeAnnotation *TSTypeAnnotation `json:"typeAnnotation,omitempty"`
}

func (n *Identifier) expressionNode() {}
func (n *Identifier) patternNode()    {}

// PrivateIdentifier represents a private identifier (#field).
//
//nolint:govet // Field order optimized for JSON output readability
type PrivateIdentifier struct {
	BaseNode
	Name string `json:"name"` // Without the # prefix
}

func (n *PrivateIdentifier) expressionNode() {}

// ==================== Literals ====================

// Literal represents a literal value.
//
//nolint:govet // Field order optimized for JSON output readability
type Literal struct {
	BaseNode
	Value interface{} `json:"value"` // Actual value (string, number, boolean, null)
	Raw   string      `json:"raw"`   // Original source text
	Regex *RegexInfo  `json:"regex,omitempty"`
	BigInt *string    `json:"bigint,omitempty"` // BigInt as string
}

func (n *Literal) expressionNode() {}

// RegexInfo contains information about a regular expression literal.
type RegexInfo struct {
	Pattern string `json:"pattern"`
	Flags   string `json:"flags"`
}

// ==================== Expressions ====================

// ThisExpression represents the 'this' keyword.
type ThisExpression struct {
	BaseNode
}

func (n *ThisExpression) expressionNode() {}

// Super represents the 'super' keyword.
type Super struct {
	BaseNode
}

func (n *Super) expressionNode() {}

// ArrayExpression represents an array literal [1, 2, 3].
//
//nolint:govet // Field order optimized for JSON output readability
type ArrayExpression struct {
	BaseNode
	Elements []Expression `json:"elements"` // Can include nil for holes in sparse arrays
}

func (n *ArrayExpression) expressionNode() {}

// ObjectExpression represents an object literal {a: 1, b: 2}.
//
//nolint:govet // Field order optimized for JSON output readability
type ObjectExpression struct {
	BaseNode
	Properties []interface{} `json:"properties"` // Property | SpreadElement
}

func (n *ObjectExpression) expressionNode() {}

// Property represents a property in an object expression or pattern.
//
//nolint:govet // Field order optimized for JSON output readability
type Property struct {
	BaseNode
	Key        Expression `json:"key"`
	Value      Expression `json:"value"`
	Kind       string     `json:"kind"` // "init" | "get" | "set"
	Method     bool       `json:"method"`
	Shorthand  bool       `json:"shorthand"`
	Computed   bool       `json:"computed"`
	Decorators []Decorator `json:"decorators,omitempty"`
	Optional   bool       `json:"optional,omitempty"`
	TypeAnnotation *TSTypeAnnotation `json:"typeAnnotation,omitempty"`
}

// FunctionExpression represents a function expression.
//
//nolint:govet // Field order optimized for JSON output readability
type FunctionExpression struct {
	BaseNode
	ID              *Identifier         `json:"id"`
	Params          []Pattern           `json:"params"`
	Body            *BlockStatement     `json:"body"`
	Generator       bool                `json:"generator"`
	Async           bool                `json:"async"`
	Expression      bool                `json:"expression"`
	TypeParameters  *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
	ReturnType      *TSTypeAnnotation   `json:"returnType,omitempty"`
	Decorators      []Decorator         `json:"decorators,omitempty"`
	Modifiers       []interface{}       `json:"modifiers,omitempty"`
}

func (n *FunctionExpression) expressionNode() {}

// ArrowFunctionExpression represents an arrow function expression.
//
//nolint:govet // Field order optimized for JSON output readability
type ArrowFunctionExpression struct {
	BaseNode
	Params         []Pattern                   `json:"params"`
	Body           interface{}                 `json:"body"` // BlockStatement | Expression
	Generator      bool                        `json:"generator"`
	Async          bool                        `json:"async"`
	Expression     bool                        `json:"expression"`
	TypeParameters *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
	ReturnType     *TSTypeAnnotation           `json:"returnType,omitempty"`
}

func (n *ArrowFunctionExpression) expressionNode() {}

// ClassExpression represents a class expression.
//
//nolint:govet // Field order optimized for JSON output readability
type ClassExpression struct {
	BaseNode
	ID             *Identifier                 `json:"id"`
	SuperClass     Expression                  `json:"superClass"`
	Body           *ClassBody                  `json:"body"`
	Decorators     []Decorator                 `json:"decorators,omitempty"`
	TypeParameters *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
	SuperTypeParameters *TSTypeParameterInstantiation `json:"superTypeParameters,omitempty"`
	Implements     []TSClassImplements         `json:"implements,omitempty"`
	Abstract       bool                        `json:"abstract,omitempty"`
	Declare        bool                        `json:"declare,omitempty"`
}

func (n *ClassExpression) expressionNode() {}

// UnaryExpression represents a unary operation (+x, -x, !x, ~x, typeof x, void x, delete x).
//
//nolint:govet // Field order optimized for JSON output readability
type UnaryExpression struct {
	BaseNode
	Operator string     `json:"operator"` // "+", "-", "!", "~", "typeof", "void", "delete"
	Prefix   bool       `json:"prefix"`
	Argument Expression `json:"argument"`
}

func (n *UnaryExpression) expressionNode() {}

// UpdateExpression represents an update expression (++x, x++, --x, x--).
//
//nolint:govet // Field order optimized for JSON output readability
type UpdateExpression struct {
	BaseNode
	Operator string     `json:"operator"` // "++" | "--"
	Argument Expression `json:"argument"`
	Prefix   bool       `json:"prefix"`
}

func (n *UpdateExpression) expressionNode() {}

// BinaryExpression represents a binary operation (x + y, x - y, x * y, etc.).
//
//nolint:govet // Field order optimized for JSON output readability
type BinaryExpression struct {
	BaseNode
	Operator string     `json:"operator"` // "+", "-", "*", "/", "%", "**", etc.
	Left     Expression `json:"left"`
	Right    Expression `json:"right"`
}

func (n *BinaryExpression) expressionNode() {}

// LogicalExpression represents a logical operation (x && y, x || y, x ?? y).
//
//nolint:govet // Field order optimized for JSON output readability
type LogicalExpression struct {
	BaseNode
	Operator string     `json:"operator"` // "&&" | "||" | "??"
	Left     Expression `json:"left"`
	Right    Expression `json:"right"`
}

func (n *LogicalExpression) expressionNode() {}

// AssignmentExpression represents an assignment (x = y, x += y, etc.).
//
//nolint:govet // Field order optimized for JSON output readability
type AssignmentExpression struct {
	BaseNode
	Operator string     `json:"operator"` // "=", "+=", "-=", etc.
	Left     Pattern    `json:"left"`     // Can be Pattern for destructuring
	Right    Expression `json:"right"`
}

func (n *AssignmentExpression) expressionNode() {}

// ConditionalExpression represents a ternary conditional (x ? y : z).
//
//nolint:govet // Field order optimized for JSON output readability
type ConditionalExpression struct {
	BaseNode
	Test       Expression `json:"test"`
	Consequent Expression `json:"consequent"`
	Alternate  Expression `json:"alternate"`
}

func (n *ConditionalExpression) expressionNode() {}

// SequenceExpression represents a sequence of expressions (x, y, z).
//
//nolint:govet // Field order optimized for JSON output readability
type SequenceExpression struct {
	BaseNode
	Expressions []Expression `json:"expressions"`
}

func (n *SequenceExpression) expressionNode() {}

// MemberExpression represents a member access (obj.prop, obj[prop]).
//
//nolint:govet // Field order optimized for JSON output readability
type MemberExpression struct {
	BaseNode
	Object   Expression `json:"object"`
	Property Expression `json:"property"`
	Computed bool       `json:"computed"`
	Optional bool       `json:"optional,omitempty"`
}

func (n *MemberExpression) expressionNode() {}
func (n *MemberExpression) patternNode()    {}

// CallExpression represents a function call.
//
//nolint:govet // Field order optimized for JSON output readability
type CallExpression struct {
	BaseNode
	Callee         Expression                    `json:"callee"`
	Arguments      []Expression                  `json:"arguments"` // Can include SpreadElement
	Optional       bool                          `json:"optional,omitempty"`
	TypeArguments  *TSTypeParameterInstantiation `json:"typeArguments,omitempty"`
	TypeParameters *TSTypeParameterInstantiation `json:"typeParameters,omitempty"` // Deprecated, use TypeArguments
}

func (n *CallExpression) expressionNode() {}

// NewExpression represents a new expression (new Foo()).
//
//nolint:govet // Field order optimized for JSON output readability
type NewExpression struct {
	BaseNode
	Callee         Expression                    `json:"callee"`
	Arguments      []Expression                  `json:"arguments"` // Can include SpreadElement
	TypeArguments  *TSTypeParameterInstantiation `json:"typeArguments,omitempty"`
	TypeParameters *TSTypeParameterInstantiation `json:"typeParameters,omitempty"` // Deprecated
}

func (n *NewExpression) expressionNode() {}

// MetaProperty represents a meta property (new.target, import.meta).
//
//nolint:govet // Field order optimized for JSON output readability
type MetaProperty struct {
	BaseNode
	Meta     *Identifier `json:"meta"`
	Property *Identifier `json:"property"`
}

func (n *MetaProperty) expressionNode() {}

// YieldExpression represents a yield expression.
//
//nolint:govet // Field order optimized for JSON output readability
type YieldExpression struct {
	BaseNode
	Argument Expression `json:"argument"`
	Delegate bool       `json:"delegate"`
}

func (n *YieldExpression) expressionNode() {}

// AwaitExpression represents an await expression.
//
//nolint:govet // Field order optimized for JSON output readability
type AwaitExpression struct {
	BaseNode
	Argument Expression `json:"argument"`
}

func (n *AwaitExpression) expressionNode() {}

// ChainExpression represents an optional chaining expression (obj?.prop).
//
//nolint:govet // Field order optimized for JSON output readability
type ChainExpression struct {
	BaseNode
	Expression Expression `json:"expression"` // MemberExpression | CallExpression
}

func (n *ChainExpression) expressionNode() {}

// ImportExpression represents a dynamic import expression import().
//
//nolint:govet // Field order optimized for JSON output readability
type ImportExpression struct {
	BaseNode
	Source     Expression      `json:"source"`
	Attributes []ImportAttribute `json:"attributes,omitempty"`
}

func (n *ImportExpression) expressionNode() {}

// SpreadElement represents a spread element (...x).
//
//nolint:govet // Field order optimized for JSON output readability
type SpreadElement struct {
	BaseNode
	Argument Expression `json:"argument"`
}

// TemplateLiteral represents a template literal `hello ${world}`.
//
//nolint:govet // Field order optimized for JSON output readability
type TemplateLiteral struct {
	BaseNode
	Quasis      []TemplateElement `json:"quasis"`
	Expressions []Expression      `json:"expressions"`
}

func (n *TemplateLiteral) expressionNode() {}

// TaggedTemplateExpression represents a tagged template expression.
//
//nolint:govet // Field order optimized for JSON output readability
type TaggedTemplateExpression struct {
	BaseNode
	Tag           Expression                    `json:"tag"`
	Quasi         *TemplateLiteral              `json:"quasi"`
	TypeArguments *TSTypeParameterInstantiation `json:"typeArguments,omitempty"`
	TypeParameters *TSTypeParameterInstantiation `json:"typeParameters,omitempty"` // Deprecated
}

func (n *TaggedTemplateExpression) expressionNode() {}

// TemplateElement represents an element in a template literal.
//
//nolint:govet // Field order optimized for JSON output readability
type TemplateElement struct {
	BaseNode
	Tail  bool               `json:"tail"`
	Value TemplateElementValue `json:"value"`
}

// TemplateElementValue contains the cooked and raw values of a template element.
type TemplateElementValue struct {
	Cooked *string `json:"cooked"` // Processed value (nil if contains invalid escape)
	Raw    string  `json:"raw"`    // Original source text
}

// ==================== Statements ====================

// BlockStatement represents a block of statements {}.
//
//nolint:govet // Field order optimized for JSON output readability
type BlockStatement struct {
	BaseNode
	Body []Statement `json:"body"`
}

func (n *BlockStatement) statementNode() {}

// ExpressionStatement represents an expression used as a statement.
//
//nolint:govet // Field order optimized for JSON output readability
type ExpressionStatement struct {
	BaseNode
	Expression Expression `json:"expression"`
	Directive  *string    `json:"directive,omitempty"` // For "use strict" etc.
}

func (n *ExpressionStatement) statementNode() {}

// EmptyStatement represents an empty statement (;).
type EmptyStatement struct {
	BaseNode
}

func (n *EmptyStatement) statementNode() {}

// DebuggerStatement represents a debugger statement.
type DebuggerStatement struct {
	BaseNode
}

func (n *DebuggerStatement) statementNode() {}

// ReturnStatement represents a return statement.
//
//nolint:govet // Field order optimized for JSON output readability
type ReturnStatement struct {
	BaseNode
	Argument Expression `json:"argument"`
}

func (n *ReturnStatement) statementNode() {}

// BreakStatement represents a break statement.
//
//nolint:govet // Field order optimized for JSON output readability
type BreakStatement struct {
	BaseNode
	Label *Identifier `json:"label"`
}

func (n *BreakStatement) statementNode() {}

// ContinueStatement represents a continue statement.
//
//nolint:govet // Field order optimized for JSON output readability
type ContinueStatement struct {
	BaseNode
	Label *Identifier `json:"label"`
}

func (n *ContinueStatement) statementNode() {}

// LabeledStatement represents a labeled statement.
//
//nolint:govet // Field order optimized for JSON output readability
type LabeledStatement struct {
	BaseNode
	Label *Identifier `json:"label"`
	Body  Statement   `json:"body"`
}

func (n *LabeledStatement) statementNode() {}

// IfStatement represents an if statement.
//
//nolint:govet // Field order optimized for JSON output readability
type IfStatement struct {
	BaseNode
	Test       Expression `json:"test"`
	Consequent Statement  `json:"consequent"`
	Alternate  Statement  `json:"alternate"`
}

func (n *IfStatement) statementNode() {}

// SwitchStatement represents a switch statement.
//
//nolint:govet // Field order optimized for JSON output readability
type SwitchStatement struct {
	BaseNode
	Discriminant Expression    `json:"discriminant"`
	Cases        []SwitchCase  `json:"cases"`
}

func (n *SwitchStatement) statementNode() {}

// SwitchCase represents a case or default clause in a switch statement.
//
//nolint:govet // Field order optimized for JSON output readability
type SwitchCase struct {
	BaseNode
	Test       Expression  `json:"test"` // nil for default case
	Consequent []Statement `json:"consequent"`
}

// WhileStatement represents a while loop.
//
//nolint:govet // Field order optimized for JSON output readability
type WhileStatement struct {
	BaseNode
	Test Expression `json:"test"`
	Body Statement  `json:"body"`
}

func (n *WhileStatement) statementNode() {}

// DoWhileStatement represents a do-while loop.
//
//nolint:govet // Field order optimized for JSON output readability
type DoWhileStatement struct {
	BaseNode
	Body Statement  `json:"body"`
	Test Expression `json:"test"`
}

func (n *DoWhileStatement) statementNode() {}

// ForStatement represents a for loop.
//
//nolint:govet // Field order optimized for JSON output readability
type ForStatement struct {
	BaseNode
	Init   interface{} `json:"init"` // VariableDeclaration | Expression | nil
	Test   Expression  `json:"test"`
	Update Expression  `json:"update"`
	Body   Statement   `json:"body"`
}

func (n *ForStatement) statementNode() {}

// ForInStatement represents a for-in loop.
//
//nolint:govet // Field order optimized for JSON output readability
type ForInStatement struct {
	BaseNode
	Left  interface{} `json:"left"`  // VariableDeclaration | Pattern
	Right Expression  `json:"right"`
	Body  Statement   `json:"body"`
}

func (n *ForInStatement) statementNode() {}

// ForOfStatement represents a for-of loop.
//
//nolint:govet // Field order optimized for JSON output readability
type ForOfStatement struct {
	BaseNode
	Await bool        `json:"await"`
	Left  interface{} `json:"left"` // VariableDeclaration | Pattern
	Right Expression  `json:"right"`
	Body  Statement   `json:"body"`
}

func (n *ForOfStatement) statementNode() {}

// ThrowStatement represents a throw statement.
//
//nolint:govet // Field order optimized for JSON output readability
type ThrowStatement struct {
	BaseNode
	Argument Expression `json:"argument"`
}

func (n *ThrowStatement) statementNode() {}

// TryStatement represents a try-catch-finally statement.
//
//nolint:govet // Field order optimized for JSON output readability
type TryStatement struct {
	BaseNode
	Block     *BlockStatement `json:"block"`
	Handler   *CatchClause    `json:"handler"`
	Finalizer *BlockStatement `json:"finalizer"`
}

func (n *TryStatement) statementNode() {}

// CatchClause represents a catch clause.
//
//nolint:govet // Field order optimized for JSON output readability
type CatchClause struct {
	BaseNode
	Param Pattern         `json:"param"`
	Body  *BlockStatement `json:"body"`
}

// WithStatement represents a with statement.
//
//nolint:govet // Field order optimized for JSON output readability
type WithStatement struct {
	BaseNode
	Object Expression `json:"object"`
	Body   Statement  `json:"body"`
}

func (n *WithStatement) statementNode() {}

// ==================== Declarations ====================

// VariableDeclaration represents a variable declaration (var, let, const).
//
//nolint:govet // Field order optimized for JSON output readability
type VariableDeclaration struct {
	BaseNode
	Declarations []VariableDeclarator `json:"declarations"`
	Kind         string               `json:"kind"` // "var" | "let" | "const" | "using" | "await using"
	Declare      bool                 `json:"declare,omitempty"`
}

func (n *VariableDeclaration) statementNode()   {}
func (n *VariableDeclaration) declarationNode() {}

// VariableDeclarator represents a variable declarator.
//
//nolint:govet // Field order optimized for JSON output readability
type VariableDeclarator struct {
	BaseNode
	ID         Pattern    `json:"id"`
	Init       Expression `json:"init"`
	Definite   bool       `json:"definite,omitempty"` // TS: definite assignment assertion (!)
}

// FunctionDeclaration represents a function declaration.
//
//nolint:govet // Field order optimized for JSON output readability
type FunctionDeclaration struct {
	BaseNode
	ID              *Identifier                 `json:"id"`
	Params          []Pattern                   `json:"params"`
	Body            *BlockStatement             `json:"body"`
	Generator       bool                        `json:"generator"`
	Async           bool                        `json:"async"`
	Expression      bool                        `json:"expression"`
	Declare         bool                        `json:"declare,omitempty"`
	TypeParameters  *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
	ReturnType      *TSTypeAnnotation           `json:"returnType,omitempty"`
	Decorators      []Decorator                 `json:"decorators,omitempty"`
	Modifiers       []interface{}               `json:"modifiers,omitempty"`
}

func (n *FunctionDeclaration) statementNode()   {}
func (n *FunctionDeclaration) declarationNode() {}

// ClassDeclaration represents a class declaration.
//
//nolint:govet // Field order optimized for JSON output readability
type ClassDeclaration struct {
	BaseNode
	ID                  *Identifier                   `json:"id"`
	SuperClass          Expression                    `json:"superClass"`
	Body                *ClassBody                    `json:"body"`
	Decorators          []Decorator                   `json:"decorators,omitempty"`
	TypeParameters      *TSTypeParameterDeclaration   `json:"typeParameters,omitempty"`
	SuperTypeParameters *TSTypeParameterInstantiation `json:"superTypeParameters,omitempty"`
	Implements          []TSClassImplements           `json:"implements,omitempty"`
	Abstract            bool                          `json:"abstract,omitempty"`
	Declare             bool                          `json:"declare,omitempty"`
}

func (n *ClassDeclaration) statementNode()   {}
func (n *ClassDeclaration) declarationNode() {}

// ClassBody represents the body of a class.
//
//nolint:govet // Field order optimized for JSON output readability
type ClassBody struct {
	BaseNode
	Body []interface{} `json:"body"` // MethodDefinition | PropertyDefinition | StaticBlock | TSIndexSignature
}

// MethodDefinition represents a method in a class.
//
//nolint:govet // Field order optimized for JSON output readability
type MethodDefinition struct {
	BaseNode
	Key            Expression              `json:"key"`
	Value          *FunctionExpression     `json:"value"`
	Kind           string                  `json:"kind"` // "constructor" | "method" | "get" | "set"
	Computed       bool                    `json:"computed"`
	Static         bool                    `json:"static"`
	Decorators     []Decorator             `json:"decorators,omitempty"`
	Optional       bool                    `json:"optional,omitempty"`
	Override       bool                    `json:"override,omitempty"`
	Accessibility  *string                 `json:"accessibility,omitempty"` // "public" | "private" | "protected"
}

// PropertyDefinition represents a property in a class.
//
//nolint:govet // Field order optimized for JSON output readability
type PropertyDefinition struct {
	BaseNode
	Key            Expression        `json:"key"`
	Value          Expression        `json:"value"`
	Computed       bool              `json:"computed"`
	Static         bool              `json:"static"`
	Declare        bool              `json:"declare,omitempty"`
	Override       bool              `json:"override,omitempty"`
	Readonly       bool              `json:"readonly,omitempty"`
	Decorators     []Decorator       `json:"decorators,omitempty"`
	Optional       bool              `json:"optional,omitempty"`
	Definite       bool              `json:"definite,omitempty"`
	TypeAnnotation *TSTypeAnnotation `json:"typeAnnotation,omitempty"`
	Accessibility  *string           `json:"accessibility,omitempty"`
}

// AccessorProperty represents an accessor property (getter/setter shorthand).
//
//nolint:govet // Field order optimized for JSON output readability
type AccessorProperty struct {
	BaseNode
	Key            Expression        `json:"key"`
	Value          Expression        `json:"value"`
	Computed       bool              `json:"computed"`
	Static         bool              `json:"static"`
	Decorators     []Decorator       `json:"decorators,omitempty"`
	TypeAnnotation *TSTypeAnnotation `json:"typeAnnotation,omitempty"`
	Accessibility  *string           `json:"accessibility,omitempty"`
	Definite       bool              `json:"definite,omitempty"`
}

// StaticBlock represents a static initialization block in a class.
//
//nolint:govet // Field order optimized for JSON output readability
type StaticBlock struct {
	BaseNode
	Body []Statement `json:"body"`
}

// ==================== Module Import/Export ====================

// ImportDeclaration represents an import declaration.
//
//nolint:govet // Field order optimized for JSON output readability
type ImportDeclaration struct {
	BaseNode
	Specifiers       []interface{}     `json:"specifiers"` // ImportSpecifier | ImportDefaultSpecifier | ImportNamespaceSpecifier
	Source           *Literal          `json:"source"`
	Attributes       []ImportAttribute `json:"attributes,omitempty"`
	ImportKind       *string           `json:"importKind,omitempty"` // "type" | "value"
	AssertionEntries []ImportAttribute `json:"assertions,omitempty"` // Deprecated, use Attributes
}

func (n *ImportDeclaration) statementNode() {}

// ImportSpecifier represents a named import specifier.
//
//nolint:govet // Field order optimized for JSON output readability
type ImportSpecifier struct {
	BaseNode
	Imported   *Identifier `json:"imported"`
	Local      *Identifier `json:"local"`
	ImportKind *string     `json:"importKind,omitempty"` // "type" | "value"
}

// ImportDefaultSpecifier represents a default import specifier.
//
//nolint:govet // Field order optimized for JSON output readability
type ImportDefaultSpecifier struct {
	BaseNode
	Local *Identifier `json:"local"`
}

// ImportNamespaceSpecifier represents a namespace import specifier (* as x).
//
//nolint:govet // Field order optimized for JSON output readability
type ImportNamespaceSpecifier struct {
	BaseNode
	Local *Identifier `json:"local"`
}

// ImportAttribute represents an import attribute (with clause).
//
//nolint:govet // Field order optimized for JSON output readability
type ImportAttribute struct {
	BaseNode
	Key   interface{} `json:"key"`   // Identifier | Literal
	Value *Literal    `json:"value"`
}

// ExportNamedDeclaration represents a named export declaration.
//
//nolint:govet // Field order optimized for JSON output readability
type ExportNamedDeclaration struct {
	BaseNode
	Declaration Declaration       `json:"declaration"`
	Specifiers  []ExportSpecifier `json:"specifiers"`
	Source      *Literal          `json:"source"`
	ExportKind  *string           `json:"exportKind,omitempty"` // "type" | "value"
	Attributes  []ImportAttribute `json:"attributes,omitempty"`
	Assertions  []ImportAttribute `json:"assertions,omitempty"` // Deprecated
}

func (n *ExportNamedDeclaration) statementNode() {}

// ExportDefaultDeclaration represents a default export declaration.
//
//nolint:govet // Field order optimized for JSON output readability
type ExportDefaultDeclaration struct {
	BaseNode
	Declaration interface{} `json:"declaration"` // Declaration | Expression
	ExportKind  *string     `json:"exportKind,omitempty"`
}

func (n *ExportDefaultDeclaration) statementNode() {}

// ExportAllDeclaration represents an export * declaration.
//
//nolint:govet // Field order optimized for JSON output readability
type ExportAllDeclaration struct {
	BaseNode
	Source     *Literal          `json:"source"`
	Exported   *Identifier       `json:"exported"` // For export * as name
	ExportKind *string           `json:"exportKind,omitempty"`
	Attributes []ImportAttribute `json:"attributes,omitempty"`
	Assertions []ImportAttribute `json:"assertions,omitempty"` // Deprecated
}

func (n *ExportAllDeclaration) statementNode() {}

// ExportSpecifier represents an export specifier.
//
//nolint:govet // Field order optimized for JSON output readability
type ExportSpecifier struct {
	BaseNode
	Local      interface{} `json:"local"`      // Identifier | Literal (for string exports)
	Exported   interface{} `json:"exported"`   // Identifier | Literal
	ExportKind *string     `json:"exportKind,omitempty"`
}

// ==================== Patterns (Destructuring) ====================

// ArrayPattern represents an array destructuring pattern.
//
//nolint:govet // Field order optimized for JSON output readability
type ArrayPattern struct {
	BaseNode
	Elements       []Pattern         `json:"elements"` // Can include nil for holes
	Decorators     []Decorator       `json:"decorators,omitempty"`
	Optional       bool              `json:"optional,omitempty"`
	TypeAnnotation *TSTypeAnnotation `json:"typeAnnotation,omitempty"`
}

func (n *ArrayPattern) patternNode() {}

// ObjectPattern represents an object destructuring pattern.
//
//nolint:govet // Field order optimized for JSON output readability
type ObjectPattern struct {
	BaseNode
	Properties     []interface{}     `json:"properties"` // Property | RestElement
	Decorators     []Decorator       `json:"decorators,omitempty"`
	Optional       bool              `json:"optional,omitempty"`
	TypeAnnotation *TSTypeAnnotation `json:"typeAnnotation,omitempty"`
}

func (n *ObjectPattern) patternNode() {}

// RestElement represents a rest element in destructuring (...rest).
//
//nolint:govet // Field order optimized for JSON output readability
type RestElement struct {
	BaseNode
	Argument       Pattern           `json:"argument"`
	Decorators     []Decorator       `json:"decorators,omitempty"`
	Optional       bool              `json:"optional,omitempty"`
	TypeAnnotation *TSTypeAnnotation `json:"typeAnnotation,omitempty"`
	Value          Expression        `json:"value,omitempty"` // For parameter properties
}

func (n *RestElement) patternNode() {}

// AssignmentPattern represents a default value in destructuring (x = 1).
//
//nolint:govet // Field order optimized for JSON output readability
type AssignmentPattern struct {
	BaseNode
	Left           Pattern           `json:"left"`
	Right          Expression        `json:"right"`
	Decorators     []Decorator       `json:"decorators,omitempty"`
	Optional       bool              `json:"optional,omitempty"`
	TypeAnnotation *TSTypeAnnotation `json:"typeAnnotation,omitempty"`
}

func (n *AssignmentPattern) patternNode() {}

// ==================== Comments and Tokens ====================

// Comment represents a comment in the source code.
//
//nolint:govet // Field order optimized for JSON output readability
type Comment struct {
	Type  string           `json:"type"`  // "Line" | "Block"
	Value string           `json:"value"`
	Loc   *SourceLocation  `json:"loc,omitempty"`
	Range *Range           `json:"range,omitempty"`
}

// Token represents a token in the source code.
//
//nolint:govet // Field order optimized for JSON output readability
type Token struct {
	Type  string          `json:"type"`
	Value string          `json:"value"`
	Loc   *SourceLocation `json:"loc,omitempty"`
	Range *Range          `json:"range,omitempty"`
}

// ==================== Decorators ====================

// Decorator represents a decorator (@decorator).
//
//nolint:govet // Field order optimized for JSON output readability
type Decorator struct {
	BaseNode
	Expression Expression `json:"expression"`
}
