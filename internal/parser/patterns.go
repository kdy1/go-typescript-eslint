package parser

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/lexer"
)

// parseBindingPattern parses a binding pattern (identifier, array pattern, or object pattern).
func (p *Parser) parseBindingPattern() (ast.Pattern, error) {
	switch p.current.Type {
	case lexer.LBRACK:
		return p.parseArrayPattern()
	case lexer.LBRACE:
		return p.parseObjectPattern()
	case lexer.IDENT:
		start := p.current.Pos
		name := p.current.Literal
		p.nextToken()
		return &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Name: name,
		}, nil
	default:
		return nil, p.errorAtCurrent("expected binding pattern")
	}
}

// parseArrayPattern parses an array destructuring pattern [a, b, c].
func (p *Parser) parseArrayPattern() (*ast.ArrayPattern, error) {
	start := p.current.Pos
	p.nextToken() // consume '['

	elements := []ast.Pattern{}

	for !p.match(lexer.RBRACK) && !p.isAtEnd() {
		// Handle holes in patterns
		if p.match(lexer.COMMA) {
			elements = append(elements, nil)
			p.nextToken()
			continue
		}

		// Handle rest element
		if p.consume(lexer.ELLIPSIS) {
			element, err := p.parseArrayRestElement()
			if err != nil {
				return nil, err
			}
			elements = append(elements, element)
			break
		}

		element, err := p.parseArrayPatternElement()
		if err != nil {
			return nil, err
		}

		elements = append(elements, element)

		if !p.consume(lexer.COMMA) {
			break
		}
	}

	if err := p.expect(lexer.RBRACK); err != nil {
		return nil, err
	}

	return &ast.ArrayPattern{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeArrayPattern.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Elements: elements,
	}, nil
}

// parseArrayRestElement parses a rest element in an array pattern.
func (p *Parser) parseArrayRestElement() (ast.Pattern, error) {
	arg, err := p.parseBindingPattern()
	if err != nil {
		return nil, err
	}
	return &ast.RestElement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeRestElement.String(),
		},
		Argument: arg,
	}, nil
}

// parseArrayPatternElement parses a single array pattern element with optional default.
func (p *Parser) parseArrayPatternElement() (ast.Pattern, error) {
	element, err := p.parseBindingPattern()
	if err != nil {
		return nil, err
	}

	// Check for default value
	if p.consume(lexer.ASSIGN) {
		right, err := p.parseAssignmentExpression()
		if err != nil {
			return nil, err
		}
		element = &ast.AssignmentPattern{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeAssignmentPattern.String(),
			},
			Left:  element,
			Right: right,
		}
	}

	return element, nil
}

// parseObjectPattern parses an object destructuring pattern {a, b, c}.
func (p *Parser) parseObjectPattern() (*ast.ObjectPattern, error) {
	start := p.current.Pos
	p.nextToken() // consume '{'

	properties := []interface{}{}

	for !p.match(lexer.RBRACE) && !p.isAtEnd() {
		// Handle rest element
		if p.consume(lexer.ELLIPSIS) {
			arg, err := p.parseBindingPattern()
			if err != nil {
				return nil, err
			}

			// Create a RestElement directly (not wrapped in Property)
			properties = append(properties, &ast.RestElement{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeRestElement.String(),
				},
				Argument: arg,
			})
			break
		}

		prop, err := p.parseObjectPatternProperty()
		if err != nil {
			return nil, err
		}
		properties = append(properties, prop)

		if !p.consume(lexer.COMMA) {
			break
		}
	}

	if err := p.expect(lexer.RBRACE); err != nil {
		return nil, err
	}

	return &ast.ObjectPattern{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeObjectPattern.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Properties: properties,
	}, nil
}

