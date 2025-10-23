package parser

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/lexer"
)

// parseFunctionDeclaration parses a function declaration.
func (p *Parser) parseFunctionDeclaration() (*ast.FunctionDeclaration, error) {
	start := p.current.Pos
	async := p.consume(lexer.ASYNC)

	if err := p.expect(lexer.FUNCTION); err != nil {
		return nil, err
	}

	generator := p.consume(lexer.MUL)
	id := p.parseOptionalIdentifier()
	typeParameters := p.parseOptionalTypeParameters()

	params, returnType, body, err := p.parseFunctionSignatureAndBody(async, generator)
	if err != nil {
		return nil, err
	}

	return &ast.FunctionDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeFunctionDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		ID:             id,
		Params:         params,
		Body:           body,
		Generator:      generator,
		Async:          async,
		ReturnType:     returnType,
		TypeParameters: typeParameters,
	}, nil
}

// parseOptionalIdentifier parses an optional identifier.
func (p *Parser) parseOptionalIdentifier() *ast.Identifier {
	if p.current.Type == lexer.IDENT {
		id := &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
		p.nextToken()
		return id
	}
	return nil
}

// parseOptionalTypeParameters parses optional TypeScript type parameters.
func (p *Parser) parseOptionalTypeParameters() *ast.TSTypeParameterDeclaration {
	if p.current.Type == lexer.LSS {
		typeParameters, err := p.parseTSTypeParameters()
		if err != nil {
			// Not type parameters, backtrack
			return nil
		}
		return typeParameters
	}
	return nil
}

// parseFunctionSignatureAndBody parses function parameters, return type, and body.
func (p *Parser) parseFunctionSignatureAndBody(async, generator bool) ([]ast.Pattern, *ast.TSTypeAnnotation, *ast.BlockStatement, error) {
	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, nil, nil, err
	}

	oldInFunction := p.inFunction
	oldAllowYield := p.allowYield
	oldAllowAwait := p.allowAwait
	p.inFunction = true
	p.allowYield = generator
	p.allowAwait = async

	params, err := p.parseFunctionParams()
	if err != nil {
		return nil, nil, nil, err
	}

	// Parse return type annotation (TypeScript)
	var returnType *ast.TSTypeAnnotation
	if p.consume(lexer.COLON) {
		returnType, err = p.parseTSTypeAnnotation()
		if err != nil {
			return nil, nil, nil, err
		}
	}

	// Parse body
	var body *ast.BlockStatement
	if p.current.Type == lexer.LBRACE {
		body, err = p.parseBlockStatement()
		if err != nil {
			return nil, nil, nil, err
		}
	}

	p.inFunction = oldInFunction
	p.allowYield = oldAllowYield
	p.allowAwait = oldAllowAwait

	return params, returnType, body, nil
}

// parseFunctionExpression parses a function expression.
func (p *Parser) parseFunctionExpression() (*ast.FunctionExpression, error) {
	start := p.current.Pos
	async := p.current.Type == lexer.ASYNC
	if async {
		p.nextToken()
	}

	if err := p.expect(lexer.FUNCTION); err != nil {
		return nil, err
	}

	generator := p.consume(lexer.MUL)
	id := p.parseOptionalIdentifier()
	typeParameters := p.parseOptionalTypeParameters()

	params, returnType, body, err := p.parseFunctionParamsAndBody(async, generator)
	if err != nil {
		return nil, err
	}

	return &ast.FunctionExpression{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeFunctionExpression.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		ID:             id,
		Params:         params,
		Body:           body,
		Generator:      generator,
		Async:          async,
		ReturnType:     returnType,
		TypeParameters: typeParameters,
	}, nil
}

// parseFunctionParamsAndBody parses function expression parameters, return type, and body (requires block).
func (p *Parser) parseFunctionParamsAndBody(async, generator bool) ([]ast.Pattern, *ast.TSTypeAnnotation, *ast.BlockStatement, error) {
	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, nil, nil, err
	}

	oldInFunction := p.inFunction
	oldAllowYield := p.allowYield
	oldAllowAwait := p.allowAwait
	p.inFunction = true
	p.allowYield = generator
	p.allowAwait = async

	params, err := p.parseFunctionParams()
	if err != nil {
		return nil, nil, nil, err
	}

	// Parse return type annotation (TypeScript)
	var returnType *ast.TSTypeAnnotation
	if p.consume(lexer.COLON) {
		returnType, err = p.parseTSTypeAnnotation()
		if err != nil {
			return nil, nil, nil, err
		}
	}

	// Parse body
	body, err := p.parseBlockStatement()
	if err != nil {
		return nil, nil, nil, err
	}

	p.inFunction = oldInFunction
	p.allowYield = oldAllowYield
	p.allowAwait = oldAllowAwait

	return params, returnType, body, nil
}

