package ast

// This file defines TypeScript-specific AST node types.
// Based on: https://typescript-eslint.io/packages/typescript-estree/ast-spec/

// ==================== TypeScript Type Keywords ====================

// TSAnyKeyword represents the 'any' type keyword.
type TSAnyKeyword struct {
	BaseNode
}

func (n *TSAnyKeyword) TSTypeNode() {}

// TSBigIntKeyword represents the 'bigint' type keyword.
type TSBigIntKeyword struct {
	BaseNode
}

func (n *TSBigIntKeyword) TSTypeNode() {}

// TSBooleanKeyword represents the 'boolean' type keyword.
type TSBooleanKeyword struct {
	BaseNode
}

func (n *TSBooleanKeyword) TSTypeNode() {}

// TSIntrinsicKeyword represents the 'intrinsic' type keyword.
type TSIntrinsicKeyword struct {
	BaseNode
}

func (n *TSIntrinsicKeyword) TSTypeNode() {}

// TSNeverKeyword represents the 'never' type keyword.
type TSNeverKeyword struct {
	BaseNode
}

func (n *TSNeverKeyword) TSTypeNode() {}

// TSNullKeyword represents the 'null' type keyword.
type TSNullKeyword struct {
	BaseNode
}

func (n *TSNullKeyword) TSTypeNode() {}

// TSNumberKeyword represents the 'number' type keyword.
type TSNumberKeyword struct {
	BaseNode
}

func (n *TSNumberKeyword) TSTypeNode() {}

// TSObjectKeyword represents the 'object' type keyword.
type TSObjectKeyword struct {
	BaseNode
}

func (n *TSObjectKeyword) TSTypeNode() {}

// TSStringKeyword represents the 'string' type keyword.
type TSStringKeyword struct {
	BaseNode
}

func (n *TSStringKeyword) TSTypeNode() {}

// TSSymbolKeyword represents the 'symbol' type keyword.
type TSSymbolKeyword struct {
	BaseNode
}

func (n *TSSymbolKeyword) TSTypeNode() {}

// TSUndefinedKeyword represents the 'undefined' type keyword.
type TSUndefinedKeyword struct {
	BaseNode
}

func (n *TSUndefinedKeyword) TSTypeNode() {}

// TSUnknownKeyword represents the 'unknown' type keyword.
type TSUnknownKeyword struct {
	BaseNode
}

func (n *TSUnknownKeyword) TSTypeNode() {}

// TSVoidKeyword represents the 'void' type keyword.
type TSVoidKeyword struct {
	BaseNode
}

func (n *TSVoidKeyword) TSTypeNode() {}

// ==================== TypeScript Type Expressions ====================

// TSArrayType represents an array type (T[]).
//
//nolint:govet // Field order optimized for JSON output readability
type TSArrayType struct {
	BaseNode
	ElementType TSNode `json:"elementType"`
}

func (n *TSArrayType) TSTypeNode() {}

// TSTupleType represents a tuple type ([T, U]).
//
//nolint:govet // Field order optimized for JSON output readability
type TSTupleType struct {
	BaseNode
	ElementTypes []TSNode `json:"elementTypes"` // Can include TSNamedTupleMember
}

func (n *TSTupleType) TSTypeNode() {}

// TSUnionType represents a union type (T | U).
//
//nolint:govet // Field order optimized for JSON output readability
type TSUnionType struct {
	BaseNode
	Types []TSNode `json:"types"`
}

func (n *TSUnionType) TSTypeNode() {}

// TSIntersectionType represents an intersection type (T & U).
//
//nolint:govet // Field order optimized for JSON output readability
type TSIntersectionType struct {
	BaseNode
	Types []TSNode `json:"types"`
}

func (n *TSIntersectionType) TSTypeNode() {}

// TSConditionalType represents a conditional type (T extends U ? X : Y).
//
//nolint:govet // Field order optimized for JSON output readability
type TSConditionalType struct {
	BaseNode
	CheckType   TSNode `json:"checkType"`
	ExtendsType TSNode `json:"extendsType"`
	TrueType    TSNode `json:"trueType"`
	FalseType   TSNode `json:"falseType"`
}

func (n *TSConditionalType) TSTypeNode() {}

// TSInferType represents an infer type (infer T).
//
//nolint:govet // Field order optimized for JSON output readability
type TSInferType struct {
	BaseNode
	TypeParameter *TSTypeParameter `json:"typeParameter"`
}