// parseObjectPatternProperty parses a property in an object pattern.
func (p *Parser) parseObjectPatternProperty() (*ast.Property, error) {
	start := p.current.Pos

	// Parse key
	computed := false
	var key ast.Expression
	var err error

	if p.consume(lexer.LBRACK) {
		computed = true
		key, err = p.parseAssignmentExpression()
		if err != nil {
			return nil, err
		}
		if err := p.expect(lexer.RBRACK); err != nil {
			return nil, err
		}
	} else if p.current.Type == lexer.IDENT {
		key = &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
		p.nextToken()
	} else if p.current.Type == lexer.STRING || p.current.Type == lexer.NUMBER {
		key = &ast.Literal{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeLiteral.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Value: p.current.Literal,
			Raw:   p.current.Literal,
		}
		p.nextToken()
	} else {
		return nil, p.errorAtCurrent("expected property key")
	}

	// Check for shorthand property
	if !computed && p.match(lexer.COMMA, lexer.RBRACE, lexer.ASSIGN) {
		if id, ok := key.(*ast.Identifier); ok {
			var value ast.Expression = id

			// Check for default value
			if p.consume(lexer.ASSIGN) {
				right, err := p.parseAssignmentExpression()
				if err != nil {
					return nil, err
				}
				value = &ast.AssignmentPattern{
					BaseNode: ast.BaseNode{
						NodeType: ast.NodeTypeAssignmentPattern.String(),
					},
					Left:  id,
					Right: right,
				}
			}

			return &ast.Property{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeProperty.String(),
					Range:    &ast.Range{start, p.current.Pos},
				},
				Key:       key,
				Value:     value,
				Kind:      "init",
				Shorthand: true,
				Computed:  false,
			}, nil
		}
	}

	// Non-shorthand property
	if err := p.expect(lexer.COLON); err != nil {
		return nil, err
	}

	valuePat, err := p.parseBindingPattern()
	if err != nil {
		return nil, err
	}

	var value ast.Expression
	// Check for default value
	if p.consume(lexer.ASSIGN) {
		right, err := p.parseAssignmentExpression()
		if err != nil {
			return nil, err
		}
		value = &ast.AssignmentPattern{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeAssignmentPattern.String(),
			},
			Left:  valuePat,
			Right: right,
		}
	} else {
		// Cast Pattern to Expression - valid since Identifier implements both
		value, _ = valuePat.(ast.Expression)
	}

	return &ast.Property{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeProperty.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Key:       key,
		Value:     value,
		Kind:      "init",
		Shorthand: false,
		Computed:  computed,
	}, nil
}

// parseTemplateLiteral parses a template literal.
func (p *Parser) parseTemplateLiteral() (*ast.TemplateLiteral, error) {
	start := p.current.Pos

	quasis := []ast.TemplateElement{}
	expressions := []ast.Expression{}

	// Handle template without substitutions
	if p.current.Type == lexer.TemplateNoSub {
		cooked := p.current.Literal // TODO: Unescape
		quasis = append(quasis, ast.TemplateElement{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTemplateElement.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Value: ast.TemplateElementValue{
				Raw:    p.current.Literal,
				Cooked: &cooked,
			},
			Tail: true,
		})
		p.nextToken()
		return &ast.TemplateLiteral{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTemplateLiteral.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Quasis:      quasis,
			Expressions: expressions,
		}, nil
	}

	// Handle template head
	if p.current.Type == lexer.TemplateHead {
		cooked := p.current.Literal // TODO: Unescape
		quasis = append(quasis, ast.TemplateElement{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTemplateElement.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Value: ast.TemplateElementValue{
				Raw:    p.current.Literal,
				Cooked: &cooked,
			},
			Tail: false,
		})
		p.nextToken()

		// Parse expressions and template parts
		for {
			// Parse expression
			expr, err := p.parseExpression()
			if err != nil {
				return nil, err
			}
			expressions = append(expressions, expr)

			// Parse template middle or tail
			if p.current.Type == lexer.TemplateMiddle {
				cooked := p.current.Literal // TODO: Unescape
				quasis = append(quasis, ast.TemplateElement{
					BaseNode: ast.BaseNode{
						NodeType: ast.NodeTypeTemplateElement.String(),
						Range:    &ast.Range{p.current.Pos, p.current.End},
					},
					Value: ast.TemplateElementValue{
						Raw:    p.current.Literal,
						Cooked: &cooked,
					},
					Tail: false,
				})
				p.nextToken()
			} else if p.current.Type == lexer.TemplateTail {
				cooked := p.current.Literal // TODO: Unescape
				quasis = append(quasis, ast.TemplateElement{
					BaseNode: ast.BaseNode{
						NodeType: ast.NodeTypeTemplateElement.String(),
						Range:    &ast.Range{p.current.Pos, p.current.End},
					},
					Value: ast.TemplateElementValue{
						Raw:    p.current.Literal,
						Cooked: &cooked,
					},
					Tail: true,
				})
				p.nextToken()
				break
			} else {
				return nil, p.errorAtCurrent("expected template middle or tail")
			}
		}
	}

	return &ast.TemplateLiteral{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTemplateLiteral.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Quasis:      quasis,
		Expressions: expressions,
	}, nil
}