// parseFunctionExpressionBody parses the body of a function expression (used for object methods).
func (p *Parser) parseFunctionExpressionBody(async, generator bool) (*ast.FunctionExpression, error) {
	start := p.current.Pos

	// Parse type parameters (TypeScript)
	var typeParameters *ast.TSTypeParameterDeclaration
	if p.current.Type == lexer.LSS {
		var err error
		typeParameters, err = p.parseTSTypeParameters()
		if err != nil {
			typeParameters = nil
		}
	}

	// Parse parameters
	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, err
	}

	oldInFunction := p.inFunction
	oldAllowYield := p.allowYield
	oldAllowAwait := p.allowAwait
	p.inFunction = true
	p.allowYield = generator
	p.allowAwait = async

	params, err := p.parseFunctionParams()
	if err != nil {
		return nil, err
	}

	// Parse return type annotation (TypeScript)
	var returnType *ast.TSTypeAnnotation
	if p.consume(lexer.COLON) {
		returnType, err = p.parseTSTypeAnnotation()
		if err != nil {
			return nil, err
		}
	}

	// Parse body
	body, err := p.parseBlockStatement()
	if err != nil {
		return nil, err
	}

	p.inFunction = oldInFunction
	p.allowYield = oldAllowYield
	p.allowAwait = oldAllowAwait

	return &ast.FunctionExpression{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeFunctionExpression.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Params:         params,
		Body:           body,
		Generator:      generator,
		Async:          async,
		ReturnType:     returnType,
		TypeParameters: typeParameters,
	}, nil
}

// parseFunctionParams parses function parameters.
func (p *Parser) parseFunctionParams() ([]ast.Pattern, error) {
	params := []ast.Pattern{}

	for !p.match(lexer.RPAREN) && !p.isAtEnd() {
		// Handle rest parameter
		if p.consume(lexer.ELLIPSIS) {
			param, err := p.parseRestParameter()
			if err != nil {
				return nil, err
			}
			params = append(params, param)
			break
		}

		param, err := p.parseSingleFunctionParam()
		if err != nil {
			return nil, err
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

// parseRestParameter parses a rest parameter (...param).
func (p *Parser) parseRestParameter() (ast.Pattern, error) {
	param, err := p.parseBindingPattern()
	if err != nil {
		return nil, err
	}
	return &ast.RestElement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeRestElement.String(),
		},
		Argument: param,
	}, nil
}

// parseSingleFunctionParam parses a single function parameter with optional type and default value.
func (p *Parser) parseSingleFunctionParam() (ast.Pattern, error) {
	param, err := p.parseBindingPattern()
	if err != nil {
		return nil, err
	}

	// Parse type annotation (TypeScript)
	param = p.parseParamTypeAnnotation(param)

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

	return param, nil
}

// parseParamTypeAnnotation parses optional type annotation for a parameter.
func (p *Parser) parseParamTypeAnnotation(param ast.Pattern) ast.Pattern {
	if id, ok := param.(*ast.Identifier); ok {
		if p.consume(lexer.QUESTION) {
			id.Optional = true
		}
		if p.consume(lexer.COLON) {
			typeAnnotation, err := p.parseTSTypeAnnotation()
			if err == nil {
				id.TypeAnnotation = typeAnnotation
			}
		}
	}
	return param
}

// parseParenthesizedOrArrowFunction parses a parenthesized expression or arrow function.
func (p *Parser) parseParenthesizedOrArrowFunction() (ast.Expression, error) {
	start := p.current.Pos
	p.nextToken() // consume '('

	// Empty parameter list for arrow function
	if p.match(lexer.RPAREN) {
		p.nextToken()
		if p.current.Type == lexer.ARROW {
			return p.parseArrowFunctionFromParams(start, []ast.Pattern{})
		}
		return nil, p.errorAtCurrent("unexpected empty parentheses")
	}

	// Try to parse as parameters or expression
	// This is ambiguous - could be (x) => x or (x + 1)

	// Save state for potential backtracking
	savedPos := p.current.Pos

	// Try parsing as parameters first
	params, err := p.parseFunctionParams()
	if err == nil && p.current.Type == lexer.ARROW {
		return p.parseArrowFunctionFromParams(start, params)
	}

	// Not arrow function, parse as parenthesized expression
	// Backtrack
	p.current.Pos = savedPos

	// Skip opening paren (already consumed)
	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.RPAREN); err != nil {
		return nil, err
	}

	// Check if it's actually an arrow function
	if p.current.Type == lexer.ARROW {
		// Convert expression to params
		// This is a simplified version - in reality, we'd need more sophisticated conversion
		if id, ok := expr.(*ast.Identifier); ok {
			return p.parseArrowFunctionFromParams(start, []ast.Pattern{id})
		}
		return nil, p.errorAtCurrent("invalid arrow function parameters")
	}

	return expr, nil
}

