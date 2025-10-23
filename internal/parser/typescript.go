package parser

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/lexer"
)

// parseTSTypeAnnotation parses a TypeScript type annotation (: Type).
func (p *Parser) parseTSTypeAnnotation() (*ast.TSTypeAnnotation, error) {
	start := p.current.Pos

	tsType, err := p.parseTSType()
	if err != nil {
		return nil, err
	}

	return &ast.TSTypeAnnotation{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSTypeAnnotation.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		TypeAnnotation: tsType,
	}, nil
}

// tryParseTSTypeAnnotation attempts to parse a type annotation, returning nil if it fails.
func (p *Parser) tryParseTSTypeAnnotation() (*ast.TSTypeAnnotation, error) {
	// This is a simplified version - in production, we'd need better lookahead
	return p.parseTSTypeAnnotation()
}

// parseTSType parses a TypeScript type.
func (p *Parser) parseTSType() (ast.TSNode, error) {
	return p.parseTSUnionOrIntersectionType()
}

// parseTSUnionOrIntersectionType parses union or intersection types (A | B or A & B).
func (p *Parser) parseTSUnionOrIntersectionType() (ast.TSNode, error) {
	// Parse first type
	typ, err := p.parseTSPrimaryType()
	if err != nil {
		return nil, err
	}

	// Check for union or intersection
	if p.current.Type == lexer.OR {
		// Union type
		types := []ast.TSNode{typ}
		for p.consume(lexer.OR) {
			t, err := p.parseTSPrimaryType()
			if err != nil {
				return nil, err
			}
			types = append(types, t)
		}
		return &ast.TSUnionType{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTSUnionType.String(),
			},
			Types: types,
		}, nil
	} else if p.current.Type == lexer.AND {
		// Intersection type
		types := []ast.TSNode{typ}
		for p.consume(lexer.AND) {
			t, err := p.parseTSPrimaryType()
			if err != nil {
				return nil, err
			}
			types = append(types, t)
		}
		return &ast.TSIntersectionType{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTSIntersectionType.String(),
			},
			Types: types,
		}, nil
	}

	return typ, nil
}

// parseTSPrimaryType parses a primary TypeScript type.
func (p *Parser) parseTSPrimaryType() (ast.TSNode, error) {
	start := p.current.Pos

	switch p.current.Type {
	case lexer.ANY:
		p.nextToken()
		return &ast.TSAnyKeyword{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTSAnyKeyword.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
		}, nil

	case lexer.UNKNOWN:
		p.nextToken()
		return &ast.TSUnknownKeyword{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTSUnknownKeyword.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
		}, nil

	case lexer.NEVER:
		p.nextToken()
		return &ast.TSNeverKeyword{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTSNeverKeyword.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
		}, nil

	case lexer.StringKeyword:
		p.nextToken()
		return &ast.TSStringKeyword{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTSStringKeyword.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
		}, nil

	case lexer.NumberKeyword:
		p.nextToken()
		return &ast.TSNumberKeyword{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTSNumberKeyword.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
		}, nil

	case lexer.BOOLEAN:
		p.nextToken()
		return &ast.TSBooleanKeyword{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTSBooleanKeyword.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
		}, nil

	case lexer.SYMBOL:
		p.nextToken()
		return &ast.TSSymbolKeyword{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTSSymbolKeyword.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
		}, nil

	case lexer.VOID:
		p.nextToken()
		return &ast.TSVoidKeyword{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTSVoidKeyword.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
		}, nil

	case lexer.UNDEFINED:
		p.nextToken()
		return &ast.TSUndefinedKeyword{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTSUndefinedKeyword.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
		}, nil

	case lexer.NULL:
		p.nextToken()
		return &ast.TSNullKeyword{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTSNullKeyword.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
		}, nil

	case lexer.THIS:
		p.nextToken()
		return &ast.TSThisType{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTSThisType.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
		}, nil

	case lexer.IDENT:
		return p.parseTSTypeReference()

	case lexer.LBRACE:
		return p.parseTSTypeLiteral()

	case lexer.LBRACK:
		return p.parseTSTupleType()

	case lexer.LPAREN:
		return p.parseTSFunctionType()

	case lexer.NEW:
		return p.parseTSConstructorType()

	case lexer.TYPEOF:
		return p.parseTSTypeQuery()

	case lexer.IMPORT:
		return p.parseTSImportType()

	case lexer.STRING, lexer.NUMBER, lexer.TRUE, lexer.FALSE:
		// Literal type
		value := p.current.Literal
		p.nextToken()
		return &ast.TSLiteralType{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTSLiteralType.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Literal: &ast.Literal{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeLiteral.String(),
				},
				Value: value,
				Raw:   value,
			},
		}, nil

	default:
		return nil, p.errorAtCurrent("expected type")
	}
}

