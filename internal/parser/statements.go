package parser

import (
	"fmt"

	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/lexer"
)

// parseStatementListItem parses a statement or declaration at the top level or in a block.
func (p *Parser) parseStatementListItem() (ast.Statement, error) {
	// Check for declarations first
	switch p.current.Type {
	case lexer.FUNCTION:
		return p.parseFunctionDeclaration()
	case lexer.CLASS:
		return p.parseClassDeclaration()
	case lexer.CONST, lexer.LET, lexer.VAR:
		return p.parseVariableStatement()
	case lexer.INTERFACE:
		return p.parseTSInterfaceDeclaration()
	case lexer.TYPE:
		return p.parseTSTypeAliasDeclaration()
	case lexer.ENUM:
		return p.parseTSEnumDeclaration()
	case lexer.NAMESPACE, lexer.MODULE:
		return p.parseTSModuleDeclaration()
	case lexer.IMPORT:
		return p.parseImportDeclaration()
	case lexer.EXPORT:
		return p.parseExportDeclaration()
	case lexer.DECLARE:
		return p.parseTSDeclareStatement()
	case lexer.ASYNC:
		// Could be async function declaration
		if p.peek.Type == lexer.FUNCTION {
			return p.parseFunctionDeclaration()
		}
		return p.parseStatement()
	default:
		return p.parseStatement()
	}
}

// parseStatement parses a statement.
func (p *Parser) parseStatement() (ast.Statement, error) {
	switch p.current.Type {
	case lexer.IF:
		return p.parseIfStatement()
	case lexer.WHILE:
		return p.parseWhileStatement()
	case lexer.DO:
		return p.parseDoWhileStatement()
	case lexer.FOR:
		return p.parseForStatement()
	case lexer.SWITCH:
		return p.parseSwitchStatement()
	case lexer.RETURN:
		return p.parseReturnStatement()
	case lexer.BREAK:
		return p.parseBreakStatement()
	case lexer.CONTINUE:
		return p.parseContinueStatement()
	case lexer.THROW:
		return p.parseThrowStatement()
	case lexer.TRY:
		return p.parseTryStatement()
	case lexer.DEBUGGER:
		return p.parseDebuggerStatement()
	case lexer.WITH:
		return p.parseWithStatement()
	case lexer.LBRACE:
		return p.parseBlockStatement()
	case lexer.SEMICOLON:
		return p.parseEmptyStatement()
	default:
		// Expression statement or labeled statement
		return p.parseExpressionOrLabeledStatement()
	}
}

// parseBlockStatement parses a block statement { ... }.
func (p *Parser) parseBlockStatement() (*ast.BlockStatement, error) {
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

	return &ast.BlockStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeBlockStatement.String(),
			Range:    &ast.Range{start, p.current.End},
		},
		Body: body,
	}, nil
}

// parseVariableStatement parses a variable declaration statement.
func (p *Parser) parseVariableStatement() (*ast.VariableDeclaration, error) {
	start := p.current.Pos
	kind := p.current.Literal // "var", "let", "const"
	p.nextToken()

	declarations := []*ast.VariableDeclarator{}

	// Parse first declarator
	declarator, err := p.parseVariableDeclarator()
	if err != nil {
		return nil, err
	}
	declarations = append(declarations, declarator)

	// Parse additional declarators
	for p.consume(lexer.COMMA) {
		declarator, err := p.parseVariableDeclarator()
		if err != nil {
			return nil, err
		}
		declarations = append(declarations, declarator)
	}

	// Consume optional semicolon
	p.consume(lexer.SEMICOLON)

	return &ast.VariableDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeVariableDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Declarations: declarations,
		Kind:         kind,
	}, nil
}

// parseVariableDeclarator parses a single variable declarator (id = init).
func (p *Parser) parseVariableDeclarator() (*ast.VariableDeclarator, error) {
	start := p.current.Pos

	// Parse the pattern (identifier or destructuring)
	id, err := p.parseBindingPattern()
	if err != nil {
		return nil, err
	}

	var init ast.Expression
	if p.consume(lexer.ASSIGN) {
		init, err = p.parseAssignmentExpression()
		if err != nil {
			return nil, err
		}
	}

	return &ast.VariableDeclarator{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeVariableDeclarator.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		ID:   id,
		Init: init,
	}, nil
}

// parseIfStatement parses an if statement.
func (p *Parser) parseIfStatement() (*ast.IfStatement, error) {
	start := p.current.Pos
	p.nextToken() // consume 'if'

	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, err
	}

	test, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.RPAREN); err != nil {
		return nil, err
	}

	consequent, err := p.parseStatement()
	if err != nil {
		return nil, err
	}

	var alternate ast.Statement
	if p.consume(lexer.ELSE) {
		alternate, err = p.parseStatement()
		if err != nil {
			return nil, err
		}
	}

	return &ast.IfStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeIfStatement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Test:       test,
		Consequent: consequent,
		Alternate:  alternate,
	}, nil
}

