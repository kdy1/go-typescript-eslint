package parser

import (
	"fmt"

	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/lexer"
)

// parseStatementListItem parses a statement or declaration at the top level or in a block.
func (p *Parser) parseStatementListItem() (ast.Statement, error) {
	// Check for declarations first
	if stmt, ok := p.tryParseDeclaration(); ok {
		return stmt()
	}
	return p.parseStatement()
}

// tryParseDeclaration tries to parse a declaration statement.
func (p *Parser) tryParseDeclaration() (func() (ast.Statement, error), bool) {
	fn := p.matchDeclarationParser()
	if fn != nil {
		return fn, true
	}
	return nil, false
}

// matchDeclarationParser returns the appropriate parser function for the current token.
func (p *Parser) matchDeclarationParser() func() (ast.Statement, error) {
	switch p.current.Type {
	case lexer.FUNCTION:
		return func() (ast.Statement, error) { return p.parseFunctionDeclaration() }
	case lexer.CLASS:
		return func() (ast.Statement, error) { return p.parseClassDeclaration() }
	case lexer.CONST, lexer.LET, lexer.VAR:
		return func() (ast.Statement, error) { return p.parseVariableStatement() }
	case lexer.INTERFACE:
		return func() (ast.Statement, error) { return p.parseTSInterfaceDeclaration() }
	case lexer.TYPE:
		return func() (ast.Statement, error) { return p.parseTSTypeAliasDeclaration() }
	case lexer.ENUM:
		return func() (ast.Statement, error) { return p.parseTSEnumDeclaration() }
	case lexer.NAMESPACE, lexer.MODULE:
		return func() (ast.Statement, error) { return p.parseTSModuleDeclaration() }
	case lexer.IMPORT:
		return func() (ast.Statement, error) { return p.parseImportDeclaration() }
	case lexer.EXPORT:
		return p.parseExportDeclaration
	case lexer.DECLARE:
		return func() (ast.Statement, error) { return p.parseTSDeclareStatement() }
	case lexer.ASYNC:
		return p.matchAsyncFunctionDeclaration()
	}
	return nil
}

// matchAsyncFunctionDeclaration checks if the current token is an async function declaration.
func (p *Parser) matchAsyncFunctionDeclaration() func() (ast.Statement, error) {
	if p.peek.Type == lexer.FUNCTION {
		return func() (ast.Statement, error) { return p.parseFunctionDeclaration() }
	}
	return nil
}

// parseStatement parses a statement.
func (p *Parser) parseStatement() (ast.Statement, error) {
	// Try control flow statements
	if stmt := p.tryParseControlFlowStatement(); stmt != nil {
		return stmt()
	}
	// Try simple statements
	if stmt := p.tryParseSimpleStatement(); stmt != nil {
		return stmt()
	}
	// Default: expression statement or labeled statement
	return p.parseExpressionOrLabeledStatement()
}

// tryParseControlFlowStatement tries to parse control flow statements.
func (p *Parser) tryParseControlFlowStatement() func() (ast.Statement, error) {
	switch p.current.Type {
	case lexer.IF:
		return func() (ast.Statement, error) { return p.parseIfStatement() }
	case lexer.WHILE:
		return func() (ast.Statement, error) { return p.parseWhileStatement() }
	case lexer.DO:
		return func() (ast.Statement, error) { return p.parseDoWhileStatement() }
	case lexer.FOR:
		return func() (ast.Statement, error) { return p.parseForStatement() }
	case lexer.SWITCH:
		return func() (ast.Statement, error) { return p.parseSwitchStatement() }
	case lexer.TRY:
		return func() (ast.Statement, error) { return p.parseTryStatement() }
	}
	return nil
}

