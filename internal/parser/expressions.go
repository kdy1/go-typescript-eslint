package parser

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/lexer"
)

// Operator precedence levels (higher number = higher precedence)
const (
	precedenceLowest = iota
	precedenceComma
	precedenceAssignment
	precedenceConditional
	precedenceNullishCoalescing
	precedenceLogicalOR
	precedenceLogicalAND
	precedenceBitwiseOR
	precedenceBitwiseXOR
	precedenceBitwiseAND
	precedenceEquality
	precedenceRelational
	precedenceShift
	precedenceAdditive
	precedenceMultiplicative
	precedenceExponentiation
	precedenceUnary
	precedencePostfix
	precedenceCall
	precedenceMember
)

// precedence returns the precedence level for a token type.
func precedence(typ lexer.TokenType) int {
	switch typ {
	case lexer.COMMA:
		return precedenceComma
	case lexer.ASSIGN, lexer.AddAssign, lexer.SubAssign, lexer.MulAssign,
		lexer.QuoAssign, lexer.RemAssign, lexer.PowerAssign,
		lexer.ShlAssign, lexer.ShrAssign, lexer.ShrUnsignedAssign,
		lexer.AndAssign, lexer.OrAssign, lexer.XorAssign,
		lexer.NullishAssign:
		return precedenceAssignment
	case lexer.QUESTION:
		return precedenceConditional
	case lexer.NULLISH:
		return precedenceNullishCoalescing
	case lexer.LOR:
		return precedenceLogicalOR
	case lexer.LAND:
		return precedenceLogicalAND
	case lexer.OR:
		return precedenceBitwiseOR
	case lexer.XOR:
		return precedenceBitwiseXOR
	case lexer.AND:
		return precedenceBitwiseAND
	case lexer.EQL, lexer.NEQ, lexer.EqlStrict, lexer.NeqStrict:
		return precedenceEquality
	case lexer.LSS, lexer.LEQ, lexer.GTR, lexer.GEQ,
		lexer.INSTANCEOF, lexer.IN:
		return precedenceRelational
	case lexer.SHL, lexer.SHR, lexer.SHRUnsigned:
		return precedenceShift
	case lexer.ADD, lexer.SUB:
		return precedenceAdditive
	case lexer.MUL, lexer.QUO, lexer.REM:
		return precedenceMultiplicative
	case lexer.POWER:
		return precedenceExponentiation
	default:
		return precedenceLowest
	}
}

// parseExpression parses a full expression (includes comma operator).
func (p *Parser) parseExpression() (ast.Expression, error) {
	return p.parseBinaryExpression(precedenceLowest)
}

// parseAssignmentExpression parses an assignment expression or higher precedence.
func (p *Parser) parseAssignmentExpression() (ast.Expression, error) {
	return p.parseBinaryExpression(precedenceAssignment)
}

// parseBinaryExpression parses binary expressions using precedence climbing.
func (p *Parser) parseBinaryExpression(minPrec int) (ast.Expression, error) {
	// Parse the left-hand side (prefix/primary)
	left, err := p.parseUnaryExpression()
	if err != nil {
		return nil, err
	}

	// Climb the precedence ladder
	for {
		// Check for ternary conditional operator
		if p.current.Type == lexer.QUESTION && precedence(lexer.QUESTION) >= minPrec {
			left, err = p.parseConditionalExpression(left)
			if err != nil {
				return nil, err
			}
			continue
		}

		// Get operator precedence
		prec := precedence(p.current.Type)
		if prec < minPrec {
			break
		}

		operator := p.current.Literal
		opType := p.current.Type
		p.nextToken()

		// Parse right-hand side
		right, err := p.parseBinaryExpression(prec + 1)
		if err != nil {
			return nil, err
		}

		// Create appropriate node type
		if isAssignmentOp(opType) {
			// Convert left Expression to Pattern
			// In JavaScript/TypeScript, assignment left side must be a valid Pattern
			var leftPattern ast.Pattern
			switch l := left.(type) {
			case ast.Pattern:
				leftPattern = l
			default:
				// If it's not a pattern, we still need to assign it
				// This might happen with member expressions, identifiers, etc.
				// which implement both Expression and Pattern
				leftPattern, _ = left.(ast.Pattern)
			}

			left = &ast.AssignmentExpression{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeAssignmentExpression.String(),
				},
				Operator: operator,
				Left:     leftPattern,
				Right:    right,
			}
		} else if isLogicalOp(opType) {
			left = &ast.LogicalExpression{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeLogicalExpression.String(),
				},
				Operator: operator,
				Left:     left,
				Right:    right,
			}
		} else {
			left = &ast.BinaryExpression{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeBinaryExpression.String(),
				},
				Operator: operator,
				Left:     left,
				Right:    right,
			}
		}
	}

	return left, nil
}

