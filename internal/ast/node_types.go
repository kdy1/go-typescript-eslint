package ast

// NodeType represents the type of an AST node.
// This is equivalent to AST_NODE_TYPES in TypeScript ESTree.
type NodeType int

// AST Node Types - Complete enumeration of all ESTree and TypeScript-specific node types.
// Based on: https://typescript-eslint.io/packages/typescript-estree/ast-spec/
const (
	// NodeTypeUnknown represents an unknown or uninitialized node type.
	NodeTypeUnknown NodeType = iota

	// ==================== Program & Core ====================

	// NodeTypeProgram represents the root node of an AST.
	NodeTypeProgram

	// ==================== Identifiers & Literals ====================

	// NodeTypeIdentifier represents an identifier (variable name, function name, etc.).
	NodeTypeIdentifier
	// NodeTypePrivateIdentifier represents a private identifier (#field).
	NodeTypePrivateIdentifier
	// NodeTypeLiteral represents a literal value (string, number, boolean, null, regex).
	NodeTypeLiteral

	// ==================== Expressions ====================

	// NodeTypeThisExpression represents the 'this' keyword.
	NodeTypeThisExpression
	// NodeTypeSuper represents the 'super' keyword.
	NodeTypeSuper

	// NodeTypeArrayExpression represents an array literal [1, 2, 3].
	NodeTypeArrayExpression
	// NodeTypeObjectExpression represents an object literal {a: 1, b: 2}.
	NodeTypeObjectExpression
	// NodeTypeProperty represents a property in an object expression.
	NodeTypeProperty

	// NodeTypeFunctionExpression represents a function expression.
	NodeTypeFunctionExpression
	// NodeTypeArrowFunctionExpression represents an arrow function expression.
	NodeTypeArrowFunctionExpression

	// NodeTypeClassExpression represents a class expression.
	NodeTypeClassExpression

	// NodeTypeUnaryExpression represents a unary operation (+x, -x, !x, ~x, typeof x, void x, delete x).
	NodeTypeUnaryExpression
	// NodeTypeUpdateExpression represents an update expression (++x, x++, --x, x--).
	NodeTypeUpdateExpression

	// NodeTypeBinaryExpression represents a binary operation (x + y, x - y, x * y, etc.).
	NodeTypeBinaryExpression
	// NodeTypeLogicalExpression represents a logical operation (x && y, x || y, x ?? y).
	NodeTypeLogicalExpression
	// NodeTypeAssignmentExpression represents an assignment (x = y, x += y, etc.).
	NodeTypeAssignmentExpression

	// NodeTypeConditionalExpression represents a ternary conditional (x ? y : z).
	NodeTypeConditionalExpression
	// NodeTypeSequenceExpression represents a sequence of expressions (x, y, z).
	NodeTypeSequenceExpression

	// NodeTypeMemberExpression represents a member access (obj.prop, obj[prop]).
	NodeTypeMemberExpression
	// NodeTypeCallExpression represents a function call.
	NodeTypeCallExpression
	// NodeTypeNewExpression represents a new expression (new Foo()).
	NodeTypeNewExpression
	// NodeTypeMetaProperty represents a meta property (new.target, import.meta).
	NodeTypeMetaProperty

	// NodeTypeTemplateLiteral represents a template literal `hello ${world}`.
	NodeTypeTemplateLiteral
	// NodeTypeTaggedTemplateExpression represents a tagged template expression.
	NodeTypeTaggedTemplateExpression
	// NodeTypeTemplateElement represents an element in a template literal.
	NodeTypeTemplateElement

	// NodeTypeYieldExpression represents a yield expression.
	NodeTypeYieldExpression
	// NodeTypeAwaitExpression represents an await expression.
	NodeTypeAwaitExpression
	// NodeTypeChainExpression represents an optional chaining expression (obj?.prop).
	NodeTypeChainExpression
	// NodeTypeImportExpression represents a dynamic import expression import().
	NodeTypeImportExpression
	// NodeTypeSpreadElement represents a spread element (...x).
	NodeTypeSpreadElement

	// ==================== Statements ====================

	// NodeTypeBlockStatement represents a block of statements {}.
	NodeTypeBlockStatement
	// NodeTypeExpressionStatement represents an expression used as a statement.
	NodeTypeExpressionStatement
	// NodeTypeEmptyStatement represents an empty statement (;).
	NodeTypeEmptyStatement
	// NodeTypeDebuggerStatement represents a debugger statement.
	NodeTypeDebuggerStatement
	// NodeTypeReturnStatement represents a return statement.
	NodeTypeReturnStatement
	// NodeTypeBreakStatement represents a break statement.
	NodeTypeBreakStatement
	// NodeTypeContinueStatement represents a continue statement.
	NodeTypeContinueStatement
	// NodeTypeLabeledStatement represents a labeled statement.
	NodeTypeLabeledStatement

	// NodeTypeIfStatement represents an if statement.
	NodeTypeIfStatement
	// NodeTypeSwitchStatement represents a switch statement.
	NodeTypeSwitchStatement
	// NodeTypeSwitchCase represents a case or default clause in a switch statement.
	NodeTypeSwitchCase

	// NodeTypeWhileStatement represents a while loop.
	NodeTypeWhileStatement
	// NodeTypeDoWhileStatement represents a do-while loop.
	NodeTypeDoWhileStatement
	// NodeTypeForStatement represents a for loop.
	NodeTypeForStatement
	// NodeTypeForInStatement represents a for-in loop.
	NodeTypeForInStatement
	// NodeTypeForOfStatement represents a for-of loop.
	NodeTypeForOfStatement

	// NodeTypeThrowStatement represents a throw statement.
	NodeTypeThrowStatement
	// NodeTypeTryStatement represents a try-catch-finally statement.
	NodeTypeTryStatement
	// NodeTypeCatchClause represents a catch clause.
	NodeTypeCatchClause

	// NodeTypeWithStatement represents a with statement.
	NodeTypeWithStatement

	// ==================== Declarations ====================

	// NodeTypeVariableDeclaration represents a variable declaration (var, let, const).
	NodeTypeVariableDeclaration
	// NodeTypeVariableDeclarator represents a variable declarator.
	NodeTypeVariableDeclarator
	// NodeTypeFunctionDeclaration represents a function declaration.
	NodeTypeFunctionDeclaration

	// NodeTypeClassDeclaration represents a class declaration.
	NodeTypeClassDeclaration
	// NodeTypeClassBody represents the body of a class.
	NodeTypeClassBody
	// NodeTypeMethodDefinition represents a method in a class.
	NodeTypeMethodDefinition
	// NodeTypePropertyDefinition represents a property in a class.
	NodeTypePropertyDefinition
	// NodeTypeAccessorProperty represents an accessor property (getter/setter shorthand).
	NodeTypeAccessorProperty
	// NodeTypeStaticBlock represents a static initialization block in a class.
	NodeTypeStaticBlock

	// NodeTypeImportDeclaration represents an import declaration.
	NodeTypeImportDeclaration
	// NodeTypeImportSpecifier represents a named import specifier.
	NodeTypeImportSpecifier
	// NodeTypeImportDefaultSpecifier represents a default import specifier.
	NodeTypeImportDefaultSpecifier
	// NodeTypeImportNamespaceSpecifier represents a namespace import specifier (* as x).
	NodeTypeImportNamespaceSpecifier
	// NodeTypeImportAttribute represents an import attribute (with clause).
	NodeTypeImportAttribute

	// NodeTypeExportNamedDeclaration represents a named export declaration.
	NodeTypeExportNamedDeclaration
	// NodeTypeExportDefaultDeclaration represents a default export declaration.
	NodeTypeExportDefaultDeclaration
	// NodeTypeExportAllDeclaration represents an export * declaration.
	NodeTypeExportAllDeclaration
	// NodeTypeExportSpecifier represents an export specifier.
	NodeTypeExportSpecifier

	// ==================== Patterns (Destructuring) ====================

	// NodeTypeArrayPattern represents an array destructuring pattern.
	NodeTypeArrayPattern
	// NodeTypeObjectPattern represents an object destructuring pattern.
	NodeTypeObjectPattern
	// NodeTypeRestElement represents a rest element in destructuring (...rest).
	NodeTypeRestElement
	// NodeTypeAssignmentPattern represents a default value in destructuring (x = 1).
	NodeTypeAssignmentPattern

	// ==================== JSX (React) ====================

	// NodeTypeJSXElement represents a JSX element.
	NodeTypeJSXElement
	// NodeTypeJSXFragment represents a JSX fragment (<>...</>).
	NodeTypeJSXFragment
	// NodeTypeJSXOpeningElement represents a JSX opening element (<div>).
	NodeTypeJSXOpeningElement
	// NodeTypeJSXClosingElement represents a JSX closing element (</div>).
	NodeTypeJSXClosingElement
	// NodeTypeJSXOpeningFragment represents a JSX opening fragment (<>).
	NodeTypeJSXOpeningFragment
	// NodeTypeJSXClosingFragment represents a JSX closing fragment (</>).
	NodeTypeJSXClosingFragment

	// NodeTypeJSXAttribute represents a JSX attribute.
	NodeTypeJSXAttribute
	// NodeTypeJSXSpreadAttribute represents a JSX spread attribute ({...props}).
	NodeTypeJSXSpreadAttribute
	// NodeTypeJSXIdentifier represents a JSX identifier.
	NodeTypeJSXIdentifier
	// NodeTypeJSXNamespacedName represents a JSX namespaced name (ns:name).
	NodeTypeJSXNamespacedName
	// NodeTypeJSXMemberExpression represents a JSX member expression (obj.prop).
	NodeTypeJSXMemberExpression
	// NodeTypeJSXExpressionContainer represents a JSX expression container {expr}.
	NodeTypeJSXExpressionContainer
	// NodeTypeJSXEmptyExpression represents an empty JSX expression {}.
	NodeTypeJSXEmptyExpression
	// NodeTypeJSXText represents JSX text content.
	NodeTypeJSXText
	// NodeTypeJSXSpreadChild represents a JSX spread child ({...children}).
	NodeTypeJSXSpreadChild

	// ==================== Decorators ====================

	// NodeTypeDecorator represents a decorator (@decorator).
	NodeTypeDecorator

	// ==================== TypeScript Type Keywords ====================

	// NodeTypeTSAnyKeyword represents the 'any' type keyword.
	NodeTypeTSAnyKeyword
	// NodeTypeTSBigIntKeyword represents the 'bigint' type keyword.
	NodeTypeTSBigIntKeyword
	// NodeTypeTSBooleanKeyword represents the 'boolean' type keyword.
	NodeTypeTSBooleanKeyword
	// NodeTypeTSIntrinsicKeyword represents the 'intrinsic' type keyword.
	NodeTypeTSIntrinsicKeyword
	// NodeTypeTSNeverKeyword represents the 'never' type keyword.
	NodeTypeTSNeverKeyword
	// NodeTypeTSNullKeyword represents the 'null' type keyword.
	NodeTypeTSNullKeyword
	// NodeTypeTSNumberKeyword represents the 'number' type keyword.
	NodeTypeTSNumberKeyword
	// NodeTypeTSObjectKeyword represents the 'object' type keyword.
	NodeTypeTSObjectKeyword
	// NodeTypeTSStringKeyword represents the 'string' type keyword.
	NodeTypeTSStringKeyword
	// NodeTypeTSSymbolKeyword represents the 'symbol' type keyword.
	NodeTypeTSSymbolKeyword
	// NodeTypeTSUndefinedKeyword represents the 'undefined' type keyword.
	NodeTypeTSUndefinedKeyword
	// NodeTypeTSUnknownKeyword represents the 'unknown' type keyword.
	NodeTypeTSUnknownKeyword
	// NodeTypeTSVoidKeyword represents the 'void' type keyword.
	NodeTypeTSVoidKeyword

	// ==================== TypeScript Type Expressions ====================

	// NodeTypeTSArrayType represents an array type (T[]).
	NodeTypeTSArrayType
	// NodeTypeTSTupleType represents a tuple type ([T, U]).
	NodeTypeTSTupleType
	// NodeTypeTSUnionType represents a union type (T | U).
	NodeTypeTSUnionType
	// NodeTypeTSIntersectionType represents an intersection type (T & U).
	NodeTypeTSIntersectionType
	// NodeTypeTSConditionalType represents a conditional type (T extends U ? X : Y).
	NodeTypeTSConditionalType
	// NodeTypeTSInferType represents an infer type (infer T).
	NodeTypeTSInferType
	// NodeTypeTSTypeReference represents a type reference (Foo, Array<T>).
	NodeTypeTSTypeReference
	// NodeTypeTSTypeQuery represents a typeof type query (typeof x).
	NodeTypeTSTypeQuery
	// NodeTypeTSTypeLiteral represents a type literal ({a: string}).
	NodeTypeTSTypeLiteral
	// NodeTypeTSFunctionType represents a function type ((x: T) => U).
	NodeTypeTSFunctionType
	// NodeTypeTSConstructorType represents a constructor type (new () => T).
	NodeTypeTSConstructorType
	// NodeTypeTSMappedType represents a mapped type ({[K in T]: U}).
	NodeTypeTSMappedType
	// NodeTypeTSLiteralType represents a literal type ('foo', 42).
	NodeTypeTSLiteralType
	// NodeTypeTSIndexedAccessType represents an indexed access type (T[K]).
	NodeTypeTSIndexedAccessType
	// NodeTypeTSOptionalType represents an optional type (T?).
	NodeTypeTSOptionalType
	// NodeTypeTSRestType represents a rest type (...T[]).
	NodeTypeTSRestType
	// NodeTypeTSThisType represents the 'this' type.
	NodeTypeTSThisType
	// NodeTypeTSTypeOperator represents a type operator (keyof T, readonly T).
	NodeTypeTSTypeOperator
	// NodeTypeTSTemplateLiteralType represents a template literal type.
	NodeTypeTSTemplateLiteralType

	// ==================== TypeScript Type Declarations ====================

	// NodeTypeTSTypeAnnotation represents a type annotation (: T).
	NodeTypeTSTypeAnnotation
	// NodeTypeTSTypeAliasDeclaration represents a type alias declaration.
	NodeTypeTSTypeAliasDeclaration
	// NodeTypeTSInterfaceDeclaration represents an interface declaration.
	NodeTypeTSInterfaceDeclaration
	// NodeTypeTSInterfaceBody represents the body of an interface.
	NodeTypeTSInterfaceBody
	// NodeTypeTSInterfaceHeritage represents an interface extends clause.
	NodeTypeTSInterfaceHeritage
	// NodeTypeTSEnumDeclaration represents an enum declaration.
	NodeTypeTSEnumDeclaration
	// NodeTypeTSEnumBody represents the body of an enum.
	NodeTypeTSEnumBody
	// NodeTypeTSEnumMember represents a member of an enum.
	NodeTypeTSEnumMember
	// NodeTypeTSModuleDeclaration represents a module or namespace declaration.
	NodeTypeTSModuleDeclaration
	// NodeTypeTSModuleBlock represents the body of a module.
	NodeTypeTSModuleBlock

	// ==================== TypeScript Type Components ====================

	// NodeTypeTSTypeParameter represents a type parameter (<T>).
	NodeTypeTSTypeParameter
	// NodeTypeTSTypeParameterDeclaration represents a type parameter declaration.
	NodeTypeTSTypeParameterDeclaration
	// NodeTypeTSTypeParameterInstantiation represents a type parameter instantiation.
	NodeTypeTSTypeParameterInstantiation
	// NodeTypeTSCallSignatureDeclaration represents a call signature.
	NodeTypeTSCallSignatureDeclaration
	// NodeTypeTSConstructSignatureDeclaration represents a construct signature.
	NodeTypeTSConstructSignatureDeclaration
	// NodeTypeTSPropertySignature represents a property signature in a type.
	NodeTypeTSPropertySignature
	// NodeTypeTSMethodSignature represents a method signature in a type.
	NodeTypeTSMethodSignature
	// NodeTypeTSIndexSignature represents an index signature.
	NodeTypeTSIndexSignature
	// NodeTypeTSNamedTupleMember represents a named tuple member.
	NodeTypeTSNamedTupleMember

	// ==================== TypeScript Type Assertions & Expressions ====================

	// NodeTypeTSAsExpression represents a type assertion using 'as' (x as T).
	NodeTypeTSAsExpression
	// NodeTypeTSTypeAssertion represents a type assertion using angle brackets (<T>x).
	NodeTypeTSTypeAssertion
	// NodeTypeTSNonNullExpression represents a non-null assertion (x!).
	NodeTypeTSNonNullExpression
	// NodeTypeTSSatisfiesExpression represents a satisfies expression (x satisfies T).
	NodeTypeTSSatisfiesExpression
	// NodeTypeTSInstantiationExpression represents a type instantiation (Foo<T>).
	NodeTypeTSInstantiationExpression

	// ==================== TypeScript Type Predicates ====================

	// NodeTypeTSTypePredicate represents a type predicate (x is T).
	NodeTypeTSTypePredicate

	// ==================== TypeScript Modifier Keywords ====================

	// NodeTypeTSAbstractKeyword represents the 'abstract' modifier keyword.
	NodeTypeTSAbstractKeyword
	// NodeTypeTSAsyncKeyword represents the 'async' modifier keyword.
	NodeTypeTSAsyncKeyword
	// NodeTypeTSDeclareKeyword represents the 'declare' modifier keyword.
	NodeTypeTSDeclareKeyword
	// NodeTypeTSExportKeyword represents the 'export' modifier keyword.
	NodeTypeTSExportKeyword
	// NodeTypeTSPrivateKeyword represents the 'private' modifier keyword.
	NodeTypeTSPrivateKeyword
	// NodeTypeTSProtectedKeyword represents the 'protected' modifier keyword.
	NodeTypeTSProtectedKeyword
	// NodeTypeTSPublicKeyword represents the 'public' modifier keyword.
	NodeTypeTSPublicKeyword
	// NodeTypeTSReadonlyKeyword represents the 'readonly' modifier keyword.
	NodeTypeTSReadonlyKeyword
	// NodeTypeTSStaticKeyword represents the 'static' modifier keyword.
	NodeTypeTSStaticKeyword

	// ==================== TypeScript Abstract Members ====================

	// NodeTypeTSAbstractAccessorProperty represents an abstract accessor property.
	NodeTypeTSAbstractAccessorProperty
	// NodeTypeTSAbstractMethodDefinition represents an abstract method.
	NodeTypeTSAbstractMethodDefinition
	// NodeTypeTSAbstractPropertyDefinition represents an abstract property.
	NodeTypeTSAbstractPropertyDefinition

	// ==================== TypeScript Import/Export ====================

	// NodeTypeTSImportEqualsDeclaration represents an import = declaration.
	NodeTypeTSImportEqualsDeclaration
	// NodeTypeTSImportType represents an import type (import('module').Type).
	NodeTypeTSImportType
	// NodeTypeTSExternalModuleReference represents an external module reference.
	NodeTypeTSExternalModuleReference
	// NodeTypeTSExportAssignment represents an export = statement.
	NodeTypeTSExportAssignment
	// NodeTypeTSNamespaceExportDeclaration represents a namespace export declaration.
	NodeTypeTSNamespaceExportDeclaration

	// ==================== TypeScript Other ====================

	// NodeTypeTSQualifiedName represents a qualified name (A.B.C).
	NodeTypeTSQualifiedName
	// NodeTypeTSParameterProperty represents a parameter property in a constructor.
	NodeTypeTSParameterProperty
	// NodeTypeTSDeclareFunction represents a declare function statement.
	NodeTypeTSDeclareFunction
	// NodeTypeTSEmptyBodyFunctionExpression represents a function with no body.
	NodeTypeTSEmptyBodyFunctionExpression
	// NodeTypeTSClassImplements represents a class implements clause.
	NodeTypeTSClassImplements
)

