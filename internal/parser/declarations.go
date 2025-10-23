package parser

import (
	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/lexer"
)

// parseImportDeclaration parses an import declaration.
func (p *Parser) parseImportDeclaration() (*ast.ImportDeclaration, error) {
	start := p.current.Pos
	p.nextToken() // consume 'import'

	importKind := p.parseImportKind()

	// Check for side-effect import: import 'module'
	if p.current.Type == lexer.STRING {
		return p.parseSideEffectImport(start, importKind)
	}

	// Parse import specifiers
	specifiers, err := p.parseImportSpecifierList()
	if err != nil {
		return nil, err
	}

	// Parse 'from' clause
	if err := p.expect(lexer.FROM); err != nil {
		return nil, err
	}

	source, err := p.parseModuleSpecifier()
	if err != nil {
		return nil, err
	}

	attributes := p.parseImportAttributesIfPresent()
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

// parseImportKind checks for and parses the 'type' keyword in TypeScript imports.
func (p *Parser) parseImportKind() string {
	if p.current.Type == lexer.TYPE {
		p.nextToken()
		return "type"
	}
	return "value"
}

// parseSideEffectImport parses a side-effect import: import 'module'
func (p *Parser) parseSideEffectImport(start int, importKind string) (*ast.ImportDeclaration, error) {
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
		Specifiers: []interface{}{},
		Source:     source,
		ImportKind: &importKind,
	}, nil
}

// parseImportSpecifierList parses the list of import specifiers.
func (p *Parser) parseImportSpecifierList() ([]interface{}, error) {
	specifiers := []interface{}{}

	// Parse default import
	if p.current.Type == lexer.IDENT {
		defaultSpec := p.parseDefaultImportSpecifier()
		specifiers = append(specifiers, defaultSpec)
		p.nextToken()

		// Check for additional imports after default
		if p.consume(lexer.COMMA) {
			additionalSpecs, err := p.parseAdditionalImportSpecifiers()
			if err != nil {
				return nil, err
			}
			specifiers = append(specifiers, additionalSpecs...)
		}
	} else {
		// Parse namespace or named imports
		specs, err := p.parseAdditionalImportSpecifiers()
		if err != nil {
			return nil, err
		}
		specifiers = append(specifiers, specs...)
	}

	return specifiers, nil
}

// parseDefaultImportSpecifier creates a default import specifier from current token.
func (p *Parser) parseDefaultImportSpecifier() *ast.ImportDefaultSpecifier {
	return &ast.ImportDefaultSpecifier{
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
	}
}

// parseAdditionalImportSpecifiers parses namespace or named imports.
func (p *Parser) parseAdditionalImportSpecifiers() ([]interface{}, error) {
	switch p.current.Type {
	case lexer.MUL:
		spec, err := p.parseImportNamespaceSpecifier()
		if err != nil {
			return nil, err
		}
		return []interface{}{spec}, nil
	case lexer.LBRACE:
		specs, err := p.parseImportSpecifiers()
		if err != nil {
			return nil, err
		}
		result := make([]interface{}, len(specs))
		for i, spec := range specs {
			result[i] = spec
		}
		return result, nil
	default:
		return nil, p.errorAtCurrent("expected import specifier")
	}
}

// parseModuleSpecifier parses the module specifier string.
func (p *Parser) parseModuleSpecifier() (*ast.Literal, error) {
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
	return source, nil
}

// parseImportAttributesIfPresent parses import attributes if 'with' clause is present.
func (p *Parser) parseImportAttributesIfPresent() []ast.ImportAttribute {
	var attributes []ast.ImportAttribute
	if p.current.Type == lexer.IDENT && p.current.Literal == "with" {
		p.nextToken()
		attributesPtr, err := p.parseImportAttributes()
		if err == nil {
			for _, attr := range attributesPtr {
				attributes = append(attributes, *attr)
			}
		}
	}
	return attributes
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
	return parseListInBraces(p, p.parseImportSpecifier)
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
	return parseListInBraces(p, p.parseImportAttribute)
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
		return p.parseExportDefaultDeclaration(start)
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
	return p.parseExportWithDeclaration(start, exportKind)
}

// parseExportWithDeclaration parses export with a declaration.
func (p *Parser) parseExportWithDeclaration(start int, exportKind string) (ast.Statement, error) {
	declaration, err := p.parseExportableDeclaration()
	if err != nil {
		return nil, err
	}

	// Convert ast.Node to ast.Declaration
	var decl ast.Declaration
	if declaration != nil {
		decl, _ = declaration.(ast.Declaration)
	}

	return &ast.ExportNamedDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeExportNamedDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Declaration: decl,
		Specifiers:  []ast.ExportSpecifier{},
		Source:      nil,
		ExportKind:  &exportKind,
	}, nil
}

// parseExportableDeclaration parses declarations that can be exported.
func (p *Parser) parseExportableDeclaration() (ast.Node, error) {
	switch p.current.Type {
	case lexer.VAR, lexer.LET, lexer.CONST:
		return p.parseVariableStatement()
	case lexer.FUNCTION:
		return p.parseFunctionDeclaration()
	case lexer.CLASS:
		return p.parseClassDeclaration()
	case lexer.INTERFACE:
		return p.parseTSInterfaceDeclaration()
	case lexer.TYPE:
		return p.parseTSTypeAliasDeclaration()
	case lexer.ENUM:
		return p.parseTSEnumDeclaration()
	default:
		return nil, p.errorAtCurrent("expected declaration")
	}
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

	// Convert []*ast.ExportSpecifier to []ast.ExportSpecifier
	var specs []ast.ExportSpecifier
	for _, spec := range specifiers {
		specs = append(specs, *spec)
	}

	return &ast.ExportNamedDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeExportNamedDeclaration.String(),
			Range:    &ast.Range{start, p.current.Pos},
		},
		Declaration: nil,
		Specifiers:  specs,
		Source:      source,
		ExportKind:  &exportKind,
	}, nil
}

// parseExportSpecifiers parses export specifiers {a, b as c}.
func (p *Parser) parseExportSpecifiers() ([]*ast.ExportSpecifier, error) {
	return parseListInBraces(p, p.parseExportSpecifier)
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