// parseTSTypeReference parses a type reference (e.g., Foo, Array<T>).
func (p *Parser) parseTSTypeReference() (*ast.TSTypeReference, error) {
	start := p.current.Pos

	typeName, err := p.parseTSEntityName()
	if err != nil {
		return nil, err
	}

	var typeParameters *ast.TSTypeParameterInstantiation
	if p.current.Type == lexer.LSS {
		typeParameters, err = p.parseTSTypeArguments()
		if err != nil {
			return nil, err
		}
	}

	return &ast.TSTypeReference{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSTypeReference.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		TypeName:       typeName,
		TypeParameters: typeParameters,
	}, nil
}

// parseTSEntityName parses a type name (identifier or qualified name).
func (p *Parser) parseTSEntityName() (ast.Node, error) {
	start := p.current.Pos

	if p.current.Type != lexer.IDENT {
		return nil, p.errorAtCurrent("expected identifier")
	}

	name := &ast.Identifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeIdentifier.String(),
			Range:    &ast.Range{start, p.current.End},
		},
		Name: p.current.Literal,
	}
	p.nextToken()

	// Check for qualified name (e.g., A.B.C)
	for p.consume(lexer.PERIOD) {
		if p.current.Type != lexer.IDENT {
			return nil, p.errorAtCurrent("expected identifier after '.'")
		}

		right := &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
		p.nextToken()

		name = &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Name: name.Name + "." + right.Name,
		}
	}

	return name, nil
}

// parseTSTypeLiteral parses a type literal {a: string, b: number}.
func (p *Parser) parseTSTypeLiteral() (*ast.TSTypeLiteral, error) {
	start := p.current.Pos
	p.nextToken() // consume '{'

	members := []interface{}{}

	for !p.match(lexer.RBRACE) && !p.isAtEnd() {
		member, err := p.parseTSTypeElement()
		if err != nil {
			p.synchronize()
			continue
		}
		members = append(members, member)

		// Consume optional separator
		p.consume(lexer.SEMICOLON)
		p.consume(lexer.COMMA)
	}

	if err := p.expect(lexer.RBRACE); err != nil {
		return nil, err
	}

	return &ast.TSTypeLiteral{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSTypeLiteral.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Members: members,
	}, nil
}

