package ast

// This file defines TypeScript-specific AST node types.
// Based on: https://typescript-eslint.io/packages/typescript-estree/ast-spec/

// ==================== TypeScript Type Keywords ====================

// TSAnyKeyword represents the 'any' type keyword.
type TSAnyKeyword struct {
	BaseNode
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSAnyKeyword) TSTypeNode() {}

// TSBigIntKeyword represents the 'bigint' type keyword.
type TSBigIntKeyword struct {
	BaseNode
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSBigIntKeyword) TSTypeNode() {}

// TSBooleanKeyword represents the 'boolean' type keyword.
type TSBooleanKeyword struct {
	BaseNode
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSBooleanKeyword) TSTypeNode() {}

// TSIntrinsicKeyword represents the 'intrinsic' type keyword.
type TSIntrinsicKeyword struct {
	BaseNode
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSIntrinsicKeyword) TSTypeNode() {}

// TSNeverKeyword represents the 'never' type keyword.
type TSNeverKeyword struct {
	BaseNode
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSNeverKeyword) TSTypeNode() {}

// TSNullKeyword represents the 'null' type keyword.
type TSNullKeyword struct {
	BaseNode
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSNullKeyword) TSTypeNode() {}

// TSNumberKeyword represents the 'number' type keyword.
type TSNumberKeyword struct {
	BaseNode
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSNumberKeyword) TSTypeNode() {}

// TSObjectKeyword represents the 'object' type keyword.
type TSObjectKeyword struct {
	BaseNode
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSObjectKeyword) TSTypeNode() {}

// TSStringKeyword represents the 'string' type keyword.
type TSStringKeyword struct {
	BaseNode
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSStringKeyword) TSTypeNode() {}

// TSSymbolKeyword represents the 'symbol' type keyword.
type TSSymbolKeyword struct {
	BaseNode
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSSymbolKeyword) TSTypeNode() {}

// TSUndefinedKeyword represents the 'undefined' type keyword.
type TSUndefinedKeyword struct {
	BaseNode
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSUndefinedKeyword) TSTypeNode() {}

// TSUnknownKeyword represents the 'unknown' type keyword.
type TSUnknownKeyword struct {
	BaseNode
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSUnknownKeyword) TSTypeNode() {}

// TSVoidKeyword represents the 'void' type keyword.
type TSVoidKeyword struct {
	BaseNode
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSVoidKeyword) TSTypeNode() {}

// ==================== TypeScript Type Expressions ====================

// TSArrayType represents an array type (T[]).
type TSArrayType struct {
	BaseNode
	ElementType TSNode `json:"elementType"`
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSArrayType) TSTypeNode() {}

// TSTupleType represents a tuple type ([T, U]).
type TSTupleType struct {
	BaseNode
	ElementTypes []TSNode `json:"elementTypes"` // Can include TSNamedTupleMember
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSTupleType) TSTypeNode() {}

// TSUnionType represents a union type (T | U).
type TSUnionType struct {
	BaseNode
	Types []TSNode `json:"types"`
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSUnionType) TSTypeNode() {}

// TSIntersectionType represents an intersection type (T & U).
type TSIntersectionType struct {
	BaseNode
	Types []TSNode `json:"types"`
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSIntersectionType) TSTypeNode() {}

// TSConditionalType represents a conditional type (T extends U ? X : Y).
type TSConditionalType struct {
	BaseNode
	CheckType   TSNode `json:"checkType"`
	ExtendsType TSNode `json:"extendsType"`
	TrueType    TSNode `json:"trueType"`
	FalseType   TSNode `json:"falseType"`
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSConditionalType) TSTypeNode() {}

// TSInferType represents an infer type (infer T).
type TSInferType struct {
	BaseNode
	TypeParameter *TSTypeParameter `json:"typeParameter"`
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSInferType) TSTypeNode() {}

