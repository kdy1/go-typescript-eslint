package converter

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
)

// convertExpressionStatement converts an ExpressionStatement node.
func (c *Converter) convertExpressionStatement(node *ast.ExpressionStatement) *ast.ExpressionStatement {
	if node == nil {
		return nil
	}

	result := &ast.ExpressionStatement{
		BaseNode:   c.copyBaseNode(&node.BaseNode),
		Expression: c.convertExpression(node.Expression),
		Directive:  node.Directive,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertBlockStatement converts a BlockStatement node.
func (c *Converter) convertBlockStatement(node *ast.BlockStatement) *ast.BlockStatement {
	if node == nil {
		return nil
	}

	result := &ast.BlockStatement{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Body:     c.convertStatements(node.Body),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertEmptyStatement converts an EmptyStatement node.
func (c *Converter) convertEmptyStatement(node *ast.EmptyStatement) *ast.EmptyStatement {
	if node == nil {
		return nil
	}

	result := &ast.EmptyStatement{
		BaseNode: c.copyBaseNode(&node.BaseNode),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertIfStatement converts an IfStatement node.
func (c *Converter) convertIfStatement(node *ast.IfStatement) *ast.IfStatement {
	if node == nil {
		return nil
	}

	result := &ast.IfStatement{
		BaseNode:   c.copyBaseNode(&node.BaseNode),
		Test:       c.convertExpression(node.Test),
		Consequent: c.convertStatement(node.Consequent),
		Alternate:  c.convertStatement(node.Alternate),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertSwitchStatement converts a SwitchStatement node.
func (c *Converter) convertSwitchStatement(node *ast.SwitchStatement) *ast.SwitchStatement {
	if node == nil {
		return nil
	}

	cases := make([]ast.SwitchCase, len(node.Cases))
	for i := range node.Cases {
		caseNode := node.Cases[i]
		if switchCase, ok := c.ConvertNode(&caseNode).(*ast.SwitchCase); ok && switchCase != nil {
			cases[i] = *switchCase
		}
	}

	result := &ast.SwitchStatement{
		BaseNode:     c.copyBaseNode(&node.BaseNode),
		Discriminant: c.convertExpression(node.Discriminant),
		Cases:        cases,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertWhileStatement converts a WhileStatement node.
func (c *Converter) convertWhileStatement(node *ast.WhileStatement) *ast.WhileStatement {
	if node == nil {
		return nil
	}

	result := &ast.WhileStatement{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Test:     c.convertExpression(node.Test),
		Body:     c.convertStatement(node.Body),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertDoWhileStatement converts a DoWhileStatement node.
func (c *Converter) convertDoWhileStatement(node *ast.DoWhileStatement) *ast.DoWhileStatement {
	if node == nil {
		return nil
	}

	result := &ast.DoWhileStatement{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Body:     c.convertStatement(node.Body),
		Test:     c.convertExpression(node.Test),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertForStatement converts a ForStatement node.
func (c *Converter) convertForStatement(node *ast.ForStatement) *ast.ForStatement {
	if node == nil {
		return nil
	}

	var init interface{}
	if node.Init != nil {
		if astNode, ok := node.Init.(ast.Node); ok {
			init = c.ConvertNode(astNode)
		}
	}

	result := &ast.ForStatement{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Init:     init,
		Test:     c.convertExpression(node.Test),
		Update:   c.convertExpression(node.Update),
		Body:     c.convertStatement(node.Body),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertForInStatement converts a ForInStatement node.
func (c *Converter) convertForInStatement(node *ast.ForInStatement) *ast.ForInStatement {
	if node == nil {
		return nil
	}

	var left interface{}
	if node.Left != nil {
		if astNode, ok := node.Left.(ast.Node); ok {
			left = c.ConvertNode(astNode)
		}
	}

	result := &ast.ForInStatement{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Left:     left,
		Right:    c.convertExpression(node.Right),
		Body:     c.convertStatement(node.Body),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertForOfStatement converts a ForOfStatement node.
func (c *Converter) convertForOfStatement(node *ast.ForOfStatement) *ast.ForOfStatement {
	if node == nil {
		return nil
	}

	var left interface{}
	if node.Left != nil {
		if astNode, ok := node.Left.(ast.Node); ok {
			left = c.ConvertNode(astNode)
		}
	}

	result := &ast.ForOfStatement{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Left:     left,
		Right:    c.convertExpression(node.Right),
		Body:     c.convertStatement(node.Body),
		Await:    node.Await,
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertBreakStatement converts a BreakStatement node.
func (c *Converter) convertBreakStatement(node *ast.BreakStatement) *ast.BreakStatement {
	if node == nil {
		return nil
	}

	result := &ast.BreakStatement{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Label:    c.convertIdentifier(node.Label),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertContinueStatement converts a ContinueStatement node.
func (c *Converter) convertContinueStatement(node *ast.ContinueStatement) *ast.ContinueStatement {
	if node == nil {
		return nil
	}

	result := &ast.ContinueStatement{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Label:    c.convertIdentifier(node.Label),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertReturnStatement converts a ReturnStatement node.
func (c *Converter) convertReturnStatement(node *ast.ReturnStatement) *ast.ReturnStatement {
	if node == nil {
		return nil
	}

	result := &ast.ReturnStatement{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Argument: c.convertExpression(node.Argument),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertThrowStatement converts a ThrowStatement node.
func (c *Converter) convertThrowStatement(node *ast.ThrowStatement) *ast.ThrowStatement {
	if node == nil {
		return nil
	}

	result := &ast.ThrowStatement{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Argument: c.convertExpression(node.Argument),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertTryStatement converts a TryStatement node.
func (c *Converter) convertTryStatement(node *ast.TryStatement) *ast.TryStatement {
	if node == nil {
		return nil
	}

	var handler *ast.CatchClause
	if node.Handler != nil {
		if cc, ok := c.ConvertNode(node.Handler).(*ast.CatchClause); ok {
			handler = cc
		}
	}

	result := &ast.TryStatement{
		BaseNode:  c.copyBaseNode(&node.BaseNode),
		Block:     c.convertBlockStatement(node.Block),
		Handler:   handler,
		Finalizer: c.convertBlockStatement(node.Finalizer),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertWithStatement converts a WithStatement node.
func (c *Converter) convertWithStatement(node *ast.WithStatement) *ast.WithStatement {
	if node == nil {
		return nil
	}

	result := &ast.WithStatement{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Object:   c.convertExpression(node.Object),
		Body:     c.convertStatement(node.Body),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertLabeledStatement converts a LabeledStatement node.
func (c *Converter) convertLabeledStatement(node *ast.LabeledStatement) *ast.LabeledStatement {
	if node == nil {
		return nil
	}

	result := &ast.LabeledStatement{
		BaseNode: c.copyBaseNode(&node.BaseNode),
		Label:    c.convertIdentifier(node.Label),
		Body:     c.convertStatement(node.Body),
	}

	c.registerNodeMapping(node, result)
	return result
}

// convertDebuggerStatement converts a DebuggerStatement node.
func (c *Converter) convertDebuggerStatement(node *ast.DebuggerStatement) *ast.DebuggerStatement {
	if node == nil {
		return nil
	}

	result := &ast.DebuggerStatement{
		BaseNode: c.copyBaseNode(&node.BaseNode),
	}

	c.registerNodeMapping(node, result)
	return result
}

// Helper methods

// convertStatement converts a single statement node.
func (c *Converter) convertStatement(stmt ast.Statement) ast.Statement {
	if stmt == nil {
		return nil
	}
	converted := c.ConvertNode(stmt)
	if converted == nil {
		return nil
	}
	if statement, ok := converted.(ast.Statement); ok {
		return statement
	}
	return nil
}

// convertStatements converts a slice of statement nodes.
func (c *Converter) convertStatements(stmts []ast.Statement) []ast.Statement {
	if stmts == nil {
		return nil
	}

	result := make([]ast.Statement, len(stmts))
	for i, stmt := range stmts {
		result[i] = c.convertStatement(stmt)
	}
	return result
}
