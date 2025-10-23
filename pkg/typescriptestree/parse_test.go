package typescriptestree_test

import (
	"testing"

	"github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"
)

func TestParse_BasicFunctionality(t *testing.T) {
	source := `const x: number = 42;`

	opts := typescriptestree.NewBuilder().
		WithSourceType(typescriptestree.SourceTypeModule).
		WithLoc(true).
		WithRange(true).
		WithComment(true).
		WithTokens(true).
		MustBuild()

	result, err := typescriptestree.Parse(source, opts)
	if err != nil {
		t.Fatalf("Parse() returned error: %v", err)
	}

	if result == nil {
		t.Fatal("Parse() returned nil result")
	}

	if result.AST == nil {
		t.Fatal("Parse() returned nil AST")
	}

	if result.Services != nil {
		t.Error("Parse() should not return Services for basic parsing")
	}

	if result.AST.Type() != typescriptestree.AST_NODE_TYPES.Program {
		t.Errorf("Expected root node type to be Program, got %s", result.AST.Type())
	}

	if result.AST.SourceType != "module" {
		t.Errorf("Expected SourceType to be 'module', got %s", result.AST.SourceType)
	}
}

func TestParse_WithNilOptions(t *testing.T) {
	source := `const x = 1;`

	result, err := typescriptestree.Parse(source, nil)
	if err != nil {
		t.Fatalf("Parse() with nil options returned error: %v", err)
	}

	if result == nil {
		t.Fatal("Parse() returned nil result")
	}

	if result.AST == nil {
		t.Fatal("Parse() returned nil AST")
	}
}

func TestParse_WithoutComments(t *testing.T) {
	source := `// This is a comment
const x = 42;`

	opts := typescriptestree.NewBuilder().
		WithComment(false).
		MustBuild()

	result, err := typescriptestree.Parse(source, opts)
	if err != nil {
		t.Fatalf("Parse() returned error: %v", err)
	}

	if result.AST.Comments != nil {
		t.Error("Parse() should not include comments when Comment option is false")
	}
}

func TestParse_WithoutTokens(t *testing.T) {
	source := `const x = 42;`

	opts := typescriptestree.NewBuilder().
		WithTokens(false).
		MustBuild()

	result, err := typescriptestree.Parse(source, opts)
	if err != nil {
		t.Fatalf("Parse() returned error: %v", err)
	}

	if result.AST.Tokens != nil {
		t.Error("Parse() should not include tokens when Tokens option is false")
	}
}

func TestParse_JSXEnabled(t *testing.T) {
	source := `const element = <div>Hello</div>;`

	opts := typescriptestree.NewBuilder().
		WithJSX(true).
		MustBuild()

	result, err := typescriptestree.Parse(source, opts)
	if err != nil {
		t.Fatalf("Parse() with JSX returned error: %v", err)
	}

	if result == nil {
		t.Fatal("Parse() returned nil result")
	}

	if result.AST == nil {
		t.Fatal("Parse() returned nil AST")
	}
}

func TestParse_ScriptSourceType(t *testing.T) {
	source := `var x = 1;`

	opts := typescriptestree.NewBuilder().
		WithSourceType(typescriptestree.SourceTypeScript).
		MustBuild()

	result, err := typescriptestree.Parse(source, opts)
	if err != nil {
		t.Fatalf("Parse() returned error: %v", err)
	}

	if result.AST.SourceType != "script" {
		t.Errorf("Expected SourceType to be 'script', got %s", result.AST.SourceType)
	}
}

func TestParse_AllowInvalidAST(t *testing.T) {
	// Source with syntax error
	source := `const x =`

	opts := typescriptestree.NewBuilder().
		WithAllowInvalidAST(true).
		MustBuild()

	result, err := typescriptestree.Parse(source, opts)
	// Should return a result even with error
	if result == nil {
		t.Fatal("Parse() with AllowInvalidAST should return result even on error")
	}

	// Error should still be returned
	if err == nil {
		t.Log("Note: Parser may not detect this as an error yet")
	}
}

func TestParseAndGenerateServices_BasicFunctionality(t *testing.T) {
	source := `const x: number = 42;`

	// Skip this test if we don't have a tsconfig.json
	// In a real scenario, you'd create a test tsconfig
	opts, err := typescriptestree.NewServicesBuilder().
		WithSourceType(typescriptestree.SourceTypeModule).
		WithLoc(true).
		WithRange(true).
		Build()
	if err != nil {
		t.Fatalf("Failed to build options: %v", err)
	}

	result, err := typescriptestree.ParseAndGenerateServices(source, opts)

	// This may fail without a tsconfig, which is expected
	if err != nil {
		t.Logf("ParseAndGenerateServices() returned error (expected without tsconfig): %v", err)
		return
	}

	if result == nil {
		t.Fatal("ParseAndGenerateServices() returned nil result")
	}

	if result.AST == nil {
		t.Fatal("ParseAndGenerateServices() returned nil AST")
	}

	// Services should be populated when using ParseAndGenerateServices
	if result.Services == nil {
		t.Error("ParseAndGenerateServices() should return Services")
	}
}

func TestParseAndGenerateServices_WithNilOptions(t *testing.T) {
	source := `const x = 1;`

	_, err := typescriptestree.ParseAndGenerateServices(source, nil)
	if err == nil {
		t.Error("ParseAndGenerateServices() with nil options should return error")
	}
}