// TSTypeReference represents a type reference (Foo, Array<T>).
type TSTypeReference struct {
	BaseNode
	TypeName       interface{}                   `json:"typeName"` // Identifier | TSQualifiedName
	TypeArguments  *TSTypeParameterInstantiation `json:"typeArguments,omitempty"`
	TypeParameters *TSTypeParameterInstantiation `json:"typeParameters,omitempty"` // Deprecated
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSTypeReference) TSTypeNode() {}

// TSTypeQuery represents a typeof type query (typeof x).
type TSTypeQuery struct {
	BaseNode
	ExprName       interface{}                   `json:"exprName"` // Identifier | TSQualifiedName | TSImportType
	TypeArguments  *TSTypeParameterInstantiation `json:"typeArguments,omitempty"`
	TypeParameters *TSTypeParameterInstantiation `json:"typeParameters,omitempty"` // Deprecated
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSTypeQuery) TSTypeNode() {}

// TSTypeLiteral represents a type literal ({a: string}).
type TSTypeLiteral struct {
	BaseNode
	// TSPropertySignature | TSMethodSignature | TSCallSignatureDeclaration |
	// TSConstructSignatureDeclaration | TSIndexSignature
	Members []interface{} `json:"members"`
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSTypeLiteral) TSTypeNode() {}

// TSFunctionType represents a function type ((x: T) => U).
type TSFunctionType struct {
	BaseNode
	Params         []Pattern                   `json:"params"`
	ReturnType     *TSTypeAnnotation           `json:"returnType"`
	TypeParameters *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSFunctionType) TSTypeNode() {}

// TSConstructorType represents a constructor type (new () => T).
type TSConstructorType struct {
	BaseNode
	Params         []Pattern                   `json:"params"`
	ReturnType     *TSTypeAnnotation           `json:"returnType"`
	TypeParameters *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
	Abstract       bool                        `json:"abstract,omitempty"`
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSConstructorType) TSTypeNode() {}

// TSMappedType represents a mapped type ({[K in T]: U}).
type TSMappedType struct {
	BaseNode
	TypeParameter  *TSTypeParameter  `json:"typeParameter"`
	NameType       TSNode            `json:"nameType,omitempty"`
	TypeAnnotation *TSTypeAnnotation `json:"typeAnnotation,omitempty"`
	Optional       interface{}       `json:"optional,omitempty"` // true | false | "+" | "-"
	Readonly       interface{}       `json:"readonly,omitempty"` // true | false | "+" | "-"
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSMappedType) TSTypeNode() {}

// TSLiteralType represents a literal type ('foo', 42).
type TSLiteralType struct {
	BaseNode
	Literal interface{} `json:"literal"` // Literal | UnaryExpression | UpdateExpression
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSLiteralType) TSTypeNode() {}

// TSIndexedAccessType represents an indexed access type (T[K]).
type TSIndexedAccessType struct {
	BaseNode
	ObjectType TSNode `json:"objectType"`
	IndexType  TSNode `json:"indexType"`
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSIndexedAccessType) TSTypeNode() {}

// TSOptionalType represents an optional type (T?).
type TSOptionalType struct {
	BaseNode
	TypeAnnotation TSNode `json:"typeAnnotation"`
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSOptionalType) TSTypeNode() {}

// TSRestType represents a rest type (...T[]).
type TSRestType struct {
	BaseNode
	TypeAnnotation TSNode `json:"typeAnnotation"`
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSRestType) TSTypeNode() {}

// TSThisType represents the 'this' type.
type TSThisType struct {
	BaseNode
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSThisType) TSTypeNode() {}

// TSTypeOperator represents a type operator (keyof T, readonly T, unique T).
type TSTypeOperator struct {
	BaseNode
	Operator       string `json:"operator"` // "keyof" | "readonly" | "unique"
	TypeAnnotation TSNode `json:"typeAnnotation"`
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSTypeOperator) TSTypeNode() {}

// TSTemplateLiteralType represents a template literal type.
type TSTemplateLiteralType struct {
	BaseNode
	Quasis []TemplateElement `json:"quasis"`
	Types  []TSNode          `json:"types"`
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSTemplateLiteralType) TSTypeNode() {}