// parseTSTypeElement parses a type element (property signature, method signature, etc.).
func (p *Parser) parseTSTypeElement() (ast.Node, error) {
	start := p.current.Pos

	// Check for index signature
	if p.current.Type == lexer.LBRACK {
		return p.parseTSIndexSignature()
	}

	// Check for call signature
	if p.current.Type == lexer.LPAREN || p.current.Type == lexer.LSS {
		return p.parseTSCallSignature()
	}

	// Check for construct signature
	if p.consume(lexer.NEW) {
		return p.parseTSConstructSignature()
	}

	// Parse property or method signature
	readonly := p.consume(lexer.READONLY)

	// Parse key
	computed := false
	var key ast.Expression
	var err error

	if p.consume(lexer.LBRACK) {
		computed = true
		key, err = p.parseExpression()
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
	} else {
		return nil, p.errorAtCurrent("expected property name")
	}

	optional := p.consume(lexer.QUESTION)

	// Check for method signature
	if p.current.Type == lexer.LPAREN || p.current.Type == lexer.LSS {
		// Method signature
		var typeParameters *ast.TSTypeParameterDeclaration
		if p.current.Type == lexer.LSS {
			typeParameters, err = p.parseTSTypeParameters()
			if err != nil {
				return nil, err
			}
		}

		if err := p.expect(lexer.LPAREN); err != nil {
			return nil, err
		}

		params, err := p.parseTSFunctionParams()
		if err != nil {
			return nil, err
		}

		var returnType *ast.TSTypeAnnotation
		if p.consume(lexer.COLON) {
			returnType, err = p.parseTSTypeAnnotation()
			if err != nil {
				return nil, err
			}
		}

		return &ast.TSMethodSignature{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeTSMethodSignature.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Key:            key,
			Computed:       computed,
			Optional:       optional,
			Params:         params,
			ReturnType:     returnType,
			TypeParameters: typeParameters,
		}, nil
	}

	// Property signature
	var typeAnnotation *ast.TSTypeAnnotation
	if p.consume(lexer.COLON) {
		typeAnnotation, err = p.parseTSTypeAnnotation()
		if err != nil {
			return nil, err
		}
	}

	return &ast.TSPropertySignature{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSPropertySignature.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Key:            key,
		Computed:       computed,
		Optional:       optional,
		Readonly:       readonly,
		TypeAnnotation: typeAnnotation,
	}, nil
}

// parseTSIndexSignature parses an index signature [key: string]: Type.
func (p *Parser) parseTSIndexSignature() (*ast.TSIndexSignature, error) {
	start := p.current.Pos
	p.nextToken() // consume '['

	// Parse parameter
	if p.current.Type != lexer.IDENT {
		return nil, p.errorAtCurrent("expected identifier")
	}

	param := &ast.Identifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeIdentifier.String(),
			Range:    &ast.Range{p.current.Pos, p.current.End},
		},
		Name: p.current.Literal,
	}
	p.nextToken()

	// Parse parameter type
	if err := p.expect(lexer.COLON); err != nil {
		return nil, err
	}

	paramType, err := p.parseTSTypeAnnotation()
	if err != nil {
		return nil, err
	}
	param.TypeAnnotation = paramType

	if err := p.expect(lexer.RBRACK); err != nil {
		return nil, err
	}

	// Parse index type
	if err := p.expect(lexer.COLON); err != nil {
		return nil, err
	}

	typeAnnotation, err := p.parseTSTypeAnnotation()
	if err != nil {
		return nil, err
	}

	return &ast.TSIndexSignature{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSIndexSignature.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Parameters:     []ast.Pattern{param},
		TypeAnnotation: typeAnnotation,
	}, nil
}

// parseTSCallSignature parses a call signature (x: string): string.
func (p *Parser) parseTSCallSignature() (*ast.TSCallSignatureDeclaration, error) {
	start := p.current.Pos

	var typeParameters *ast.TSTypeParameterDeclaration
	var err error

	if p.current.Type == lexer.LSS {
		typeParameters, err = p.parseTSTypeParameters()
		if err != nil {
			return nil, err
		}
	}

	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, err
	}

	params, err := p.parseTSFunctionParams()
	if err != nil {
		return nil, err
	}

	var returnType *ast.TSTypeAnnotation
	if p.consume(lexer.COLON) {
		returnType, err = p.parseTSTypeAnnotation()
		if err != nil {
			return nil, err
		}
	}

	return &ast.TSCallSignatureDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSCallSignatureDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Params:         params,
		ReturnType:     returnType,
		TypeParameters: typeParameters,
	}, nil
}

