package program

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kdy1/go-typescript-eslint/internal/ast"
)

func TestCreateProgram(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a tsconfig.json
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")
	tsconfigContent := `{
		"compilerOptions": {
			"target": "ES2020",
			"strict": true
		}
	}`

	if err := os.WriteFile(tsconfigPath, []byte(tsconfigContent), 0600); err != nil {
		t.Fatalf("Failed to write tsconfig: %v", err)
	}

	// Create program
	opts := &ProgramOptions{
		TSConfigPath: tsconfigPath,
		RootDir:      tmpDir,
	}

	program, err := CreateProgram(opts)
	if err != nil {
		t.Fatalf("Failed to create program: %v", err)
	}

	if program == nil {
		t.Fatal("Expected non-nil program")
	}

	if program.Config == nil {
		t.Fatal("Expected non-nil config")
	}

	if program.Config.CompilerOptions.Target != "ES2020" {
		t.Errorf("Expected target ES2020, got %s", program.Config.CompilerOptions.Target)
	}
}

func TestProgramGetSourceFile(t *testing.T) {
	program := &Program{
		SourceFiles: make(map[string]*ast.Program),
	}

	// Initially empty
	_, ok := program.GetSourceFile("test.ts")
	if ok {
		t.Error("Expected no source file initially")
	}

	// Add a source file
	testAST := &ast.Program{}
	program.AddSourceFile("test.ts", testAST)

	// Should now retrieve it
	retrieved, ok := program.GetSourceFile("test.ts")
	if !ok {
		t.Error("Expected to find source file")
	}

	if retrieved != testAST {
		t.Error("Retrieved AST does not match added AST")
	}
}

func TestProgramGetCompilerOptions(t *testing.T) {
	program := &Program{
		Config: &TSConfig{
			CompilerOptions: CompilerOptions{
				Target: "ES2020",
				Strict: boolPtr(true),
			},
		},
	}

	opts := program.GetCompilerOptions()
	if opts == nil {
		t.Fatal("Expected non-nil compiler options")
	}

	if opts.Target != "ES2020" {
		t.Errorf("Expected target ES2020, got %s", opts.Target)
	}

	if opts.Strict == nil || !*opts.Strict {
		t.Error("Expected strict to be true")
	}
}

func TestFindConfigForFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Create nested directory structure
	srcDir := filepath.Join(tmpDir, "src")
	if err := os.MkdirAll(srcDir, 0750); err != nil {
		t.Fatalf("Failed to create src dir: %v", err)
	}

	// Create tsconfig in root
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")
	if err := os.WriteFile(tsconfigPath, []byte("{}"), 0600); err != nil {
		t.Fatalf("Failed to write tsconfig: %v", err)
	}

	// Create a file in src
	filePath := filepath.Join(srcDir, "test.ts")
	if err := os.WriteFile(filePath, []byte("const x = 1;"), 0600); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Find config for file
	foundConfig, err := FindConfigForFile(filePath)
	if err != nil {
		t.Fatalf("Failed to find config: %v", err)
	}

	if foundConfig != tsconfigPath {
		t.Errorf("Expected config %s, got %s", tsconfigPath, foundConfig)
	}
}

// Helper function
func boolPtr(b bool) *bool {
	return &b
}