// TSImportType represents an import type (import('module').Type).
type TSImportType struct {
	BaseNode
	Argument       *TSLiteralType                `json:"argument"`            // Module specifier
	Qualifier      interface{}                   `json:"qualifier,omitempty"` // Identifier | TSQualifiedName
	TypeArguments  *TSTypeParameterInstantiation `json:"typeArguments,omitempty"`
	TypeParameters *TSTypeParameterInstantiation `json:"typeParameters,omitempty"` // Deprecated
	Attributes     *ImportAttribute              `json:"attributes,omitempty"`
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSImportType) TSTypeNode() {}

// ==================== TypeScript Type Annotations ====================

// TSTypeAnnotation represents a type annotation (: T).
type TSTypeAnnotation struct {
	BaseNode
	TypeAnnotation TSNode `json:"typeAnnotation"`
}

// ==================== TypeScript Type Declarations ====================

// TSTypeAliasDeclaration represents a type alias declaration.
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
type TSInterfaceBody struct {
	BaseNode
	// TSPropertySignature | TSMethodSignature | TSCallSignatureDeclaration |
	// TSConstructSignatureDeclaration | TSIndexSignature
	Body []interface{} `json:"body"`
}

// TSInterfaceHeritage represents an interface extends clause.
type TSInterfaceHeritage struct {
	BaseNode
	Expression     Expression                    `json:"expression"`
	TypeArguments  *TSTypeParameterInstantiation `json:"typeArguments,omitempty"`
	TypeParameters *TSTypeParameterInstantiation `json:"typeParameters,omitempty"` // Deprecated
}

// TSEnumDeclaration represents an enum declaration.
type TSEnumDeclaration struct {
	BaseNode
	ID      *Identifier    `json:"id"`
	Members []TSEnumMember `json:"members"`
	Const   bool           `json:"const,omitempty"`
	Declare bool           `json:"declare,omitempty"`
}

func (n *TSEnumDeclaration) statementNode()   {}
func (n *TSEnumDeclaration) declarationNode() {}

// TSEnumMember represents a member of an enum.
type TSEnumMember struct {
	BaseNode
	ID          interface{} `json:"id"` // Identifier | Literal
	Initializer Expression  `json:"initializer,omitempty"`
	Computed    bool        `json:"computed,omitempty"`
}

// TSModuleDeclaration represents a module or namespace declaration.
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
type TSModuleBlock struct {
	BaseNode
	Body []Statement `json:"body"`
}

func (n *TSModuleBlock) statementNode() {}

// ==================== TypeScript Type Parameters ====================

// TSTypeParameter represents a type parameter (<T>).
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
type TSTypeParameterDeclaration struct {
	BaseNode
	Params []TSTypeParameter `json:"params"`
}

// TSTypeParameterInstantiation represents a type parameter instantiation.
type TSTypeParameterInstantiation struct {
	BaseNode
	Params []TSNode `json:"params"`
}

// ==================== TypeScript Type Signatures ====================

// TSCallSignatureDeclaration represents a call signature.
type TSCallSignatureDeclaration struct {
	BaseNode
	Params         []Pattern                   `json:"params"`
	ReturnType     *TSTypeAnnotation           `json:"returnType,omitempty"`
	TypeParameters *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
}

// TSConstructSignatureDeclaration represents a construct signature.
type TSConstructSignatureDeclaration struct {
	BaseNode
	Params         []Pattern                   `json:"params"`
	ReturnType     *TSTypeAnnotation           `json:"returnType,omitempty"`
	TypeParameters *TSTypeParameterDeclaration `json:"typeParameters,omitempty"`
}

// TSPropertySignature represents a property signature in a type.
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
type TSNamedTupleMember struct {
	BaseNode
	Label       *Identifier `json:"label"`
	ElementType TSNode      `json:"elementType"`
	Optional    bool        `json:"optional,omitempty"`
}

// ==================== TypeScript Type Assertions & Expressions ====================

