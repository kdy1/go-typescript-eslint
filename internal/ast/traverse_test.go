package ast

import "testing"

func TestWalk(t *testing.T) {
	// Create a simple AST: BinaryExpression with two Identifiers
	left := &Identifier{
		BaseNode: BaseNode{NodeType: "Identifier", Start: 0, EndPos: 1},
		Name:     "a",
	}
	right := &Identifier{
		BaseNode: BaseNode{NodeType: "Identifier", Start: 4, EndPos: 5},
		Name:     "b",
	}
	binary := &BinaryExpression{
		BaseNode: BaseNode{NodeType: "BinaryExpression", Start: 0, EndPos: 5},
		Operator: "+",
		Left:     left,
		Right:    right,
	}

	visited := []string{}
	Walk(binary, VisitorFunc(func(node Node) bool {
		visited = append(visited, node.Type())
		return true
	}))

	expected := []string{"BinaryExpression", "Identifier", "Identifier"}
	if len(visited) != len(expected) {
		t.Errorf("expected %d nodes visited, got %d", len(expected), len(visited))
		return
	}
	for i, nodeType := range expected {
		if visited[i] != nodeType {
			t.Errorf("expected node type %s at index %d, got %s", nodeType, i, visited[i])
		}
	}
}

func TestWalkSkipChildren(t *testing.T) {
	// Create a simple AST
	left := &Identifier{
		BaseNode: BaseNode{NodeType: "Identifier"},
		Name:     "a",
	}
	right := &Identifier{
		BaseNode: BaseNode{NodeType: "Identifier"},
		Name:     "b",
	}
	binary := &BinaryExpression{
		BaseNode: BaseNode{NodeType: "BinaryExpression"},
		Operator: "+",
		Left:     left,
		Right:    right,
	}

	visited := []string{}
	Walk(binary, VisitorFunc(func(node Node) bool {
		visited = append(visited, node.Type())
		// Skip children of BinaryExpression
		return node.Type() != "BinaryExpression"
	}))

	if len(visited) != 1 {
		t.Errorf("expected 1 node visited, got %d", len(visited))
	}
	if visited[0] != "BinaryExpression" {
		t.Errorf("expected BinaryExpression, got %s", visited[0])
	}
}

func TestFindFirst(t *testing.T) {
	// Create AST with multiple identifiers
	id1 := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "foo"}
	id2 := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "bar"}
	binary := &BinaryExpression{
		BaseNode: BaseNode{NodeType: "BinaryExpression"},
		Left:     id1,
		Right:    id2,
	}

	result := FindFirst(binary, func(node Node) bool {
		if id, ok := node.(*Identifier); ok {
			return id.Name == "bar"
		}
		return false
	})

	if result == nil {
		t.Fatal("expected to find node, got nil")
	}
	if id, ok := result.(*Identifier); !ok || id.Name != "bar" {
		t.Errorf("expected to find identifier 'bar'")
	}
}

func TestFindAll(t *testing.T) {
	// Create AST with multiple identifiers
	id1 := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "foo"}
	id2 := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "bar"}
	id3 := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "baz"}

	binary1 := &BinaryExpression{
		BaseNode: BaseNode{NodeType: "BinaryExpression"},
		Left:     id1,
		Right:    id2,
	}
	binary2 := &BinaryExpression{
		BaseNode: BaseNode{NodeType: "BinaryExpression"},
		Left:     binary1,
		Right:    id3,
	}

	results := FindAll(binary2, func(node Node) bool {
		return node.Type() == "Identifier"
	})

	if len(results) != 3 {
		t.Errorf("expected 3 identifiers, got %d", len(results))
	}
}

func TestFindByType(t *testing.T) {
	id1 := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "foo"}
	id2 := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "bar"}
	binary := &BinaryExpression{
		BaseNode: BaseNode{NodeType: "BinaryExpression"},
		Left:     id1,
		Right:    id2,
	}

	results := FindByType(binary, "Identifier")
	if len(results) != 2 {
		t.Errorf("expected 2 identifiers, got %d", len(results))
	}
}

func TestGetParent(t *testing.T) {
	left := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "a"}
	right := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "b"}
	binary := &BinaryExpression{
		BaseNode: BaseNode{NodeType: "BinaryExpression"},
		Left:     left,
		Right:    right,
	}

	parent := GetParent(binary, left)
	if parent != binary {
		t.Error("expected parent to be the binary expression")
	}

	parent = GetParent(binary, binary)
	if parent != nil {
		t.Error("expected root node to have no parent")
	}
}

func TestGetAncestors(t *testing.T) {
	id := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "a"}
	unary := &UnaryExpression{
		BaseNode: BaseNode{NodeType: "UnaryExpression"},
		Operator: "-",
		Argument: id,
	}
	binary := &BinaryExpression{
		BaseNode: BaseNode{NodeType: "BinaryExpression"},
		Left:     unary,
		Right:    &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "b"},
	}

	ancestors := GetAncestors(binary, id)
	if len(ancestors) != 2 {
		t.Errorf("expected 2 ancestors, got %d", len(ancestors))
	}
}

func TestContains(t *testing.T) {
	left := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "a"}
	right := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "b"}
	binary := &BinaryExpression{
		BaseNode: BaseNode{NodeType: "BinaryExpression"},
		Left:     left,
		Right:    right,
	}

	if !Contains(binary, left) {
		t.Error("expected binary to contain left")
	}

	other := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "c"}
	if Contains(binary, other) {
		t.Error("expected binary not to contain other")
	}
}

func TestCountNodes(t *testing.T) {
	id1 := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "a"}
	id2 := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "b"}
	binary := &BinaryExpression{
		BaseNode: BaseNode{NodeType: "BinaryExpression"},
		Left:     id1,
		Right:    id2,
	}

	count := CountNodes(binary)
	if count != 3 {
		t.Errorf("expected 3 nodes, got %d", count)
	}
}

func TestTraverseWithContext(t *testing.T) {
	left := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "a"}
	right := &Identifier{BaseNode: BaseNode{NodeType: "Identifier"}, Name: "b"}
	binary := &BinaryExpression{
		BaseNode: BaseNode{NodeType: "BinaryExpression"},
		Left:     left,
		Right:    right,
	}

	var foundParent Node
	TraverseWithContext(binary, func(node Node, ctx *TraverseContext) bool {
		if node == left {
			foundParent = ctx.Parent
			return false
		}
		return true
	})

	if foundParent != binary {
		t.Error("expected parent to be the binary expression")
	}
}
