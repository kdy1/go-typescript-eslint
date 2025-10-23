package ast

// guards.go provides type guards and assertions for AST nodes.
// These functions allow safe type checking and conversion of nodes.

// IsExpression returns true if the node is an Expression.
func IsExpression(node Node) bool {
	_, ok := node.(Expression)
	return ok
}

// IsStatement returns true if the node is a Statement.
func IsStatement(node Node) bool {
	_, ok := node.(Statement)
	return ok
}

// IsPattern returns true if the node is a Pattern.
func IsPattern(node Node) bool {
	_, ok := node.(Pattern)
	return ok
}

// IsDeclaration returns true if the node is a Declaration.
func IsDeclaration(node Node) bool {
	_, ok := node.(Declaration)
	return ok
}

// IsTSNode returns true if the node is a TypeScript-specific node.
func IsTSNode(node Node) bool {
	_, ok := node.(TSNode)
	return ok
}

// AsExpression attempts to convert the node to an Expression.
// Returns the expression and true if successful, nil and false otherwise.
//
//nolint:ireturn // Interface types are intentional for AST node conversion
func AsExpression(node Node) (Expression, bool) {
	expr, ok := node.(Expression)
	return expr, ok
}

// AsStatement attempts to convert the node to a Statement.
// Returns the statement and true if successful, nil and false otherwise.
//
//nolint:ireturn // Interface types are intentional for AST node conversion
func AsStatement(node Node) (Statement, bool) {
	stmt, ok := node.(Statement)
	return stmt, ok
}

// AsPattern attempts to convert the node to a Pattern.
// Returns the pattern and true if successful, nil and false otherwise.
//
//nolint:ireturn // Interface types are intentional for AST node conversion
func AsPattern(node Node) (Pattern, bool) {
	pat, ok := node.(Pattern)
	return pat, ok
}

// AsDeclaration attempts to convert the node to a Declaration.
// Returns the declaration and true if successful, nil and false otherwise.
//
//nolint:ireturn // Interface types are intentional for AST node conversion
func AsDeclaration(node Node) (Declaration, bool) {
	decl, ok := node.(Declaration)
	return decl, ok
}

// AsTypeScriptNode attempts to convert the node to a TSNode.
// Returns the TS node and true if successful, nil and false otherwise.
//
//nolint:ireturn // Interface types are intentional for AST node conversion
func AsTypeScriptNode(node Node) (TSNode, bool) {
	tsNode, ok := node.(TSNode)
	return tsNode, ok
}

// Type-specific guards for common node types

// IsIdentifier returns true if the node is an Identifier.
func IsIdentifier(node Node) bool {
	return node != nil && node.Type() == "Identifier"
}

// IsLiteral returns true if the node is a Literal.
func IsLiteral(node Node) bool {
	return node != nil && node.Type() == "Literal"
}

// IsMemberExpression returns true if the node is a MemberExpression.
func IsMemberExpression(node Node) bool {
	return node != nil && node.Type() == "MemberExpression"
}

// IsCallExpression returns true if the node is a CallExpression.
func IsCallExpression(node Node) bool {
	return node != nil && node.Type() == "CallExpression"
}

// IsFunctionExpression returns true if the node is a FunctionExpression.
func IsFunctionExpression(node Node) bool {
	return node != nil && node.Type() == "FunctionExpression"
}

// IsArrowFunctionExpression returns true if the node is an ArrowFunctionExpression.
func IsArrowFunctionExpression(node Node) bool {
	return node != nil && node.Type() == "ArrowFunctionExpression"
}

// IsClassExpression returns true if the node is a ClassExpression.
func IsClassExpression(node Node) bool {
	return node != nil && node.Type() == "ClassExpression"
}

// IsObjectExpression returns true if the node is an ObjectExpression.
func IsObjectExpression(node Node) bool {
	return node != nil && node.Type() == "ObjectExpression"
}

// IsArrayExpression returns true if the node is an ArrayExpression.
func IsArrayExpression(node Node) bool {
	return node != nil && node.Type() == "ArrayExpression"
}