// parseTSConstructSignature parses a construct signature new (x: string): Type.
func (p *Parser) parseTSConstructSignature() (*ast.TSConstructSignatureDeclaration, error) {
	start := p.current.Pos

	var typeParameters *ast.TSTypeParameterDeclaration
	var err error

	if p.current.Type == lexer.LSS {
		typeParameters, err = p.parseTSTypeParameters()
		if err != nil {
			return nil, err
		}
	}

	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, err
	}

	params, err := p.parseTSFunctionParams()
	if err != nil {
		return nil, err
	}

	var returnType *ast.TSTypeAnnotation
	if p.consume(lexer.COLON) {
		returnType, err = p.parseTSTypeAnnotation()
		if err != nil {
			return nil, err
		}
	}

	return &ast.TSConstructSignatureDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSConstructSignatureDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Params:         params,
		ReturnType:     returnType,
		TypeParameters: typeParameters,
	}, nil
}

// parseTSTupleType parses a tuple type [string, number].
func (p *Parser) parseTSTupleType() (*ast.TSTupleType, error) {
	start := p.current.Pos
	p.nextToken() // consume '['

	elementTypes := []ast.TSNode{}

	for !p.match(lexer.RBRACK) && !p.isAtEnd() {
		// Handle rest element
		if p.consume(lexer.ELLIPSIS) {
			elemType, err := p.parseTSType()
			if err != nil {
				return nil, err
			}
			elementTypes = append(elementTypes, &ast.TSRestType{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeTSRestType.String(),
				},
				TypeAnnotation: elemType,
			})
			break
		}

		elemType, err := p.parseTSType()
		if err != nil {
			return nil, err
		}

		// Check for optional element
		if p.consume(lexer.QUESTION) {
			elemType = &ast.TSOptionalType{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeTSOptionalType.String(),
				},
				TypeAnnotation: elemType,
			}
		}

		elementTypes = append(elementTypes, elemType)

		if !p.consume(lexer.COMMA) {
			break
		}
	}

	if err := p.expect(lexer.RBRACK); err != nil {
		return nil, err
	}

	return &ast.TSTupleType{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSTupleType.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		ElementTypes: elementTypes,
	}, nil
}

// parseTSFunctionType parses a function type (x: string) => string.
func (p *Parser) parseTSFunctionType() (*ast.TSFunctionType, error) {
	start := p.current.Pos

	var typeParameters *ast.TSTypeParameterDeclaration
	var err error

	if p.current.Type == lexer.LSS {
		typeParameters, err = p.parseTSTypeParameters()
		if err != nil {
			return nil, err
		}
	}

	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, err
	}

	params, err := p.parseTSFunctionParams()
	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.ARROW); err != nil {
		return nil, err
	}

	returnType, err := p.parseTSTypeAnnotation()
	if err != nil {
		return nil, err
	}

	return &ast.TSFunctionType{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSFunctionType.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Params:         params,
		ReturnType:     returnType,
		TypeParameters: typeParameters,
	}, nil
}

// parseTSConstructorType parses a constructor type new (x: string) => Type.
func (p *Parser) parseTSConstructorType() (*ast.TSConstructorType, error) {
	start := p.current.Pos
	p.nextToken() // consume 'new'

	var typeParameters *ast.TSTypeParameterDeclaration
	var err error

	if p.current.Type == lexer.LSS {
		typeParameters, err = p.parseTSTypeParameters()
		if err != nil {
			return nil, err
		}
	}

	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, err
	}

	params, err := p.parseTSFunctionParams()
	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.ARROW); err != nil {
		return nil, err
	}

	returnType, err := p.parseTSTypeAnnotation()
	if err != nil {
		return nil, err
	}

	return &ast.TSConstructorType{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSConstructorType.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Params:         params,
		ReturnType:     returnType,
		TypeParameters: typeParameters,
	}, nil
}