// parseWhileStatement parses a while statement.
func (p *Parser) parseWhileStatement() (*ast.WhileStatement, error) {
	start := p.current.Pos
	p.nextToken() // consume 'while'

	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, err
	}

	test, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.RPAREN); err != nil {
		return nil, err
	}

	oldInLoop := p.inLoop
	p.inLoop = true
	body, err := p.parseStatement()
	p.inLoop = oldInLoop

	if err != nil {
		return nil, err
	}

	return &ast.WhileStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeWhileStatement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Test: test,
		Body: body,
	}, nil
}

// parseDoWhileStatement parses a do-while statement.
func (p *Parser) parseDoWhileStatement() (*ast.DoWhileStatement, error) {
	start := p.current.Pos
	p.nextToken() // consume 'do'

	oldInLoop := p.inLoop
	p.inLoop = true
	body, err := p.parseStatement()
	p.inLoop = oldInLoop

	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.WHILE); err != nil {
		return nil, err
	}

	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, err
	}

	test, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.RPAREN); err != nil {
		return nil, err
	}

	p.consume(lexer.SEMICOLON)

	return &ast.DoWhileStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeDoWhileStatement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Body: body,
		Test: test,
	}, nil
}

// parseForStatement parses a for, for-in, or for-of statement.
func (p *Parser) parseForStatement() (ast.Statement, error) {
	start := p.current.Pos
	p.nextToken() // consume 'for'

	await := false
	if p.consume(lexer.AWAIT) {
		await = true
	}

	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, err
	}

	// Parse init
	var init ast.Node
	var err error

	if p.match(lexer.VAR, lexer.LET, lexer.CONST) {
		// Variable declaration
		kind := p.current.Literal
		p.nextToken()

		// Check if this is a for-in or for-of loop
		if p.current.Type == lexer.IDENT {
			idStart := p.current.Pos
			id := &ast.Identifier{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeIdentifier.String(),
					Range:    &ast.Range{idStart, p.current.End},
				},
				Name: p.current.Literal,
			}
			p.nextToken()

			// Parse type annotation if present
			if p.consume(lexer.COLON) {
				typeAnnotation, err := p.parseTSTypeAnnotation()
				if err != nil {
					return nil, err
				}
				id.TypeAnnotation = typeAnnotation
			}

			if p.match(lexer.IN, lexer.OF) {
				return p.parseForInOfStatement(start, kind, id, await)
			}

			// Not for-in/of, parse as regular variable declarator
			var initExpr ast.Expression
			if p.consume(lexer.ASSIGN) {
				initExpr, err = p.parseAssignmentExpression()
				if err != nil {
					return nil, err
				}
			}

			declarator := &ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeVariableDeclarator.String(),
					Range:    &ast.Range{idStart, p.current.Pos},
				},
				ID:   id,
				Init: initExpr,
			}

			init = &ast.VariableDeclaration{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeVariableDeclaration.String(),
					Range:    &ast.Range{start, p.current.Pos},
				},
				Declarations: []*ast.VariableDeclarator{declarator},
				Kind:         kind,
			}
		}
	} else if !p.consume(lexer.SEMICOLON) {
		// Expression
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		// Check for for-in or for-of
		if p.match(lexer.IN, lexer.OF) {
			return p.parseForInOfStatement(start, "", expr, await)
		}

		init = expr
		p.consume(lexer.SEMICOLON)
	}

	// Regular for loop
	var test ast.Expression
	if !p.match(lexer.SEMICOLON) {
		test, err = p.parseExpression()
		if err != nil {
			return nil, err
		}
	}
	if err := p.expect(lexer.SEMICOLON); err != nil {
		return nil, err
	}

	var update ast.Expression
	if !p.match(lexer.RPAREN) {
		update, err = p.parseExpression()
		if err != nil {
			return nil, err
		}
	}

	if err := p.expect(lexer.RPAREN); err != nil {
		return nil, err
	}

	oldInLoop := p.inLoop
	p.inLoop = true
	body, err := p.parseStatement()
	p.inLoop = oldInLoop

	if err != nil {
		return nil, err
	}

	return &ast.ForStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeForStatement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Init:   init,
		Test:   test,
		Update: update,
		Body:   body,
	}, nil
}