// parseConditionalExpression parses a ternary conditional expression.
func (p *Parser) parseConditionalExpression(test ast.Expression) (ast.Expression, error) {
	p.nextToken() // consume '?'

	consequent, err := p.parseAssignmentExpression()
	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.COLON); err != nil {
		return nil, err
	}

	alternate, err := p.parseAssignmentExpression()
	if err != nil {
		return nil, err
	}

	return &ast.ConditionalExpression{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeConditionalExpression.String(),
		},
		Test:       test,
		Consequent: consequent,
		Alternate:  alternate,
	}, nil
}

// parseUnaryExpression parses unary expressions (prefix operators).
func (p *Parser) parseUnaryExpression() (ast.Expression, error) {
	start := p.current.Pos

	// Check for unary operators
	switch p.current.Type {
	case lexer.INC, lexer.DEC:
		operator := p.current.Literal
		p.nextToken()
		argument, err := p.parseUnaryExpression()
		if err != nil {
			return nil, err
		}
		return &ast.UpdateExpression{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeUpdateExpression.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Operator: operator,
			Argument: argument,
			Prefix:   true,
		}, nil

	case lexer.ADD, lexer.SUB, lexer.NOT, lexer.BNOT:
		operator := p.current.Literal
		p.nextToken()
		argument, err := p.parseUnaryExpression()
		if err != nil {
			return nil, err
		}
		return &ast.UnaryExpression{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeUnaryExpression.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Operator: operator,
			Argument: argument,
			Prefix:   true,
		}, nil

	case lexer.TYPEOF, lexer.VOID, lexer.DELETE:
		operator := p.current.Literal
		p.nextToken()
		argument, err := p.parseUnaryExpression()
		if err != nil {
			return nil, err
		}
		return &ast.UnaryExpression{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeUnaryExpression.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Operator: operator,
			Argument: argument,
			Prefix:   true,
		}, nil

	case lexer.AWAIT:
		if !p.allowAwait {
			return nil, p.errorAtCurrent("await is only allowed in async functions")
		}
		p.nextToken()
		argument, err := p.parseUnaryExpression()
		if err != nil {
			return nil, err
		}
		return &ast.AwaitExpression{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeAwaitExpression.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Argument: argument,
		}, nil

	case lexer.LSS:
		// Type assertion (TypeScript) or JSX
		if p.jsxEnabled {
			return p.parseJSXElement()
		}
		return p.parseTSTypeAssertion()

	default:
		return p.parsePostfixExpression()
	}
}

// parsePostfixExpression parses postfix expressions (e.g., x++, x--).
func (p *Parser) parsePostfixExpression() (ast.Expression, error) {
	expr, err := p.parseLeftHandSideExpression()
	if err != nil {
		return nil, err
	}

	// Check for postfix operators
	if p.match(lexer.INC, lexer.DEC) {
		operator := p.current.Literal
		start := p.current.Pos
		p.nextToken()
		return &ast.UpdateExpression{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeUpdateExpression.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Operator: operator,
			Argument: expr,
			Prefix:   false,
		}, nil
	}

	return expr, nil
}

// parseLeftHandSideExpression parses left-hand side expressions (member, call, etc.).
func (p *Parser) parseLeftHandSideExpression() (ast.Expression, error) {
	return p.parseMemberOrCallExpression()
}