// parseTSTypeQuery parses a typeof query typeof x.
func (p *Parser) parseTSTypeQuery() (*ast.TSTypeQuery, error) {
	start := p.current.Pos
	p.nextToken() // consume 'typeof'

	exprName, err := p.parseTSEntityName()
	if err != nil {
		return nil, err
	}

	return &ast.TSTypeQuery{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSTypeQuery.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		ExprName: exprName,
	}, nil
}

// parseTSImportType parses an import type import('module').Type.
func (p *Parser) parseTSImportType() (*ast.TSImportType, error) {
	start := p.current.Pos
	p.nextToken() // consume 'import'

	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, err
	}

	if p.current.Type != lexer.STRING {
		return nil, p.errorAtCurrent("expected string literal")
	}

	literal := &ast.Literal{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeLiteral.String(),
		},
		Value: p.current.Literal,
		Raw:   p.current.Literal,
	}
	p.nextToken()

	if err := p.expect(lexer.RPAREN); err != nil {
		return nil, err
	}

	argument := &ast.TSLiteralType{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSLiteralType.String(),
		},
		Literal: literal,
	}

	var qualifier ast.Node
	var err error
	if p.consume(lexer.PERIOD) {
		qualifier, err = p.parseTSEntityName()
		if err != nil {
			return nil, err
		}
	}

	var typeParameters *ast.TSTypeParameterInstantiation
	if p.current.Type == lexer.LSS {
		typeParameters, err = p.parseTSTypeArguments()
		if err != nil {
			return nil, err
		}
	}

	return &ast.TSImportType{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSImportType.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Argument:       argument,
		Qualifier:      qualifier,
		TypeParameters: typeParameters,
	}, nil
}

// parseTSFunctionParams parses TypeScript function parameters.
func (p *Parser) parseTSFunctionParams() ([]ast.Pattern, error) {
	params := []ast.Pattern{}

	for !p.match(lexer.RPAREN) && !p.isAtEnd() {
		// Handle rest parameter
		if p.consume(lexer.ELLIPSIS) {
			param, err := p.parseBindingPattern()
			if err != nil {
				return nil, err
			}

			if id, ok := param.(*ast.Identifier); ok {
				if p.consume(lexer.COLON) {
					typeAnnotation, err := p.parseTSTypeAnnotation()
					if err != nil {
						return nil, err
					}
					id.TypeAnnotation = typeAnnotation
				}
			}

			params = append(params, &ast.RestElement{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeRestElement.String(),
				},
				Argument: param,
			})
			break
		}

		param, err := p.parseBindingPattern()
		if err != nil {
			return nil, err
		}

		// Parse type annotation (TypeScript)
		if id, ok := param.(*ast.Identifier); ok {
			if p.consume(lexer.QUESTION) {
				id.Optional = true
			}
			if p.consume(lexer.COLON) {
				typeAnnotation, err := p.parseTSTypeAnnotation()
				if err != nil {
					return nil, err
				}
				id.TypeAnnotation = typeAnnotation
			}
		}

		// Parse default value
		if p.consume(lexer.ASSIGN) {
			init, err := p.parseAssignmentExpression()
			if err != nil {
				return nil, err
			}
			param = &ast.AssignmentPattern{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeAssignmentPattern.String(),
				},
				Left:  param,
				Right: init,
			}
		}

		params = append(params, param)

		if !p.consume(lexer.COMMA) {
			break
		}
	}

	if err := p.expect(lexer.RPAREN); err != nil {
		return nil, err
	}

	return params, nil
}

// parseTSTypeParameters parses type parameter declaration <T, U>.
func (p *Parser) parseTSTypeParameters() (*ast.TSTypeParameterDeclaration, error) {
	start := p.current.Pos
	if err := p.expect(lexer.LSS); err != nil {
		return nil, err
	}

	params := []ast.TSTypeParameter{}

	for !p.match(lexer.GTR) && !p.isAtEnd() {
		param, err := p.parseTSTypeParameter()
		if err != nil {
			return nil, err
		}
		params = append(params, *param)

		if !p.consume(lexer.COMMA) {
			break
		}
	}

	if err := p.expect(lexer.GTR); err != nil {
		return nil, err
	}

	return &ast.TSTypeParameterDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSTypeParameterDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Params: params,
	}, nil
}

