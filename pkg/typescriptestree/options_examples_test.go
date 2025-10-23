package typescriptestree_test

import (
	"fmt"

	"github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"
)

// Example_basicParseOptions demonstrates basic usage of ParseOptions.
func Example_basicParseOptions() {
	// Create options with defaults
	opts := typescriptestree.NewParseOptions()
	opts.SourceType = typescriptestree.SourceTypeModule
	opts.Loc = true
	opts.Range = true

	// Validate the options
	if err := opts.Validate(); err != nil {
		fmt.Printf("Invalid options: %v\n", err)
		return
	}

	fmt.Println("Options configured successfully")
	// Output: Options configured successfully
}

// Example_builderPattern demonstrates using the builder pattern for options.
func Example_builderPattern() {
	// Use the builder for a fluent API
	opts, err := typescriptestree.NewBuilder().
		WithSourceType(typescriptestree.SourceTypeModule).
		WithLoc(true).
		WithRange(true).
		WithComment(true).
		WithTokens(true).
		Build()

	if err != nil {
		fmt.Printf("Error building options: %v\n", err)
		return
	}

	fmt.Printf("Source type: %s\n", opts.SourceType)
	fmt.Printf("Location tracking: %v\n", opts.Loc)
	fmt.Printf("Range tracking: %v\n", opts.Range)
	fmt.Printf("Comments: %v\n", opts.Comment)
	// Output:
	// Source type: module
	// Location tracking: true
	// Range tracking: true
	// Comments: true
}

// Example_filePathInference demonstrates automatic JSX inference from file path.
func Example_filePathInference() {
	// JSX is automatically enabled for .tsx files
	opts := typescriptestree.NewBuilder().
		WithFilePath("components/MyComponent.tsx").
		MustBuild()

	fmt.Printf("JSX enabled: %v\n", opts.JSX)
	fmt.Printf("File path: %s\n", opts.FilePath)

	// For .ts files, JSX is not enabled
	opts2 := typescriptestree.NewBuilder().
		WithFilePath("utils/helper.ts").
		MustBuild()

	fmt.Printf("JSX enabled for .ts: %v\n", opts2.JSX)
	// Output:
	// JSX enabled: true
	// File path: components/MyComponent.tsx
	// JSX enabled for .ts: false
}

// Example_jsDocParsing demonstrates JSDoc parsing mode configuration.
func Example_jsDocParsing() {
	// Parse all JSDoc comments (default)
	opts1, _ := typescriptestree.NewBuilder().
		WithJSDocParsingMode(typescriptestree.JSDocParsingModeAll).
		Build()
	fmt.Printf("JSDoc mode (all): %s\n", opts1.JSDocParsingMode)

	// Skip JSDoc parsing for performance
	opts2, _ := typescriptestree.NewBuilder().
		WithJSDocParsingMode(typescriptestree.JSDocParsingModeNone).
		Build()
	fmt.Printf("JSDoc mode (none): %s\n", opts2.JSDocParsingMode)

	// Only parse type information
	opts3, _ := typescriptestree.NewBuilder().
		WithJSDocParsingMode(typescriptestree.JSDocParsingModeTypeInfo).
		Build()
	fmt.Printf("JSDoc mode (type-info): %s\n", opts3.JSDocParsingMode)
	// Output:
	// JSDoc mode (all): all
	// JSDoc mode (none): none
	// JSDoc mode (type-info): type-info
}

// Example_servicesOptions demonstrates type-aware parsing configuration.
func Example_servicesOptions() {
	// Configure for type-aware parsing
	opts, err := typescriptestree.NewServicesBuilder().
		WithProject("./tsconfig.json").
		WithTSConfigRootDir("/path/to/project").
		Build()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Project paths: %v\n", opts.Project)
	fmt.Printf("TSConfig root: %s\n", opts.TSConfigRootDir)
	// Output:
	// Project paths: [./tsconfig.json]
	// TSConfig root: /path/to/project
}

// Example_multipleProjects demonstrates configuring multiple TypeScript projects.
func Example_multipleProjects() {
	opts, _ := typescriptestree.NewServicesBuilder().
		WithProject(
			"./packages/*/tsconfig.json",
			"./apps/*/tsconfig.json",
		).
		WithProjectFolderIgnoreList(
			"**/node_modules/**",
			"**/dist/**",
			"**/build/**",
		).
		Build()

	fmt.Printf("Projects: %d\n", len(opts.Project))
	fmt.Printf("Ignore patterns: %d\n", len(opts.ProjectFolderIgnoreList))
	// Output:
	// Projects: 2
	// Ignore patterns: 3
}

