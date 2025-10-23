package converter

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
)

// Converter transforms TypeScript AST nodes into ESTree-compatible format.
// It maintains bidirectional mappings between TypeScript and ESTree nodes
// for use by ParserServices in type-aware linting.
type Converter struct {
	// Source code being converted
	source string

	// Conversion options
	options *Options

	// Bidirectional node mappings for ParserServices
	esTreeNodeToTSNodeMap map[ast.Node]ast.Node
	tsNodeToESTreeNodeMap map[ast.Node]ast.Node

	// Context state during conversion
	allowPattern bool // Whether patterns (destructuring) are allowed in current context
}

// Options configures the AST conversion process.
type Options struct {
	// PreserveNodeMaps determines whether to maintain bidirectional node mappings.
	// These mappings are required for type-aware linting rules.
	PreserveNodeMaps bool

	// UseJSDocParsingMode enables JSDoc comment parsing for type information.
	UseJSDocParsingMode bool

	// SuppressDeprecatedPropertyWarnings disables warnings for deprecated properties.
	SuppressDeprecatedPropertyWarnings bool
}

// NodeMaps contains the bidirectional mappings between TypeScript and ESTree nodes.
// These maps are used by ParserServices to correlate ESTree nodes with TypeScript's
// type checker and program information.
type NodeMaps struct {
	// ESTreeNodeToTSNodeMap maps ESTree nodes to their original TypeScript nodes.
	ESTreeNodeToTSNodeMap map[ast.Node]ast.Node

	// TSNodeToESTreeNodeMap maps TypeScript nodes to their converted ESTree nodes.
	TSNodeToESTreeNodeMap map[ast.Node]ast.Node
}

// NewConverter creates a new AST converter with the given source code and options.
func NewConverter(source string, opts *Options) *Converter {
	if opts == nil {
		opts = &Options{
			PreserveNodeMaps: true,
		}
	}

	return &Converter{
		source:                source,
		options:               opts,
		esTreeNodeToTSNodeMap: make(map[ast.Node]ast.Node),
		tsNodeToESTreeNodeMap: make(map[ast.Node]ast.Node),
		allowPattern:          false,
	}
}

// ConvertProgram converts a TypeScript Program node to ESTree format.
// This is the main entry point for AST conversion.
func (c *Converter) ConvertProgram(program *ast.Program) *ast.Program {
	if program == nil {
		return nil
	}

	// The program node is already in ESTree format from our parser,
	// but we need to ensure all child nodes are properly converted
	// and mappings are registered.

	result := &ast.Program{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeProgram.String(),
			Loc:      program.Loc,
			Range:    program.Range,
			Start:    program.Start,
			EndPos:   program.EndPos,
		},
		SourceType: program.SourceType,
		Body:       c.convertStatements(program.Body),
		Comments:   c.convertComments(program.Comments),
		Tokens:     c.convertTokens(program.Tokens),
		Decorators: c.convertDecorators(program.Decorators),
	}

	// Register the program node mapping
	c.registerNodeMapping(program, result)

	return result
}