// tryParseSimpleStatement tries to parse simple statements.
func (p *Parser) tryParseSimpleStatement() func() (ast.Statement, error) {
	switch p.current.Type {
	case lexer.RETURN:
		return func() (ast.Statement, error) { return p.parseReturnStatement() }
	case lexer.BREAK:
		return func() (ast.Statement, error) { return p.parseBreakStatement() }
	case lexer.CONTINUE:
		return func() (ast.Statement, error) { return p.parseContinueStatement() }
	case lexer.THROW:
		return func() (ast.Statement, error) { return p.parseThrowStatement() }
	case lexer.DEBUGGER:
		return func() (ast.Statement, error) { return p.parseDebuggerStatement() }
	case lexer.WITH:
		return func() (ast.Statement, error) { return p.parseWithStatement() }
	case lexer.LBRACE:
		return func() (ast.Statement, error) { return p.parseBlockStatement() }
	case lexer.SEMICOLON:
		return func() (ast.Statement, error) { return p.parseEmptyStatement() }
	}
	return nil
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

	declarations := []ast.VariableDeclarator{}

	// Parse first declarator
	declarator, err := p.parseVariableDeclarator()
	if err != nil {
		return nil, err
	}
	declarations = append(declarations, *declarator)

	// Parse additional declarators
	for p.consume(lexer.COMMA) {
		declarator, err := p.parseVariableDeclarator()
		if err != nil {
			return nil, err
		}
		declarations = append(declarations, *declarator)
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

	await := p.consume(lexer.AWAIT)

	if err := p.expect(lexer.LPAREN); err != nil {
		return nil, err
	}

	// Parse init part and check if it's for-in/for-of
	init, isForInOf, err := p.parseForInit(start, await)
	if err != nil {
		return nil, err
	}

	// If it's for-in or for-of, return the result
	if stmt, ok := init.(ast.Statement); ok && isForInOf {
		return stmt, nil
	}

	// Regular for loop - parse test and update
	return p.parseRegularForLoop(start, init)
}

// parseForInit parses the init part of a for statement and returns whether it's for-in/of.
func (p *Parser) parseForInit(start int, await bool) (ast.Node, bool, error) {
	if p.match(lexer.VAR, lexer.LET, lexer.CONST) {
		return p.parseForVarInit(start, await)
	}

	if p.consume(lexer.SEMICOLON) {
		return nil, false, nil
	}

	// Expression init
	expr, err := p.parseExpression()
	if err != nil {
		return nil, false, err
	}

	// Check for for-in or for-of
	if p.match(lexer.IN, lexer.OF) {
		stmt, err := p.parseForInOfStatement(start, "", expr, await)
		return stmt, true, err
	}

	p.consume(lexer.SEMICOLON)
	return expr, false, nil
}

// parseForVarInit parses variable declaration in for statement init.
func (p *Parser) parseForVarInit(start int, await bool) (ast.Node, bool, error) {
	kind := p.current.Literal
	p.nextToken()

	if p.current.Type != lexer.IDENT {
		return nil, false, p.errorAtCurrent("expected identifier")
	}

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
			return nil, false, err
		}
		id.TypeAnnotation = typeAnnotation
	}

	// Check if it's for-in or for-of
	if p.match(lexer.IN, lexer.OF) {
		stmt, err := p.parseForInOfStatement(start, kind, id, await)
		return stmt, true, err
	}

	// Regular variable declaration
	return p.createForVarDeclaration(start, kind, idStart, id)
}

// createForVarDeclaration creates a variable declaration for regular for loop.
func (p *Parser) createForVarDeclaration(start int, kind string, idStart int, id *ast.Identifier) (ast.Node, bool, error) {
	var initExpr ast.Expression
	var err error
	if p.consume(lexer.ASSIGN) {
		initExpr, err = p.parseAssignmentExpression()
		if err != nil {
			return nil, false, err
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

	init := &ast.VariableDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeVariableDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Declarations: []ast.VariableDeclarator{*declarator},
		Kind:         kind,
	}

	return init, false, nil
}

// parseRegularForLoop parses the test, update, and body of a regular for loop.
func (p *Parser) parseRegularForLoop(start int, init ast.Node) (*ast.ForStatement, error) {
	var test ast.Expression
	var err error
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
	leftNode := left
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
			Declarations: []ast.VariableDeclarator{
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

	cases := []ast.SwitchCase{}
	for !p.match(lexer.RBRACE) && !p.isAtEnd() {
		switchCase, err := p.parseSwitchCase()
		if err != nil {
			p.synchronize()
			continue
		}
		cases = append(cases, *switchCase)
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
	} else if !p.consume(lexer.DEFAULT) {
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
//
//nolint:unparam // error return required to match statement parsing interface
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
//
//nolint:unparam // error return required to match statement parsing interface
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