// parseMemberOrCallExpression parses member access and call expressions.
func (p *Parser) parseMemberOrCallExpression() (ast.Expression, error) {
	expr, err := p.parsePrimaryExpression()
	if err != nil {
		return nil, err
	}

	for {
		switch p.current.Type {
		case lexer.PERIOD:
			p.nextToken()
			if p.current.Type != lexer.IDENT {
				return nil, p.errorAtCurrent("expected property name")
			}
			property := &ast.Identifier{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeIdentifier.String(),
					Range:    &ast.Range{p.current.Pos, p.current.End},
				},
				Name: p.current.Literal,
			}
			p.nextToken()
			expr = &ast.MemberExpression{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeMemberExpression.String(),
				},
				Object:   expr,
				Property: property,
				Computed: false,
			}

		case lexer.LBRACK:
			p.nextToken()
			property, err := p.parseExpression()
			if err != nil {
				return nil, err
			}
			if err := p.expect(lexer.RBRACK); err != nil {
				return nil, err
			}
			expr = &ast.MemberExpression{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeMemberExpression.String(),
				},
				Object:   expr,
				Property: property,
				Computed: true,
			}

		case lexer.OPTIONAL:
			// Optional chaining
			p.nextToken()
			if p.consume(lexer.LBRACK) {
				property, err := p.parseExpression()
				if err != nil {
					return nil, err
				}
				if err := p.expect(lexer.RBRACK); err != nil {
					return nil, err
				}
				expr = &ast.ChainExpression{
					BaseNode: ast.BaseNode{
						NodeType: ast.NodeTypeChainExpression.String(),
					},
					Expression: &ast.MemberExpression{
						BaseNode: ast.BaseNode{
							NodeType: ast.NodeTypeMemberExpression.String(),
						},
						Object:   expr,
						Property: property,
						Computed: true,
						Optional: true,
					},
				}
			} else if p.consume(lexer.LPAREN) {
				args, err := p.parseArguments()
				if err != nil {
					return nil, err
				}
				expr = &ast.ChainExpression{
					BaseNode: ast.BaseNode{
						NodeType: ast.NodeTypeChainExpression.String(),
					},
					Expression: &ast.CallExpression{
						BaseNode: ast.BaseNode{
							NodeType: ast.NodeTypeCallExpression.String(),
						},
						Callee:    expr,
						Arguments: args,
						Optional:  true,
					},
				}
			} else {
				if p.current.Type != lexer.IDENT {
					return nil, p.errorAtCurrent("expected property name")
				}
				property := &ast.Identifier{
					BaseNode: ast.BaseNode{
						NodeType: ast.NodeTypeIdentifier.String(),
						Range:    &ast.Range{p.current.Pos, p.current.End},
					},
					Name: p.current.Literal,
				}
				p.nextToken()
				expr = &ast.ChainExpression{
					BaseNode: ast.BaseNode{
						NodeType: ast.NodeTypeChainExpression.String(),
					},
					Expression: &ast.MemberExpression{
						BaseNode: ast.BaseNode{
							NodeType: ast.NodeTypeMemberExpression.String(),
						},
						Object:   expr,
						Property: property,
						Computed: false,
						Optional: true,
					},
				}
			}

		case lexer.LPAREN:
			p.nextToken()
			args, err := p.parseArguments()
			if err != nil {
				return nil, err
			}
			expr = &ast.CallExpression{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeCallExpression.String(),
				},
				Callee:    expr,
				Arguments: args,
			}

		case lexer.TEMPLATE, lexer.TemplateHead:
			// Tagged template expression
			template, err := p.parseTemplateLiteral()
			if err != nil {
				return nil, err
			}
			expr = &ast.TaggedTemplateExpression{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeTaggedTemplateExpression.String(),
				},
				Tag:   expr,
				Quasi: template,
			}

		case lexer.NOT:
			// Non-null assertion (TypeScript)
			p.nextToken()
			expr = &ast.TSNonNullExpression{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeTSNonNullExpression.String(),
				},
				Expression: expr,
			}

		default:
			return expr, nil
		}
	}
}

