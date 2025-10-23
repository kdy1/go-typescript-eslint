package ast

import "testing"

func TestGetNodeRange(t *testing.T) {
	node := &Identifier{
		BaseNode: BaseNode{NodeType: "Identifier", Start: 10, EndPos: 15},
		Name:     "test",
	}

	rng := GetNodeRange(node)
	if rng[0] != 10 || rng[1] != 15 {
		t.Errorf("expected range [10, 15], got [%d, %d]", rng[0], rng[1])
	}

	nilRng := GetNodeRange(nil)
	if nilRng[0] != 0 || nilRng[1] != 0 {
		t.Error("expected [0, 0] for nil node")
	}
}

func TestNodeSpan(t *testing.T) {
	node := &Identifier{
		BaseNode: BaseNode{NodeType: "Identifier", Start: 10, EndPos: 15},
		Name:     "test",
	}

	span := NodeSpan(node)
	if span != 5 {
		t.Errorf("expected span 5, got %d", span)
	}

	if NodeSpan(nil) != 0 {
		t.Error("expected 0 for nil node")
	}
}

func TestNodesOverlap(t *testing.T) {
	tests := []struct {
		name     string
		a        Node
		b        Node
		expected bool
	}{
		{
			name:     "overlapping nodes",
			a:        &Identifier{BaseNode: BaseNode{Start: 0, EndPos: 10}},
			b:        &Identifier{BaseNode: BaseNode{Start: 5, EndPos: 15}},
			expected: true,
		},
		{
			name:     "non-overlapping nodes",
			a:        &Identifier{BaseNode: BaseNode{Start: 0, EndPos: 10}},
			b:        &Identifier{BaseNode: BaseNode{Start: 10, EndPos: 20}},
			expected: false,
		},
		{
			name:     "contained node",
			a:        &Identifier{BaseNode: BaseNode{Start: 0, EndPos: 20}},
			b:        &Identifier{BaseNode: BaseNode{Start: 5, EndPos: 15}},
			expected: true,
		},
		{
			name:     "nil nodes",
			a:        nil,
			b:        &Identifier{BaseNode: BaseNode{Start: 0, EndPos: 10}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NodesOverlap(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestNodeContains(t *testing.T) {
	outer := &Identifier{BaseNode: BaseNode{Start: 0, EndPos: 20}}
	inner := &Identifier{BaseNode: BaseNode{Start: 5, EndPos: 15}}
	outside := &Identifier{BaseNode: BaseNode{Start: 25, EndPos: 30}}

	if !NodeContains(outer, inner) {
		t.Error("outer should contain inner")
	}

	if NodeContains(outer, outside) {
		t.Error("outer should not contain outside")
	}

	if NodeContains(nil, inner) {
		t.Error("nil should not contain anything")
	}
}

func TestIsBefore(t *testing.T) {
	before := &Identifier{BaseNode: BaseNode{Start: 0, EndPos: 10}}
	after := &Identifier{BaseNode: BaseNode{Start: 10, EndPos: 20}}

	if !IsBefore(before, after) {
		t.Error("before should be before after")
	}

	if IsBefore(after, before) {
		t.Error("after should not be before before")
	}
}

func TestIsAfter(t *testing.T) {
	before := &Identifier{BaseNode: BaseNode{Start: 0, EndPos: 10}}
	after := &Identifier{BaseNode: BaseNode{Start: 10, EndPos: 20}}

	if !IsAfter(after, before) {
		t.Error("after should be after before")
	}

	if IsAfter(before, after) {
		t.Error("before should not be after after")
	}
}

func TestGetNodeText(t *testing.T) {
	source := "const foo = 42;"
	node := &Identifier{
		BaseNode: BaseNode{Start: 6, EndPos: 9},
		Name:     "foo",
	}

	text := GetNodeText(node, source)
	if text != "foo" {
		t.Errorf("expected 'foo', got '%s'", text)
	}

	if GetNodeText(nil, source) != "" {
		t.Error("expected empty string for nil node")
	}

	if GetNodeText(node, "") != "" {
		t.Error("expected empty string for empty source")
	}
}

func TestIsInRange(t *testing.T) {
	node := &Identifier{BaseNode: BaseNode{Start: 10, EndPos: 20}}

	if !IsInRange(node, 15) {
		t.Error("15 should be in range")
	}

	if IsInRange(node, 5) {
		t.Error("5 should not be in range")
	}

	if IsInRange(node, 20) {
		t.Error("20 should not be in range (exclusive)")
	}

	if IsInRange(nil, 10) {
		t.Error("nil node should not contain any position")
	}
}

func TestIsNullLiteral(t *testing.T) {
	nullLit := &Literal{
		BaseNode: BaseNode{NodeType: "Literal"},
		Value:    nil,
		Raw:      "null",
	}

	if !IsNullLiteral(nullLit) {
		t.Error("should be null literal")
	}

	stringLit := &Literal{
		BaseNode: BaseNode{NodeType: "Literal"},
		Value:    "test",
		Raw:      "\"test\"",
	}

	if IsNullLiteral(stringLit) {
		t.Error("should not be null literal")
	}
}

func TestIsBooleanLiteral(t *testing.T) {
	trueLit := &Literal{
		BaseNode: BaseNode{NodeType: "Literal"},
		Value:    true,
		Raw:      "true",
	}

	if !IsBooleanLiteral(trueLit) {
		t.Error("should be boolean literal")
	}

	stringLit := &Literal{
		BaseNode: BaseNode{NodeType: "Literal"},
		Value:    "test",
		Raw:      "\"test\"",
	}

	if IsBooleanLiteral(stringLit) {
		t.Error("should not be boolean literal")
	}
}

func TestIsStringLiteral(t *testing.T) {
	stringLit := &Literal{
		BaseNode: BaseNode{NodeType: "Literal"},
		Value:    "test",
		Raw:      "\"test\"",
	}

	if !IsStringLiteral(stringLit) {
		t.Error("should be string literal")
	}

	numLit := &Literal{
		BaseNode: BaseNode{NodeType: "Literal"},
		Value:    42,
		Raw:      "42",
	}

	if IsStringLiteral(numLit) {
		t.Error("should not be string literal")
	}
}

func TestIsNumberLiteral(t *testing.T) {
	numLit := &Literal{
		BaseNode: BaseNode{NodeType: "Literal"},
		Value:    42,
		Raw:      "42",
	}

	if !IsNumberLiteral(numLit) {
		t.Error("should be number literal")
	}

	stringLit := &Literal{
		BaseNode: BaseNode{NodeType: "Literal"},
		Value:    "test",
		Raw:      "\"test\"",
	}

	if IsNumberLiteral(stringLit) {
		t.Error("should not be number literal")
	}
}

func TestIsAsyncFunction(t *testing.T) {
	asyncFunc := &FunctionExpression{
		BaseNode: BaseNode{NodeType: "FunctionExpression"},
		Async:    true,
	}

	if !IsAsyncFunction(asyncFunc) {
		t.Error("should be async function")
	}

	syncFunc := &FunctionExpression{
		BaseNode: BaseNode{NodeType: "FunctionExpression"},
		Async:    false,
	}

	if IsAsyncFunction(syncFunc) {
		t.Error("should not be async function")
	}
}

func TestIsGeneratorFunction(t *testing.T) {
	genFunc := &FunctionExpression{
		BaseNode:  BaseNode{NodeType: "FunctionExpression"},
		Generator: true,
	}

	if !IsGeneratorFunction(genFunc) {
		t.Error("should be generator function")
	}

	normalFunc := &FunctionExpression{
		BaseNode:  BaseNode{NodeType: "FunctionExpression"},
		Generator: false,
	}

	if IsGeneratorFunction(normalFunc) {
		t.Error("should not be generator function")
	}
}

func TestGetFunctionName(t *testing.T) {
	namedFunc := &FunctionExpression{
		BaseNode: BaseNode{NodeType: "FunctionExpression"},
		ID:       &Identifier{Name: "myFunc"},
	}

	name := GetFunctionName(namedFunc)
	if name != "myFunc" {
		t.Errorf("expected 'myFunc', got '%s'", name)
	}

	anonFunc := &FunctionExpression{
		BaseNode: BaseNode{NodeType: "FunctionExpression"},
		ID:       nil,
	}

	if GetFunctionName(anonFunc) != "" {
		t.Error("expected empty string for anonymous function")
	}
}

func TestGetClassName(t *testing.T) {
	namedClass := &ClassExpression{
		BaseNode: BaseNode{NodeType: "ClassExpression"},
		ID:       &Identifier{Name: "MyClass"},
	}

	name := GetClassName(namedClass)
	if name != "MyClass" {
		t.Errorf("expected 'MyClass', got '%s'", name)
	}

	anonClass := &ClassExpression{
		BaseNode: BaseNode{NodeType: "ClassExpression"},
		ID:       nil,
	}

	if GetClassName(anonClass) != "" {
		t.Error("expected empty string for anonymous class")
	}
}

func TestGetAllIdentifiers(t *testing.T) {
	id1 := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "a"}
	id2 := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "b"}
	binary := &BinaryExpression{
		BaseNode: BaseNode{NodeType: "BinaryExpression"},
		Left:     id1,
		Right:    id2,
	}

	identifiers := GetAllIdentifiers(binary)
	if len(identifiers) != 2 {
		t.Errorf("expected 2 identifiers, got %d", len(identifiers))
	}
}
