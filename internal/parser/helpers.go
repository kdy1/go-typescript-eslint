package parser

import "github.com/kdy1/go-typescript-eslint/internal/ast"

// makeRange creates a pointer to a Range from start and end positions.
func makeRange(start, end int) *ast.Range {
	r := ast.Range{start, end}
	return &r
}

// makeBaseNode creates a BaseNode with the given type and range.
func makeBaseNode(nodeType ast.NodeType, start, end int) ast.BaseNode {
	return ast.BaseNode{
		NodeType: nodeType.String(),
		Range:    makeRange(start, end),
		Start:    start,
		EndPos:   end,
	}
}