// parsePrimaryExpression parses primary expressions (literals, identifiers, etc.).
func (p *Parser) parsePrimaryExpression() (ast.Expression, error) {
	start := p.current.Pos

	switch p.current.Type {
	case lexer.THIS:
		p.nextToken()
		return &ast.ThisExpression{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeThisExpression.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
		}, nil

	case lexer.SUPER:
		p.nextToken()
		return &ast.Super{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeSuper.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
		}, nil

	case lexer.IDENT:
		name := p.current.Literal
		p.nextToken()
		id := &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Name: name,
		}

		// Check for type arguments (TypeScript)
		if p.current.Type == lexer.LSS {
			typeArgs, err := p.parseTSTypeArguments()
			if err == nil && typeArgs != nil {
				return &ast.TSInstantiationExpression{
					BaseNode: ast.BaseNode{
						NodeType: ast.NodeTypeTSInstantiationExpression.String(),
					},
					Expression:     id,
					TypeParameters: typeArgs,
				}, nil
			}
		}

		return id, nil

	case lexer.NULL:
		p.nextToken()
		return &ast.Literal{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeLiteral.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Value: nil,
			Raw:   "null",
		}, nil

	case lexer.TRUE:
		p.nextToken()
		return &ast.Literal{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeLiteral.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Value: true,
			Raw:   "true",
		}, nil

	case lexer.FALSE:
		p.nextToken()
		return &ast.Literal{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeLiteral.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Value: false,
			Raw:   "false",
		}, nil

	case lexer.NUMBER:
		value := p.current.Literal
		p.nextToken()
		return &ast.Literal{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeLiteral.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Value: value, // TODO: Parse actual numeric value
			Raw:   value,
		}, nil

	case lexer.STRING:
		value := p.current.Literal
		p.nextToken()
		return &ast.Literal{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeLiteral.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Value: value, // TODO: Unescape string
			Raw:   value,
		}, nil

	case lexer.REGEXP:
		value := p.current.Literal
		p.nextToken()
		return &ast.Literal{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeLiteral.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Value: value,
			Raw:   value,
			Regex: &ast.RegexInfo{
				Pattern: value, // TODO: Parse pattern and flags
				Flags:   "",
			},
		}, nil

	case lexer.TEMPLATE, lexer.TemplateHead, lexer.TemplateNoSub:
		return p.parseTemplateLiteral()

	case lexer.LBRACK:
		return p.parseArrayExpression()

	case lexer.LBRACE:
		return p.parseObjectExpression()

	case lexer.LPAREN:
		return p.parseParenthesizedOrArrowFunction()

	case lexer.FUNCTION:
		return p.parseFunctionExpression()

	case lexer.CLASS:
		return p.parseClassExpression()

	case lexer.NEW:
		return p.parseNewExpression()

	case lexer.IMPORT:
		return p.parseImportExpression()

	case lexer.ASYNC:
		return p.parseAsyncExpression()

	case lexer.YIELD:
		return p.parseYieldExpression()

	case lexer.JSXTagStart, lexer.LSS:
		if p.jsxEnabled {
			return p.parseJSXElement()
		}
		return nil, p.errorAtCurrent("unexpected token '<'")

	default:
		return nil, p.errorAtCurrent("unexpected token in expression")
	}
}

// parseArguments parses function call arguments.
func (p *Parser) parseArguments() ([]ast.Expression, error) {
	args := []ast.Expression{}

	for !p.match(lexer.RPAREN) && !p.isAtEnd() {
		// Handle spread arguments
		if p.consume(lexer.ELLIPSIS) {
			arg, err := p.parseAssignmentExpression()
			if err != nil {
				return nil, err
			}
			args = append(args, &ast.SpreadElement{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeSpreadElement.String(),
				},
				Argument: arg,
			})
		} else {
			arg, err := p.parseAssignmentExpression()
			if err != nil {
				return nil, err
			}
			args = append(args, arg)
		}

		if !p.consume(lexer.COMMA) {
			break
		}
	}

	if err := p.expect(lexer.RPAREN); err != nil {
		return nil, err
	}

	return args, nil
}

