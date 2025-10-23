package ast

import "testing"

// helper function to reduce duplication in tests
func runNodeTypeTests(t *testing.T, testName string, checkFunc func(Node) bool, tests []struct {
	name     string
	node     Node
	expected bool
}) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkFunc(tt.node)
			if result != tt.expected {
				t.Errorf("%s: expected %v, got %v", testName, tt.expected, result)
			}
		})
	}
}

func TestIsExpression(t *testing.T) {
	id := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "test"}
	if !IsExpression(id) {
		t.Error("Identifier should be an Expression")
	}

	stmt := &BlockStatement{BaseNode: BaseNode{NodeType: "BlockStatement"}}
	if IsExpression(stmt) {
		t.Error("BlockStatement should not be an Expression")
	}
}

func TestIsStatement(t *testing.T) {
	stmt := &BlockStatement{BaseNode: BaseNode{NodeType: "BlockStatement"}}
	if !IsStatement(stmt) {
		t.Error("BlockStatement should be a Statement")
	}

	expr := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "test"}
	if IsStatement(expr) {
		t.Error("Identifier should not be a Statement")
	}
}

func TestIsIdentifier(t *testing.T) {
	id := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "test"}
	if !IsIdentifier(id) {
		t.Error("Identifier node should return true")
	}

	lit := &Literal{BaseNode: BaseNode{NodeType: "Literal"}}
	if IsIdentifier(lit) {
		t.Error("Literal node should return false")
	}

	if IsIdentifier(nil) {
		t.Error("nil should return false")
	}
}

func TestIsLiteral(t *testing.T) {
	lit := &Literal{BaseNode: BaseNode{NodeType: "Literal"}}
	if !IsLiteral(lit) {
		t.Error("Literal node should return true")
	}

	id := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "test"}
	if IsLiteral(id) {
		t.Error("Identifier node should return false")
	}
}

func TestIsFunction(t *testing.T) {
	tests := []struct {
		name     string
		node     Node
		expected bool
	}{
		{
			name:     "FunctionExpression",
			node:     &FunctionExpression{BaseNode: BaseNode{NodeType: "FunctionExpression"}},
			expected: true,
		},
		{
			name:     "ArrowFunctionExpression",
			node:     &ArrowFunctionExpression{BaseNode: BaseNode{NodeType: "ArrowFunctionExpression"}},
			expected: true,
		},
		{
			name:     "FunctionDeclaration",
			node:     &FunctionDeclaration{BaseNode: BaseNode{NodeType: "FunctionDeclaration"}},
			expected: true,
		},
		{
			name:     "Identifier",
			node:     &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}},
			expected: false,
		},
		{
			name:     "nil",
			node:     nil,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsFunction(tt.node)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsClass(t *testing.T) {
	tests := []struct {
		name     string
		node     Node
		expected bool
	}{
		{
			name:     "ClassExpression",
			node:     &ClassExpression{BaseNode: BaseNode{NodeType: "ClassExpression"}},
			expected: true,
		},
		{
			name:     "ClassDeclaration",
			node:     &ClassDeclaration{BaseNode: BaseNode{NodeType: "ClassDeclaration"}},
			expected: true,
		},
		{
			name:     "FunctionDeclaration",
			node:     &FunctionDeclaration{BaseNode: BaseNode{NodeType: "FunctionDeclaration"}},
			expected: false,
		},
		{
			name:     "nil",
			node:     nil,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsClass(tt.node)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsLoop(t *testing.T) {
	tests := []struct {
		name     string
		node     Node
		expected bool
	}{
		{
			name:     "ForStatement",
			node:     &ForStatement{BaseNode: BaseNode{NodeType: "ForStatement"}},
			expected: true,
		},
		{
			name:     "WhileStatement",
			node:     &WhileStatement{BaseNode: BaseNode{NodeType: "WhileStatement"}},
			expected: true,
		},
		{
			name:     "DoWhileStatement",
			node:     &DoWhileStatement{BaseNode: BaseNode{NodeType: "DoWhileStatement"}},
			expected: true,
		},
		{
			name:     "ForInStatement",
			node:     &ForInStatement{BaseNode: BaseNode{NodeType: "ForInStatement"}},
			expected: true,
		},
		{
			name:     "ForOfStatement",
			node:     &ForOfStatement{BaseNode: BaseNode{NodeType: "ForOfStatement"}},
			expected: true,
		},
		{
			name:     "IfStatement",
			node:     &IfStatement{BaseNode: BaseNode{NodeType: "IfStatement"}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsLoop(tt.node)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsJSX(t *testing.T) {
	tests := []struct {
		name     string
		node     Node
		expected bool
	}{
		{
			name:     "JSXElement",
			node:     &JSXElement{BaseNode: BaseNode{NodeType: "JSXElement"}},
			expected: true,
		},
		{
			name:     "JSXFragment",
			node:     &JSXFragment{BaseNode: BaseNode{NodeType: "JSXFragment"}},
			expected: true,
		},
		{
			name:     "JSXText",
			node:     &JSXText{BaseNode: BaseNode{NodeType: "JSXText"}},
			expected: true,
		},
		{
			name:     "Identifier",
			node:     &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsJSX(tt.node)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsTypeScript(t *testing.T) {
	tests := []struct {
		name     string
		node     Node
		expected bool
	}{
		{
			name:     "TSTypeAnnotation",
			node:     &TSTypeAnnotation{BaseNode: BaseNode{NodeType: "TSTypeAnnotation"}},
			expected: true,
		},
		{
			name:     "TSInterfaceDeclaration",
			node:     &TSInterfaceDeclaration{BaseNode: BaseNode{NodeType: "TSInterfaceDeclaration"}},
			expected: true,
		},
		{
			name:     "Identifier",
			node:     &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}},
			expected: false,
		},
	}

	runNodeTypeTests(t, "IsTypeScript", IsTypeScript, tests)
}

func TestIsImport(t *testing.T) {
	tests := []struct {
		name     string
		node     Node
		expected bool
	}{
		{
			name:     "ImportDeclaration",
			node:     &ImportDeclaration{BaseNode: BaseNode{NodeType: "ImportDeclaration"}},
			expected: true,
		},
		{
			name:     "ImportSpecifier",
			node:     &ImportSpecifier{BaseNode: BaseNode{NodeType: "ImportSpecifier"}},
			expected: true,
		},
		{
			name:     "ExportDeclaration",
			node:     &ExportNamedDeclaration{BaseNode: BaseNode{NodeType: "ExportNamedDeclaration"}},
			expected: false,
		},
	}

	runNodeTypeTests(t, "IsImport", IsImport, tests)
}

func TestIsExport(t *testing.T) {
	tests := []struct {
		name     string
		node     Node
		expected bool
	}{
		{
			name:     "ExportNamedDeclaration",
			node:     &ExportNamedDeclaration{BaseNode: BaseNode{NodeType: "ExportNamedDeclaration"}},
			expected: true,
		},
		{
			name:     "ExportDefaultDeclaration",
			node:     &ExportDefaultDeclaration{BaseNode: BaseNode{NodeType: "ExportDefaultDeclaration"}},
			expected: true,
		},
		{
			name:     "ImportDeclaration",
			node:     &ImportDeclaration{BaseNode: BaseNode{NodeType: "ImportDeclaration"}},
			expected: false,
		},
	}

	runNodeTypeTests(t, "IsExport", IsExport, tests)
}
