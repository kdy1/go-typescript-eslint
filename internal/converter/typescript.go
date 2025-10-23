package converter

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
)

// convertTSTypeAnnotation converts a TSTypeAnnotation node.
func (c *Converter) convertTSTypeAnnotation(node *ast.TSTypeAnnotation) *ast.TSTypeAnnotation {
	if node == nil {
		return nil
	}

	var typeAnnotation ast.TSNode
	if node.TypeAnnotation != nil {
		if ts, ok := c.ConvertNode(node.TypeAnnotation).(ast.TSNode); ok {
			typeAnnotation = ts
		}
	}

	result := &ast.TSTypeAnnotation{
		BaseNode:       c.copyBaseNode(&node.BaseNode),
		TypeAnnotation: typeAnnotation,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertTSInterfaceDeclaration converts a TSInterfaceDeclaration node.
func (c *Converter) convertTSInterfaceDeclaration(node *ast.TSInterfaceDeclaration) *ast.TSInterfaceDeclaration {
	if node == nil {
		return nil
	}

	extends := make([]ast.TSInterfaceHeritage, len(node.Extends))
	for i := range node.Extends {
		ext := node.Extends[i]
		if ih, ok := c.ConvertNode(&ext).(*ast.TSInterfaceHeritage); ok && ih != nil {
			extends[i] = *ih
		}
	}

	result := &ast.TSInterfaceDeclaration{
		BaseNode:       c.copyBaseNode(&node.BaseNode),
		ID:             c.convertIdentifier(node.ID),
		Body:           c.convertTSInterfaceBody(node.Body),
		Extends:        extends,
		TypeParameters: c.convertTSTypeParameterDeclaration(node.TypeParameters),
		Declare:        node.Declare,
		Abstract:       node.Abstract,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertTSInterfaceBody converts a TSInterfaceBody node.
func (c *Converter) convertTSInterfaceBody(node *ast.TSInterfaceBody) *ast.TSInterfaceBody {
	if node == nil {
		return nil
	}

	body := make([]interface{}, len(node.Body))
	for i, member := range node.Body {
		if astNode, ok := member.(ast.Node); ok {
			body[i] = c.ConvertNode(astNode)
		}
	}

	result := &ast.TSInterfaceBody{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Body:     body,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertTSTypeAliasDeclaration converts a TSTypeAliasDeclaration node.
func (c *Converter) convertTSTypeAliasDeclaration(node *ast.TSTypeAliasDeclaration) *ast.TSTypeAliasDeclaration {
	if node == nil {
		return nil
	}

	var typeAnnotation ast.TSNode
	if node.TypeAnnotation != nil {
		if ts, ok := c.ConvertNode(node.TypeAnnotation).(ast.TSNode); ok {
			typeAnnotation = ts
		}
	}

	result := &ast.TSTypeAliasDeclaration{
		BaseNode:       c.copyBaseNode(&node.BaseNode),
		ID:             c.convertIdentifier(node.ID),
		TypeAnnotation: typeAnnotation,
		TypeParameters: c.convertTSTypeParameterDeclaration(node.TypeParameters),
		Declare:        node.Declare,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertTSEnumDeclaration converts a TSEnumDeclaration node.
func (c *Converter) convertTSEnumDeclaration(node *ast.TSEnumDeclaration) *ast.TSEnumDeclaration {
	if node == nil {
		return nil
	}

	members := make([]ast.TSEnumMember, len(node.Members))
	for i := range node.Members {
		member := node.Members[i]
		if em, ok := c.ConvertNode(&member).(*ast.TSEnumMember); ok && em != nil {
			members[i] = *em
		}
	}

	result := &ast.TSEnumDeclaration{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		ID:       c.convertIdentifier(node.ID),
		Members:  members,
		Const:    node.Const,
		Declare:  node.Declare,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertTSModuleDeclaration converts a TSModuleDeclaration node.
func (c *Converter) convertTSModuleDeclaration(node *ast.TSModuleDeclaration) *ast.TSModuleDeclaration {
	if node == nil {
		return nil
	}

	var id interface{}
	if node.ID != nil {
		if astNode, ok := node.ID.(ast.Node); ok {
			id = c.ConvertNode(astNode)
		}
	}

	var body interface{}
	if node.Body != nil {
		if astNode, ok := node.Body.(ast.Node); ok {
			body = c.ConvertNode(astNode)
		}
	}

	result := &ast.TSModuleDeclaration{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		ID:       id,
		Body:     body,
		Global:   node.Global,
		Declare:  node.Declare,
		Kind:     node.Kind,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertTSAsExpression converts a TSAsExpression node.
func (c *Converter) convertTSAsExpression(node *ast.TSAsExpression) *ast.TSAsExpression {
	if node == nil {
		return nil
	}

	var typeAnnotation ast.TSNode
	if node.TypeAnnotation != nil {
		if ts, ok := c.ConvertNode(node.TypeAnnotation).(ast.TSNode); ok {
			typeAnnotation = ts
		}
	}

	result := &ast.TSAsExpression{
		BaseNode:       c.copyBaseNode(&node.BaseNode),
		Expression:     c.convertExpression(node.Expression),
		TypeAnnotation: typeAnnotation,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertTSTypeAssertion converts a TSTypeAssertion node.
func (c *Converter) convertTSTypeAssertion(node *ast.TSTypeAssertion) *ast.TSTypeAssertion {
	if node == nil {
		return nil
	}

	var typeAnnotation ast.TSNode
	if node.TypeAnnotation != nil {
		if ts, ok := c.ConvertNode(node.TypeAnnotation).(ast.TSNode); ok {
			typeAnnotation = ts
		}
	}

	result := &ast.TSTypeAssertion{
		BaseNode:       c.copyBaseNode(&node.BaseNode),
		TypeAnnotation: typeAnnotation,
		Expression:     c.convertExpression(node.Expression),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertTSNonNullExpression converts a TSNonNullExpression node.
func (c *Converter) convertTSNonNullExpression(node *ast.TSNonNullExpression) *ast.TSNonNullExpression {
	if node == nil {
		return nil
	}

	result := &ast.TSNonNullExpression{
		BaseNode:   c.copyBaseNode(&node.BaseNode),
		Expression: c.convertExpression(node.Expression),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertTSTypeParameterDeclaration converts a TSTypeParameterDeclaration node.
func (c *Converter) convertTSTypeParameterDeclaration(node *ast.TSTypeParameterDeclaration) *ast.TSTypeParameterDeclaration {
	if node == nil {
		return nil
	}

	params := make([]ast.TSTypeParameter, len(node.Params))
	for i := range node.Params {
		param := node.Params[i]
		if tp, ok := c.ConvertNode(&param).(*ast.TSTypeParameter); ok && tp != nil {
			params[i] = *tp
		}
	}

	result := &ast.TSTypeParameterDeclaration{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Params:   params,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertTSTypeParameterInstantiation converts a TSTypeParameterInstantiation node.
func (c *Converter) convertTSTypeParameterInstantiation(node *ast.TSTypeParameterInstantiation) *ast.TSTypeParameterInstantiation {
	if node == nil {
		return nil
	}

	params := make([]ast.TSNode, len(node.Params))
	for i, param := range node.Params {
		if ts, ok := c.ConvertNode(param).(ast.TSNode); ok {
			params[i] = ts
		}
	}

	result := &ast.TSTypeParameterInstantiation{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Params:   params,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertTSClassImplements converts a slice of TSClassImplements nodes.
//
//nolint:unused // TODO: Will be used when needed for specific TypeScript conversions
func (c *Converter) convertTSClassImplements(implements []ast.TSClassImplements) []ast.TSClassImplements {
	if implements == nil {
		return nil
	}

	result := make([]ast.TSClassImplements, len(implements))
	for i := range implements {
		impl := implements[i]
		if ci, ok := c.ConvertNode(&impl).(*ast.TSClassImplements); ok && ci != nil {
			result[i] = *ci
		} else {
			result[i] = impl
		}
	}
	return result
}