// IsBlockStatement returns true if the node is a BlockStatement.
func IsBlockStatement(node Node) bool {
	return node != nil && node.Type() == "BlockStatement"
}

// IsVariableDeclaration returns true if the node is a VariableDeclaration.
func IsVariableDeclaration(node Node) bool {
	return node != nil && node.Type() == "VariableDeclaration"
}

// IsFunctionDeclaration returns true if the node is a FunctionDeclaration.
func IsFunctionDeclaration(node Node) bool {
	return node != nil && node.Type() == "FunctionDeclaration"
}

// IsClassDeclaration returns true if the node is a ClassDeclaration.
func IsClassDeclaration(node Node) bool {
	return node != nil && node.Type() == "ClassDeclaration"
}

// IsIfStatement returns true if the node is an IfStatement.
func IsIfStatement(node Node) bool {
	return node != nil && node.Type() == "IfStatement"
}

// IsForStatement returns true if the node is a ForStatement.
func IsForStatement(node Node) bool {
	return node != nil && node.Type() == "ForStatement"
}

// IsWhileStatement returns true if the node is a WhileStatement.
func IsWhileStatement(node Node) bool {
	return node != nil && node.Type() == "WhileStatement"
}

// IsReturnStatement returns true if the node is a ReturnStatement.
func IsReturnStatement(node Node) bool {
	return node != nil && node.Type() == "ReturnStatement"
}

// IsThrowStatement returns true if the node is a ThrowStatement.
func IsThrowStatement(node Node) bool {
	return node != nil && node.Type() == "ThrowStatement"
}

// IsTryStatement returns true if the node is a TryStatement.
func IsTryStatement(node Node) bool {
	return node != nil && node.Type() == "TryStatement"
}

// IsFunction returns true if the node is any kind of function
// (FunctionExpression, ArrowFunctionExpression, FunctionDeclaration).
func IsFunction(node Node) bool {
	if node == nil {
		return false
	}
	t := node.Type()
	return t == "FunctionExpression" || t == "ArrowFunctionExpression" || t == "FunctionDeclaration"
}

// IsClass returns true if the node is any kind of class
// (ClassExpression, ClassDeclaration).
func IsClass(node Node) bool {
	if node == nil {
		return false
	}
	t := node.Type()
	return t == "ClassExpression" || t == "ClassDeclaration"
}

// IsLoop returns true if the node is any kind of loop
// (ForStatement, ForInStatement, ForOfStatement, WhileStatement, DoWhileStatement).
func IsLoop(node Node) bool {
	if node == nil {
		return false
	}
	t := node.Type()
	return t == "ForStatement" || t == "ForInStatement" || t == "ForOfStatement" ||
		t == "WhileStatement" || t == "DoWhileStatement"
}

// IsJSX returns true if the node is a JSX node.
func IsJSX(node Node) bool {
	if node == nil {
		return false
	}
	t := node.Type()
	return len(t) >= 3 && t[:3] == "JSX"
}

// IsTypeScript returns true if the node is a TypeScript-specific node.
func IsTypeScript(node Node) bool {
	if node == nil {
		return false
	}
	t := node.Type()
	return len(t) >= 2 && t[:2] == "TS"
}

// IsImport returns true if the node is an import-related node.
func IsImport(node Node) bool {
	if node == nil {
		return false
	}
	t := node.Type()
	return t == "ImportDeclaration" || t == "ImportSpecifier" ||
		t == "ImportDefaultSpecifier" || t == "ImportNamespaceSpecifier" ||
		t == "ImportExpression" || t == "TSImportEqualsDeclaration"
}

// IsExport returns true if the node is an export-related node.
func IsExport(node Node) bool {
	if node == nil {
		return false
	}
	t := node.Type()
	return t == "ExportNamedDeclaration" || t == "ExportDefaultDeclaration" ||
		t == "ExportAllDeclaration" || t == "ExportSpecifier" ||
		t == "TSExportAssignment" || t == "TSNamespaceExportDeclaration"
}

// IsModuleDeclaration returns true if the node is a module-level declaration
// (import or export).
func IsModuleDeclaration(node Node) bool {
	return IsImport(node) || IsExport(node)
}