// ConvertNode converts any AST node to ESTree format.
// This method dispatches to specific conversion methods based on node type.
func (c *Converter) ConvertNode(node ast.Node) ast.Node {
	if node == nil {
		return nil
	}

	// Check if we've already converted this node
	if converted, exists := c.tsNodeToESTreeNodeMap[node]; exists {
		return converted
	}

	var result ast.Node

	// Dispatch to appropriate conversion method based on node type
	switch n := node.(type) {
	// Program
	case *ast.Program:
		result = c.ConvertProgram(n)

	// Identifiers
	case *ast.Identifier:
		result = c.convertIdentifier(n)
	case *ast.PrivateIdentifier:
		result = c.convertPrivateIdentifier(n)

	// Literals
	case *ast.Literal:
		result = c.convertLiteral(n)

	// Expressions
	case *ast.ThisExpression:
		result = c.convertThisExpression(n)
	case *ast.ArrayExpression:
		result = c.convertArrayExpression(n)
	case *ast.ObjectExpression:
		result = c.convertObjectExpression(n)
	case *ast.FunctionExpression:
		result = c.convertFunctionExpression(n)
	case *ast.ArrowFunctionExpression:
		result = c.convertArrowFunctionExpression(n)
	case *ast.ClassExpression:
		result = c.convertClassExpression(n)
	case *ast.UnaryExpression:
		result = c.convertUnaryExpression(n)
	case *ast.UpdateExpression:
		result = c.convertUpdateExpression(n)
	case *ast.BinaryExpression:
		result = c.convertBinaryExpression(n)
	case *ast.LogicalExpression:
		result = c.convertLogicalExpression(n)
	case *ast.AssignmentExpression:
		result = c.convertAssignmentExpression(n)
	case *ast.MemberExpression:
		result = c.convertMemberExpression(n)
	case *ast.ConditionalExpression:
		result = c.convertConditionalExpression(n)
	case *ast.CallExpression:
		result = c.convertCallExpression(n)
	case *ast.NewExpression:
		result = c.convertNewExpression(n)
	case *ast.SequenceExpression:
		result = c.convertSequenceExpression(n)
	case *ast.TemplateLiteral:
		result = c.convertTemplateLiteral(n)
	case *ast.TaggedTemplateExpression:
		result = c.convertTaggedTemplateExpression(n)
	case *ast.YieldExpression:
		result = c.convertYieldExpression(n)
	case *ast.AwaitExpression:
		result = c.convertAwaitExpression(n)
	case *ast.ChainExpression:
		result = c.convertChainExpression(n)

	// Statements
	case *ast.ExpressionStatement:
		result = c.convertExpressionStatement(n)
	case *ast.BlockStatement:
		result = c.convertBlockStatement(n)
	case *ast.EmptyStatement:
		result = c.convertEmptyStatement(n)
	case *ast.IfStatement:
		result = c.convertIfStatement(n)
	case *ast.SwitchStatement:
		result = c.convertSwitchStatement(n)
	case *ast.WhileStatement:
		result = c.convertWhileStatement(n)
	case *ast.DoWhileStatement:
		result = c.convertDoWhileStatement(n)
	case *ast.ForStatement:
		result = c.convertForStatement(n)
	case *ast.ForInStatement:
		result = c.convertForInStatement(n)
	case *ast.ForOfStatement:
		result = c.convertForOfStatement(n)
	case *ast.BreakStatement:
		result = c.convertBreakStatement(n)
	case *ast.ContinueStatement:
		result = c.convertContinueStatement(n)
	case *ast.ReturnStatement:
		result = c.convertReturnStatement(n)
	case *ast.ThrowStatement:
		result = c.convertThrowStatement(n)
	case *ast.TryStatement:
		result = c.convertTryStatement(n)
	case *ast.WithStatement:
		result = c.convertWithStatement(n)
	case *ast.LabeledStatement:
		result = c.convertLabeledStatement(n)
	case *ast.DebuggerStatement:
		result = c.convertDebuggerStatement(n)

	// Declarations
	case *ast.VariableDeclaration:
		result = c.convertVariableDeclaration(n)
	case *ast.FunctionDeclaration:
		result = c.convertFunctionDeclaration(n)
	case *ast.ClassDeclaration:
		result = c.convertClassDeclaration(n)

	// Module imports/exports
	case *ast.ImportDeclaration:
		result = c.convertImportDeclaration(n)
	case *ast.ExportNamedDeclaration:
		result = c.convertExportNamedDeclaration(n)
	case *ast.ExportDefaultDeclaration:
		result = c.convertExportDefaultDeclaration(n)
	case *ast.ExportAllDeclaration:
		result = c.convertExportAllDeclaration(n)

	// Patterns
	case *ast.ArrayPattern:
		result = c.convertArrayPattern(n)
	case *ast.ObjectPattern:
		result = c.convertObjectPattern(n)
	case *ast.RestElement:
		result = c.convertRestElement(n)
	case *ast.AssignmentPattern:
		result = c.convertAssignmentPattern(n)

	// TypeScript-specific nodes
	case *ast.TSTypeAnnotation:
		result = c.convertTSTypeAnnotation(n)
	case *ast.TSInterfaceDeclaration:
		result = c.convertTSInterfaceDeclaration(n)
	case *ast.TSTypeAliasDeclaration:
		result = c.convertTSTypeAliasDeclaration(n)
	case *ast.TSEnumDeclaration:
		result = c.convertTSEnumDeclaration(n)
	case *ast.TSModuleDeclaration:
		result = c.convertTSModuleDeclaration(n)
	case *ast.TSAsExpression:
		result = c.convertTSAsExpression(n)
	case *ast.TSTypeAssertion:
		result = c.convertTSTypeAssertion(n)
	case *ast.TSNonNullExpression:
		result = c.convertTSNonNullExpression(n)

	default:
		// For nodes that don't need conversion, return as-is
		result = node
	}

	// Register the mapping if node mappings are enabled
	if c.options.PreserveNodeMaps && result != nil && result != node {
		c.registerNodeMapping(node, result)
	}

	return result
}

// registerNodeMapping registers bidirectional mappings between TypeScript and ESTree nodes.
func (c *Converter) registerNodeMapping(tsNode, esTreeNode ast.Node) {
	if !c.options.PreserveNodeMaps {
		return
	}

	c.tsNodeToESTreeNodeMap[tsNode] = esTreeNode
	c.esTreeNodeToTSNodeMap[esTreeNode] = tsNode
}

// GetNodeMaps returns the bidirectional node mappings for ParserServices.
func (c *Converter) GetNodeMaps() *NodeMaps {
	return &NodeMaps{
		ESTreeNodeToTSNodeMap: c.esTreeNodeToTSNodeMap,
		TSNodeToESTreeNodeMap: c.tsNodeToESTreeNodeMap,
	}
}

// withAllowPattern executes a function with allowPattern set to the given value,
// then restores the previous value.
func (c *Converter) withAllowPattern(allow bool, fn func()) {
	old := c.allowPattern
	c.allowPattern = allow
	fn()
	c.allowPattern = old
}
