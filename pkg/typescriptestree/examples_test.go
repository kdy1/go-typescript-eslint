package typescriptestree_test

import (
	"fmt"
	"log"

	"github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"
)

// Example demonstrates basic parsing of TypeScript code.
func Example() {
	source := `const greeting: string = "Hello, World!";`

	opts := typescriptestree.NewBuilder().
		WithSourceType(typescriptestree.SourceTypeModule).
		WithLoc(true).
		WithRange(true).
		MustBuild()

	result, err := typescriptestree.Parse(source, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Parsed AST node type: %s\n", result.AST.Type())
	fmt.Printf("Source type: %s\n", result.AST.SourceType)
	// Output:
	// Parsed AST node type: Program
	// Source type: module
}

// Example_parse demonstrates the Parse function with various options.
func Example_parse() {
	source := `
		// This is a TypeScript file
		interface User {
			name: string;
			age: number;
		}

		const user: User = {
			name: "John",
			age: 30
		};
	`

	opts := typescriptestree.NewBuilder().
		WithSourceType(typescriptestree.SourceTypeModule).
		WithComment(true).
		WithTokens(true).
		WithLoc(true).
		WithRange(true).
		MustBuild()

	result, err := typescriptestree.Parse(source, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Root node: %s\n", result.AST.Type())
	fmt.Printf("Has comments: %t\n", len(result.AST.Comments) > 0)
	fmt.Printf("Has tokens: %t\n", len(result.AST.Tokens) > 0)
	// Output:
	// Root node: Program
	// Has comments: true
	// Has tokens: true
}

// Example_jsxParsing demonstrates parsing JSX/TSX code.
func Example_jsxParsing() {
	source := `
		const Component = () => {
			return <div className="container">Hello, React!</div>;
		};
	`

	opts := typescriptestree.NewBuilder().
		WithJSX(true).
		WithSourceType(typescriptestree.SourceTypeModule).
		MustBuild()

	result, err := typescriptestree.Parse(source, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Parsed JSX successfully: %t\n", result.AST != nil)
	// Output:
	// Parsed JSX successfully: true
}

// Example_parseAndGenerateServices demonstrates type-aware parsing.
func Example_parseAndGenerateServices() {
	source := `
		function add(a: number, b: number): number {
			return a + b;
		}
	`

	opts, err := typescriptestree.NewServicesBuilder().
		WithSourceType(typescriptestree.SourceTypeModule).
		WithLoc(true).
		WithRange(true).
		Build()
	if err != nil {
		log.Fatal(err)
	}

	result, err := typescriptestree.ParseAndGenerateServices(source, opts)
	if err != nil {
		// May fail without a tsconfig.json file
		fmt.Println("Note: ParseAndGenerateServices requires a TypeScript project configuration")
		return
	}

	fmt.Printf("Has AST: %t\n", result.AST != nil)
	fmt.Printf("Has Services: %t\n", result.Services != nil)
}

// Example_nodeTypes demonstrates using AST_NODE_TYPES constants.
func Example_nodeTypes() {
	// Access node type constants
	fmt.Printf("Program node type: %s\n", typescriptestree.AST_NODE_TYPES.Program)
	fmt.Printf("Function declaration: %s\n", typescriptestree.AST_NODE_TYPES.FunctionDeclaration)
	fmt.Printf("Interface declaration: %s\n", typescriptestree.AST_NODE_TYPES.TSInterfaceDeclaration)
	fmt.Printf("Arrow function: %s\n", typescriptestree.AST_NODE_TYPES.ArrowFunctionExpression)

	// Output:
	// Program node type: Program
	// Function declaration: FunctionDeclaration
	// Interface declaration: TSInterfaceDeclaration
	// Arrow function: ArrowFunctionExpression
}

// Example_tokenTypes demonstrates using AST_TOKEN_TYPES constants.
func Example_tokenTypes() {
	// Access token type constants
	fmt.Printf("Const keyword: %s\n", typescriptestree.AST_TOKEN_TYPES.Const)
	fmt.Printf("Arrow token: %s\n", typescriptestree.AST_TOKEN_TYPES.Arrow)
	fmt.Printf("Left brace: %s\n", typescriptestree.AST_TOKEN_TYPES.LeftBrace)

	// Output:
	// Const keyword: CONST
	// Arrow token: ARROW
	// Left brace: LBRACE
}

// Example_builder demonstrates using the options builder pattern.
func Example_builder() {
	opts := typescriptestree.NewBuilder().
		WithSourceType(typescriptestree.SourceTypeModule).
		WithJSX(true).
		WithComment(true).
		WithTokens(true).
		WithLoc(true).
		WithRange(true).
		WithFilePath("component.tsx").
		MustBuild()

	source := `const App = () => <div>Hello</div>;`

	result, err := typescriptestree.Parse(source, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully parsed: %t\n", result.AST != nil)
	// Output:
	// Successfully parsed: true
}

// Example_clearCache demonstrates clearing the program cache.
func Example_clearCache() {
	// Parse some TypeScript code (which may cache programs)
	source := `const x: number = 42;`
	opts := typescriptestree.NewBuilder().MustBuild()
	_, _ = typescriptestree.Parse(source, opts)

	// Clear the cache (useful for testing)
	typescriptestree.ClearProgramCache()

	fmt.Println("Cache cleared successfully")
	// Output:
	// Cache cleared successfully
}

// Example_errorHandlingBasic demonstrates error handling with AllowInvalidAST.
func Example_errorHandlingBasic() {
	// Source with syntax error
	source := `const x =` // incomplete

	// With AllowInvalidAST, we get a partial AST even on error
	opts := typescriptestree.NewBuilder().
		WithAllowInvalidAST(true).
		MustBuild()

	result, err := typescriptestree.Parse(source, opts)

	fmt.Printf("Got result: %t\n", result != nil)
	fmt.Printf("Got error: %t\n", err != nil)

	// Without AllowInvalidAST, parsing fails
	opts2 := typescriptestree.NewBuilder().
		WithAllowInvalidAST(false).
		MustBuild()

	result2, err2 := typescriptestree.Parse(source, opts2)

	fmt.Printf("Strict mode result: %t\n", result2 == nil)
	fmt.Printf("Strict mode error: %t\n", err2 != nil)
}

// Example_scriptVsModule demonstrates the difference between script and module source types.
func Example_scriptVsModule() {
	source := `var x = 1;`

	// Parse as script
	scriptOpts := typescriptestree.NewBuilder().
		WithSourceType(typescriptestree.SourceTypeScript).
		MustBuild()

	scriptResult, _ := typescriptestree.Parse(source, scriptOpts)
	fmt.Printf("Script source type: %s\n", scriptResult.AST.SourceType)

	// Parse as module
	moduleOpts := typescriptestree.NewBuilder().
		WithSourceType(typescriptestree.SourceTypeModule).
		MustBuild()

	moduleResult, _ := typescriptestree.Parse(source, moduleOpts)
	fmt.Printf("Module source type: %s\n", moduleResult.AST.SourceType)

	// Output:
	// Script source type: script
	// Module source type: module
}

// Example_minimalParsing demonstrates parsing with minimal options for performance.
func Example_minimalParsing() {
	source := `const x: number = 42;`

	// Minimal parsing - no location, range, comments, or tokens
	opts := typescriptestree.NewBuilder().
		WithComment(false).
		WithTokens(false).
		WithLoc(false).
		WithRange(false).
		MustBuild()

	result, err := typescriptestree.Parse(source, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("AST only, minimal overhead: %t\n", result.AST != nil)
	fmt.Printf("No comments: %t\n", result.AST.Comments == nil)
	fmt.Printf("No tokens: %t\n", result.AST.Tokens == nil)
	// Output:
	// AST only, minimal overhead: true
	// No comments: true
	// No tokens: true
}