// Example_extraFileExtensions demonstrates configuring additional file extensions.
func Example_extraFileExtensions() {
	// Support .vue files as TypeScript
	opts, _ := typescriptestree.NewServicesBuilder().
		WithExtraFileExtensions(".vue", ".svelte").
		Build()

	fmt.Printf("Extra extensions: %v\n", opts.ExtraFileExtensions)
	// Output:
	// Extra extensions: [.vue .svelte]
}

// Example_caching demonstrates cache configuration for performance.
func Example_caching() {
	// Configure glob cache lifetime
	opts, _ := typescriptestree.NewServicesBuilder().
		WithGlobCacheLifetime(300). // 5 minutes
		Build()

	if opts.CacheLifetime != nil && opts.CacheLifetime.Glob != nil {
		fmt.Printf("Glob cache lifetime: %d seconds\n", *opts.CacheLifetime.Glob)
	}
	// Output:
	// Glob cache lifetime: 300 seconds
}

// Example_debugging demonstrates enabling debug output.
func Example_debugging() {
	opts, _ := typescriptestree.NewBuilder().
		WithDebugLevel("typescript-estree", "parser").
		Build()

	fmt.Printf("Debug modules: %v\n", opts.DebugLevel)
	// Output:
	// Debug modules: [typescript-estree parser]
}

// Example_errorHandling demonstrates error handling configuration.
func Example_errorHandling() {
	// Strict mode: error on TypeScript issues
	strictOpts, _ := typescriptestree.NewServicesBuilder().
		WithErrorOnTypeScriptSyntacticAndSemanticIssues(true).
		WithErrorOnUnknownASTType(true).
		Build()

	fmt.Printf("Error on TS issues: %v\n", strictOpts.ErrorOnTypeScriptSyntacticAndSemanticIssues)
	fmt.Printf("Error on unknown AST: %v\n", strictOpts.ErrorOnUnknownASTType)

	// Lenient mode: allow invalid AST
	lenientOpts, _ := typescriptestree.NewBuilder().
		WithAllowInvalidAST(true).
		Build()

	fmt.Printf("Allow invalid AST: %v\n", lenientOpts.AllowInvalidAST)
	// Output:
	// Error on TS issues: true
	// Error on unknown AST: true
	// Allow invalid AST: true
}

// Example_customLogger demonstrates custom logging configuration.
func Example_customLogger() {
	messages := []string{}
	customLogger := func(msg string) {
		messages = append(messages, msg)
	}

	opts, _ := typescriptestree.NewBuilder().
		WithLoggerFn(customLogger).
		Build()

	// Logger is configured
	if opts.LoggerFn != nil {
		opts.LoggerFn("Test message")
		fmt.Printf("Logged messages: %d\n", len(messages))
	}
	// Output:
	// Logged messages: 1
}

// Example_validation demonstrates option validation.
func Example_validation() {
	// Invalid source type
	opts := typescriptestree.NewParseOptions()
	opts.SourceType = "invalid"

	if err := opts.Validate(); err != nil {
		fmt.Println("Validation failed as expected")
	}

	// Valid options
	opts.SourceType = typescriptestree.SourceTypeModule
	if err := opts.Validate(); err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
	} else {
		fmt.Println("Validation passed")
	}
	// Output:
	// Validation failed as expected
	// Validation passed
}

// Example_realWorldConfiguration demonstrates a realistic configuration
// for a TypeScript project with ESLint.
func Example_realWorldConfiguration() {
	opts, err := typescriptestree.NewServicesBuilder().
		// Base parse options
		WithParseOptions(typescriptestree.ParseOptions{
			SourceType:       typescriptestree.SourceTypeModule,
			Loc:              true,
			Range:            true,
			Comment:          true,
			Tokens:           true,
			JSDocParsingMode: typescriptestree.JSDocParsingModeTypeInfo,
		}).
		// Type-aware features
		WithProject("./tsconfig.json").
		WithTSConfigRootDir(".").
		// Performance
		WithGlobCacheLifetime(300).
		// Error handling
		WithErrorOnTypeScriptSyntacticAndSemanticIssues(false).
		// Build
		Build()

	if err != nil {
		fmt.Printf("Configuration error: %v\n", err)
		return
	}

	fmt.Println("Real-world configuration created successfully")
	fmt.Printf("Source type: %s\n", opts.SourceType)
	fmt.Printf("Type-aware: %v\n", len(opts.Project) > 0)
	// Output:
	// Real-world configuration created successfully
	// Source type: module
	// Type-aware: true
}
