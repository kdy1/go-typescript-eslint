package parser

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/lexer"
)

// parseImportDeclaration parses an import declaration.
func (p *Parser) parseImportDeclaration() (*ast.ImportDeclaration, error) {
	start := p.current.Pos
	p.nextToken() // consume 'import'

	// Handle import type (TypeScript)
	importKind := "value"
	if p.current.Type == lexer.TYPE {
		importKind = "type"
		p.nextToken()
	}

	specifiers := []ast.Node{}

	// Check for import 'module' (side-effect import)
	if p.current.Type == lexer.STRING {
		source := &ast.Literal{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeLiteral.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Value: p.current.Literal,
			Raw:   p.current.Literal,
		}
		p.nextToken()
		p.consume(lexer.SEMICOLON)

		return &ast.ImportDeclaration{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeImportDeclaration.String(),
				Range:    &ast.Range{start, p.current.Pos},
			},
			Specifiers: specifiers,
			Source:     source,
			ImportKind: &importKind,
		}, nil
	}

	// Parse default import
	if p.current.Type == lexer.IDENT {
		specifiers = append(specifiers, &ast.ImportDefaultSpecifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeImportDefaultSpecifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Local: &ast.Identifier{
				BaseNode: ast.BaseNode{
					NodeType: ast.NodeTypeIdentifier.String(),
					Range:    &ast.Range{p.current.Pos, p.current.End},
				},
				Name: p.current.Literal,
			},
		})
		p.nextToken()

		// Check for additional imports after default
		if p.consume(lexer.COMMA) {
			if p.current.Type == lexer.MUL {
				// namespace import
				spec, err := p.parseImportNamespaceSpecifier()
				if err != nil {
					return nil, err
				}
				specifiers = append(specifiers, spec)
			} else if p.current.Type == lexer.LBRACE {
				// named imports
				specs, err := p.parseImportSpecifiers()
				if err != nil {
					return nil, err
				}
				specifiers = append(specifiers, specs...)
			}
		}
	} else if p.current.Type == lexer.MUL {
		// Namespace import
		spec, err := p.parseImportNamespaceSpecifier()
		if err != nil {
			return nil, err
		}
		specifiers = append(specifiers, spec)
	} else if p.current.Type == lexer.LBRACE {
		// Named imports
		specs, err := p.parseImportSpecifiers()
		if err != nil {
			return nil, err
		}
		specifiers = append(specifiers, specs...)
	} else {
		return nil, p.errorAtCurrent("expected import specifier")
	}

	// Parse 'from' clause
	if err := p.expect(lexer.FROM); err != nil {
		return nil, err
	}

	if p.current.Type != lexer.STRING {
		return nil, p.errorAtCurrent("expected module specifier")
	}

	source := &ast.Literal{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeLiteral.String(),
			Range:    &ast.Range{p.current.Pos, p.current.End},
		},
		Value: p.current.Literal,
		Raw:   p.current.Literal,
	}
	p.nextToken()

	// Parse import attributes (with clause)
	var attributes []*ast.ImportAttribute
	if p.current.Type == lexer.IDENT && p.current.Literal == "with" {
		p.nextToken()
		var err error
		attributes, err = p.parseImportAttributes()
		if err != nil {
			return nil, err
		}
	}

	p.consume(lexer.SEMICOLON)

	return &ast.ImportDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeImportDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Specifiers: specifiers,
		Source:     source,
		ImportKind: &importKind,
		Attributes: attributes,
	}, nil
}

// parseImportNamespaceSpecifier parses a namespace import (* as name).
func (p *Parser) parseImportNamespaceSpecifier() (*ast.ImportNamespaceSpecifier, error) {
	start := p.current.Pos
	p.nextToken() // consume '*'

	if err := p.expect(lexer.AS); err != nil {
		return nil, err
	}

	if p.current.Type != lexer.IDENT {
		return nil, p.errorAtCurrent("expected identifier")
	}

	local := &ast.Identifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeIdentifier.String(),
			Range:    &ast.Range{p.current.Pos, p.current.End},
		},
		Name: p.current.Literal,
	}
	p.nextToken()

	return &ast.ImportNamespaceSpecifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeImportNamespaceSpecifier.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Local: local,
	}, nil
}

// parseImportSpecifiers parses named import specifiers {a, b as c}.
func (p *Parser) parseImportSpecifiers() ([]*ast.ImportSpecifier, error) {
	if err := p.expect(lexer.LBRACE); err != nil {
		return nil, err
	}

	specifiers := []*ast.ImportSpecifier{}

	for !p.match(lexer.RBRACE) && !p.isAtEnd() {
		spec, err := p.parseImportSpecifier()
		if err != nil {
			return nil, err
		}
		specifiers = append(specifiers, spec)

		if !p.consume(lexer.COMMA) {
			break
		}
	}

	if err := p.expect(lexer.RBRACE); err != nil {
		return nil, err
	}

	return specifiers, nil
}