// parseArrayExpression parses an array literal [1, 2, 3].
func (p *Parser) parseArrayExpression() (*ast.ArrayExpression, error) {
	start := p.current.Pos
	p.nextToken() // consume '['

	elements := []ast.Expression{}

	for !p.match(lexer.RBRACK) && !p.isAtEnd() {
		// Handle holes in sparse arrays
		if p.match(lexer.COMMA) {
			elements = append(elements, nil)
			p.nextToken()
			continue
		}

		// Handle spread elements
		if p.consume(lexer.ELLIPSIS) {
			arg, err := p.parseAssignmentExpression()
			if err != nil {
				return nil, err
			}
			elements = append(elements, &ast.SpreadElement{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeSpreadElement.String(),
				},
				Argument: arg,
			})
		} else {
			element, err := p.parseAssignmentExpression()
			if err != nil {
				return nil, err
			}
			elements = append(elements, element)
		}

		if !p.consume(lexer.COMMA) {
			break
		}
	}

	if err := p.expect(lexer.RBRACK); err != nil {
		return nil, err
	}

	return &ast.ArrayExpression{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeArrayExpression.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Elements: elements,
	}, nil
}

// parseObjectExpression parses an object literal {a: 1, b: 2}.
func (p *Parser) parseObjectExpression() (*ast.ObjectExpression, error) {
	start := p.current.Pos
	p.nextToken() // consume '{'

	properties := []interface{}{}

	for !p.match(lexer.RBRACE) && !p.isAtEnd() {
		// Handle spread properties
		if p.consume(lexer.ELLIPSIS) {
			arg, err := p.parseAssignmentExpression()
			if err != nil {
				return nil, err
			}
			properties = append(properties, &ast.SpreadElement{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeSpreadElement.String(),
				},
				Argument: arg,
			})
		} else {
			prop, err := p.parseProperty()
			if err != nil {
				return nil, err
			}
			properties = append(properties, prop)
		}

		if !p.consume(lexer.COMMA) {
			break
		}
	}

	if err := p.expect(lexer.RBRACE); err != nil {
		return nil, err
	}

	return &ast.ObjectExpression{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeObjectExpression.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Properties: properties,
	}, nil
}

// parseProperty parses an object property.
func (p *Parser) parseProperty() (*ast.Property, error) {
	start := p.current.Pos

	// Check for getter/setter
	kind := "init"
	if p.current.Type == lexer.GET || p.current.Type == lexer.SET {
		kind = p.current.Literal
		p.nextToken()
	}

	// Check for async method
	async := p.consume(lexer.ASYNC)

	// Check for generator method
	generator := p.consume(lexer.MUL)

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
	if !computed && p.match(lexer.COMMA, lexer.RBRACE) {
		if id, ok := key.(*ast.Identifier); ok {
			return &ast.Property{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeProperty.String(),
					Range:    &ast.Range{start, p.current.Pos},
				},
				Key:       key,
				Value:     id,
				Kind:      kind,
				Method:    false,
				Shorthand: true,
				Computed:  false,
			}, nil
		}
	}

	// Parse type annotation if present (TypeScript)
	if p.consume(lexer.QUESTION) {
		// Optional property
	}
	if p.consume(lexer.COLON) && p.current.Type != lexer.COLON {
		// Could be type annotation - try to parse it
		_, _ = p.tryParseTSTypeAnnotation()
	}

	// Check for method
	method := false
	var value ast.Expression

	if p.consume(lexer.LPAREN) || async || generator {
		// Method
		method = true
		value, err = p.parseFunctionExpressionBody(async, generator)
		if err != nil {
			return nil, err
		}
	} else {
		// Regular property
		if err := p.expect(lexer.COLON); err != nil {
			return nil, err
		}
		value, err = p.parseAssignmentExpression()
		if err != nil {
			return nil, err
		}
	}

	return &ast.Property{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeProperty.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Key:       key,
		Value:     value,
		Kind:      kind,
		Method:    method,
		Shorthand: false,
		Computed:  computed,
	}, nil
}

