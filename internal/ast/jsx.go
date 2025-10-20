package ast

// This file defines JSX-specific AST node types.
// Based on: https://github.com/facebook/jsx

// ==================== JSX Elements ====================

// JSXElement represents a JSX element.
//
//nolint:govet // Field order optimized for JSON output readability
type JSXElement struct {
	BaseNode
	OpeningElement  *JSXOpeningElement  `json:"openingElement"`
	Children        []interface{}       `json:"children"` // JSXText | JSXExpressionContainer | JSXSpreadChild | JSXElement | JSXFragment
	ClosingElement  *JSXClosingElement  `json:"closingElement"`
}

func (n *JSXElement) expressionNode() {}

// JSXFragment represents a JSX fragment (<>...</>).
//
//nolint:govet // Field order optimized for JSON output readability
type JSXFragment struct {
	BaseNode
	OpeningFragment *JSXOpeningFragment `json:"openingFragment"`
	Children        []interface{}       `json:"children"`
	ClosingFragment *JSXClosingFragment `json:"closingFragment"`
}

func (n *JSXFragment) expressionNode() {}

// JSXOpeningElement represents a JSX opening element (<div>).
//
//nolint:govet // Field order optimized for JSON output readability
type JSXOpeningElement struct {
	BaseNode
	SelfClosing    bool                          `json:"selfClosing"`
	Name           interface{}                   `json:"name"` // JSXIdentifier | JSXMemberExpression | JSXNamespacedName
	Attributes     []interface{}                 `json:"attributes"` // JSXAttribute | JSXSpreadAttribute
	TypeArguments  *TSTypeParameterInstantiation `json:"typeArguments,omitempty"`
	TypeParameters *TSTypeParameterInstantiation `json:"typeParameters,omitempty"` // Deprecated
}

// JSXClosingElement represents a JSX closing element (</div>).
//
//nolint:govet // Field order optimized for JSON output readability
type JSXClosingElement struct {
	BaseNode
	Name interface{} `json:"name"` // JSXIdentifier | JSXMemberExpression | JSXNamespacedName
}

// JSXOpeningFragment represents a JSX opening fragment (<>).
type JSXOpeningFragment struct {
	BaseNode
}

// JSXClosingFragment represents a JSX closing fragment (</>).
type JSXClosingFragment struct {
	BaseNode
}

// ==================== JSX Attributes ====================

// JSXAttribute represents a JSX attribute.
//
//nolint:govet // Field order optimized for JSON output readability
type JSXAttribute struct {
	BaseNode
	Name  interface{} `json:"name"`  // JSXIdentifier | JSXNamespacedName
	Value interface{} `json:"value"` // Literal | JSXExpressionContainer | JSXElement | JSXFragment | nil
}

// JSXSpreadAttribute represents a JSX spread attribute ({...props}).
//
//nolint:govet // Field order optimized for JSON output readability
type JSXSpreadAttribute struct {
	BaseNode
	Argument Expression `json:"argument"`
}

// ==================== JSX Names ====================

// JSXIdentifier represents a JSX identifier.
//
//nolint:govet // Field order optimized for JSON output readability
type JSXIdentifier struct {
	BaseNode
	Name string `json:"name"`
}

// JSXNamespacedName represents a JSX namespaced name (ns:name).
//
//nolint:govet // Field order optimized for JSON output readability
type JSXNamespacedName struct {
	BaseNode
	Namespace *JSXIdentifier `json:"namespace"`
	Name      *JSXIdentifier `json:"name"`
}

// JSXMemberExpression represents a JSX member expression (obj.prop).
//
//nolint:govet // Field order optimized for JSON output readability
type JSXMemberExpression struct {
	BaseNode
	Object   interface{}    `json:"object"`   // JSXIdentifier | JSXMemberExpression
	Property *JSXIdentifier `json:"property"`
}

// ==================== JSX Content ====================

// JSXExpressionContainer represents a JSX expression container {expr}.
//
//nolint:govet // Field order optimized for JSON output readability
type JSXExpressionContainer struct {
	BaseNode
	Expression interface{} `json:"expression"` // Expression | JSXEmptyExpression
}

func (n *JSXExpressionContainer) expressionNode() {}

// JSXEmptyExpression represents an empty JSX expression {}.
type JSXEmptyExpression struct {
	BaseNode
}

func (n *JSXEmptyExpression) expressionNode() {}

// JSXText represents JSX text content.
//
//nolint:govet // Field order optimized for JSON output readability
type JSXText struct {
	BaseNode
	Value string `json:"value"`
	Raw   string `json:"raw"`
}

// JSXSpreadChild represents a JSX spread child ({...children}).
//
//nolint:govet // Field order optimized for JSON output readability
type JSXSpreadChild struct {
	BaseNode
	Expression Expression `json:"expression"`
}
