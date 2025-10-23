package ast

import (
	"encoding/json"
	"fmt"
)

// NodeEquals compares two nodes for structural equality.
// It performs a deep comparison of all node properties.
func NodeEquals(a, b Node) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Type() != b.Type() {
		return false
	}

	// Use JSON serialization for deep comparison
	// This is simple but may not be the most efficient approach
	aJSON, err := json.Marshal(a)
	if err != nil {
		return false
	}
	bJSON, err := json.Marshal(b)
	if err != nil {
		return false
	}

	return string(aJSON) == string(bJSON)
}

// CloneNode creates a deep copy of a node.
// This uses JSON serialization/deserialization for simplicity.
// For production use, consider implementing a more efficient cloning mechanism.
func CloneNode(node Node) (Node, error) {
	if node == nil {
		return nil, nil
	}

	// Serialize to JSON
	data, err := json.Marshal(node)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal node: %w", err)
	}

	// Deserialize back to a node
	// Note: This will create a generic map structure.
	// For full functionality, you'd need to implement type-specific deserialization.
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal node: %w", err)
	}

	// This is a simplified version. In a real implementation,
	// you'd want to reconstruct the proper node type.
	return nil, fmt.Errorf("full node cloning not yet implemented")
}

// GetNodeRange returns the source range of a node as [start, end].
func GetNodeRange(node Node) [2]int {
	if node == nil {
		return [2]int{0, 0}
	}
	return [2]int{node.Pos(), node.End()}
}

// GetNodeLocation returns the source location of a node.
func GetNodeLocation(node Node) *SourceLocation {
	if node == nil {
		return nil
	}

	// Try to access BaseNode's Loc field
	if bn, ok := node.(interface{ GetLoc() *SourceLocation }); ok {
		return bn.GetLoc()
	}

	return nil
}

// NodeSpan returns the length of the node in characters.
func NodeSpan(node Node) int {
	if node == nil {
		return 0
	}
	return node.End() - node.Pos()
}

// NodesOverlap checks if two nodes overlap in source code.
func NodesOverlap(a, b Node) bool {
	if a == nil || b == nil {
		return false
	}
	aStart, aEnd := a.Pos(), a.End()
	bStart, bEnd := b.Pos(), b.End()
	return (aStart <= bStart && bStart < aEnd) || (bStart <= aStart && aStart < bEnd)
}

// NodeContains checks if node 'a' contains node 'b' in source code.
func NodeContains(a, b Node) bool {
	if a == nil || b == nil {
		return false
	}
	return a.Pos() <= b.Pos() && b.End() <= a.End()
}

// IsBefore checks if node 'a' appears before node 'b' in source code.
func IsBefore(a, b Node) bool {
	if a == nil || b == nil {
		return false
	}
	return a.End() <= b.Pos()
}

// IsAfter checks if node 'a' appears after node 'b' in source code.
func IsAfter(a, b Node) bool {
	if a == nil || b == nil {
		return false
	}
	return a.Pos() >= b.End()
}

// GetNodeText returns the source text of a node given the full source code.
func GetNodeText(node Node, source string) string {
	if node == nil || source == "" {
		return ""
	}
	start, end := node.Pos(), node.End()
	if start < 0 || end > len(source) || start > end {
		return ""
	}
	return source[start:end]
}

// IsInRange checks if a position is within the node's range.
func IsInRange(node Node, pos int) bool {
	if node == nil {
		return false
	}
	return node.Pos() <= pos && pos < node.End()
}

// GetNodeAtPosition returns the deepest node that contains the given position.
func GetNodeAtPosition(root Node, pos int) Node {
	var result Node
	Traverse(root, func(node Node) bool {
		if IsInRange(node, pos) {
			result = node
			return true // Continue to find deeper nodes
		}
		// If we're past this node, don't traverse its children
		return node.Pos() <= pos
	})
	return result
}

// GetNodesInRange returns all nodes that overlap with the given range.
func GetNodesInRange(root Node, start, end int) []Node {
	return FindAll(root, func(node Node) bool {
		nodeStart, nodeEnd := node.Pos(), node.End()
		// Check if ranges overlap
		return (nodeStart < end && nodeEnd > start)
	})
}

// IsEmptyStatement checks if a node is an empty statement.
func IsEmptyStatement(node Node) bool {
	return node != nil && node.Type() == "EmptyStatement"
}

// IsNullLiteral checks if a node is a null literal.
func IsNullLiteral(node Node) bool {
	if !IsLiteral(node) {
		return false
	}
	lit, ok := node.(*Literal)
	return ok && lit.Value == nil && lit.Raw == "null"
}

// IsBooleanLiteral checks if a node is a boolean literal.
func IsBooleanLiteral(node Node) bool {
	if !IsLiteral(node) {
		return false
	}
	lit, ok := node.(*Literal)
	if !ok {
		return false
	}
	_, isBool := lit.Value.(bool)
	return isBool
}

// IsStringLiteral checks if a node is a string literal.
func IsStringLiteral(node Node) bool {
	if !IsLiteral(node) {
		return false
	}
	lit, ok := node.(*Literal)
	if !ok {
		return false
	}
	_, isString := lit.Value.(string)
	return isString
}

// IsNumberLiteral checks if a node is a number literal.
func IsNumberLiteral(node Node) bool {
	if !IsLiteral(node) {
		return false
	}
	lit, ok := node.(*Literal)
	if !ok {
		return false
	}
	switch lit.Value.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return true
	}
	return false
}

// IsRegExpLiteral checks if a node is a regular expression literal.
func IsRegExpLiteral(node Node) bool {
	if !IsLiteral(node) {
		return false
	}
	lit, ok := node.(*Literal)
	return ok && lit.Regex != nil
}

