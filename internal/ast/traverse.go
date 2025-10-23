package ast

import "reflect"

// Visitor is the interface for AST node visitors.
// Implementations can define Visit methods to handle specific node types.
type Visitor interface {
	// Visit is called for each node in the AST.
	// If it returns false, traversal of this node's children is skipped.
	Visit(node Node) bool
}

// VisitorFunc is a function type that implements the Visitor interface.
type VisitorFunc func(node Node) bool

// Visit implements the Visitor interface for VisitorFunc.
func (f VisitorFunc) Visit(node Node) bool {
	return f(node)
}

// Walk traverses an AST in depth-first order, calling the visitor's Visit method
// for each node. If Visit returns false, the node's children are not traversed.
func Walk(node Node, visitor Visitor) {
	if node == nil || visitor == nil {
		return
	}

	// Visit the current node
	if !visitor.Visit(node) {
		return
	}

	// Get visitor keys for this node type
	keys := GetVisitorKeys(node.Type())
	if len(keys) == 0 {
		return
	}

	// Use reflection to access child nodes
	nodeValue := reflect.ValueOf(node)
	if nodeValue.Kind() == reflect.Ptr {
		nodeValue = nodeValue.Elem()
	}

	// Traverse each child property
	for _, key := range keys {
		field := nodeValue.FieldByName(capitalizeFirst(key))
		if !field.IsValid() {
			continue
		}

		traverseField(field, visitor)
	}
}

// traverseField handles traversal of a field value, which may be a node,
// a slice of nodes, or a pointer to a node.
//
//nolint:cyclop // Complexity is inherent to reflection-based traversal
//nolint:exhaustive // Only specific reflection kinds need handling
func traverseField(field reflect.Value, visitor Visitor) {
	switch field.Kind() {
	case reflect.Ptr:
		traversePointerField(field, visitor)
	case reflect.Slice:
		traverseSliceField(field, visitor)
	case reflect.Interface:
		traverseInterfaceField(field, visitor)
	default:
		// Other types (Bool, Int, String, etc.) are not traversable nodes
		return
	}
}

func traversePointerField(field reflect.Value, visitor Visitor) {
	if !field.IsNil() {
		if node, ok := field.Interface().(Node); ok {
			Walk(node, visitor)
		}
	}
}

func traverseSliceField(field reflect.Value, visitor Visitor) {
	for i := 0; i < field.Len(); i++ {
		elem := field.Index(i)
		if elem.Kind() == reflect.Ptr && !elem.IsNil() {
			if node, ok := elem.Interface().(Node); ok {
				Walk(node, visitor)
			}
		} else if node, ok := elem.Interface().(Node); ok {
			Walk(node, visitor)
		}
	}
}

func traverseInterfaceField(field reflect.Value, visitor Visitor) {
	if !field.IsNil() {
		if node, ok := field.Interface().(Node); ok {
			Walk(node, visitor)
		}
	}
}

// capitalizeFirst capitalizes the first letter of a string (for field name lookup).
func capitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	if s[0] >= 'a' && s[0] <= 'z' {
		return string(s[0]-'a'+'A') + s[1:]
	}
	return s
}

// TraverseContext holds context information during AST traversal.
type TraverseContext struct {
	// Parent is the parent node of the current node
	Parent Node
	// Ancestors is the list of ancestor nodes (excluding parent)
	Ancestors []Node
	// Key is the property name of the current node in its parent
	Key string
	// Index is the index if the current node is in an array
	Index *int
}

// ContextVisitor is a visitor that receives context information during traversal.
type ContextVisitor interface {
	// VisitWithContext is called for each node with its context.
	// If it returns false, traversal of this node's children is skipped.
	VisitWithContext(node Node, ctx *TraverseContext) bool
}

// ContextVisitorFunc is a function type that implements ContextVisitor.
type ContextVisitorFunc func(node Node, ctx *TraverseContext) bool

// VisitWithContext implements ContextVisitor for ContextVisitorFunc.
func (f ContextVisitorFunc) VisitWithContext(node Node, ctx *TraverseContext) bool {
	return f(node, ctx)
}

