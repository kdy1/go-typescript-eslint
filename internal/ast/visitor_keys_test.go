package ast

import (
	"testing"
)

func TestGetVisitorKeys(t *testing.T) {
	tests := []struct {
		name     string
		nodeType string
		expected []string
	}{
		{
			name:     "Program",
			nodeType: "Program",
			expected: []string{"body"},
		},
		{
			name:     "Identifier has no children",
			nodeType: "Identifier",
			expected: []string{},
		},
		{
			name:     "BinaryExpression",
			nodeType: "BinaryExpression",
			expected: []string{"left", "right"},
		},
		{
			name:     "FunctionDeclaration",
			nodeType: "FunctionDeclaration",
			expected: []string{"id", "typeParameters", "params", "returnType", "body"},
		},
		{
			name:     "ClassDeclaration",
			nodeType: "ClassDeclaration",
			expected: []string{"decorators", "id", "typeParameters", "superClass", "superTypeArguments", "implements", "body"},
		},
		{
			name:     "Unknown node type",
			nodeType: "UnknownNodeType",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys := GetVisitorKeys(tt.nodeType)
			if len(keys) != len(tt.expected) {
				t.Errorf("expected %d keys, got %d", len(tt.expected), len(keys))
				return
			}
			for i, key := range keys {
				if key != tt.expected[i] {
					t.Errorf("expected key %s at index %d, got %s", tt.expected[i], i, key)
				}
			}
		})
	}
}

func TestHasVisitorKeys(t *testing.T) {
	tests := []struct {
		name     string
		nodeType string
		expected bool
	}{
		{
			name:     "Program exists",
			nodeType: "Program",
			expected: true,
		},
		{
			name:     "Identifier exists",
			nodeType: "Identifier",
			expected: true,
		},
		{
			name:     "Unknown type does not exist",
			nodeType: "UnknownType",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasVisitorKeys(tt.nodeType)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAllNodeTypesHaveVisitorKeys(t *testing.T) {
	// Test that all major node types have visitor keys defined
	nodeTypes := []string{
		"Program", "Identifier", "Literal",
		"FunctionExpression", "ArrowFunctionExpression",
		"BinaryExpression", "UnaryExpression",
		"IfStatement", "ForStatement", "WhileStatement",
		"VariableDeclaration", "FunctionDeclaration", "ClassDeclaration",
		"JSXElement", "JSXOpeningElement",
		"TSTypeAnnotation", "TSInterfaceDeclaration",
	}

	for _, nodeType := range nodeTypes {
		if !HasVisitorKeys(nodeType) {
			t.Errorf("node type %s does not have visitor keys defined", nodeType)
		}
	}
}
