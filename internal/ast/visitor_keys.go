package ast

// VisitorKeys defines the child properties that should be traversed for each AST node type.
// Keys are ordered based on their appearance in source code, not alphabetically.
// This is used by AST traversal utilities to visit all relevant child nodes.
//
// Based on: https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/visitor-keys/src/visitor-keys.ts
var VisitorKeys = map[string][]string{
	// ==================== Program & Core ====================

	"Program": {"body"},

	// ==================== Identifiers & Literals ====================

	"Identifier":        {}, // No child nodes
	"PrivateIdentifier": {}, // No child nodes
	"Literal":           {}, // No child nodes

	// ==================== Expressions ====================

	"ThisExpression": {}, // No child nodes
	"Super":          {}, // No child nodes

	"ArrayExpression":  {"elements"},
	"ObjectExpression": {"properties"},
	"Property":         {"key", "value"},

	"FunctionExpression":      {"id", "typeParameters", "params", "returnType", "body"},
	"ArrowFunctionExpression": {"typeParameters", "params", "returnType", "body"},

	"ClassExpression": {"decorators", "id", "typeParameters", "superClass", "superTypeArguments", "implements", "body"},

	"UnaryExpression":      {"argument"},
	"UpdateExpression":     {"argument"},
	"BinaryExpression":     {"left", "right"},
	"LogicalExpression":    {"left", "right"},
	"AssignmentExpression": {"left", "right"},

	"ConditionalExpression": {"test", "consequent", "alternate"},
	"SequenceExpression":    {"expressions"},

	"MemberExpression": {"object", "property"},
	"CallExpression":   {"callee", "typeArguments", "arguments"},
	"NewExpression":    {"callee", "typeArguments", "arguments"},
	"MetaProperty":     {"meta", "property"},

	"TemplateLiteral":            {"quasis", "expressions"},
	"TaggedTemplateExpression":   {"tag", "typeArguments", "quasi"},
	"TemplateElement":            {}, // No child nodes
	"YieldExpression":            {"argument"},
	"AwaitExpression":            {"argument"},
	"ChainExpression":            {"expression"},
	"ImportExpression":           {"source"},
	"SpreadElement":              {"argument"},

	// ==================== Statements ====================

	"BlockStatement":      {"body"},
	"ExpressionStatement": {"expression"},
	"EmptyStatement":      {}, // No child nodes

	"IfStatement":      {"test", "consequent", "alternate"},
	"SwitchStatement":  {"discriminant", "cases"},
	"SwitchCase":       {"test", "consequent"},
	"LabeledStatement": {"label", "body"},

	"WhileStatement":   {"test", "body"},
	"DoWhileStatement": {"body", "test"},
	"ForStatement":     {"init", "test", "update", "body"},
	"ForInStatement":   {"left", "right", "body"},
	"ForOfStatement":   {"left", "right", "body"},

	"TryStatement":     {"block", "handler", "finalizer"},
	"CatchClause":      {"param", "body"},
	"ThrowStatement":   {"argument"},
	"ReturnStatement":  {"argument"},
	"BreakStatement":   {"label"},
	"ContinueStatement": {"label"},

	"DebuggerStatement": {}, // No child nodes
	"WithStatement":     {"object", "body"},

	// ==================== Declarations ====================

	"VariableDeclaration": {"declarations"},
	"VariableDeclarator":  {"id", "init"},

	"FunctionDeclaration": {"id", "typeParameters", "params", "returnType", "body"},

	"ClassDeclaration": {"decorators", "id", "typeParameters", "superClass", "superTypeArguments", "implements", "body"},
	"ClassBody":        {"body"},
	"MethodDefinition": {"decorators", "key", "value"},
	"PropertyDefinition": {"decorators", "key", "typeAnnotation", "value"},
	"AccessorProperty": {"decorators", "key", "typeAnnotation", "value"},
	"StaticBlock":      {"body"},

	"ImportDeclaration":          {"specifiers", "source", "attributes"},
	"ImportSpecifier":            {"imported", "local"},
	"ImportDefaultSpecifier":     {"local"},
	"ImportNamespaceSpecifier":   {"local"},
	"ImportAttribute":            {"key", "value"},

	"ExportNamedDeclaration":    {"declaration", "specifiers", "source"},
	"ExportDefaultDeclaration":  {"declaration"},
	"ExportAllDeclaration":      {"exported", "source"},
	"ExportSpecifier":           {"exported", "local"},

	// ==================== Patterns ====================

	"ArrayPattern":       {"elements"},
	"ObjectPattern":      {"properties"},
	"RestElement":        {"argument"},
	"AssignmentPattern":  {"left", "right"},

	// ==================== Other ====================

	"Decorator": {"expression"},

	// ==================== JSX ====================

	"JSXElement":         {"openingElement", "children", "closingElement"},
	"JSXFragment":        {"openingFragment", "children", "closingFragment"},
	"JSXOpeningElement":  {"name", "typeArguments", "attributes"},
	"JSXClosingElement":  {"name"},
	"JSXOpeningFragment": {}, // No child nodes
	"JSXClosingFragment": {}, // No child nodes

	"JSXAttribute":       {"name", "value"},
	"JSXSpreadAttribute": {"argument"},

	"JSXIdentifier":        {}, // No child nodes
	"JSXNamespacedName":    {"namespace", "name"},
	"JSXMemberExpression":  {"object", "property"},

	"JSXText":                {}, // No child nodes
	"JSXExpressionContainer": {"expression"},
	"JSXEmptyExpression":     {}, // No child nodes
	"JSXSpreadChild":         {"expression"},

	// ==================== TypeScript Types ====================

	// Type Keywords
	"TSAnyKeyword":       {}, // No child nodes
	"TSUnknownKeyword":   {}, // No child nodes
	"TSNeverKeyword":     {}, // No child nodes
	"TSStringKeyword":    {}, // No child nodes
	"TSNumberKeyword":    {}, // No child nodes
	"TSBooleanKeyword":   {}, // No child nodes
	"TSBigIntKeyword":    {}, // No child nodes
	"TSSymbolKeyword":    {}, // No child nodes
	"TSObjectKeyword":    {}, // No child nodes
	"TSVoidKeyword":      {}, // No child nodes
	"TSUndefinedKeyword": {}, // No child nodes
	"TSNullKeyword":      {}, // No child nodes
	"TSIntrinsicKeyword": {}, // No child nodes

	// Type Expressions
	"TSArrayType":           {"elementType"},
	"TSTupleType":           {"elementTypes"},
	"TSUnionType":           {"types"},
	"TSIntersectionType":    {"types"},
	"TSConditionalType":     {"checkType", "extendsType", "trueType", "falseType"},
	"TSInferType":           {"typeParameter"},
	"TSMappedType":          {"typeParameter", "nameType", "typeAnnotation"},
	"TSIndexedAccessType":   {"objectType", "indexType"},
	"TSTemplateLiteralType": {"quasis", "types"},
	"TSTypeReference":       {"typeName", "typeArguments"},
	"TSTypeQuery":           {"exprName", "typeArguments"},
	"TSTypeLiteral":         {"members"},
	"TSImportType":          {"argument", "qualifier", "typeArguments"},
	"TSFunctionType":        {"typeParameters", "params", "returnType"},
	"TSConstructorType":     {"typeParameters", "params", "returnType"},
	"TSLiteralType":         {"literal"},
	"TSOptionalType":        {"typeAnnotation"},
	"TSRestType":            {"typeAnnotation"},
	"TSThisType":            {}, // No child nodes
	"TSTypeOperator":        {"typeAnnotation"},
	"TSParenthesizedType":   {"typeAnnotation"},

	// Type Declarations
	"TSTypeAnnotation":        {"typeAnnotation"},
	"TSTypeAliasDeclaration":  {"id", "typeParameters", "typeAnnotation"},
	"TSInterfaceDeclaration":  {"id", "typeParameters", "extends", "body"},
	"TSInterfaceBody":         {"body"},
	"TSInterfaceHeritage":     {"expression", "typeArguments"},
	"TSEnumDeclaration":       {"id", "members"},
	"TSEnumMember":            {"id", "initializer"},
	"TSModuleDeclaration":     {"id", "body"},
	"TSModuleBlock":           {"body"},

	// Type Components
	"TSTypeParameter":                {"name", "constraint", "default"},
	"TSTypeParameterDeclaration":     {"params"},
	"TSTypeParameterInstantiation":   {"params"},
	"TSCallSignatureDeclaration":     {"typeParameters", "params", "returnType"},
	"TSConstructSignatureDeclaration": {"typeParameters", "params", "returnType"},
	"TSPropertySignature":            {"key", "typeAnnotation"},
	"TSMethodSignature":              {"key", "typeParameters", "params", "returnType"},
	"TSIndexSignature":               {"parameters", "typeAnnotation"},
	"TSNamedTupleMember":             {"label", "elementType"},

	// Type Assertions
	"TSAsExpression":            {"expression", "typeAnnotation"},
	"TSTypeAssertion":           {"typeAnnotation", "expression"},
	"TSNonNullExpression":       {"expression"},
	"TSSatisfiesExpression":     {"expression", "typeAnnotation"},
	"TSInstantiationExpression": {"expression", "typeArguments"},

	// Type Predicates
	"TSTypePredicate": {"parameterName", "typeAnnotation"},

	// Abstract Members
	"TSAbstractAccessorProperty":   {"decorators", "key", "typeAnnotation", "value"},
	"TSAbstractMethodDefinition":   {"decorators", "key", "value"},
	"TSAbstractPropertyDefinition": {"decorators", "key", "typeAnnotation", "value"},

	// Import/Export
	"TSImportEqualsDeclaration":     {"id", "moduleReference"},
	"TSExternalModuleReference":     {"expression"},
	"TSExportAssignment":            {"expression"},
	"TSNamespaceExportDeclaration":  {"id"},

	// Other
	"TSQualifiedName":                  {"left", "right"},
	"TSParameterProperty":              {"decorators", "parameter"},
	"TSDeclareFunction":                {"id", "typeParameters", "params", "returnType", "body"},
	"TSEmptyBodyFunctionExpression":    {"id", "typeParameters", "params", "returnType"},
	"TSClassImplements":                {"expression", "typeArguments"},
}

// GetVisitorKeys returns the visitor keys for a given node type.
// If the node type is not found, it returns an empty slice.
func GetVisitorKeys(nodeType string) []string {
	if keys, ok := VisitorKeys[nodeType]; ok {
		return keys
	}
	return []string{}
}

// HasVisitorKeys returns true if the node type has visitor keys defined.
func HasVisitorKeys(nodeType string) bool {
	_, ok := VisitorKeys[nodeType]
	return ok
}