// parseArrowFunctionFromParams parses an arrow function given the parameters.
func (p *Parser) parseArrowFunctionFromParams(start int, params []ast.Pattern) (*ast.ArrowFunctionExpression, error) {
	if err := p.expect(lexer.ARROW); err != nil {
		return nil, err
	}

	// Parse body
	var body ast.Node
	var err error

	oldInFunction := p.inFunction
	p.inFunction = true

	if p.current.Type == lexer.LBRACE {
		body, err = p.parseBlockStatement()
	} else {
		// Expression body
		expr, err := p.parseAssignmentExpression()
		if err != nil {
			return nil, err
		}
		body = expr
	}

	p.inFunction = oldInFunction

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
		Async:  false,
	}, nil
}

// parseClassDeclaration parses a class declaration.
func (p *Parser) parseClassDeclaration() (*ast.ClassDeclaration, error) {
	start := p.current.Pos
	p.nextToken() // consume 'class'

	// Parse class name
	var id *ast.Identifier
	if p.current.Type == lexer.IDENT {
		id = &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
		p.nextToken()
	}

	// Parse type parameters (TypeScript)
	var typeParameters *ast.TSTypeParameterDeclaration
	if p.current.Type == lexer.LSS {
		var err error
		typeParameters, err = p.parseTSTypeParameters()
		if err != nil {
			typeParameters = nil
		}
	}

	// Parse extends clause
	var superClass ast.Expression
	var superTypeParameters *ast.TSTypeParameterInstantiation
	if p.consume(lexer.EXTENDS) {
		var err error
		superClass, err = p.parseLeftHandSideExpression()
		if err != nil {
			return nil, err
		}

		// Parse type arguments for super class (TypeScript)
		if p.current.Type == lexer.LSS {
			superTypeParameters, err = p.parseTSTypeArguments()
			if err != nil {
				superTypeParameters = nil
			}
		}
	}

	// Parse implements clause (TypeScript)
	var implements []ast.TSClassImplements
	if p.consume(lexer.IMPLEMENTS) {
		for {
			impl, err := p.parseTSClassImplements()
			if err != nil {
				return nil, err
			}
			implements = append(implements, *impl)

			if !p.consume(lexer.COMMA) {
				break
			}
		}
	}

	// Parse class body
	body, err := p.parseClassBody()
	if err != nil {
		return nil, err
	}

	return &ast.ClassDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeClassDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		ID:                  id,
		SuperClass:          superClass,
		Body:                body,
		TypeParameters:      typeParameters,
		SuperTypeParameters: superTypeParameters,
		Implements:          implements,
	}, nil
}

// parseClassExpression parses a class expression.
func (p *Parser) parseClassExpression() (*ast.ClassExpression, error) {
	start := p.current.Pos
	p.nextToken() // consume 'class'

	// Parse optional class name
	var id *ast.Identifier
	if p.current.Type == lexer.IDENT {
		id = &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
		p.nextToken()
	}

	// Parse type parameters (TypeScript)
	var typeParameters *ast.TSTypeParameterDeclaration
	if p.current.Type == lexer.LSS {
		var err error
		typeParameters, err = p.parseTSTypeParameters()
		if err != nil {
			typeParameters = nil
		}
	}

	// Parse extends clause
	var superClass ast.Expression
	var superTypeParameters *ast.TSTypeParameterInstantiation
	if p.consume(lexer.EXTENDS) {
		var err error
		superClass, err = p.parseLeftHandSideExpression()
		if err != nil {
			return nil, err
		}

		// Parse type arguments for super class (TypeScript)
		if p.current.Type == lexer.LSS {
			superTypeParameters, err = p.parseTSTypeArguments()
			if err != nil {
				superTypeParameters = nil
			}
		}
	}

	// Parse implements clause (TypeScript)
	var implements []ast.TSClassImplements
	if p.consume(lexer.IMPLEMENTS) {
		for {
			impl, err := p.parseTSClassImplements()
			if err != nil {
				return nil, err
			}
			implements = append(implements, *impl)

			if !p.consume(lexer.COMMA) {
				break
			}
		}
	}

	// Parse class body
	body, err := p.parseClassBody()
	if err != nil {
		return nil, err
	}

	return &ast.ClassExpression{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeClassExpression.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		ID:                  id,
		SuperClass:          superClass,
		Body:                body,
		TypeParameters:      typeParameters,
		SuperTypeParameters: superTypeParameters,
		Implements:          implements,
	}, nil
}