// parseForInOfStatement parses for-in or for-of statement.
func (p *Parser) parseForInOfStatement(start int, kind string, left ast.Node, await bool) (ast.Statement, error) {
	isForOf := p.current.Type == lexer.OF
	p.nextToken() // consume 'in' or 'of'

	right, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.RPAREN); err != nil {
		return nil, err
	}

	oldInLoop := p.inLoop
	p.inLoop = true
	body, err := p.parseStatement()
	p.inLoop = oldInLoop

	if err != nil {
		return nil, err
	}

	// Convert left to proper format
	var leftNode ast.Node = left
	if kind != "" {
		// Wrap in variable declaration
		var pattern ast.Pattern
		if id, ok := left.(*ast.Identifier); ok {
			pattern = id
		} else if p, ok := left.(ast.Pattern); ok {
			pattern = p
		}

		leftNode = &ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeVariableDeclaration.String(),
			},
			Declarations: []*ast.VariableDeclarator{
				{
					BaseNode: ast.BaseNode{
						NodeType: ast.NodeTypeVariableDeclarator.String(),
					},
					ID: pattern,
				},
			},
			Kind: kind,
		}
	}

	if isForOf {
		return &ast.ForOfStatement{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeForOfStatement.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Left:  leftNode,
			Right: right,
			Body:  body,
			Await: await,
		}, nil
	}

	return &ast.ForInStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeForInStatement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Left:  leftNode,
		Right: right,
		Body:  body,
	}, nil
}

// parseReturnStatement parses a return statement.
func (p *Parser) parseReturnStatement() (*ast.ReturnStatement, error) {
	start := p.current.Pos
	p.nextToken() // consume 'return'

	var argument ast.Expression
	if !p.match(lexer.SEMICOLON) && !p.isAtEnd() && p.current.Line == p.peek.Line {
		var err error
		argument, err = p.parseExpression()
		if err != nil {
			return nil, err
		}
	}

	p.consume(lexer.SEMICOLON)

	return &ast.ReturnStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeReturnStatement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Argument: argument,
	}, nil
}

// parseBreakStatement parses a break statement.
func (p *Parser) parseBreakStatement() (*ast.BreakStatement, error) {
	start := p.current.Pos
	p.nextToken() // consume 'break'

	var label *ast.Identifier
	if p.current.Type == lexer.IDENT {
		label = &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
		p.nextToken()
	}

	p.consume(lexer.SEMICOLON)

	return &ast.BreakStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeBreakStatement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Label: label,
	}, nil
}

// parseContinueStatement parses a continue statement.
func (p *Parser) parseContinueStatement() (*ast.ContinueStatement, error) {
	start := p.current.Pos
	p.nextToken() // consume 'continue'

	var label *ast.Identifier
	if p.current.Type == lexer.IDENT {
		label = &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
		p.nextToken()
	}

	p.consume(lexer.SEMICOLON)

	return &ast.ContinueStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeContinueStatement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Label: label,
	}, nil
}

// parseThrowStatement parses a throw statement.
func (p *Parser) parseThrowStatement() (*ast.ThrowStatement, error) {
	start := p.current.Pos
	p.nextToken() // consume 'throw'

	// No line terminator allowed between throw and its expression
	if p.current.Line != p.peek.Line {
		return nil, p.errorAtCurrent("line break is not allowed between 'throw' and its expression")
	}

	argument, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	p.consume(lexer.SEMICOLON)

	return &ast.ThrowStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeThrowStatement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Argument: argument,
	}, nil
}

// parseTryStatement parses a try-catch-finally statement.
func (p *Parser) parseTryStatement() (*ast.TryStatement, error) {
	start := p.current.Pos
	p.nextToken() // consume 'try'

	block, err := p.parseBlockStatement()
	if err != nil {
		return nil, err
	}

	var handler *ast.CatchClause
	if p.consume(lexer.CATCH) {
		handler, err = p.parseCatchClause()
		if err != nil {
			return nil, err
		}
	}

	var finalizer *ast.BlockStatement
	if p.consume(lexer.FINALLY) {
		finalizer, err = p.parseBlockStatement()
		if err != nil {
			return nil, err
		}
	}

	if handler == nil && finalizer == nil {
		return nil, p.errorAtCurrent("try statement must have either catch or finally")
	}

	return &ast.TryStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTryStatement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Block:     block,
		Handler:   handler,
		Finalizer: finalizer,
	}, nil
}

// parseCatchClause parses a catch clause.
func (p *Parser) parseCatchClause() (*ast.CatchClause, error) {
	start := p.current.Pos

	var param ast.Pattern
	if p.consume(lexer.LPAREN) {
		var err error
		param, err = p.parseBindingPattern()
		if err != nil {
			return nil, err
		}
		if err := p.expect(lexer.RPAREN); err != nil {
			return nil, err
		}
	}

	body, err := p.parseBlockStatement()
	if err != nil {
		return nil, err
	}

	return &ast.CatchClause{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeCatchClause.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Param: param,
		Body:  body,
	}, nil
}

