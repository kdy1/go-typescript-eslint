package program

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/kdy1/go-typescript-eslint/internal/ast"
)

// Program represents a TypeScript program with type checking capabilities.
// This is a placeholder for future integration with TypeScript's type checker.
type Program struct {
	// Config is the resolved TypeScript configuration
	Config *TSConfig

	// RootNames contains the list of root file names in the program
	RootNames []string

	// SourceFiles maps file paths to their parsed ASTs
	SourceFiles map[string]*ast.Program

	// CreatedAt tracks when this program was created (for cache invalidation)
	CreatedAt time.Time

	// mu protects concurrent access to the program
	mu sync.RWMutex
}

// ProgramOptions configures program creation behavior.
type ProgramOptions struct {
	// TSConfigPath is the path to tsconfig.json
	TSConfigPath string

	// RootDir is the root directory for resolving relative paths
	RootDir string

	// SourceFiles is a list of source files to include in the program
	SourceFiles []string

	// AllowJS enables JavaScript file parsing
	AllowJS bool
}

// CreateProgram creates a TypeScript program from a tsconfig.json file.
// This function handles tsconfig resolution, inheritance, and program initialization.
func CreateProgram(opts *ProgramOptions) (*Program, error) {
	if opts == nil {
		return nil, fmt.Errorf("program options cannot be nil")
	}

	// Resolve tsconfig.json path
	tsconfigPath := opts.TSConfigPath
	if tsconfigPath == "" {
		// Default to tsconfig.json in root directory
		if opts.RootDir != "" {
			tsconfigPath = filepath.Join(opts.RootDir, "tsconfig.json")
		} else {
			tsconfigPath = "tsconfig.json"
		}
	}

	// Parse and resolve tsconfig with inheritance
	config, err := ResolveTSConfig(tsconfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve tsconfig: %w", err)
	}

	// Determine root file names
	rootNames := opts.SourceFiles
	if len(rootNames) == 0 {
		// Use files from tsconfig
		rootNames = config.Files
	}

	// Create program instance
	program := &Program{
		Config:      config,
		RootNames:   rootNames,
		SourceFiles: make(map[string]*ast.Program),
		CreatedAt:   time.Now(),
	}

	return program, nil
}

// GetSourceFile retrieves the AST for a specific file in the program.
func (p *Program) GetSourceFile(filePath string) (*ast.Program, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	ast, ok := p.SourceFiles[filePath]
	return ast, ok
}

// AddSourceFile adds a parsed AST to the program.
func (p *Program) AddSourceFile(filePath string, ast *ast.Program) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.SourceFiles[filePath] = ast
}

// GetCompilerOptions returns the TypeScript compiler options for this program.
func (p *Program) GetCompilerOptions() *CompilerOptions {
	return &p.Config.CompilerOptions
}

// FindConfigForFile finds the appropriate tsconfig.json for a given file.
// It searches up the directory tree from the file's location.
func FindConfigForFile(filePath string) (string, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to resolve absolute path: %w", err)
	}

	dir := filepath.Dir(absPath)

	// Search up the directory tree
	for {
		tsconfigPath := filepath.Join(dir, "tsconfig.json")
		// Check if file exists
		if _, err := os.Stat(tsconfigPath); err == nil {
			return tsconfigPath, nil
		}

		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached the root
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("no tsconfig.json found for file %s", filePath)
}