// parseClassBody parses a class body.
func (p *Parser) parseClassBody() (*ast.ClassBody, error) {
	start := p.current.Pos
	if err := p.expect(lexer.LBRACE); err != nil {
		return nil, err
	}

	oldInClass := p.inClass
	p.inClass = true

	body := []interface{}{}

	for !p.match(lexer.RBRACE) && !p.isAtEnd() {
		// Parse class element
		element, err := p.parseClassElement()
		if err != nil {
			p.synchronize()
			continue
		}
		if element != nil {
			body = append(body, element)
		}
	}

	p.inClass = oldInClass

	if err := p.expect(lexer.RBRACE); err != nil {
		return nil, err
	}

	return &ast.ClassBody{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeClassBody.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Body: body,
	}, nil
}

// parseClassElement parses a class member (method, property, etc.).
func (p *Parser) parseClassElement() (ast.Node, error) {
	start := p.current.Pos

	// Skip semicolons
	if p.consume(lexer.SEMICOLON) {
		return nil, nil
	}

	// Check for static block
	if p.current.Type == lexer.STATIC && p.peek.Type == lexer.LBRACE {
		p.nextToken()
		body, err := p.parseBlockStatement()
		if err != nil {
			return nil, err
		}
		return &ast.StaticBlock{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeStaticBlock.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Body: []ast.Statement{body},
		}, nil
	}

	// Parse modifiers
	isStatic := p.consume(lexer.STATIC)
	_ = false // isAbstract - not used in PropertyDefinition
	isReadonly := false
	isDeclare := false

	if p.current.Type == lexer.READONLY {
		isReadonly = true
		p.nextToken()
	}

	// Check for async/generator
	async := p.consume(lexer.ASYNC)
	generator := false
	if async && p.consume(lexer.MUL) {
		generator = true
	} else if p.consume(lexer.MUL) {
		generator = true
	}

	// Parse accessor type (get/set)
	kind := "method"
	switch p.current.Type {
	case lexer.GET:
		kind = "get"
		p.nextToken()
	case lexer.SET:
		kind = "set"
		p.nextToken()
	case lexer.CONSTRUCTOR:
		kind = "constructor"
	}

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
		return nil, p.errorAtCurrent("expected class member key")
	}

	// Check for property vs method
	optional := p.consume(lexer.QUESTION)

	// Try to parse type annotation
	var typeAnnotation *ast.TSTypeAnnotation
	if p.current.Type == lexer.COLON {
		p.nextToken()
		typeAnnotation, err = p.parseTSTypeAnnotation()
		if err != nil {
			return nil, err
		}
	}

	// Check if it's a method (has parameters)
	if p.current.Type == lexer.LPAREN || kind == "get" || kind == "set" || kind == "constructor" || generator {
		// Method
		value, err := p.parseFunctionExpressionBody(async, generator)
		if err != nil {
			return nil, err
		}

		return &ast.MethodDefinition{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeMethodDefinition.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Key:      key,
			Value:    value,
			Kind:     kind,
			Computed: computed,
			Static:   isStatic,
		}, nil
	}

	// Property
	var value ast.Expression
	if p.consume(lexer.ASSIGN) {
		value, err = p.parseAssignmentExpression()
		if err != nil {
			return nil, err
		}
	}

	p.consume(lexer.SEMICOLON)

	return &ast.PropertyDefinition{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypePropertyDefinition.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Key:            key,
		Value:          value,
		Computed:       computed,
		Static:         isStatic,
		TypeAnnotation: typeAnnotation,
		Optional:       optional,
		Readonly:       isReadonly,
		Declare:        isDeclare,
	}, nil
}
