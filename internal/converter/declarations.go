package converter

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
)

// convertVariableDeclaration converts a VariableDeclaration node.
func (c *Converter) convertVariableDeclaration(node *ast.VariableDeclaration) *ast.VariableDeclaration {
	if node == nil {
		return nil
	}

	declarators := make([]ast.VariableDeclarator, len(node.Declarations))
	for i, decl := range node.Declarations {
		if vd, ok := c.ConvertNode(&decl).(*ast.VariableDeclarator); ok && vd != nil {
			declarators[i] = *vd
		} else {
			declarators[i] = decl
		}
	}

	result := &ast.VariableDeclaration{
		BaseNode:     c.copyBaseNode(&node.BaseNode),
		Declarations: declarators,
		Kind:         node.Kind,
		Declare:      node.Declare,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertFunctionDeclaration converts a FunctionDeclaration node.
func (c *Converter) convertFunctionDeclaration(node *ast.FunctionDeclaration) *ast.FunctionDeclaration {
	if node == nil {
		return nil
	}

	result := &ast.FunctionDeclaration{
		BaseNode:       c.copyBaseNode(&node.BaseNode),
		ID:             c.convertIdentifier(node.ID),
		Params:         c.convertPatterns(node.Params),
		Body:           c.convertBlockStatement(node.Body),
		Generator:      node.Generator,
		Async:          node.Async,
		Expression:     node.Expression,
		TypeParameters: c.convertTSTypeParameterDeclaration(node.TypeParameters),
		ReturnType:     c.convertTSTypeAnnotation(node.ReturnType),
		Declare:        node.Declare,
		Decorators:     c.convertDecorators(node.Decorators),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertClassDeclaration converts a ClassDeclaration node.
func (c *Converter) convertClassDeclaration(node *ast.ClassDeclaration) *ast.ClassDeclaration {
	if node == nil {
		return nil
	}

	implements := make([]ast.TSClassImplements, len(node.Implements))
	for i, impl := range node.Implements {
		if ci, ok := c.ConvertNode(&impl).(*ast.TSClassImplements); ok && ci != nil {
			implements[i] = *ci
		} else {
			implements[i] = impl
		}
	}

	result := &ast.ClassDeclaration{
		BaseNode:            c.copyBaseNode(&node.BaseNode),
		ID:                  c.convertIdentifier(node.ID),
		SuperClass:          c.convertExpression(node.SuperClass),
		Body:                c.convertClassBody(node.Body),
		TypeParameters:      c.convertTSTypeParameterDeclaration(node.TypeParameters),
		SuperTypeParameters: c.convertTSTypeParameterInstantiation(node.SuperTypeParameters),
		Implements:          implements,
		Decorators:          c.convertDecorators(node.Decorators),
		Abstract:            node.Abstract,
		Declare:             node.Declare,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertImportDeclaration converts an ImportDeclaration node.
func (c *Converter) convertImportDeclaration(node *ast.ImportDeclaration) *ast.ImportDeclaration {
	if node == nil {
		return nil
	}

	specifiers := make([]interface{}, len(node.Specifiers))
	for i, spec := range node.Specifiers {
		specifiers[i] = c.ConvertNode(spec.(ast.Node))
	}

	attributes := make([]ast.ImportAttribute, len(node.Attributes))
	for i, attr := range node.Attributes {
		if ia, ok := c.ConvertNode(&attr).(*ast.ImportAttribute); ok && ia != nil {
			attributes[i] = *ia
		} else {
			attributes[i] = attr
		}
	}

	result := &ast.ImportDeclaration{
		BaseNode:         c.copyBaseNode(&node.BaseNode),
		Specifiers:       specifiers,
		Source:           c.convertLiteral(node.Source),
		Attributes:       attributes,
		ImportKind:       node.ImportKind,
		AssertionEntries: attributes, // deprecated field, same as Attributes
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertExportNamedDeclaration converts an ExportNamedDeclaration node.
func (c *Converter) convertExportNamedDeclaration(node *ast.ExportNamedDeclaration) *ast.ExportNamedDeclaration {
	if node == nil {
		return nil
	}

	var declaration ast.Declaration
	if node.Declaration != nil {
		declaration = c.ConvertNode(node.Declaration).(ast.Declaration)
	}

	specifiers := make([]ast.ExportSpecifier, len(node.Specifiers))
	for i, spec := range node.Specifiers {
		if es, ok := c.ConvertNode(&spec).(*ast.ExportSpecifier); ok && es != nil {
			specifiers[i] = *es
		} else {
			specifiers[i] = spec
		}
	}

	attributes := make([]ast.ImportAttribute, len(node.Attributes))
	for i, attr := range node.Attributes {
		if ia, ok := c.ConvertNode(&attr).(*ast.ImportAttribute); ok && ia != nil {
			attributes[i] = *ia
		} else {
			attributes[i] = attr
		}
	}

	result := &ast.ExportNamedDeclaration{
		BaseNode:    c.copyBaseNode(&node.BaseNode),
		Declaration: declaration,
		Specifiers:  specifiers,
		Source:      c.convertLiteral(node.Source),
		ExportKind:  node.ExportKind,
		Attributes:  attributes,
		Assertions:  attributes, // deprecated field
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertExportDefaultDeclaration converts an ExportDefaultDeclaration node.
func (c *Converter) convertExportDefaultDeclaration(node *ast.ExportDefaultDeclaration) *ast.ExportDefaultDeclaration {
	if node == nil {
		return nil
	}

	var declaration interface{}
	if node.Declaration != nil {
		declaration = c.ConvertNode(node.Declaration.(ast.Node))
	}

	result := &ast.ExportDefaultDeclaration{
		BaseNode:    c.copyBaseNode(&node.BaseNode),
		Declaration: declaration,
		ExportKind:  node.ExportKind,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertExportAllDeclaration converts an ExportAllDeclaration node.
func (c *Converter) convertExportAllDeclaration(node *ast.ExportAllDeclaration) *ast.ExportAllDeclaration {
	if node == nil {
		return nil
	}

	attributes := make([]ast.ImportAttribute, len(node.Attributes))
	for i, attr := range node.Attributes {
		if ia, ok := c.ConvertNode(&attr).(*ast.ImportAttribute); ok && ia != nil {
			attributes[i] = *ia
		} else {
			attributes[i] = attr
		}
	}

	result := &ast.ExportAllDeclaration{
		BaseNode:   c.copyBaseNode(&node.BaseNode),
		Source:     c.convertLiteral(node.Source),
		Exported:   c.convertIdentifier(node.Exported),
		Attributes: attributes,
		ExportKind: node.ExportKind,
		Assertions: attributes, // deprecated field
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertClassBody converts a ClassBody node.
func (c *Converter) convertClassBody(node *ast.ClassBody) *ast.ClassBody {
	if node == nil {
		return nil
	}

	body := make([]interface{}, len(node.Body))
	for i, member := range node.Body {
		body[i] = c.ConvertNode(member.(ast.Node))
	}

	result := &ast.ClassBody{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Body:     body,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertDecorators converts a slice of Decorator nodes.
func (c *Converter) convertDecorators(decorators []ast.Decorator) []ast.Decorator {
	if decorators == nil {
		return nil
	}

	result := make([]ast.Decorator, len(decorators))
	for i, decorator := range decorators {
		if d, ok := c.ConvertNode(&decorator).(*ast.Decorator); ok && d != nil {
			result[i] = *d
		} else {
			result[i] = decorator
		}
	}
	return result
}