func (n *TSInferType) TSTypeNode() {}

// TSTypeReference represents a type reference (Foo, Array<T>).
//
//nolint:govet // Field order optimized for JSON output readability
type TSTypeReference struct {
	BaseNode
	TypeName       interface{}                   `json:"typeName"` // Identifier | TSQualifiedName
	TypeArguments  *TSTypeParameterInstantiation `json:"typeArguments,omitempty"`
	TypeParameters *TSTypeParameterInstantiation `json:"typeParameters,omitempty"` // Deprecated
}

func (n *TSTypeReference) TSTypeNode() {}

// TSTypeQuery represents a typeof type query (typeof x).
//
//nolint:govet // Field order optimized for JSON output readability
type TSTypeQuery struct {
	BaseNode
	ExprName       interface{}                   `json:"exprName"` // Identifier | TSQualifiedName | TSImportType
	TypeArguments  *TSTypeParameterInstantiation `json:"typeArguments,omitempty"`
	TypeParameters *TSTypeParameterInstantiation `json:"typeParameters,omitempty"` // Deprecated
}

func (n *TSTypeQuery) TSTypeNode() {}

// TSTypeLiteral represents a type literal ({a: string}).
//
//nolint:govet // Field order optimized for JSON output readability
type TSTypeLiteral struct {
	BaseNode
	Members []interface{} `json:"members"` // TSPropertySignature | TSMethodSignature | TSCallSignatureDeclaration | TSConstructSignatureDeclaration | TSIndexSignature
}

func (n *TSTypeLiteral) TSTypeNode() {}

// TSFunctionType represents a function type ((x: T) => U).
//
//nolint:govet // Field order optimized for JSON output readability
type TSFunctionType struct {
	BaseNode
	Params         []Pattern                   `json:"params"`
	ReturnType     *TSTypeAnnotation           `json:"returnType"`
	TypeParameters *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
}

func (n *TSFunctionType) TSTypeNode() {}

// TSConstructorType represents a constructor type (new () => T).
//
//nolint:govet // Field order optimized for JSON output readability
type TSConstructorType struct {
	BaseNode
	Params         []Pattern                   `json:"params"`
	ReturnType     *TSTypeAnnotation           `json:"returnType"`
	TypeParameters *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
	Abstract       bool                        `json:"abstract,omitempty"`
}

func (n *TSConstructorType) TSTypeNode() {}

// TSMappedType represents a mapped type ({[K in T]: U}).
//
//nolint:govet // Field order optimized for JSON output readability
type TSMappedType struct {
	BaseNode
	TypeParameter  *TSTypeParameter `json:"typeParameter"`
	NameType       TSNode           `json:"nameType,omitempty"`
	TypeAnnotation *TSTypeAnnotation `json:"typeAnnotation,omitempty"`
	Optional       interface{}      `json:"optional,omitempty"` // true | false | "+" | "-"
	Readonly       interface{}      `json:"readonly,omitempty"` // true | false | "+" | "-"
}

func (n *TSMappedType) TSTypeNode() {}

// TSLiteralType represents a literal type ('foo', 42).
//
//nolint:govet // Field order optimized for JSON output readability
type TSLiteralType struct {
	BaseNode
	Literal interface{} `json:"literal"` // Literal | UnaryExpression | UpdateExpression
}

func (n *TSLiteralType) TSTypeNode() {}

// TSIndexedAccessType represents an indexed access type (T[K]).
//
//nolint:govet // Field order optimized for JSON output readability
type TSIndexedAccessType struct {
	BaseNode
	ObjectType TSNode `json:"objectType"`
	IndexType  TSNode `json:"indexType"`
}

func (n *TSIndexedAccessType) TSTypeNode() {}

// TSOptionalType represents an optional type (T?).
//
//nolint:govet // Field order optimized for JSON output readability
type TSOptionalType struct {
	BaseNode
	TypeAnnotation TSNode `json:"typeAnnotation"`
}

func (n *TSOptionalType) TSTypeNode() {}

// TSRestType represents a rest type (...T[]).
//
//nolint:govet // Field order optimized for JSON output readability
type TSRestType struct {
	BaseNode
	TypeAnnotation TSNode `json:"typeAnnotation"`
}

func (n *TSRestType) TSTypeNode() {}