// WalkWithContext traverses an AST with context information about parent nodes.
func WalkWithContext(node Node, visitor ContextVisitor) {
	if node == nil || visitor == nil {
		return
	}
	walkWithContextInternal(node, visitor, &TraverseContext{
		Ancestors: []Node{},
	})
}

func walkWithContextInternal(node Node, visitor ContextVisitor, ctx *TraverseContext) {
	if !visitor.VisitWithContext(node, ctx) {
		return
	}

	// Get visitor keys for this node type
	keys := GetVisitorKeys(node.Type())
	if len(keys) == 0 {
		return
	}

	// Use reflection to access child nodes
	nodeValue := reflect.ValueOf(node)
	if nodeValue.Kind() == reflect.Ptr {
		nodeValue = nodeValue.Elem()
	}

	// Update context for children
	var newAncestors []Node
	if ctx.Parent != nil {
		newAncestors = make([]Node, len(ctx.Ancestors)+1)
		copy(newAncestors, ctx.Ancestors)
		newAncestors[len(ctx.Ancestors)] = ctx.Parent
	} else {
		newAncestors = ctx.Ancestors
	}

	// Traverse each child property
	for _, key := range keys {
		field := nodeValue.FieldByName(capitalizeFirst(key))
		if !field.IsValid() {
			continue
		}

		childCtx := &TraverseContext{
			Parent:    node,
			Ancestors: newAncestors,
			Key:       key,
		}

		traverseFieldWithContext(field, visitor, childCtx)
	}
}

//nolint:cyclop // Complexity is inherent to reflection-based traversal
//nolint:exhaustive // Only specific reflection kinds need handling
func traverseFieldWithContext(field reflect.Value, visitor ContextVisitor, ctx *TraverseContext) {
	switch field.Kind() {
	case reflect.Ptr:
		traversePointerWithContext(field, visitor, ctx)
	case reflect.Slice:
		traverseSliceWithContext(field, visitor, ctx)
	case reflect.Interface:
		traverseInterfaceWithContext(field, visitor, ctx)
	default:
		// Other types (Bool, Int, String, etc.) are not traversable nodes
		return
	}
}

func traversePointerWithContext(field reflect.Value, visitor ContextVisitor, ctx *TraverseContext) {
	if !field.IsNil() {
		if node, ok := field.Interface().(Node); ok {
			walkWithContextInternal(node, visitor, ctx)
		}
	}
}

func traverseSliceWithContext(field reflect.Value, visitor ContextVisitor, ctx *TraverseContext) {
	for i := 0; i < field.Len(); i++ {
		elem := field.Index(i)
		idx := i // Create a new variable for each iteration
		childCtx := &TraverseContext{
			Parent:    ctx.Parent,
			Ancestors: ctx.Ancestors,
			Key:       ctx.Key,
			Index:     &idx,
		}
		if elem.Kind() == reflect.Ptr && !elem.IsNil() {
			if node, ok := elem.Interface().(Node); ok {
				walkWithContextInternal(node, visitor, childCtx)
			}
		} else if node, ok := elem.Interface().(Node); ok {
			walkWithContextInternal(node, visitor, childCtx)
		}
	}
}

func traverseInterfaceWithContext(field reflect.Value, visitor ContextVisitor, ctx *TraverseContext) {
	if !field.IsNil() {
		if node, ok := field.Interface().(Node); ok {
			walkWithContextInternal(node, visitor, ctx)
		}
	}
}

// Traverse is a convenience function that traverses an AST with a simple callback.
func Traverse(node Node, fn func(node Node) bool) {
	Walk(node, VisitorFunc(fn))
}

// TraverseWithContext is a convenience function that traverses an AST with context.
func TraverseWithContext(node Node, fn func(node Node, ctx *TraverseContext) bool) {
	WalkWithContext(node, ContextVisitorFunc(fn))
}

// FindFirst traverses the AST and returns the first node for which the predicate returns true.
//
//nolint:ireturn // Interface types are intentional for generic AST traversal
func FindFirst(root Node, predicate func(node Node) bool) Node {
	var result Node
	Traverse(root, func(node Node) bool {
		if predicate(node) {
			result = node
			return false // Stop traversal
		}
		return true
	})
	return result
}