// parseSwitchStatement parses a switch statement.
func (p *Parser) parseSwitchStatement() (*ast.SwitchStatement, error) {
	start := p.current.Pos
	p.nextToken() // consume 'switch'

	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, err
	}

	discriminant, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.RPAREN); err != nil {
		return nil, err
	}

	if err := p.expect(lexer.LBRACE); err != nil {
		return nil, err
	}

	oldInSwitch := p.inSwitch
	p.inSwitch = true

	cases := []*ast.SwitchCase{}
	for !p.match(lexer.RBRACE) && !p.isAtEnd() {
		switchCase, err := p.parseSwitchCase()
		if err != nil {
			p.synchronize()
			continue
		}
		cases = append(cases, switchCase)
	}

	p.inSwitch = oldInSwitch

	if err := p.expect(lexer.RBRACE); err != nil {
		return nil, err
	}

	return &ast.SwitchStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeSwitchStatement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Discriminant: discriminant,
		Cases:        cases,
	}, nil
}

// parseSwitchCase parses a switch case or default clause.
func (p *Parser) parseSwitchCase() (*ast.SwitchCase, error) {
	start := p.current.Pos

	var test ast.Expression
	if p.consume(lexer.CASE) {
		var err error
		test, err = p.parseExpression()
		if err != nil {
			return nil, err
		}
	} else if p.consume(lexer.DEFAULT) {
		// test remains nil for default case
	} else {
		return nil, p.errorAtCurrent("expected 'case' or 'default'")
	}

	if err := p.expect(lexer.COLON); err != nil {
		return nil, err
	}

	consequent := []ast.Statement{}
	for !p.match(lexer.CASE, lexer.DEFAULT, lexer.RBRACE) && !p.isAtEnd() {
		stmt, err := p.parseStatementListItem()
		if err != nil {
			p.synchronize()
			continue
		}
		consequent = append(consequent, stmt)
	}

	return &ast.SwitchCase{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeSwitchCase.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Test:       test,
		Consequent: consequent,
	}, nil
}

// parseDebuggerStatement parses a debugger statement.
func (p *Parser) parseDebuggerStatement() (*ast.DebuggerStatement, error) {
	start := p.current.Pos
	p.nextToken() // consume 'debugger'
	p.consume(lexer.SEMICOLON)

	return &ast.DebuggerStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeDebuggerStatement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
	}, nil
}

// parseWithStatement parses a with statement (discouraged in strict mode).
func (p *Parser) parseWithStatement() (*ast.WithStatement, error) {
	start := p.current.Pos
	p.nextToken() // consume 'with'

	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, err
	}

	object, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if err := p.expect(lexer.RPAREN); err != nil {
		return nil, err
	}

	body, err := p.parseStatement()
	if err != nil {
		return nil, err
	}

	return &ast.WithStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeWithStatement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Object: object,
		Body:   body,
	}, nil
}

// parseEmptyStatement parses an empty statement (;).
func (p *Parser) parseEmptyStatement() (*ast.EmptyStatement, error) {
	start := p.current.Pos
	p.nextToken() // consume ';'

	return &ast.EmptyStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeEmptyStatement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
	}, nil
}

// parseExpressionOrLabeledStatement parses an expression statement or labeled statement.
func (p *Parser) parseExpressionOrLabeledStatement() (ast.Statement, error) {
	start := p.current.Pos
	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	// Check for labeled statement
	if id, ok := expr.(*ast.Identifier); ok && p.consume(lexer.COLON) {
		body, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		return &ast.LabeledStatement{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeLabeledStatement.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Label: id,
			Body:  body,
		}, nil
	}

	p.consume(lexer.SEMICOLON)

	return &ast.ExpressionStatement{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeExpressionStatement.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Expression: expr,
	}, nil
}

// parseTSDeclareStatement parses a TypeScript declare statement.
func (p *Parser) parseTSDeclareStatement() (ast.Statement, error) {
	p.nextToken() // consume 'declare'

	// Parse the declaration after 'declare'
	switch p.current.Type {
	case lexer.VAR, lexer.LET, lexer.CONST:
		return p.parseVariableStatement()
	case lexer.FUNCTION:
		return p.parseFunctionDeclaration()
	case lexer.CLASS:
		return p.parseClassDeclaration()
	case lexer.ENUM:
		return p.parseTSEnumDeclaration()
	case lexer.NAMESPACE, lexer.MODULE:
		return p.parseTSModuleDeclaration()
	default:
		return nil, p.errorAtCurrent(fmt.Sprintf("unexpected token after 'declare': %v", p.current.Type))
	}
}
