package converter

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
)

// convertIdentifier converts an Identifier node.
func (c *Converter) convertIdentifier(node *ast.Identifier) *ast.Identifier {
	if node == nil {
		return nil
	}

	result := &ast.Identifier{
		BaseNode:       c.copyBaseNode(&node.BaseNode),
		Name:           node.Name,
		Optional:       node.Optional,
		TypeAnnotation: c.convertTSTypeAnnotation(node.TypeAnnotation),
		Decorators:     c.convertDecorators(node.Decorators),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertPrivateIdentifier converts a PrivateIdentifier node.
func (c *Converter) convertPrivateIdentifier(node *ast.PrivateIdentifier) *ast.PrivateIdentifier {
	if node == nil {
		return nil
	}

	result := &ast.PrivateIdentifier{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Name:     node.Name,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertLiteral converts a Literal node.
func (c *Converter) convertLiteral(node *ast.Literal) *ast.Literal {
	if node == nil {
		return nil
	}

	result := &ast.Literal{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Value:    node.Value,
		Raw:      node.Raw,
		Regex:    node.Regex,
		BigInt:   node.BigInt,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertThisExpression converts a ThisExpression node.
func (c *Converter) convertThisExpression(node *ast.ThisExpression) *ast.ThisExpression {
	if node == nil {
		return nil
	}

	result := &ast.ThisExpression{
		BaseNode: c.copyBaseNode(&node.BaseNode),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertArrayExpression converts an ArrayExpression node.
func (c *Converter) convertArrayExpression(node *ast.ArrayExpression) *ast.ArrayExpression {
	if node == nil {
		return nil
	}

	result := &ast.ArrayExpression{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Elements: c.convertExpressions(node.Elements),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertObjectExpression converts an ObjectExpression node.
func (c *Converter) convertObjectExpression(node *ast.ObjectExpression) *ast.ObjectExpression {
	if node == nil {
		return nil
	}

	properties := make([]interface{}, len(node.Properties))
	for i, prop := range node.Properties {
		if astNode, ok := prop.(ast.Node); ok {
			properties[i] = c.ConvertNode(astNode)
		}
	}

	result := &ast.ObjectExpression{
		BaseNode:   c.copyBaseNode(&node.BaseNode),
		Properties: properties,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertFunctionExpression converts a FunctionExpression node.
func (c *Converter) convertFunctionExpression(node *ast.FunctionExpression) *ast.FunctionExpression {
	if node == nil {
		return nil
	}

	result := &ast.FunctionExpression{
		BaseNode:       c.copyBaseNode(&node.BaseNode),
		ID:             c.convertIdentifier(node.ID),
		Params:         c.convertPatterns(node.Params),
		Body:           c.convertBlockStatement(node.Body),
		Generator:      node.Generator,
		Async:          node.Async,
		Expression:     node.Expression,
		TypeParameters: c.convertTSTypeParameterDeclaration(node.TypeParameters),
		ReturnType:     c.convertTSTypeAnnotation(node.ReturnType),
		Decorators:     c.convertDecorators(node.Decorators),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertArrowFunctionExpression converts an ArrowFunctionExpression node.
func (c *Converter) convertArrowFunctionExpression(node *ast.ArrowFunctionExpression) *ast.ArrowFunctionExpression {
	if node == nil {
		return nil
	}

	var body interface{}
	if node.Body != nil {
		if astNode, ok := node.Body.(ast.Node); ok {
			body = c.ConvertNode(astNode)
		}
	}

	result := &ast.ArrowFunctionExpression{
		BaseNode:       c.copyBaseNode(&node.BaseNode),
		Params:         c.convertPatterns(node.Params),
		Body:           body,
		Async:          node.Async,
		Expression:     node.Expression,
		TypeParameters: c.convertTSTypeParameterDeclaration(node.TypeParameters),
		ReturnType:     c.convertTSTypeAnnotation(node.ReturnType),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertClassExpression converts a ClassExpression node.
func (c *Converter) convertClassExpression(node *ast.ClassExpression) *ast.ClassExpression {
	if node == nil {
		return nil
	}

	implements := make([]ast.TSClassImplements, len(node.Implements))
	for i := range node.Implements {
		impl := node.Implements[i]
		if ci, ok := c.ConvertNode(&impl).(*ast.TSClassImplements); ok && ci != nil {
			implements[i] = *ci
		} else {
			implements[i] = impl
		}
	}

	result := &ast.ClassExpression{
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

// convertUnaryExpression converts a UnaryExpression node.
func (c *Converter) convertUnaryExpression(node *ast.UnaryExpression) *ast.UnaryExpression {
	if node == nil {
		return nil
	}

	result := &ast.UnaryExpression{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Operator: node.Operator,
		Prefix:   node.Prefix,
		Argument: c.convertExpression(node.Argument),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertUpdateExpression converts an UpdateExpression node.
func (c *Converter) convertUpdateExpression(node *ast.UpdateExpression) *ast.UpdateExpression {
	if node == nil {
		return nil
	}

	result := &ast.UpdateExpression{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Operator: node.Operator,
		Argument: c.convertExpression(node.Argument),
		Prefix:   node.Prefix,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertBinaryExpression converts a BinaryExpression node.
func (c *Converter) convertBinaryExpression(node *ast.BinaryExpression) *ast.BinaryExpression {
	if node == nil {
		return nil
	}

	result := &ast.BinaryExpression{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Operator: node.Operator,
		Left:     c.convertExpression(node.Left),
		Right:    c.convertExpression(node.Right),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertLogicalExpression converts a LogicalExpression node.
func (c *Converter) convertLogicalExpression(node *ast.LogicalExpression) *ast.LogicalExpression {
	if node == nil {
		return nil
	}

	result := &ast.LogicalExpression{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Operator: node.Operator,
		Left:     c.convertExpression(node.Left),
		Right:    c.convertExpression(node.Right),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertAssignmentExpression converts an AssignmentExpression node.
func (c *Converter) convertAssignmentExpression(node *ast.AssignmentExpression) *ast.AssignmentExpression {
	if node == nil {
		return nil
	}

	var left ast.Pattern
	if convertedLeft := c.ConvertNode(node.Left); convertedLeft != nil {
		if pattern, ok := convertedLeft.(ast.Pattern); ok {
			left = pattern
		}
	}

	result := &ast.AssignmentExpression{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Operator: node.Operator,
		Left:     left,
		Right:    c.convertExpression(node.Right),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertMemberExpression converts a MemberExpression node.
func (c *Converter) convertMemberExpression(node *ast.MemberExpression) *ast.MemberExpression {
	if node == nil {
		return nil
	}

	result := &ast.MemberExpression{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Object:   c.convertExpression(node.Object),
		Property: c.convertExpression(node.Property),
		Computed: node.Computed,
		Optional: node.Optional,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertConditionalExpression converts a ConditionalExpression node.
func (c *Converter) convertConditionalExpression(node *ast.ConditionalExpression) *ast.ConditionalExpression {
	if node == nil {
		return nil
	}

	result := &ast.ConditionalExpression{
		BaseNode:   c.copyBaseNode(&node.BaseNode),
		Test:       c.convertExpression(node.Test),
		Consequent: c.convertExpression(node.Consequent),
		Alternate:  c.convertExpression(node.Alternate),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertCallExpression converts a CallExpression node.
func (c *Converter) convertCallExpression(node *ast.CallExpression) *ast.CallExpression {
	if node == nil {
		return nil
	}

	result := &ast.CallExpression{
		BaseNode:      c.copyBaseNode(&node.BaseNode),
		Callee:        c.convertExpression(node.Callee),
		Arguments:     c.convertExpressions(node.Arguments),
		Optional:      node.Optional,
		TypeArguments: c.convertTSTypeParameterInstantiation(node.TypeArguments),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertNewExpression converts a NewExpression node.
func (c *Converter) convertNewExpression(node *ast.NewExpression) *ast.NewExpression {
	if node == nil {
		return nil
	}

	result := &ast.NewExpression{
		BaseNode:      c.copyBaseNode(&node.BaseNode),
		Callee:        c.convertExpression(node.Callee),
		Arguments:     c.convertExpressions(node.Arguments),
		TypeArguments: c.convertTSTypeParameterInstantiation(node.TypeArguments),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertSequenceExpression converts a SequenceExpression node.
func (c *Converter) convertSequenceExpression(node *ast.SequenceExpression) *ast.SequenceExpression {
	if node == nil {
		return nil
	}

	result := &ast.SequenceExpression{
		BaseNode:    c.copyBaseNode(&node.BaseNode),
		Expressions: c.convertExpressions(node.Expressions),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertTemplateLiteral converts a TemplateLiteral node.
func (c *Converter) convertTemplateLiteral(node *ast.TemplateLiteral) *ast.TemplateLiteral {
	if node == nil {
		return nil
	}

	quasis := make([]ast.TemplateElement, len(node.Quasis))
	for i := range node.Quasis {
		quasi := node.Quasis[i]
		if tpl, ok := c.ConvertNode(&quasi).(*ast.TemplateElement); ok && tpl != nil {
			quasis[i] = *tpl
		}
	}

	result := &ast.TemplateLiteral{
		BaseNode:    c.copyBaseNode(&node.BaseNode),
		Quasis:      quasis,
		Expressions: c.convertExpressions(node.Expressions),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertTaggedTemplateExpression converts a TaggedTemplateExpression node.
func (c *Converter) convertTaggedTemplateExpression(node *ast.TaggedTemplateExpression) *ast.TaggedTemplateExpression {
	if node == nil {
		return nil
	}

	result := &ast.TaggedTemplateExpression{
		BaseNode:      c.copyBaseNode(&node.BaseNode),
		Tag:           c.convertExpression(node.Tag),
		Quasi:         c.convertTemplateLiteral(node.Quasi),
		TypeArguments: c.convertTSTypeParameterInstantiation(node.TypeArguments),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertYieldExpression converts a YieldExpression node.
func (c *Converter) convertYieldExpression(node *ast.YieldExpression) *ast.YieldExpression {
	if node == nil {
		return nil
	}

	result := &ast.YieldExpression{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Argument: c.convertExpression(node.Argument),
		Delegate: node.Delegate,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertAwaitExpression converts an AwaitExpression node.
func (c *Converter) convertAwaitExpression(node *ast.AwaitExpression) *ast.AwaitExpression {
	if node == nil {
		return nil
	}

	result := &ast.AwaitExpression{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Argument: c.convertExpression(node.Argument),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertChainExpression converts a ChainExpression node.
func (c *Converter) convertChainExpression(node *ast.ChainExpression) *ast.ChainExpression {
	if node == nil {
		return nil
	}

	result := &ast.ChainExpression{
		BaseNode:   c.copyBaseNode(&node.BaseNode),
		Expression: c.convertExpression(node.Expression),
	}

	c.registerNodeMapping(node, result)
	return result
}

// Helper methods

// convertExpression converts a single expression node.
func (c *Converter) convertExpression(expr ast.Expression) ast.Expression {
	if expr == nil {
		return nil
	}
	converted := c.ConvertNode(expr)
	if converted == nil {
		return nil
	}
	if expression, ok := converted.(ast.Expression); ok {
		return expression
	}
	return nil
}

// convertExpressions converts a slice of expression nodes.
func (c *Converter) convertExpressions(exprs []ast.Expression) []ast.Expression {
	if exprs == nil {
		return nil
	}

	result := make([]ast.Expression, len(exprs))
	for i, expr := range exprs {
		result[i] = c.convertExpression(expr)
	}
	return result
}