// TSThisType represents the 'this' type.
type TSThisType struct {
	BaseNode
}

func (n *TSThisType) TSTypeNode() {}

// TSTypeOperator represents a type operator (keyof T, readonly T, unique T).
//
//nolint:govet // Field order optimized for JSON output readability
type TSTypeOperator struct {
	BaseNode
	Operator       string `json:"operator"` // "keyof" | "readonly" | "unique"
	TypeAnnotation TSNode `json:"typeAnnotation"`
}

func (n *TSTypeOperator) TSTypeNode() {}

// TSTemplateLiteralType represents a template literal type.
//
//nolint:govet // Field order optimized for JSON output readability
type TSTemplateLiteralType struct {
	BaseNode
	Quasis []TemplateElement `json:"quasis"`
	Types  []TSNode          `json:"types"`
}

func (n *TSTemplateLiteralType) TSTypeNode() {}

// TSImportType represents an import type (import('module').Type).
//
//nolint:govet // Field order optimized for JSON output readability
type TSImportType struct {
	BaseNode
	Argument       *TSLiteralType                `json:"argument"` // Module specifier
	Qualifier      interface{}                   `json:"qualifier,omitempty"` // Identifier | TSQualifiedName
	TypeArguments  *TSTypeParameterInstantiation `json:"typeArguments,omitempty"`
	TypeParameters *TSTypeParameterInstantiation `json:"typeParameters,omitempty"` // Deprecated
	Attributes     *ImportAttribute              `json:"attributes,omitempty"`
}

func (n *TSImportType) TSTypeNode() {}

// ==================== TypeScript Type Annotations ====================

// TSTypeAnnotation represents a type annotation (: T).
//
//nolint:govet // Field order optimized for JSON output readability
type TSTypeAnnotation struct {
	BaseNode
	TypeAnnotation TSNode `json:"typeAnnotation"`
}

// ==================== TypeScript Type Declarations ====================

// TSTypeAliasDeclaration represents a type alias declaration.
//
//nolint:govet // Field order optimized for JSON output readability
type TSTypeAliasDeclaration struct {
	BaseNode
	ID             *Identifier                 `json:"id"`
	TypeAnnotation TSNode                      `json:"typeAnnotation"`
	TypeParameters *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
	Declare        bool                        `json:"declare,omitempty"`
}

func (n *TSTypeAliasDeclaration) statementNode()   {}
func (n *TSTypeAliasDeclaration) declarationNode() {}

// TSInterfaceDeclaration represents an interface declaration.
//
//nolint:govet // Field order optimized for JSON output readability
type TSInterfaceDeclaration struct {
	BaseNode
	ID             *Identifier                 `json:"id"`
	Body           *TSInterfaceBody            `json:"body"`
	TypeParameters *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
	Extends        []TSInterfaceHeritage       `json:"extends,omitempty"`
	Implements     []TSInterfaceHeritage       `json:"implements,omitempty"`
	Abstract       bool                        `json:"abstract,omitempty"`
	Declare        bool                        `json:"declare,omitempty"`
}

func (n *TSInterfaceDeclaration) statementNode()   {}
func (n *TSInterfaceDeclaration) declarationNode() {}

// TSInterfaceBody represents the body of an interface.
//
//nolint:govet // Field order optimized for JSON output readability
type TSInterfaceBody struct {
	BaseNode
	Body []interface{} `json:"body"` // TSPropertySignature | TSMethodSignature | TSCallSignatureDeclaration | TSConstructSignatureDeclaration | TSIndexSignature
}

// TSInterfaceHeritage represents an interface extends clause.
//
//nolint:govet // Field order optimized for JSON output readability
type TSInterfaceHeritage struct {
	BaseNode
	Expression     Expression                    `json:"expression"`
	TypeArguments  *TSTypeParameterInstantiation `json:"typeArguments,omitempty"`
	TypeParameters *TSTypeParameterInstantiation `json:"typeParameters,omitempty"` // Deprecated
}

// TSEnumDeclaration represents an enum declaration.
//
//nolint:govet // Field order optimized for JSON output readability
type TSEnumDeclaration struct {
	BaseNode
	ID      *Identifier  `json:"id"`
	Members []TSEnumMember `json:"members"`
	Const   bool         `json:"const,omitempty"`
	Declare bool         `json:"declare,omitempty"`
}

