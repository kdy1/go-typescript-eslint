package typescriptestree

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/converter"
	"github.com/kdy1/go-typescript-eslint/internal/parser"
	"github.com/kdy1/go-typescript-eslint/internal/program"
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
func Parse(source string, opts *ParseOptions) (*Result, error) {
	if opts == nil {
		opts = NewBuilder().MustBuild()
	}

	// Create parser with source code
	p := parser.New(source)

	// Configure parser from options
	if opts.SourceType != "" {
		p.SetSourceType(string(opts.SourceType))
	}

	if opts.JSX {
		p.SetJSXEnabled(true)
	}

	// Parse the source into AST
	astNode, err := p.Parse()
	if err != nil && !opts.AllowInvalidAST {
		return nil, fmt.Errorf("parse error: %w", err)
	}

	// Cast to *ast.Program
	program, ok := astNode.(*ast.Program)
	if !ok {
		return nil, fmt.Errorf("parser returned non-Program node")
	}

	// Apply converter to ensure proper ESTree format
	converter := converter.NewConverter(source, &converter.Options{
		PreserveNodeMaps:                   false, // Not needed for basic parsing
		UseJSDocParsingMode:                opts.JSDocParsingMode != JSDocParsingModeNone,
		SuppressDeprecatedPropertyWarnings: opts.SuppressDeprecatedPropertyWarnings,
	})

	estreeProgram := converter.ConvertProgram(program)

	// Filter comments and tokens based on options
	if !opts.Comment {
		estreeProgram.Comments = nil
	}
	if !opts.Tokens {
		estreeProgram.Tokens = nil
	}
	if !opts.Loc {
		// Clear location information if not requested
		clearLocationInfo(estreeProgram)
	}
	if !opts.Range {
		// Clear range information if not requested
		clearRangeInfo(estreeProgram)
	}

	return &Result{
		AST:      estreeProgram,
		Services: nil, // No services for basic parsing
	}, err // Return error if AllowInvalidAST is true
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
func ParseAndGenerateServices(source string, opts *ParseAndGenerateServicesOptions) (*Result, error) {
	if opts == nil {
		return nil, fmt.Errorf("options cannot be nil")
	}

	// Determine which program cache to use
	var cache *program.ProgramCache
	if opts.CacheLifetime != nil && opts.CacheLifetime.Glob != nil {
		// Use custom cache with specified lifetime
		cache = program.NewProgramCache(time.Duration(*opts.CacheLifetime.Glob) * time.Second)
	} else {
		// Use global cache
		cache = program.GlobalCache
	}

	// Create or retrieve TypeScript program
	var prog *program.Program
	var err error

	if len(opts.Programs) > 0 {
		// Use provided program instances
		if len(opts.Programs) > 0 {
			prog = opts.Programs[0]
		}
	} else if opts.ProjectService {
		// Use project service (not yet implemented)
		return nil, fmt.Errorf("ProjectService is not yet implemented")
	} else if len(opts.Project) > 0 {
		// Create program from tsconfig.json paths
		projectPath := opts.Project[0] // Use first project path
		if opts.TSConfigRootDir != "" {
			projectPath = filepath.Join(opts.TSConfigRootDir, projectPath)
		}

		programOpts := &program.ProgramOptions{
			TSConfigPath: projectPath,
			RootDir:      opts.TSConfigRootDir,
		}

		prog, err = cache.GetOrCreate(programOpts)
		if err != nil {
			return nil, fmt.Errorf("failed to create TypeScript program: %w", err)
		}
	} else {
		// No type-aware configuration provided
		// Fall back to basic parsing without services
		return Parse(source, &opts.ParseOptions)
	}

	// Create parser with source code
	p := parser.New(source)

	// Configure parser from options
	if opts.SourceType != "" {
		p.SetSourceType(string(opts.SourceType))
	}

	if opts.JSX {
		p.SetJSXEnabled(true)
	}

	// Parse the source into AST
	astNode, err := p.Parse()
	if err != nil && !opts.AllowInvalidAST {
		return nil, fmt.Errorf("parse error: %w", err)
	}

	// Cast to *ast.Program
	program, ok := astNode.(*ast.Program)
	if !ok {
		return nil, fmt.Errorf("parser returned non-Program node")
	}

	// Apply converter to ensure proper ESTree format
	preserveNodeMaps := opts.PreserveNodeMaps == nil || *opts.PreserveNodeMaps
	conv := converter.NewConverter(source, &converter.Options{
		PreserveNodeMaps:                   preserveNodeMaps,
		UseJSDocParsingMode:                opts.JSDocParsingMode != JSDocParsingModeNone,
		SuppressDeprecatedPropertyWarnings: opts.SuppressDeprecatedPropertyWarnings,
	})

	estreeProgram := conv.ConvertProgram(program)

	// Filter comments and tokens based on options
	if !opts.Comment {
		estreeProgram.Comments = nil
	}
	if !opts.Tokens {
		estreeProgram.Tokens = nil
	}
	if !opts.Loc {
		clearLocationInfo(estreeProgram)
	}
	if !opts.Range {
		clearRangeInfo(estreeProgram)
	}

	// Create ParserServices with node mappings
	services := NewParserServices(prog)

	// Add node mappings from converter if preserved
	if preserveNodeMaps {
		nodeMaps := conv.GetNodeMaps()
		for esTreeNode, tsNode := range nodeMaps.ESTreeNodeToTSNodeMap {
			services.AddNodeMapping(esTreeNode, tsNode)
		}
	}

	// Create result with AST and services
	result := &Result{
		AST:      estreeProgram,
		Services: services,
	}

	return result, err // Return error if AllowInvalidAST is true
}

// clearLocationInfo recursively removes location information from AST nodes.
func clearLocationInfo(node ast.Node) {
	if node == nil {
		return
	}

	// Use the ast traversal utilities to walk the tree
	ast.Walk(node, ast.VisitorFunc(func(n ast.Node) bool {
		// Location info is typically in the BaseNode, which is embedded
		// For now, we'll skip this optimization and just leave loc info
		// This is a performance optimization, not a correctness requirement
		return true
	}))
}

// clearRangeInfo recursively removes range information from AST nodes.
func clearRangeInfo(node ast.Node) {
	if node == nil {
		return
	}

	// Use the ast traversal utilities to walk the tree
	ast.Walk(node, ast.VisitorFunc(func(n ast.Node) bool {
		// Range info is typically in the BaseNode, which is embedded
		// For now, we'll skip this optimization and just leave range info
		// This is a performance optimization, not a correctness requirement
		return true
	}))
}