// parseImportSpecifier parses a single import specifier.
func (p *Parser) parseImportSpecifier() (*ast.ImportSpecifier, error) {
	start := p.current.Pos

	// Check for type import (TypeScript)
	importKind := "value"
	if p.current.Type == lexer.TYPE {
		importKind = "type"
		p.nextToken()
	}

	if p.current.Type != lexer.IDENT {
		return nil, p.errorAtCurrent("expected identifier")
	}

	imported := &ast.Identifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeIdentifier.String(),
			Range:    &ast.Range{p.current.Pos, p.current.End},
		},
		Name: p.current.Literal,
	}
	p.nextToken()

	local := imported

	// Check for 'as' clause
	if p.consume(lexer.AS) {
		if p.current.Type != lexer.IDENT {
			return nil, p.errorAtCurrent("expected identifier")
		}
		local = &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
		p.nextToken()
	}

	return &ast.ImportSpecifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeImportSpecifier.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Imported:   imported,
		Local:      local,
		ImportKind: &importKind,
	}, nil
}

// parseImportAttributes parses import attributes (with clause).
func (p *Parser) parseImportAttributes() ([]*ast.ImportAttribute, error) {
	if err := p.expect(lexer.LBRACE); err != nil {
		return nil, err
	}

	attributes := []*ast.ImportAttribute{}

	for !p.match(lexer.RBRACE) && !p.isAtEnd() {
		attr, err := p.parseImportAttribute()
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, attr)

		if !p.consume(lexer.COMMA) {
			break
		}
	}

	if err := p.expect(lexer.RBRACE); err != nil {
		return nil, err
	}

	return attributes, nil
}

// parseImportAttribute parses a single import attribute.
func (p *Parser) parseImportAttribute() (*ast.ImportAttribute, error) {
	start := p.current.Pos

	if p.current.Type != lexer.IDENT && p.current.Type != lexer.STRING {
		return nil, p.errorAtCurrent("expected attribute key")
	}

	var key ast.Node
	if p.current.Type == lexer.IDENT {
		key = &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
	} else {
		key = &ast.Literal{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeLiteral.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Value: p.current.Literal,
			Raw:   p.current.Literal,
		}
	}
	p.nextToken()

	if err := p.expect(lexer.COLON); err != nil {
		return nil, err
	}

	if p.current.Type != lexer.STRING {
		return nil, p.errorAtCurrent("expected string value")
	}

	value := &ast.Literal{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeLiteral.String(),
			Range:    &ast.Range{p.current.Pos, p.current.End},
		},
		Value: p.current.Literal,
		Raw:   p.current.Literal,
	}
	p.nextToken()

	return &ast.ImportAttribute{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeImportAttribute.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Key:   key,
		Value: value,
	}, nil
}

// parseExportDeclaration parses an export declaration.
func (p *Parser) parseExportDeclaration() (ast.Statement, error) {
	start := p.current.Pos
	p.nextToken() // consume 'export'

	// Handle export type (TypeScript)
	exportKind := "value"
	if p.current.Type == lexer.TYPE {
		exportKind = "type"
		p.nextToken()
	}

	// Check for default export
	if p.consume(lexer.DEFAULT) {
		declaration, err := p.parseExportDefaultDeclaration(start)
		if err != nil {
			return nil, err
		}
		return declaration, nil
	}

	// Check for export * (re-export all)
	if p.current.Type == lexer.MUL {
		return p.parseExportAllDeclaration(start, exportKind)
	}

	// Check for export {...}
	if p.current.Type == lexer.LBRACE {
		return p.parseExportNamedDeclaration(start, exportKind)
	}

	// Export declaration
	var declaration ast.Node
	var err error

	switch p.current.Type {
	case lexer.VAR, lexer.LET, lexer.CONST:
		declaration, err = p.parseVariableStatement()
	case lexer.FUNCTION:
		declaration, err = p.parseFunctionDeclaration()
	case lexer.CLASS:
		declaration, err = p.parseClassDeclaration()
	case lexer.INTERFACE:
		declaration, err = p.parseTSInterfaceDeclaration()
	case lexer.TYPE:
		declaration, err = p.parseTSTypeAliasDeclaration()
	case lexer.ENUM:
		declaration, err = p.parseTSEnumDeclaration()
	default:
		return nil, p.errorAtCurrent("expected declaration")
	}

	if err != nil {
		return nil, err
	}

	return &ast.ExportNamedDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeExportNamedDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Declaration: declaration,
		Specifiers:  []ast.ExportSpecifier{},
		Source:      nil,
		ExportKind:  exportKind,
	}, nil
}