func (n *TSEnumDeclaration) statementNode()   {}
func (n *TSEnumDeclaration) declarationNode() {}

// TSEnumMember represents a member of an enum.
//
//nolint:govet // Field order optimized for JSON output readability
type TSEnumMember struct {
	BaseNode
	ID          interface{} `json:"id"` // Identifier | Literal
	Initializer Expression  `json:"initializer,omitempty"`
	Computed    bool        `json:"computed,omitempty"`
}

// TSModuleDeclaration represents a module or namespace declaration.
//
//nolint:govet // Field order optimized for JSON output readability
type TSModuleDeclaration struct {
	BaseNode
	ID      interface{} `json:"id"`   // Identifier | Literal (for string module names)
	Body    interface{} `json:"body"` // TSModuleBlock | TSModuleDeclaration
	Global  bool        `json:"global,omitempty"`
	Declare bool        `json:"declare,omitempty"`
	Kind    string      `json:"kind,omitempty"` // "module" | "namespace"
}

func (n *TSModuleDeclaration) statementNode()   {}
func (n *TSModuleDeclaration) declarationNode() {}

// TSModuleBlock represents the body of a module.
//
//nolint:govet // Field order optimized for JSON output readability
type TSModuleBlock struct {
	BaseNode
	Body []Statement `json:"body"`
}

func (n *TSModuleBlock) statementNode() {}

// ==================== TypeScript Type Parameters ====================

// TSTypeParameter represents a type parameter (<T>).
//
//nolint:govet // Field order optimized for JSON output readability
type TSTypeParameter struct {
	BaseNode
	Name       *Identifier `json:"name"`
	Constraint TSNode      `json:"constraint,omitempty"`
	Default    TSNode      `json:"default,omitempty"`
	In         bool        `json:"in,omitempty"`
	Out        bool        `json:"out,omitempty"`
	Const      bool        `json:"const,omitempty"`
}

// TSTypeParameterDeclaration represents a type parameter declaration.
//
//nolint:govet // Field order optimized for JSON output readability
type TSTypeParameterDeclaration struct {
	BaseNode
	Params []TSTypeParameter `json:"params"`
}

// TSTypeParameterInstantiation represents a type parameter instantiation.
//
//nolint:govet // Field order optimized for JSON output readability
type TSTypeParameterInstantiation struct {
	BaseNode
	Params []TSNode `json:"params"`
}

// ==================== TypeScript Type Signatures ====================

// TSCallSignatureDeclaration represents a call signature.
//
//nolint:govet // Field order optimized for JSON output readability
type TSCallSignatureDeclaration struct {
	BaseNode
	Params         []Pattern                   `json:"params"`
	ReturnType     *TSTypeAnnotation           `json:"returnType,omitempty"`
	TypeParameters *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
}

// TSConstructSignatureDeclaration represents a construct signature.
//
//nolint:govet // Field order optimized for JSON output readability
type TSConstructSignatureDeclaration struct {
	BaseNode
	Params         []Pattern                   `json:"params"`
	ReturnType     *TSTypeAnnotation           `json:"returnType,omitempty"`
	TypeParameters *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
}

// TSPropertySignature represents a property signature in a type.
//
//nolint:govet // Field order optimized for JSON output readability
type TSPropertySignature struct {
	BaseNode
	Key            Expression        `json:"key"`
	TypeAnnotation *TSTypeAnnotation `json:"typeAnnotation,omitempty"`
	Initializer    Expression        `json:"initializer,omitempty"`
	Computed       bool              `json:"computed"`
	Optional       bool              `json:"optional,omitempty"`
	Readonly       bool              `json:"readonly,omitempty"`
	Static         bool              `json:"static,omitempty"`
	Export         bool              `json:"export,omitempty"`
	Accessibility  *string           `json:"accessibility,omitempty"`
}

// TSMethodSignature represents a method signature in a type.
//
//nolint:govet // Field order optimized for JSON output readability
type TSMethodSignature struct {
	BaseNode
	Key            Expression                  `json:"key"`
	Params         []Pattern                   `json:"params"`
	ReturnType     *TSTypeAnnotation           `json:"returnType,omitempty"`
	Computed       bool                        `json:"computed"`
	Optional       bool                        `json:"optional,omitempty"`
	Static         bool                        `json:"static,omitempty"`
	Readonly       bool                        `json:"readonly,omitempty"`
	TypeParameters *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
	Accessibility  *string                     `json:"accessibility,omitempty"`
	Export         bool                        `json:"export,omitempty"`
	Kind           string                      `json:"kind"` // "method" | "get" | "set"
}