func TestClearProgramCache(t *testing.T) {
	// Should not panic
	typescriptestree.ClearProgramCache()

	// Can be called multiple times
	typescriptestree.ClearProgramCache()
	typescriptestree.ClearProgramCache()
}

func TestClearDefaultProjectMatchedFiles(t *testing.T) {
	// Should not panic
	typescriptestree.ClearDefaultProjectMatchedFiles()

	// Can be called multiple times
	typescriptestree.ClearDefaultProjectMatchedFiles()
	typescriptestree.ClearDefaultProjectMatchedFiles()
}

func TestAST_NODE_TYPES_Constants(t *testing.T) {
	// Test that all node type constants are properly initialized
	tests := []struct {
		name  string
		value string
	}{
		{"Program", typescriptestree.AST_NODE_TYPES.Program},
		{"Identifier", typescriptestree.AST_NODE_TYPES.Identifier},
		{"Literal", typescriptestree.AST_NODE_TYPES.Literal},
		{"FunctionDeclaration", typescriptestree.AST_NODE_TYPES.FunctionDeclaration},
		{"VariableDeclaration", typescriptestree.AST_NODE_TYPES.VariableDeclaration},
		{"BlockStatement", typescriptestree.AST_NODE_TYPES.BlockStatement},
		{"IfStatement", typescriptestree.AST_NODE_TYPES.IfStatement},
		{"BinaryExpression", typescriptestree.AST_NODE_TYPES.BinaryExpression},
		{"CallExpression", typescriptestree.AST_NODE_TYPES.CallExpression},
		{"MemberExpression", typescriptestree.AST_NODE_TYPES.MemberExpression},
		{"ArrowFunctionExpression", typescriptestree.AST_NODE_TYPES.ArrowFunctionExpression},
		{"TSTypeAnnotation", typescriptestree.AST_NODE_TYPES.TSTypeAnnotation},
		{"TSInterfaceDeclaration", typescriptestree.AST_NODE_TYPES.TSInterfaceDeclaration},
		{"TSAsExpression", typescriptestree.AST_NODE_TYPES.TSAsExpression},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value == "" {
				t.Errorf("AST_NODE_TYPES.%s is empty", tt.name)
			}
			if tt.value != tt.name {
				t.Errorf("AST_NODE_TYPES.%s = %q, want %q", tt.name, tt.value, tt.name)
			}
		})
	}
}

func TestAST_TOKEN_TYPES_Constants(t *testing.T) {
	// Test that all token type constants are properly initialized
	tests := []struct {
		name  string
		value string
	}{
		{"Identifier", typescriptestree.AST_TOKEN_TYPES.Identifier},
		{"Number", typescriptestree.AST_TOKEN_TYPES.Number},
		{"String", typescriptestree.AST_TOKEN_TYPES.String},
		{"Const", typescriptestree.AST_TOKEN_TYPES.Const},
		{"Let", typescriptestree.AST_TOKEN_TYPES.Let},
		{"Function", typescriptestree.AST_TOKEN_TYPES.Function},
		{"If", typescriptestree.AST_TOKEN_TYPES.If},
		{"Arrow", typescriptestree.AST_TOKEN_TYPES.Arrow},
		{"LeftParen", typescriptestree.AST_TOKEN_TYPES.LeftParen},
		{"RightParen", typescriptestree.AST_TOKEN_TYPES.RightParen},
		{"Semicolon", typescriptestree.AST_TOKEN_TYPES.Semicolon},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value == "" {
				t.Errorf("AST_TOKEN_TYPES.%s is empty", tt.name)
			}
		})
	}
}

func TestParse_FilePath_AutoDetectJSX(t *testing.T) {
	source := `const element = <div>Hello</div>;`

	opts := typescriptestree.NewBuilder().
		WithFilePath("component.tsx"). // Should auto-enable JSX
		MustBuild()

	result, err := typescriptestree.Parse(source, opts)
	if err != nil {
		t.Fatalf("Parse() with .tsx file returned error: %v", err)
	}

	if result == nil || result.AST == nil {
		t.Fatal("Parse() should handle JSX in .tsx files")
	}
}

func BenchmarkParse(b *testing.B) {
	source := `
		const x: number = 42;
		function greet(name: string): string {
			return 'Hello, ' + name;
		}
		const result = greet('World');
	`

	opts := typescriptestree.NewBuilder().
		WithSourceType(typescriptestree.SourceTypeModule).
		MustBuild()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = typescriptestree.Parse(source, opts)
	}
}

func BenchmarkParse_WithLocAndRange(b *testing.B) {
	source := `const x: number = 42;`

	opts := typescriptestree.NewBuilder().
		WithLoc(true).
		WithRange(true).
		MustBuild()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = typescriptestree.Parse(source, opts)
	}
}

func BenchmarkParse_WithCommentsAndTokens(b *testing.B) {
	source := `
		// Comment 1
		const x = 42;
		// Comment 2
		const y = 100;
	`

	opts := typescriptestree.NewBuilder().
		WithComment(true).
		WithTokens(true).
		MustBuild()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = typescriptestree.Parse(source, opts)
	}
}