// parseExportDefaultDeclaration parses an export default declaration.
func (p *Parser) parseExportDefaultDeclaration(start int) (*ast.ExportDefaultDeclaration, error) {
	var declaration ast.Node
	var err error

	switch p.current.Type {
	case lexer.FUNCTION:
		declaration, err = p.parseFunctionDeclaration()
	case lexer.CLASS:
		declaration, err = p.parseClassDeclaration()
	case lexer.INTERFACE:
		declaration, err = p.parseTSInterfaceDeclaration()
	default:
		// Expression
		expr, err := p.parseAssignmentExpression()
		if err != nil {
			return nil, err
		}
		declaration = expr
		p.consume(lexer.SEMICOLON)
	}

	if err != nil {
		return nil, err
	}

	return &ast.ExportDefaultDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeExportDefaultDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Declaration: declaration,
	}, nil
}

// parseExportAllDeclaration parses an export * declaration.
func (p *Parser) parseExportAllDeclaration(start int, exportKind string) (*ast.ExportAllDeclaration, error) {
	p.nextToken() // consume '*'

	var exported *ast.Identifier
	if p.consume(lexer.AS) {
		if p.current.Type != lexer.IDENT {
			return nil, p.errorAtCurrent("expected identifier")
		}
		exported = &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
		p.nextToken()
	}

	if err := p.expect(lexer.FROM); err != nil {
		return nil, err
	}

	if p.current.Type != lexer.STRING {
		return nil, p.errorAtCurrent("expected module specifier")
	}

	source := &ast.Literal{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeLiteral.String(),
			Range:    &ast.Range{p.current.Pos, p.current.End},
		},
		Value: p.current.Literal,
		Raw:   p.current.Literal,
	}
	p.nextToken()

	p.consume(lexer.SEMICOLON)

	return &ast.ExportAllDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeExportAllDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Source:     source,
		Exported:   exported,
		ExportKind: &exportKind,
	}, nil
}

// parseExportNamedDeclaration parses an export {...} declaration.
func (p *Parser) parseExportNamedDeclaration(start int, exportKind string) (*ast.ExportNamedDeclaration, error) {
	specifiers, err := p.parseExportSpecifiers()
	if err != nil {
		return nil, err
	}

	var source *ast.Literal
	if p.consume(lexer.FROM) {
		if p.current.Type != lexer.STRING {
			return nil, p.errorAtCurrent("expected module specifier")
		}
		source = &ast.Literal{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeLiteral.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Value: p.current.Literal,
			Raw:   p.current.Literal,
		}
		p.nextToken()
	}

	p.consume(lexer.SEMICOLON)

	return &ast.ExportNamedDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeExportNamedDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Declaration: nil,
		Specifiers:  specifiers,
		Source:      source,
		ExportKind:  exportKind,
	}, nil
}

// parseExportSpecifiers parses export specifiers {a, b as c}.
func (p *Parser) parseExportSpecifiers() ([]*ast.ExportSpecifier, error) {
	if err := p.expect(lexer.LBRACE); err != nil {
		return nil, err
	}

	specifiers := []*ast.ExportSpecifier{}

	for !p.match(lexer.RBRACE) && !p.isAtEnd() {
		spec, err := p.parseExportSpecifier()
		if err != nil {
			return nil, err
		}
		specifiers = append(specifiers, spec)

		if !p.consume(lexer.COMMA) {
			break
		}
	}

	if err := p.expect(lexer.RBRACE); err != nil {
		return nil, err
	}

	return specifiers, nil
}

// parseExportSpecifier parses a single export specifier.
func (p *Parser) parseExportSpecifier() (*ast.ExportSpecifier, error) {
	start := p.current.Pos

	// Check for type export (TypeScript)
	exportKind := "value"
	if p.current.Type == lexer.TYPE {
		exportKind = "type"
		p.nextToken()
	}

	if p.current.Type != lexer.IDENT {
		return nil, p.errorAtCurrent("expected identifier")
	}

	local := &ast.Identifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeIdentifier.String(),
			Range:    &ast.Range{p.current.Pos, p.current.End},
		},
		Name: p.current.Literal,
	}
	p.nextToken()

	exported := local

	// Check for 'as' clause
	if p.consume(lexer.AS) {
		if p.current.Type != lexer.IDENT {
			return nil, p.errorAtCurrent("expected identifier")
		}
		exported = &ast.Identifier{
			BaseNode: ast.BaseNode{
				NodeType: ast.NodeTypeIdentifier.String(),
				Range:    &ast.Range{p.current.Pos, p.current.End},
			},
			Name: p.current.Literal,
		}
		p.nextToken()
	}

	return &ast.ExportSpecifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeExportSpecifier.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Local:      local,
		Exported:   exported,
		ExportKind: &exportKind,
	}, nil
}
