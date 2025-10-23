package converter

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
)

// copyBaseNode creates a copy of a BaseNode.
func (c *Converter) copyBaseNode(node *ast.BaseNode) ast.BaseNode {
	if node == nil {
		return ast.BaseNode{}
	}

	return ast.BaseNode{
		NodeType: node.NodeType,
		Loc:      node.Loc,
		Range:    node.Range,
		Start:    node.Start,
		EndPos:   node.EndPos,
	}
}

// convertComments converts a slice of Comment nodes.
func (c *Converter) convertComments(comments []ast.Comment) []ast.Comment {
	if comments == nil {
		return nil
	}

	// Comments are already in the correct format, but we create a new slice
	// to avoid modifying the original
	result := make([]ast.Comment, len(comments))
	copy(result, comments)
	return result
}

// convertTokens converts a slice of Token nodes.
func (c *Converter) convertTokens(tokens []ast.Token) []ast.Token {
	if tokens == nil {
		return nil
	}

	// Tokens are already in ESTree format from the parser
	// Create a new slice to avoid modifying the original
	result := make([]ast.Token, len(tokens))
	copy(result, tokens)
	return result
}

// attachCommentsToNodes attaches comments to appropriate AST nodes.
// This implements a simple algorithm that associates comments with
// the nearest following or preceding node based on position.
func (c *Converter) attachCommentsToNodes(program *ast.Program) {
	if program == nil || len(program.Comments) == 0 {
		return
	}

	// TODO: Implement comment attachment logic
	// This is a placeholder for the actual implementation
	// which would:
	// 1. Sort comments by position
	// 2. Traverse the AST and attach comments to nodes
	// 3. Handle leading, trailing, and inner comments
	// 4. Handle special cases like JSDoc comments
}

// createSourceLocation creates a SourceLocation from start and end positions.
func (c *Converter) createSourceLocation(start, end int) *ast.SourceLocation {
	// Convert byte positions to line/column positions
	// This is a simplified version - actual implementation would
	// need to track line breaks in the source code

	return &ast.SourceLocation{
		Start: ast.Position{
			Line:   1, // TODO: Calculate actual line
			Column: start,
		},
		End: ast.Position{
			Line:   1, // TODO: Calculate actual line
			Column: end,
		},
	}
}

// isPatternContext returns true if we're in a context where patterns are allowed.
func (c *Converter) isPatternContext() bool {
	return c.allowPattern
}