// IsBigIntLiteral checks if a node is a bigint literal.
func IsBigIntLiteral(node Node) bool {
	if !IsLiteral(node) {
		return false
	}
	lit, ok := node.(*Literal)
	return ok && lit.BigInt != nil
}

// IsStaticMemberExpression checks if a node is a static member expression (obj.prop).
func IsStaticMemberExpression(node Node) bool {
	if !IsMemberExpression(node) {
		return false
	}
	mem, ok := node.(*MemberExpression)
	return ok && !mem.Computed
}

// IsComputedMemberExpression checks if a node is a computed member expression (obj[prop]).
func IsComputedMemberExpression(node Node) bool {
	if !IsMemberExpression(node) {
		return false
	}
	mem, ok := node.(*MemberExpression)
	return ok && mem.Computed
}

// IsMethodCall checks if a node is a method call (obj.method()).
func IsMethodCall(node Node) bool {
	if !IsCallExpression(node) {
		return false
	}
	call, ok := node.(*CallExpression)
	if !ok {
		return false
	}
	return IsMemberExpression(call.Callee)
}

// IsNewCall checks if a node is a new expression with a call (new Class()).
func IsNewCall(node Node) bool {
	return node != nil && node.Type() == "NewExpression"
}

// IsThisExpression checks if a node is 'this'.
func IsThisExpression(node Node) bool {
	return node != nil && node.Type() == "ThisExpression"
}

// IsSuperExpression checks if a node is 'super'.
func IsSuperExpression(node Node) bool {
	return node != nil && node.Type() == "Super"
}

// IsAsyncFunction checks if a node is an async function.
func IsAsyncFunction(node Node) bool {
	if !IsFunction(node) {
		return false
	}

	switch n := node.(type) {
	case *FunctionExpression:
		return n.Async
	case *ArrowFunctionExpression:
		return n.Async
	case *FunctionDeclaration:
		return n.Async
	}
	return false
}

// IsGeneratorFunction checks if a node is a generator function.
func IsGeneratorFunction(node Node) bool {
	if !IsFunction(node) {
		return false
	}

	switch n := node.(type) {
	case *FunctionExpression:
		return n.Generator
	case *FunctionDeclaration:
		return n.Generator
	}
	return false
}

// IsArrowFunction checks if a node is an arrow function.
func IsArrowFunction(node Node) bool {
	return IsArrowFunctionExpression(node)
}

// HasAwait checks if a node tree contains an await expression.
func HasAwait(node Node) bool {
	found := false
	Traverse(node, func(n Node) bool {
		if n.Type() == "AwaitExpression" {
			found = true
			return false
		}
		return true
	})
	return found
}

// HasYield checks if a node tree contains a yield expression.
func HasYield(node Node) bool {
	found := false
	Traverse(node, func(n Node) bool {
		if n.Type() == "YieldExpression" {
			found = true
			return false
		}
		return true
	})
	return found
}

// CountNodes returns the total number of nodes in the tree.
func CountNodes(root Node) int {
	count := 0
	Traverse(root, func(node Node) bool {
		count++
		return true
	})
	return count
}

// GetAllIdentifiers returns all identifier nodes in the tree.
func GetAllIdentifiers(root Node) []*Identifier {
	nodes := FindByType(root, "Identifier")
	identifiers := make([]*Identifier, 0, len(nodes))
	for _, node := range nodes {
		if id, ok := node.(*Identifier); ok {
			identifiers = append(identifiers, id)
		}
	}
	return identifiers
}

// GetIdentifierNames returns all identifier names in the tree.
func GetIdentifierNames(root Node) []string {
	identifiers := GetAllIdentifiers(root)
	names := make([]string, 0, len(identifiers))
	seen := make(map[string]bool)
	for _, id := range identifiers {
		if !seen[id.Name] {
			names = append(names, id.Name)
			seen[id.Name] = true
		}
	}
	return names
}

// IsDeclarationStatement checks if a node is a declaration statement.
func IsDeclarationStatement(node Node) bool {
	if node == nil {
		return false
	}
	t := node.Type()
	return t == "FunctionDeclaration" || t == "ClassDeclaration" ||
		t == "VariableDeclaration" || IsTypeScript(node) &&
		(t == "TSInterfaceDeclaration" || t == "TSTypeAliasDeclaration" ||
		t == "TSEnumDeclaration" || t == "TSModuleDeclaration")
}

// GetFunctionName returns the name of a function node, if it has one.
func GetFunctionName(node Node) string {
	if !IsFunction(node) {
		return ""
	}

	switch n := node.(type) {
	case *FunctionExpression:
		if n.ID != nil {
			return n.ID.Name
		}
	case *FunctionDeclaration:
		if n.ID != nil {
			return n.ID.Name
		}
	}
	return ""
}

// GetClassName returns the name of a class node, if it has one.
func GetClassName(node Node) string {
	if !IsClass(node) {
		return ""
	}

	switch n := node.(type) {
	case *ClassExpression:
		if n.ID != nil {
			return n.ID.Name
		}
	case *ClassDeclaration:
		if n.ID != nil {
			return n.ID.Name
		}
	}
	return ""
}

// IsExported checks if a declaration is exported.
func IsExported(node Node) bool {
	if node == nil {
		return false
	}
	t := node.Type()
	return t == "ExportNamedDeclaration" || t == "ExportDefaultDeclaration" ||
		t == "ExportAllDeclaration"
}

// IsDefaultExport checks if a node is a default export.
func IsDefaultExport(node Node) bool {
	return node != nil && node.Type() == "ExportDefaultDeclaration"
}

// IsNamedExport checks if a node is a named export.
func IsNamedExport(node Node) bool {
	return node != nil && node.Type() == "ExportNamedDeclaration"
}