// parseTSTypeParameter parses a single type parameter T extends Constraint = Default.
func (p *Parser) parseTSTypeParameter() (*ast.TSTypeParameter, error) {
	start := p.current.Pos

	if p.current.Type != lexer.IDENT {
		return nil, p.errorAtCurrent("expected type parameter name")
	}

	name := &ast.Identifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeIdentifier.String(),
			Range:    &ast.Range{p.current.Pos, p.current.End},
		},
		Name: p.current.Literal,
	}
	p.nextToken()

	var constraint ast.TSNode
	if p.consume(lexer.EXTENDS) {
		var err error
		constraint, err = p.parseTSType()
		if err != nil {
			return nil, err
		}
	}

	var defaultType ast.TSNode
	if p.consume(lexer.ASSIGN) {
		var err error
		defaultType, err = p.parseTSType()
		if err != nil {
			return nil, err
		}
	}

	return &ast.TSTypeParameter{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSTypeParameter.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Name:       name,
		Constraint: constraint,
		Default:    defaultType,
	}, nil
}

// parseTSTypeArguments parses type arguments <string, number>.
func (p *Parser) parseTSTypeArguments() (*ast.TSTypeParameterInstantiation, error) {
	start := p.current.Pos
	if err := p.expect(lexer.LSS); err != nil {
		return nil, err
	}

	params := []ast.TSNode{}

	for !p.match(lexer.GTR) && !p.isAtEnd() {
		param, err := p.parseTSType()
		if err != nil {
			return nil, err
		}
		params = append(params, param)

		if !p.consume(lexer.COMMA) {
			break
		}
	}

	if err := p.expect(lexer.GTR); err != nil {
		return nil, err
	}

	return &ast.TSTypeParameterInstantiation{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSTypeParameterInstantiation.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Params: params,
	}, nil
}

// parseTSTypeAssertion parses a type assertion <Type>expr.
func (p *Parser) parseTSTypeAssertion() (*ast.TSTypeAssertion, error) {
	start := p.current.Pos
	p.nextToken() // consume '<'

	typeAnnotation, err := p.parseTSType()
	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.GTR); err != nil {
		return nil, err
	}

	expression, err := p.parseUnaryExpression()
	if err != nil {
		return nil, err
	}

	return &ast.TSTypeAssertion{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSTypeAssertion.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		TypeAnnotation: typeAnnotation,
		Expression:     expression,
	}, nil
}

// parseTSInterfaceDeclaration parses an interface declaration.
func (p *Parser) parseTSInterfaceDeclaration() (*ast.TSInterfaceDeclaration, error) {
	start := p.current.Pos
	p.nextToken() // consume 'interface'

	if p.current.Type != lexer.IDENT {
		return nil, p.errorAtCurrent("expected interface name")
	}

	id := &ast.Identifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeIdentifier.String(),
			Range:    &ast.Range{p.current.Pos, p.current.End},
		},
		Name: p.current.Literal,
	}
	p.nextToken()

	// Parse type parameters
	var typeParameters *ast.TSTypeParameterDeclaration
	if p.current.Type == lexer.LSS {
		var err error
		typeParameters, err = p.parseTSTypeParameters()
		if err != nil {
			return nil, err
		}
	}

	// Parse extends clause
	var extends []ast.TSInterfaceHeritage
	if p.consume(lexer.EXTENDS) {
		for {
			heritage, err := p.parseTSInterfaceHeritage()
			if err != nil {
				return nil, err
			}
			extends = append(extends, *heritage)

			if !p.consume(lexer.COMMA) {
				break
			}
		}
	}

	// Parse body
	body, err := p.parseTSInterfaceBody()
	if err != nil {
		return nil, err
	}

	return &ast.TSInterfaceDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSInterfaceDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		ID:             id,
		TypeParameters: typeParameters,
		Extends:        extends,
		Body:           body,
	}, nil
}