//nolint:gochecknoglobals // Map is used for efficient string conversion
var nodeTypeNames = map[NodeType]string{
	NodeTypeUnknown:                         "Unknown",
	NodeTypeProgram:                         "Program",
	NodeTypeIdentifier:                      "Identifier",
	NodeTypePrivateIdentifier:               "PrivateIdentifier",
	NodeTypeLiteral:                         "Literal",
	NodeTypeThisExpression:                  "ThisExpression",
	NodeTypeSuper:                           "Super",
	NodeTypeArrayExpression:                 "ArrayExpression",
	NodeTypeObjectExpression:                "ObjectExpression",
	NodeTypeProperty:                        "Property",
	NodeTypeFunctionExpression:              "FunctionExpression",
	NodeTypeArrowFunctionExpression:         "ArrowFunctionExpression",
	NodeTypeClassExpression:                 "ClassExpression",
	NodeTypeUnaryExpression:                 "UnaryExpression",
	NodeTypeUpdateExpression:                "UpdateExpression",
	NodeTypeBinaryExpression:                "BinaryExpression",
	NodeTypeLogicalExpression:               "LogicalExpression",
	NodeTypeAssignmentExpression:            "AssignmentExpression",
	NodeTypeConditionalExpression:           "ConditionalExpression",
	NodeTypeSequenceExpression:              "SequenceExpression",
	NodeTypeMemberExpression:                "MemberExpression",
	NodeTypeCallExpression:                  "CallExpression",
	NodeTypeNewExpression:                   "NewExpression",
	NodeTypeMetaProperty:                    "MetaProperty",
	NodeTypeTemplateLiteral:                 "TemplateLiteral",
	NodeTypeTaggedTemplateExpression:        "TaggedTemplateExpression",
	NodeTypeTemplateElement:                 "TemplateElement",
	NodeTypeYieldExpression:                 "YieldExpression",
	NodeTypeAwaitExpression:                 "AwaitExpression",
	NodeTypeChainExpression:                 "ChainExpression",
	NodeTypeImportExpression:                "ImportExpression",
	NodeTypeSpreadElement:                   "SpreadElement",
	NodeTypeBlockStatement:                  "BlockStatement",
	NodeTypeExpressionStatement:             "ExpressionStatement",
	NodeTypeEmptyStatement:                  "EmptyStatement",
	NodeTypeDebuggerStatement:               "DebuggerStatement",
	NodeTypeReturnStatement:                 "ReturnStatement",
	NodeTypeBreakStatement:                  "BreakStatement",
	NodeTypeContinueStatement:               "ContinueStatement",
	NodeTypeLabeledStatement:                "LabeledStatement",
	NodeTypeIfStatement:                     "IfStatement",
	NodeTypeSwitchStatement:                 "SwitchStatement",
	NodeTypeSwitchCase:                      "SwitchCase",
	NodeTypeWhileStatement:                  "WhileStatement",
	NodeTypeDoWhileStatement:                "DoWhileStatement",
	NodeTypeForStatement:                    "ForStatement",
	NodeTypeForInStatement:                  "ForInStatement",
	NodeTypeForOfStatement:                  "ForOfStatement",
	NodeTypeThrowStatement:                  "ThrowStatement",
	NodeTypeTryStatement:                    "TryStatement",
	NodeTypeCatchClause:                     "CatchClause",
	NodeTypeWithStatement:                   "WithStatement",
	NodeTypeVariableDeclaration:             "VariableDeclaration",
	NodeTypeVariableDeclarator:              "VariableDeclarator",
	NodeTypeFunctionDeclaration:             "FunctionDeclaration",
	NodeTypeClassDeclaration:                "ClassDeclaration",
	NodeTypeClassBody:                       "ClassBody",
	NodeTypeMethodDefinition:                "MethodDefinition",
	NodeTypePropertyDefinition:              "PropertyDefinition",
	NodeTypeAccessorProperty:                "AccessorProperty",
	NodeTypeStaticBlock:                     "StaticBlock",
	NodeTypeImportDeclaration:               "ImportDeclaration",
	NodeTypeImportSpecifier:                 "ImportSpecifier",
	NodeTypeImportDefaultSpecifier:          "ImportDefaultSpecifier",
	NodeTypeImportNamespaceSpecifier:        "ImportNamespaceSpecifier",
	NodeTypeImportAttribute:                 "ImportAttribute",
	NodeTypeExportNamedDeclaration:          "ExportNamedDeclaration",
	NodeTypeExportDefaultDeclaration:        "ExportDefaultDeclaration",
	NodeTypeExportAllDeclaration:            "ExportAllDeclaration",
	NodeTypeExportSpecifier:                 "ExportSpecifier",
	NodeTypeArrayPattern:                    "ArrayPattern",
	NodeTypeObjectPattern:                   "ObjectPattern",
	NodeTypeRestElement:                     "RestElement",
	NodeTypeAssignmentPattern:               "AssignmentPattern",
	NodeTypeJSXElement:                      "JSXElement",
	NodeTypeJSXFragment:                     "JSXFragment",
	NodeTypeJSXOpeningElement:               "JSXOpeningElement",
	NodeTypeJSXClosingElement:               "JSXClosingElement",
	NodeTypeJSXOpeningFragment:              "JSXOpeningFragment",
	NodeTypeJSXClosingFragment:              "JSXClosingFragment",
	NodeTypeJSXAttribute:                    "JSXAttribute",
	NodeTypeJSXSpreadAttribute:              "JSXSpreadAttribute",
	NodeTypeJSXIdentifier:                   "JSXIdentifier",
	NodeTypeJSXNamespacedName:               "JSXNamespacedName",
	NodeTypeJSXMemberExpression:             "JSXMemberExpression",
	NodeTypeJSXExpressionContainer:          "JSXExpressionContainer",
	NodeTypeJSXEmptyExpression:              "JSXEmptyExpression",
	NodeTypeJSXText:                         "JSXText",
	NodeTypeJSXSpreadChild:                  "JSXSpreadChild",
	NodeTypeDecorator:                       "Decorator",
	NodeTypeTSAnyKeyword:                    "TSAnyKeyword",
	NodeTypeTSBigIntKeyword:                 "TSBigIntKeyword",
	NodeTypeTSBooleanKeyword:                "TSBooleanKeyword",
	NodeTypeTSIntrinsicKeyword:              "TSIntrinsicKeyword",
	NodeTypeTSNeverKeyword:                  "TSNeverKeyword",
	NodeTypeTSNullKeyword:                   "TSNullKeyword",
	NodeTypeTSNumberKeyword:                 "TSNumberKeyword",
	NodeTypeTSObjectKeyword:                 "TSObjectKeyword",
	NodeTypeTSStringKeyword:                 "TSStringKeyword",
	NodeTypeTSSymbolKeyword:                 "TSSymbolKeyword",
	NodeTypeTSUndefinedKeyword:              "TSUndefinedKeyword",
	NodeTypeTSUnknownKeyword:                "TSUnknownKeyword",
	NodeTypeTSVoidKeyword:                   "TSVoidKeyword",
	NodeTypeTSArrayType:                     "TSArrayType",
	NodeTypeTSTupleType:                     "TSTupleType",
	NodeTypeTSUnionType:                     "TSUnionType",
	NodeTypeTSIntersectionType:              "TSIntersectionType",
	NodeTypeTSConditionalType:               "TSConditionalType",
	NodeTypeTSInferType:                     "TSInferType",
	NodeTypeTSTypeReference:                 "TSTypeReference",
	NodeTypeTSTypeQuery:                     "TSTypeQuery",
	NodeTypeTSTypeLiteral:                   "TSTypeLiteral",
	NodeTypeTSFunctionType:                  "TSFunctionType",
	NodeTypeTSConstructorType:               "TSConstructorType",
	NodeTypeTSMappedType:                    "TSMappedType",
	NodeTypeTSLiteralType:                   "TSLiteralType",
	NodeTypeTSIndexedAccessType:             "TSIndexedAccessType",
	NodeTypeTSOptionalType:                  "TSOptionalType",
	NodeTypeTSRestType:                      "TSRestType",
	NodeTypeTSThisType:                      "TSThisType",
	NodeTypeTSTypeOperator:                  "TSTypeOperator",
	NodeTypeTSTemplateLiteralType:           "TSTemplateLiteralType",
	NodeTypeTSTypeAnnotation:                "TSTypeAnnotation",
	NodeTypeTSTypeAliasDeclaration:          "TSTypeAliasDeclaration",
	NodeTypeTSInterfaceDeclaration:          "TSInterfaceDeclaration",
	NodeTypeTSInterfaceBody:                 "TSInterfaceBody",
	NodeTypeTSInterfaceHeritage:             "TSInterfaceHeritage",
	NodeTypeTSEnumDeclaration:               "TSEnumDeclaration",
	NodeTypeTSEnumBody:                      "TSEnumBody",
	NodeTypeTSEnumMember:                    "TSEnumMember",
	NodeTypeTSModuleDeclaration:             "TSModuleDeclaration",
	NodeTypeTSModuleBlock:                   "TSModuleBlock",
	NodeTypeTSTypeParameter:                 "TSTypeParameter",
	NodeTypeTSTypeParameterDeclaration:      "TSTypeParameterDeclaration",
	NodeTypeTSTypeParameterInstantiation:    "TSTypeParameterInstantiation",
	NodeTypeTSCallSignatureDeclaration:      "TSCallSignatureDeclaration",
	NodeTypeTSConstructSignatureDeclaration: "TSConstructSignatureDeclaration",
	NodeTypeTSPropertySignature:             "TSPropertySignature",
	NodeTypeTSMethodSignature:               "TSMethodSignature",
	NodeTypeTSIndexSignature:                "TSIndexSignature",
	NodeTypeTSNamedTupleMember:              "TSNamedTupleMember",
	NodeTypeTSAsExpression:                  "TSAsExpression",
	NodeTypeTSTypeAssertion:                 "TSTypeAssertion",
	NodeTypeTSNonNullExpression:             "TSNonNullExpression",
	NodeTypeTSSatisfiesExpression:           "TSSatisfiesExpression",
	NodeTypeTSInstantiationExpression:       "TSInstantiationExpression",
	NodeTypeTSTypePredicate:                 "TSTypePredicate",
	NodeTypeTSAbstractKeyword:               "TSAbstractKeyword",
	NodeTypeTSAsyncKeyword:                  "TSAsyncKeyword",
	NodeTypeTSDeclareKeyword:                "TSDeclareKeyword",
	NodeTypeTSExportKeyword:                 "TSExportKeyword",
	NodeTypeTSPrivateKeyword:                "TSPrivateKeyword",
	NodeTypeTSProtectedKeyword:              "TSProtectedKeyword",
	NodeTypeTSPublicKeyword:                 "TSPublicKeyword",
	NodeTypeTSReadonlyKeyword:               "TSReadonlyKeyword",
	NodeTypeTSStaticKeyword:                 "TSStaticKeyword",
	NodeTypeTSAbstractAccessorProperty:      "TSAbstractAccessorProperty",
	NodeTypeTSAbstractMethodDefinition:      "TSAbstractMethodDefinition",
	NodeTypeTSAbstractPropertyDefinition:    "TSAbstractPropertyDefinition",
	NodeTypeTSImportEqualsDeclaration:       "TSImportEqualsDeclaration",
	NodeTypeTSImportType:                    "TSImportType",
	NodeTypeTSExternalModuleReference:       "TSExternalModuleReference",
	NodeTypeTSExportAssignment:              "TSExportAssignment",
	NodeTypeTSNamespaceExportDeclaration:    "TSNamespaceExportDeclaration",
	NodeTypeTSQualifiedName:                 "TSQualifiedName",
	NodeTypeTSParameterProperty:             "TSParameterProperty",
	NodeTypeTSDeclareFunction:               "TSDeclareFunction",
	NodeTypeTSEmptyBodyFunctionExpression:   "TSEmptyBodyFunctionExpression",
	NodeTypeTSClassImplements:               "TSClassImplements",
}

// String returns the string representation of a NodeType.
func (nt NodeType) String() string {
	if name, ok := nodeTypeNames[nt]; ok {
		return name
	}
	return "Unknown"
}