// TSAsExpression represents a type assertion using 'as' (x as T).
type TSAsExpression struct {
	BaseNode
	Expression     Expression `json:"expression"`
	TypeAnnotation TSNode     `json:"typeAnnotation"`
}

func (n *TSAsExpression) expressionNode() {}

// TSTypeAssertion represents a type assertion using angle brackets (<T>x).
type TSTypeAssertion struct {
	BaseNode
	Expression     Expression `json:"expression"`
	TypeAnnotation TSNode     `json:"typeAnnotation"`
}

func (n *TSTypeAssertion) expressionNode() {}

// TSNonNullExpression represents a non-null assertion (x!).
type TSNonNullExpression struct {
	BaseNode
	Expression Expression `json:"expression"`
}

func (n *TSNonNullExpression) expressionNode() {}

// TSSatisfiesExpression represents a satisfies expression (x satisfies T).
type TSSatisfiesExpression struct {
	BaseNode
	Expression     Expression `json:"expression"`
	TypeAnnotation TSNode     `json:"typeAnnotation"`
}

func (n *TSSatisfiesExpression) expressionNode() {}

// TSInstantiationExpression represents a type instantiation (Foo<T>).
type TSInstantiationExpression struct {
	BaseNode
	Expression     Expression                    `json:"expression"`
	TypeArguments  *TSTypeParameterInstantiation `json:"typeArguments"`
	TypeParameters *TSTypeParameterInstantiation `json:"typeParameters,omitempty"` // Deprecated
}

func (n *TSInstantiationExpression) expressionNode() {}

// ==================== TypeScript Type Predicates ====================

// TSTypePredicate represents a type predicate (x is T).
type TSTypePredicate struct {
	BaseNode
	ParameterName  interface{}       `json:"parameterName"` // Identifier | TSThisType
	TypeAnnotation *TSTypeAnnotation `json:"typeAnnotation,omitempty"`
	Asserts        bool              `json:"asserts,omitempty"`
}

// TSTypeNode marks this as a TypeScript type node.
func (n *TSTypePredicate) TSTypeNode() {}

// ==================== TypeScript Abstract Members ====================

// TSAbstractAccessorProperty represents an abstract accessor property.
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
type TSImportEqualsDeclaration struct {
	BaseNode
	ID              *Identifier `json:"id"`
	ModuleReference interface{} `json:"moduleReference"` // TSExternalModuleReference | Identifier | TSQualifiedName
	IsExport        bool        `json:"isExport,omitempty"`
	ImportKind      string      `json:"importKind,omitempty"` // "type" | "value"
}

func (n *TSImportEqualsDeclaration) statementNode() {}

// TSExternalModuleReference represents an external module reference.
type TSExternalModuleReference struct {
	BaseNode
	Expression Expression `json:"expression"`
}

// TSExportAssignment represents an export = statement.
type TSExportAssignment struct {
	BaseNode
	Expression Expression `json:"expression"`
}

func (n *TSExportAssignment) statementNode() {}

// TSNamespaceExportDeclaration represents a namespace export declaration.
type TSNamespaceExportDeclaration struct {
	BaseNode
	ID *Identifier `json:"id"`
}

func (n *TSNamespaceExportDeclaration) statementNode() {}

// ==================== TypeScript Other ====================

// TSQualifiedName represents a qualified name (A.B.C).
type TSQualifiedName struct {
	BaseNode
	Left  interface{} `json:"left"` // Identifier | TSQualifiedName
	Right *Identifier `json:"right"`
}

// TSParameterProperty represents a parameter property in a constructor.
type TSParameterProperty struct {
	BaseNode
	Parameter     Pattern     `json:"parameter"`               // Identifier | AssignmentPattern
	Accessibility *string     `json:"accessibility,omitempty"` // "public" | "private" | "protected"
	Readonly      bool        `json:"readonly,omitempty"`
	Static        bool        `json:"static,omitempty"`
	Override      bool        `json:"override,omitempty"`
	Decorators    []Decorator `json:"decorators,omitempty"`
}

func (n *TSParameterProperty) patternNode() {}

// TSDeclareFunction represents a declare function statement.
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