// parseNewExpression parses a new expression.
func (p *Parser) parseNewExpression() (*ast.NewExpression, error) {
	start := p.current.Pos
	p.nextToken() // consume 'new'

	// Check for new.target
	if p.current.Type == lexer.PERIOD {
		p.nextToken()
		if p.current.Type != lexer.IDENT || p.current.Literal != "target" {
			return nil, p.errorAtCurrent("expected 'target' after 'new.'")
		}
		p.nextToken()
		return nil, p.errorAtCurrent("new.target not yet implemented")
	}

	callee, err := p.parseMemberOrCallExpression()
	if err != nil {
		return nil, err
	}

	var args []ast.Expression
	if p.consume(lexer.LPAREN) {
		args, err = p.parseArguments()
		if err != nil {
			return nil, err
		}
	}

	return &ast.NewExpression{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeNewExpression.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Callee:    callee,
		Arguments: args,
	}, nil
}

// parseImportExpression parses a dynamic import() expression.
func (p *Parser) parseImportExpression() (*ast.ImportExpression, error) {
	start := p.current.Pos
	p.nextToken() // consume 'import'

	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, err
	}

	source, err := p.parseAssignmentExpression()
	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.RPAREN); err != nil {
		return nil, err
	}

	return &ast.ImportExpression{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeImportExpression.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Source: source,
	}, nil
}

// parseAsyncExpression parses an async function or arrow function.
func (p *Parser) parseAsyncExpression() (ast.Expression, error) {
	start := p.current.Pos
	p.nextToken() // consume 'async'

	if p.current.Type == lexer.FUNCTION {
		return p.parseFunctionExpression()
	}

	// Async arrow function
	oldAllowAwait := p.allowAwait
	p.allowAwait = true
	defer func() { p.allowAwait = oldAllowAwait }()

	// Parse parameters
	var params []ast.Pattern
	if p.consume(lexer.LPAREN) {
		var err error
		params, err = p.parseFunctionParams()
		if err != nil {
			return nil, err
		}
	} else if p.current.Type == lexer.IDENT {
		params = []ast.Pattern{
			&ast.Identifier{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeIdentifier.String(),
					Range:    &ast.Range{p.current.Pos, p.current.End},
				},
				Name: p.current.Literal,
			},
		}
		p.nextToken()
	} else {
		return nil, p.errorAtCurrent("expected parameters for async arrow function")
	}

	if err := p.expect(lexer.ARROW); err != nil {
		return nil, err
	}

	// Parse body
	var body ast.Node
	var err error
	if p.current.Type == lexer.LBRACE {
		body, err = p.parseBlockStatement()
	} else {
		expr, err := p.parseAssignmentExpression()
		if err != nil {
			return nil, err
		}
		body = expr
	}

	if err != nil {
		return nil, err
	}

	return &ast.ArrowFunctionExpression{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeArrowFunctionExpression.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Params: params,
		Body:   body,
		Async:  true,
	}, nil
}

// parseYieldExpression parses a yield expression.
func (p *Parser) parseYieldExpression() (*ast.YieldExpression, error) {
	if !p.allowYield {
		return nil, p.errorAtCurrent("yield is only allowed in generator functions")
	}

	start := p.current.Pos
	p.nextToken() // consume 'yield'

	delegate := p.consume(lexer.MUL)

	var argument ast.Expression
	if !p.match(lexer.SEMICOLON, lexer.RBRACE) && !p.isAtEnd() {
		var err error
		argument, err = p.parseAssignmentExpression()
		if err != nil {
			return nil, err
		}
	}

	return &ast.YieldExpression{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeYieldExpression.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Argument: argument,
		Delegate: delegate,
	}, nil
}

// Helper functions

func isAssignmentOp(typ lexer.TokenType) bool {
	switch typ {
	case lexer.ASSIGN, lexer.AddAssign, lexer.SubAssign, lexer.MulAssign,
		lexer.QuoAssign, lexer.RemAssign, lexer.PowerAssign,
		lexer.ShlAssign, lexer.ShrAssign, lexer.ShrUnsignedAssign,
		lexer.AndAssign, lexer.OrAssign, lexer.XorAssign,
		lexer.NullishAssign:
		return true
	}
	return false
}

func isLogicalOp(typ lexer.TokenType) bool {
	return typ == lexer.LAND || typ == lexer.LOR || typ == lexer.NULLISH
}