// parseTSInterfaceHeritage parses an interface heritage clause.
func (p *Parser) parseTSInterfaceHeritage() (*ast.TSInterfaceHeritage, error) {
	start := p.current.Pos

	expressionNode, err := p.parseTSEntityName()
	if err != nil {
		return nil, err
	}

	// Cast to Expression - TSEntityName returns Identifier/TSQualifiedName which implement Expression
	expression, _ := expressionNode.(ast.Expression)

	var typeParameters *ast.TSTypeParameterInstantiation
	if p.current.Type == lexer.LSS {
		typeParameters, err = p.parseTSTypeArguments()
		if err != nil {
			return nil, err
		}
	}

	return &ast.TSInterfaceHeritage{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSInterfaceHeritage.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Expression:     expression,
		TypeParameters: typeParameters,
	}, nil
}

// parseTSInterfaceBody parses an interface body.
func (p *Parser) parseTSInterfaceBody() (*ast.TSInterfaceBody, error) {
	start := p.current.Pos
	if err := p.expect(lexer.LBRACE); err != nil {
		return nil, err
	}

	body := []interface{}{}

	for !p.match(lexer.RBRACE) && !p.isAtEnd() {
		member, err := p.parseTSTypeElement()
		if err != nil {
			p.synchronize()
			continue
		}
		body = append(body, member)

		// Consume optional separator
		p.consume(lexer.SEMICOLON)
		p.consume(lexer.COMMA)
	}

	if err := p.expect(lexer.RBRACE); err != nil {
		return nil, err
	}

	return &ast.TSInterfaceBody{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSInterfaceBody.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Body: body,
	}, nil
}

// parseTSTypeAliasDeclaration parses a type alias declaration.
func (p *Parser) parseTSTypeAliasDeclaration() (*ast.TSTypeAliasDeclaration, error) {
	start := p.current.Pos
	p.nextToken() // consume 'type'

	if p.current.Type != lexer.IDENT {
		return nil, p.errorAtCurrent("expected type alias name")
	}

	id := &ast.Identifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeIdentifier.String(),
			Range:    &ast.Range{p.current.Pos, p.current.End},
		},
		Name: p.current.Literal,
	}
	p.nextToken()

	// Parse type parameters
	var typeParameters *ast.TSTypeParameterDeclaration
	if p.current.Type == lexer.LSS {
		var err error
		typeParameters, err = p.parseTSTypeParameters()
		if err != nil {
			return nil, err
		}
	}

	if err := p.expect(lexer.ASSIGN); err != nil {
		return nil, err
	}

	typeAnnotation, err := p.parseTSType()
	if err != nil {
		return nil, err
	}

	p.consume(lexer.SEMICOLON)

	return &ast.TSTypeAliasDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSTypeAliasDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		ID:             id,
		TypeAnnotation: typeAnnotation,
		TypeParameters: typeParameters,
	}, nil
}

// parseTSEnumDeclaration parses an enum declaration.
func (p *Parser) parseTSEnumDeclaration() (*ast.TSEnumDeclaration, error) {
	start := p.current.Pos
	p.nextToken() // consume 'enum'

	if p.current.Type != lexer.IDENT {
		return nil, p.errorAtCurrent("expected enum name")
	}

	id := &ast.Identifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeIdentifier.String(),
			Range:    &ast.Range{p.current.Pos, p.current.End},
		},
		Name: p.current.Literal,
	}
	p.nextToken()

	if err := p.expect(lexer.LBRACE); err != nil {
		return nil, err
	}

	members := []ast.TSEnumMember{}

	for !p.match(lexer.RBRACE) && !p.isAtEnd() {
		member, err := p.parseTSEnumMember()
		if err != nil {
			p.synchronize()
			continue
		}
		members = append(members, *member)

		if !p.consume(lexer.COMMA) {
			break
		}
	}

	if err := p.expect(lexer.RBRACE); err != nil {
		return nil, err
	}

	return &ast.TSEnumDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSEnumDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		ID:      id,
		Members: members,
	}, nil
}

