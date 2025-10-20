package ast

// Node is the base interface for all AST nodes.
// All AST node types must implement this interface.
type Node interface {
	// Type returns the type of the node (e.g., "Program", "Identifier").
	Type() string

	// Pos returns the start position of the node in the source.
	Pos() int

	// End returns the end position of the node in the source.
	End() int
}

// BaseNode provides common fields for all AST nodes.
// It should be embedded in all concrete node types.
//
//nolint:govet // Field order optimized for JSON output readability, not memory alignment
type BaseNode struct {
	NodeType string          `json:"type"`
	Loc      *SourceLocation `json:"loc,omitempty"`
	Range    *Range          `json:"range,omitempty"`
	Start    int             `json:"-"` // Internal use, not serialized
	EndPos   int             `json:"-"` // Internal use, not serialized
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
// It contains the start and end positions with line and column information.
type SourceLocation struct {
	Start    Position `json:"start"`
	End      Position `json:"end"`
	Filename string   `json:"source,omitempty"`
}

// Position represents a position in source code.
type Position struct {
	Line   int `json:"line"`   // 1-based line number
	Column int `json:"column"` // 0-based column number
}

// Expression is the interface for expression nodes.
// Expressions are nodes that produce a value.
type Expression interface {
	Node
	expressionNode()
}

// Statement is the interface for statement nodes.
// Statements are nodes that perform actions.
type Statement interface {
	Node
	statementNode()
}

// Pattern is the interface for pattern nodes (used in destructuring).
// Patterns can appear in variable declarations, function parameters, and assignments.
type Pattern interface {
	Node
	patternNode()
}

// Declaration is the interface for declaration nodes.
// Declarations are a subset of statements that declare variables, functions, or classes.
type Declaration interface {
	Statement
	declarationNode()
}