// TSIndexSignature represents an index signature.
//
//nolint:govet // Field order optimized for JSON output readability
type TSIndexSignature struct {
	BaseNode
	Parameters     []Pattern         `json:"parameters"`
	TypeAnnotation *TSTypeAnnotation `json:"typeAnnotation,omitempty"`
	Readonly       bool              `json:"readonly,omitempty"`
	Static         bool              `json:"static,omitempty"`
	Export         bool              `json:"export,omitempty"`
	Accessibility  *string           `json:"accessibility,omitempty"`
}

// TSNamedTupleMember represents a named tuple member.
//
//nolint:govet // Field order optimized for JSON output readability
type TSNamedTupleMember struct {
	BaseNode
	Label          *Identifier  `json:"label"`
	ElementType    TSNode       `json:"elementType"`
	Optional       bool         `json:"optional,omitempty"`
}

// ==================== TypeScript Type Assertions & Expressions ====================

// TSAsExpression represents a type assertion using 'as' (x as T).
//
//nolint:govet // Field order optimized for JSON output readability
type TSAsExpression struct {
	BaseNode
	Expression     Expression `json:"expression"`
	TypeAnnotation TSNode     `json:"typeAnnotation"`
}

func (n *TSAsExpression) expressionNode() {}

// TSTypeAssertion represents a type assertion using angle brackets (<T>x).
//
//nolint:govet // Field order optimized for JSON output readability
type TSTypeAssertion struct {
	BaseNode
	Expression     Expression `json:"expression"`
	TypeAnnotation TSNode     `json:"typeAnnotation"`
}

func (n *TSTypeAssertion) expressionNode() {}

// TSNonNullExpression represents a non-null assertion (x!).
//
//nolint:govet // Field order optimized for JSON output readability
type TSNonNullExpression struct {
	BaseNode
	Expression Expression `json:"expression"`
}

func (n *TSNonNullExpression) expressionNode() {}

// TSSatisfiesExpression represents a satisfies expression (x satisfies T).
//
//nolint:govet // Field order optimized for JSON output readability
type TSSatisfiesExpression struct {
	BaseNode
	Expression     Expression `json:"expression"`
	TypeAnnotation TSNode     `json:"typeAnnotation"`
}

func (n *TSSatisfiesExpression) expressionNode() {}

// TSInstantiationExpression represents a type instantiation (Foo<T>).
//
//nolint:govet // Field order optimized for JSON output readability
type TSInstantiationExpression struct {
	BaseNode
	Expression     Expression                    `json:"expression"`
	TypeArguments  *TSTypeParameterInstantiation `json:"typeArguments"`
	TypeParameters *TSTypeParameterInstantiation `json:"typeParameters,omitempty"` // Deprecated
}

func (n *TSInstantiationExpression) expressionNode() {}

// ==================== TypeScript Type Predicates ====================

// TSTypePredicate represents a type predicate (x is T).
//
//nolint:govet // Field order optimized for JSON output readability
type TSTypePredicate struct {
	BaseNode
	ParameterName  interface{}       `json:"parameterName"` // Identifier | TSThisType
	TypeAnnotation *TSTypeAnnotation `json:"typeAnnotation,omitempty"`
	Asserts        bool              `json:"asserts,omitempty"`
}

func (n *TSTypePredicate) TSTypeNode() {}

// ==================== TypeScript Abstract Members ====================

