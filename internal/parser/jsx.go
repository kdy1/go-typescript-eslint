package parser

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/lexer"
)

// parseJSXElement parses a JSX element <div>...</div>.
func (p *Parser) parseJSXElement() (*ast.JSXElement, error) {
	start := p.current.Pos

	// Parse opening element
	opening, err := p.parseJSXOpeningElement()
	if err != nil {
		return nil, err
	}

	// Check for self-closing
	if opening.SelfClosing {
		return &ast.JSXElement{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeJSXElement.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			OpeningElement: opening,
			ClosingElement: nil,
			Children:       []interface{}{},
		}, nil
	}

	// Parse children
	children := []interface{}{}
	for !p.isAtEnd() {
		// Check for closing tag
		if p.current.Type == lexer.LSS && p.peek.Type == lexer.QUO {
			break
		}

		child, err := p.parseJSXChild()
		if err != nil {
			return nil, err
		}
		if child != nil {
			children = append(children, child)
		}
	}

	// Parse closing element
	closing, err := p.parseJSXClosingElement()
	if err != nil {
		return nil, err
	}

	return &ast.JSXElement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeJSXElement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		OpeningElement: opening,
		ClosingElement: closing,
		Children:       children,
	}, nil
}

// parseJSXOpeningElement parses a JSX opening element <div attr="value">.
func (p *Parser) parseJSXOpeningElement() (*ast.JSXOpeningElement, error) {
	start := p.current.Pos
	p.nextToken() // consume '<'

	// Parse name
	name, err := p.parseJSXElementName()
	if err != nil {
		return nil, err
	}

	// Parse type parameters (TypeScript)
	var typeParameters *ast.TSTypeParameterInstantiation
	if p.current.Type == lexer.LSS {
		typeParameters, err = p.parseTSTypeArguments()
		if err != nil {
			// Not type parameters, ignore
			typeParameters = nil
		}
	}

	// Parse attributes
	attributes := []interface{}{}
	for !p.match(lexer.GTR, lexer.JSXSelfClosingEnd) && !p.isAtEnd() {
		attr, err := p.parseJSXAttribute()
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, attr)
	}

	selfClosing := false
	if p.consume(lexer.JSXSelfClosingEnd) {
		selfClosing = true
	} else {
		if err := p.expect(lexer.GTR); err != nil {
			return nil, err
		}
	}

	return &ast.JSXOpeningElement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeJSXOpeningElement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Name:           name,
		Attributes:     attributes,
		SelfClosing:    selfClosing,
		TypeParameters: typeParameters,
	}, nil
}

// parseJSXClosingElement parses a JSX closing element </div>.
func (p *Parser) parseJSXClosingElement() (*ast.JSXClosingElement, error) {
	start := p.current.Pos
	p.nextToken() // consume '<'

	if err := p.expect(lexer.QUO); err != nil {
		return nil, err
	}

	name, err := p.parseJSXElementName()
	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.GTR); err != nil {
		return nil, err
	}

	return &ast.JSXClosingElement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeJSXClosingElement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Name: name,
	}, nil
}

// parseJSXElementName parses a JSX element name (identifier, member expression, or namespaced name).
func (p *Parser) parseJSXElementName() (ast.Node, error) {
	start := p.current.Pos

	if p.current.Type != lexer.IDENT {
		return nil, p.errorAtCurrent("expected JSX element name")
	}

	var name ast.Node = &ast.JSXIdentifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeJSXIdentifier.String(),
			Range:    &ast.Range{p.current.Pos, p.current.End},
		},
		Name: p.current.Literal,
	}
	p.nextToken()

	// Check for namespaced name (ns:name)
	if p.consume(lexer.COLON) {
		if p.current.Type != lexer.IDENT {
			return nil, p.errorAtCurrent("expected identifier after ':'")
		}

		namespaceName := &ast.JSXIdentifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeJSXIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
		p.nextToken()

		return &ast.JSXNamespacedName{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeJSXNamespacedName.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Namespace: name.(*ast.JSXIdentifier),
			Name:      namespaceName,
		}, nil
	}

	// Check for member expression (Obj.Prop)
	for p.consume(lexer.PERIOD) {
		if p.current.Type != lexer.IDENT {
			return nil, p.errorAtCurrent("expected identifier after '.'")
		}

		property := &ast.JSXIdentifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeJSXIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
		p.nextToken()

		name = &ast.JSXMemberExpression{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeJSXMemberExpression.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Object:   name,
			Property: property,
		}
	}

	return name, nil
}

