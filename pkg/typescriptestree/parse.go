package typescriptestree

import (
	"errors"

	"github.com/kdy1/go-typescript-eslint/internal/ast"
)

// ErrNotImplemented is returned when a feature is not yet implemented.
var ErrNotImplemented = errors.New("feature not yet implemented")

// Result represents the result of parsing TypeScript source code.
type Result struct {
	// AST is the parsed Abstract Syntax Tree.
	AST *ast.Program

	// Services provides TypeScript language services for type-aware operations.
	// This is only populated when using ParseAndGenerateServices.
	Services *Services
}

// Services provides TypeScript language services for type-aware linting and analysis.
type Services struct {
	// Program is the TypeScript program instance (placeholder for now).
	Program interface{}

	// ESTreeNodeToTSNodeMap maps ESTree nodes to TypeScript AST nodes.
	ESTreeNodeToTSNodeMap map[ast.Node]interface{}

	// TSNodeToESTreeNodeMap maps TypeScript AST nodes to ESTree nodes.
	TSNodeToESTreeNodeMap map[interface{}]ast.Node
}

// Parse parses TypeScript source code into an ESTree-compatible AST.
// This is the main entry point for parsing TypeScript code.
//
// Example:
//
//	opts := typescriptestree.NewBuilder().
//		WithSourceType(typescriptestree.SourceTypeModule).
//		WithLoc(true).
//		WithRange(true).
//		MustBuild()
//	result, err := typescriptestree.Parse("const x: number = 42;", opts)
//	if err != nil {
//		// handle error
//	}
//	// use result.AST
func Parse(_ string, _ *ParseOptions) (*Result, error) {
	//nolint:godox // TODO is intentional for unimplemented feature
	// TODO: Implement full TypeScript parsing
	// This will use the internal lexer and parser packages
	return nil, ErrNotImplemented
}

// ParseAndGenerateServices parses TypeScript source code and generates
// TypeScript program services for type-aware linting and analysis.
//
// This function is required for type-aware ESLint rules that need access to
// TypeScript's type checker and program information.
//
// Example:
//
//	opts := typescriptestree.NewServicesBuilder().
//		WithProject("./tsconfig.json").
//		WithTSConfigRootDir(".").
//		Build()
//	result, err := typescriptestree.ParseAndGenerateServices("const x: number = 42;", opts)
//	if err != nil {
//		// handle error
//	}
//	// use result.AST and result.Services
func ParseAndGenerateServices(_ string, _ *ParseAndGenerateServicesOptions) (*Result, error) {
	//nolint:godox // TODO is intentional for unimplemented feature
	// TODO: Implement with TypeScript services support
	return nil, ErrNotImplemented
}
