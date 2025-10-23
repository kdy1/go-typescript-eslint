package typescriptestree

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/kdy1/go-typescript-eslint/internal/ast"
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
func Parse(_ string, _ *ParseOptions) (*Result, error) {
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

	// Parse the source code (placeholder - actual parsing not yet implemented)
	// TODO: Implement actual parsing using internal/parser
	_ = source

	// Create ParserServices with node mappings
	services := NewParserServices(prog)

	// Preserve node maps if requested
	if opts.PreserveNodeMaps != nil && !*opts.PreserveNodeMaps {
		// Don't preserve node maps - they can be cleared after rules run
		defer services.ClearNodeMappings()
	}

	// Create result with services
	result := &Result{
		AST:      nil, // TODO: Populate with actual parsed AST
		Services: services,
	}

	return result, nil
}