// parseTSEnumMember parses an enum member.
func (p *Parser) parseTSEnumMember() (*ast.TSEnumMember, error) {
	start := p.current.Pos

	var id ast.Node
	if p.current.Type == lexer.IDENT {
		id = &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
		p.nextToken()
	} else if p.current.Type == lexer.STRING {
		id = &ast.Literal{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeLiteral.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Value: p.current.Literal,
			Raw:   p.current.Literal,
		}
		p.nextToken()
	} else {
		return nil, p.errorAtCurrent("expected enum member name")
	}

	var initializer ast.Expression
	if p.consume(lexer.ASSIGN) {
		var err error
		initializer, err = p.parseAssignmentExpression()
		if err != nil {
			return nil, err
		}
	}

	return &ast.TSEnumMember{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSEnumMember.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		ID:          id,
		Initializer: initializer,
	}, nil
}

// parseTSModuleDeclaration parses a module/namespace declaration.
func (p *Parser) parseTSModuleDeclaration() (*ast.TSModuleDeclaration, error) {
	start := p.current.Pos
	p.nextToken() // consume 'namespace' or 'module'

	if p.current.Type != lexer.IDENT && p.current.Type != lexer.STRING {
		return nil, p.errorAtCurrent("expected module name")
	}

	var id ast.Node
	if p.current.Type == lexer.IDENT {
		id = &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
	} else {
		id = &ast.Literal{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeLiteral.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Value: p.current.Literal,
			Raw:   p.current.Literal,
		}
	}
	p.nextToken()

	// Parse body
	var body ast.Node
	if p.current.Type == lexer.LBRACE {
		bodyBlock, err := p.parseTSModuleBlock()
		if err != nil {
			return nil, err
		}
		body = bodyBlock
	} else {
		return nil, p.errorAtCurrent("expected module body")
	}

	return &ast.TSModuleDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSModuleDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		ID:   id,
		Body: body,
	}, nil
}

// parseTSModuleBlock parses a module block.
func (p *Parser) parseTSModuleBlock() (*ast.TSModuleBlock, error) {
	start := p.current.Pos
	if err := p.expect(lexer.LBRACE); err != nil {
		return nil, err
	}

	body := []ast.Statement{}

	for !p.match(lexer.RBRACE) && !p.isAtEnd() {
		stmt, err := p.parseStatementListItem()
		if err != nil {
			p.synchronize()
			continue
		}
		body = append(body, stmt)
	}

	if err := p.expect(lexer.RBRACE); err != nil {
		return nil, err
	}

	return &ast.TSModuleBlock{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSModuleBlock.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Body: body,
	}, nil
}

// parseTSClassImplements parses a class implements clause.
func (p *Parser) parseTSClassImplements() (*ast.TSClassImplements, error) {
	start := p.current.Pos

	expressionNode, err := p.parseTSEntityName()
	if err != nil {
		return nil, err
	}

	// Cast to Expression - TSEntityName returns Identifier/TSQualifiedName which implement Expression
	expression, _ := expressionNode.(ast.Expression)

	var typeParameters *ast.TSTypeParameterInstantiation
	if p.current.Type == lexer.LSS {
		typeParameters, err = p.parseTSTypeArguments()
		if err != nil {
			return nil, err
		}
	}

	return &ast.TSClassImplements{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSClassImplements.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Expression:     expression,
		TypeParameters: typeParameters,
	}, nil
}
