package ast

// Node is the base interface for all AST nodes.
type Node interface {
	// Type returns the type of the node (e.g., "Program", "Identifier").
	Type() string

	// Pos returns the start position of the node in the source.
	Pos() int

	// End returns the end position of the node in the source.
	End() int
}

// BaseNode provides common fields for all AST nodes.
type BaseNode struct {
	NodeType  string
	Start     int
	EndPos    int
	Loc       *SourceLocation
	Range     *[2]int
}

// Type returns the type of the node.
func (n *BaseNode) Type() string {
	return n.NodeType
}

// Pos returns the start position of the node.
func (n *BaseNode) Pos() int {
	return n.Start
}

// End returns the end position of the node.
func (n *BaseNode) End() int {
	return n.EndPos
}

// SourceLocation represents the location of a node in source code.
type SourceLocation struct {
	Start    Position
	End      Position
	Filename string
}

// Position represents a position in source code.
type Position struct {
	Line   int // 1-based line number
	Column int // 0-based column number
}

// Expression is the interface for expression nodes.
type Expression interface {
	Node
	expressionNode()
}

// Statement is the interface for statement nodes.
type Statement interface {
	Node
	statementNode()
}

// Pattern is the interface for pattern nodes (used in destructuring).
type Pattern interface {
	Node
	patternNode()
}

// Declaration is the interface for declaration nodes.
type Declaration interface {
	Statement
	declarationNode()
}