// parseJSXAttribute parses a JSX attribute or spread attribute.
func (p *Parser) parseJSXAttribute() (ast.Node, error) {
	start := p.current.Pos

	// Check for spread attribute {...props}
	if p.consume(lexer.LBRACE) {
		if !p.consume(lexer.ELLIPSIS) {
			return nil, p.errorAtCurrent("expected '...' in JSX spread attribute")
		}

		argument, err := p.parseAssignmentExpression()
		if err != nil {
			return nil, err
		}

		if err := p.expect(lexer.RBRACE); err != nil {
			return nil, err
		}

		return &ast.JSXSpreadAttribute{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeJSXSpreadAttribute.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Argument: argument,
		}, nil
	}

	// Parse attribute name
	if p.current.Type != lexer.IDENT {
		return nil, p.errorAtCurrent("expected JSX attribute name")
	}

	var name interface{} = &ast.JSXIdentifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeJSXIdentifier.String(),
			Range:    &ast.Range{p.current.Pos, p.current.End},
		},
		Name: p.current.Literal,
	}
	p.nextToken()

	// Check for namespaced attribute (ns:name)
	if p.consume(lexer.COLON) {
		if p.current.Type != lexer.IDENT {
			return nil, p.errorAtCurrent("expected identifier after ':'")
		}

		namespaceName := &ast.JSXIdentifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeJSXIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
		p.nextToken()

		name = &ast.JSXNamespacedName{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeJSXNamespacedName.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Namespace: name.(*ast.JSXIdentifier),
			Name:      namespaceName,
		}
	}

	// Parse attribute value
	var value ast.Node
	if p.consume(lexer.ASSIGN) {
		var err error
		value, err = p.parseJSXAttributeValue()
		if err != nil {
			return nil, err
		}
	}

	return &ast.JSXAttribute{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeJSXAttribute.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Name:  name,
		Value: value,
	}, nil
}

// parseJSXAttributeValue parses a JSX attribute value.
func (p *Parser) parseJSXAttributeValue() (ast.Node, error) {
	switch p.current.Type {
	case lexer.STRING, lexer.JSXAttributeString:
		value := &ast.Literal{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeLiteral.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Value: p.current.Literal,
			Raw:   p.current.Literal,
		}
		p.nextToken()
		return value, nil

	case lexer.LBRACE:
		return p.parseJSXExpressionContainer()

	case lexer.LSS:
		return p.parseJSXElement()

	default:
		return nil, p.errorAtCurrent("expected JSX attribute value")
	}
}

// parseJSXExpressionContainer parses a JSX expression container {expr}.
func (p *Parser) parseJSXExpressionContainer() (*ast.JSXExpressionContainer, error) {
	start := p.current.Pos
	p.nextToken() // consume '{'

	// Check for empty expression
	if p.match(lexer.RBRACE) {
		p.nextToken()
		return &ast.JSXExpressionContainer{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeJSXExpressionContainer.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Expression: &ast.JSXEmptyExpression{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeJSXEmptyExpression.String(),
				},
			},
		}, nil
	}

	// Check for spread children
	if p.consume(lexer.ELLIPSIS) {
		expr, err := p.parseAssignmentExpression()
		if err != nil {
			return nil, err
		}

		if err := p.expect(lexer.RBRACE); err != nil {
			return nil, err
		}

		return &ast.JSXExpressionContainer{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeJSXExpressionContainer.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Expression: &ast.JSXSpreadChild{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeJSXSpreadChild.String(),
				},
				Expression: expr,
			},
		}, nil
	}

	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.RBRACE); err != nil {
		return nil, err
	}

	return &ast.JSXExpressionContainer{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeJSXExpressionContainer.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Expression: expr,
	}, nil
}

// parseJSXChild parses a JSX child (element, text, expression, etc.).
func (p *Parser) parseJSXChild() (ast.Node, error) {
	switch p.current.Type {
	case lexer.LSS:
		// Check for JSX fragment
		if p.peek.Type == lexer.GTR {
			return p.parseJSXFragment()
		}
		return p.parseJSXElement()

	case lexer.LBRACE:
		return p.parseJSXExpressionContainer()

	case lexer.JSXText:
		text := &ast.JSXText{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeJSXText.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Value: p.current.Literal,
			Raw:   p.current.Literal,
		}
		p.nextToken()
		return text, nil

	default:
		return nil, nil
	}
}

// parseJSXFragment parses a JSX fragment <>...</>.
func (p *Parser) parseJSXFragment() (*ast.JSXFragment, error) {
	start := p.current.Pos
	p.nextToken() // consume '<'

	opening := &ast.JSXOpeningFragment{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeJSXOpeningFragment.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
	}

	if err := p.expect(lexer.GTR); err != nil {
		return nil, err
	}

	// Parse children
	children := []interface{}{}
	for !p.isAtEnd() {
		// Check for closing fragment
		if p.current.Type == lexer.LSS && p.peek.Type == lexer.QUO {
			break
		}

		child, err := p.parseJSXChild()
		if err != nil {
			return nil, err
		}
		if child != nil {
			children = append(children, child)
		}
	}

	// Parse closing fragment
	closingStart := p.current.Pos
	p.nextToken() // consume '<'
	if err := p.expect(lexer.QUO); err != nil {
		return nil, err
	}

	closing := &ast.JSXClosingFragment{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeJSXClosingFragment.String(),
			Range:    &ast.Range{closingStart, p.current.Pos},
		},
	}

	if err := p.expect(lexer.GTR); err != nil {
		return nil, err
	}

	return &ast.JSXFragment{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeJSXFragment.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		OpeningFragment: opening,
		ClosingFragment: closing,
		Children:        children,
	}, nil
}