// TSAbstractAccessorProperty represents an abstract accessor property.
//
//nolint:govet // Field order optimized for JSON output readability
type TSAbstractAccessorProperty struct {
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

// TSAbstractMethodDefinition represents an abstract method.
//
//nolint:govet // Field order optimized for JSON output readability
type TSAbstractMethodDefinition struct {
	BaseNode
	Key           Expression          `json:"key"`
	Value         *FunctionExpression `json:"value"`
	Kind          string              `json:"kind"` // "method" | "get" | "set"
	Computed      bool                `json:"computed"`
	Static        bool                `json:"static"`
	Decorators    []Decorator         `json:"decorators,omitempty"`
	Optional      bool                `json:"optional,omitempty"`
	Override      bool                `json:"override,omitempty"`
	Accessibility *string             `json:"accessibility,omitempty"`
}

// TSAbstractPropertyDefinition represents an abstract property.
//
//nolint:govet // Field order optimized for JSON output readability
type TSAbstractPropertyDefinition struct {
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

// ==================== TypeScript Import/Export ====================

// TSImportEqualsDeclaration represents an import = declaration.
//
//nolint:govet // Field order optimized for JSON output readability
type TSImportEqualsDeclaration struct {
	BaseNode
	ID            *Identifier `json:"id"`
	ModuleReference interface{} `json:"moduleReference"` // TSExternalModuleReference | Identifier | TSQualifiedName
	IsExport      bool        `json:"isExport,omitempty"`
	ImportKind    string      `json:"importKind,omitempty"` // "type" | "value"
}

func (n *TSImportEqualsDeclaration) statementNode() {}

// TSExternalModuleReference represents an external module reference.
//
//nolint:govet // Field order optimized for JSON output readability
type TSExternalModuleReference struct {
	BaseNode
	Expression Expression `json:"expression"`
}

// TSExportAssignment represents an export = statement.
//
//nolint:govet // Field order optimized for JSON output readability
type TSExportAssignment struct {
	BaseNode
	Expression Expression `json:"expression"`
}

func (n *TSExportAssignment) statementNode() {}

// TSNamespaceExportDeclaration represents a namespace export declaration.
//
//nolint:govet // Field order optimized for JSON output readability
type TSNamespaceExportDeclaration struct {
	BaseNode
	ID *Identifier `json:"id"`
}

func (n *TSNamespaceExportDeclaration) statementNode() {}

// ==================== TypeScript Other ====================

// TSQualifiedName represents a qualified name (A.B.C).
//
//nolint:govet // Field order optimized for JSON output readability
type TSQualifiedName struct {
	BaseNode
	Left  interface{} `json:"left"`  // Identifier | TSQualifiedName
	Right *Identifier `json:"right"`
}

// TSParameterProperty represents a parameter property in a constructor.
//
//nolint:govet // Field order optimized for JSON output readability
type TSParameterProperty struct {
	BaseNode
	Parameter      Pattern     `json:"parameter"` // Identifier | AssignmentPattern
	Accessibility  *string     `json:"accessibility,omitempty"` // "public" | "private" | "protected"
	Readonly       bool        `json:"readonly,omitempty"`
	Static         bool        `json:"static,omitempty"`
	Override       bool        `json:"override,omitempty"`
	Decorators     []Decorator `json:"decorators,omitempty"`
}

func (n *TSParameterProperty) patternNode() {}

// TSDeclareFunction represents a declare function statement.
//
//nolint:govet // Field order optimized for JSON output readability
type TSDeclareFunction struct {
	BaseNode
	ID             *Identifier                 `json:"id"`
	Params         []Pattern                   `json:"params"`
	ReturnType     *TSTypeAnnotation           `json:"returnType,omitempty"`
	Generator      bool                        `json:"generator"`
	Async          bool                        `json:"async"`
	Declare        bool                        `json:"declare,omitempty"`
	TypeParameters *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
}

func (n *TSDeclareFunction) statementNode() {}

// TSEmptyBodyFunctionExpression represents a function with no body.
//
//nolint:govet // Field order optimized for JSON output readability
type TSEmptyBodyFunctionExpression struct {
	BaseNode
	ID             *Identifier                 `json:"id"`
	Params         []Pattern                   `json:"params"`
	ReturnType     *TSTypeAnnotation           `json:"returnType,omitempty"`
	Generator      bool                        `json:"generator"`
	Async          bool                        `json:"async"`
	TypeParameters *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
}

func (n *TSEmptyBodyFunctionExpression) expressionNode() {}

// TSClassImplements represents a class implements clause.
//
//nolint:govet // Field order optimized for JSON output readability
type TSClassImplements struct {
	BaseNode
	Expression     Expression                    `json:"expression"`
	TypeArguments  *TSTypeParameterInstantiation `json:"typeArguments,omitempty"`
	TypeParameters *TSTypeParameterInstantiation `json:"typeParameters,omitempty"` // Deprecated
}

// TSEnumBody is an alias kept for compatibility.
type TSEnumBody struct {
	BaseNode
	Members []TSEnumMember `json:"members"`
}