// FindAll traverses the AST and returns all nodes for which the predicate returns true.
func FindAll(root Node, predicate func(node Node) bool) []Node {
	var results []Node
	Traverse(root, func(node Node) bool {
		if predicate(node) {
			results = append(results, node)
		}
		return true
	})
	return results
}

// FindByType returns all nodes of the specified type.
func FindByType(root Node, nodeType string) []Node {
	return FindAll(root, func(node Node) bool {
		return node.Type() == nodeType
	})
}

// GetParent returns the parent node of the target node, or nil if not found.
//
//nolint:ireturn // Interface types are intentional for generic AST node retrieval
func GetParent(root, target Node) Node {
	var parent Node
	TraverseWithContext(root, func(node Node, ctx *TraverseContext) bool {
		if node == target {
			parent = ctx.Parent
			return false // Stop traversal
		}
		return true
	})
	return parent
}

// GetAncestors returns all ancestor nodes of the target node.
func GetAncestors(root, target Node) []Node {
	var ancestors []Node
	TraverseWithContext(root, func(node Node, ctx *TraverseContext) bool {
		if node == target {
			if ctx.Parent != nil {
				ancestors = make([]Node, 0, len(ctx.Ancestors)+1)
				ancestors = append(ancestors, ctx.Ancestors...)
				ancestors = append(ancestors, ctx.Parent)
			} else {
				ancestors = make([]Node, len(ctx.Ancestors))
				copy(ancestors, ctx.Ancestors)
			}
			return false // Stop traversal
		}
		return true
	})
	return ancestors
}

// GetSiblings returns all sibling nodes of the target node.
// If the target is not found or has no siblings, returns an empty slice.
//
//nolint:cyclop // Complexity is inherent to sibling extraction logic
func GetSiblings(root, target Node) []Node {
	var siblings []Node
	targetCtx := findTargetContext(root, target)
	if targetCtx == nil || targetCtx.Parent == nil {
		return siblings
	}

	return extractSiblings(targetCtx.Parent, targetCtx.Key, targetCtx.Index)
}

func findTargetContext(root, target Node) *TraverseContext {
	var result *TraverseContext
	TraverseWithContext(root, func(node Node, ctx *TraverseContext) bool {
		if node == target {
			result = ctx
			return false
		}
		return true
	})
	return result
}

func extractSiblings(parent Node, key string, targetIndex *int) []Node {
	var siblings []Node
	parentValue := reflect.ValueOf(parent)
	if parentValue.Kind() == reflect.Ptr {
		parentValue = parentValue.Elem()
	}

	field := parentValue.FieldByName(capitalizeFirst(key))
	if !field.IsValid() || field.Kind() != reflect.Slice {
		return siblings
	}

	for i := 0; i < field.Len(); i++ {
		if targetIndex != nil && i == *targetIndex {
			continue // Skip the target itself
		}
		elem := field.Index(i)
		if node := extractNodeFromElement(elem); node != nil {
			siblings = append(siblings, node)
		}
	}

	return siblings
}

func extractNodeFromElement(elem reflect.Value) Node {
	if elem.Kind() == reflect.Ptr && !elem.IsNil() {
		if node, ok := elem.Interface().(Node); ok {
			return node
		}
	} else if node, ok := elem.Interface().(Node); ok {
		return node
	}
	return nil
}

// Contains checks if the AST rooted at 'root' contains 'target' node.
func Contains(root, target Node) bool {
	found := false
	Traverse(root, func(node Node) bool {
		if node == target {
			found = true
			return false
		}
		return true
	})
	return found
}

// GetDepth returns the depth of the target node in the AST (0 for root).
func GetDepth(root, target Node) int {
	depth := -1
	TraverseWithContext(root, func(node Node, ctx *TraverseContext) bool {
		if node == target {
			depth = len(ctx.Ancestors)
			if ctx.Parent != nil {
				depth++
			}
			return false
		}
		return true
	})
	return depth
}
