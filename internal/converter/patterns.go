package converter

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
)

// convertArrayPattern converts an ArrayPattern node.
func (c *Converter) convertArrayPattern(node *ast.ArrayPattern) *ast.ArrayPattern {
	if node == nil {
		return nil
	}

	elements := make([]ast.Pattern, len(node.Elements))
	for i, elem := range node.Elements {
		if elem != nil {
			if pattern, ok := c.ConvertNode(elem).(ast.Pattern); ok {
				elements[i] = pattern
			}
		}
	}

	result := &ast.ArrayPattern{
		BaseNode:       c.copyBaseNode(&node.BaseNode),
		Elements:       elements,
		TypeAnnotation: c.convertTSTypeAnnotation(node.TypeAnnotation),
		Optional:       node.Optional,
		Decorators:     c.convertDecorators(node.Decorators),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertObjectPattern converts an ObjectPattern node.
func (c *Converter) convertObjectPattern(node *ast.ObjectPattern) *ast.ObjectPattern {
	if node == nil {
		return nil
	}

	properties := make([]interface{}, len(node.Properties))
	for i, prop := range node.Properties {
		properties[i] = c.ConvertNode(prop.(ast.Node))
	}

	result := &ast.ObjectPattern{
		BaseNode:       c.copyBaseNode(&node.BaseNode),
		Properties:     properties,
		TypeAnnotation: c.convertTSTypeAnnotation(node.TypeAnnotation),
		Optional:       node.Optional,
		Decorators:     c.convertDecorators(node.Decorators),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertRestElement converts a RestElement node.
func (c *Converter) convertRestElement(node *ast.RestElement) *ast.RestElement {
	if node == nil {
		return nil
	}

	var argument ast.Pattern
	if node.Argument != nil {
		if pattern, ok := c.ConvertNode(node.Argument).(ast.Pattern); ok {
			argument = pattern
		}
	}

	result := &ast.RestElement{
		BaseNode:       c.copyBaseNode(&node.BaseNode),
		Argument:       argument,
		TypeAnnotation: c.convertTSTypeAnnotation(node.TypeAnnotation),
		Optional:       node.Optional,
		Value:          c.convertExpression(node.Value),
		Decorators:     c.convertDecorators(node.Decorators),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertAssignmentPattern converts an AssignmentPattern node.
func (c *Converter) convertAssignmentPattern(node *ast.AssignmentPattern) *ast.AssignmentPattern {
	if node == nil {
		return nil
	}

	var left ast.Pattern
	if node.Left != nil {
		if pattern, ok := c.ConvertNode(node.Left).(ast.Pattern); ok {
			left = pattern
		}
	}

	result := &ast.AssignmentPattern{
		BaseNode:       c.copyBaseNode(&node.BaseNode),
		Left:           left,
		Right:          c.convertExpression(node.Right),
		TypeAnnotation: c.convertTSTypeAnnotation(node.TypeAnnotation),
		Optional:       node.Optional,
		Decorators:     c.convertDecorators(node.Decorators),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertPatterns converts a slice of Pattern nodes.
func (c *Converter) convertPatterns(patterns []ast.Pattern) []ast.Pattern {
	if patterns == nil {
		return nil
	}

	result := make([]ast.Pattern, len(patterns))
	for i, pattern := range patterns {
		if pattern != nil {
			if p, ok := c.ConvertNode(pattern).(ast.Pattern); ok {
				result[i] = p
			}
		}
	}
	return result
}
