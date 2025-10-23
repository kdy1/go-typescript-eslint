# Parser Options

This document describes all available parser options for `go-typescript-eslint`, matching the configuration API of `@typescript-eslint/typescript-estree`.

## Table of Contents

- [ParseOptions](#parseoptions)
- [ParseAndGenerateServicesOptions](#parseandgenerateservicesoptions)
- [Builder Pattern](#builder-pattern)
- [Examples](#examples)

## ParseOptions

Basic options for parsing TypeScript/JavaScript code into an ESTree-compatible AST.

### Fields

#### `SourceType`
- **Type**: `SourceType` ("script" | "module")
- **Default**: "script"
- **Description**: Specifies whether the code should be parsed as a script or module. Modules support `import`/`export` statements.

```go
opts := typescriptestree.NewParseOptions()
opts.SourceType = typescriptestree.SourceTypeModule
```

#### `AllowInvalidAST`
- **Type**: `bool`
- **Default**: `false`
- **Description**: Prevents the parser from throwing an error if it receives an invalid AST from TypeScript. Useful for parsing malformed code.

```go
opts.AllowInvalidAST = true
```

#### `Comment`
- **Type**: `bool`
- **Default**: `false`
- **Description**: When enabled, creates a top-level comments array containing all comments found in the source code.

```go
opts.Comment = true
```

#### `DebugLevel`
- **Type**: `DebugLevel` ([]string)
- **Default**: `nil`
- **Description**: Enables detailed debugging output for specific modules. Pass module names to enable logging.

```go
opts.DebugLevel = []string{"typescript-estree", "parser"}
```

#### `ErrorOnUnknownASTType`
- **Type**: `bool`
- **Default**: `false`
- **Description**: Causes the parser to throw an error if it encounters an unknown AST node type.

```go
opts.ErrorOnUnknownASTType = true
```

#### `FilePath`
- **Type**: `string`
- **Default**: `""`
- **Description**: Absolute or relative path to the file being parsed. Used for error messages and automatic JSX detection (.tsx files).

```go
opts.FilePath = "src/components/App.tsx"
opts.InferJSXFromFilePath() // Automatically enables JSX for .tsx files
```

#### `JSDocParsingMode`
- **Type**: `JSDocParsingMode` ("all" | "none" | "type-info")
- **Default**: "all"
- **Description**: Controls how JSDoc comments are parsed.
  - `all`: Parse all JSDoc comments
  - `none`: Skip JSDoc parsing
  - `type-info`: Only parse JSDoc needed for type information

```go
opts.JSDocParsingMode = typescriptestree.JSDocParsingModeTypeInfo
```

#### `JSX`
- **Type**: `bool`
- **Default**: `false` (automatically `true` for .tsx files)
- **Description**: Enables parsing of JSX syntax.

```go
opts.JSX = true
```

#### `Loc`
- **Type**: `bool`
- **Default**: `false`
- **Description**: Adds line/column location information to AST nodes via the `loc` property.

```go
opts.Loc = true
```

#### `LoggerFn`
- **Type**: `LoggerFn` (func(message string))
- **Default**: Logs to stderr
- **Description**: Custom logging function. Set to `nil` to disable logging.

```go
opts.LoggerFn = func(msg string) {
    log.Println("[parser]", msg)
}
```

#### `Range`
- **Type**: `bool`
- **Default**: `false`
- **Description**: Adds `range` information to AST nodes indicating character start/end positions.

```go
opts.Range = true
```

#### `Tokens`
- **Type**: `bool`
- **Default**: `false`
- **Description**: Creates a top-level tokens array containing all tokens found during lexical analysis.

```go
opts.Tokens = true
```

#### `SuppressDeprecatedPropertyWarnings`
- **Type**: `bool`
- **Default**: `false`
- **Description**: Prevents warnings about deprecated AST properties from being logged.

```go
opts.SuppressDeprecatedPropertyWarnings = true
```

## ParseAndGenerateServicesOptions

Extended options for type-aware parsing with TypeScript language services support.

### Additional Fields (beyond ParseOptions)

#### `CacheLifetime`
- **Type**: `*CacheLifetime`
- **Default**: `nil`
- **Description**: Controls internal cache expiry times for performance optimization.

```go
opts.CacheLifetime = &typescriptestree.CacheLifetime{
    Glob: &typescriptestree.CacheDurationSeconds(300), // 5 minutes
}
```

#### `DisallowAutomaticSingleRunInference`
- **Type**: `bool`
- **Default**: `false`
- **Description**: Disables the performance heuristic that infers whether the parser is being used for a single run or multiple runs.

```go
opts.DisallowAutomaticSingleRunInference = true
```

#### `ErrorOnTypeScriptSyntacticAndSemanticIssues`
- **Type**: `bool`
- **Default**: `false`
- **Description**: Causes the parser to throw an error if TypeScript reports any syntactic or semantic issues.

```go
opts.ErrorOnTypeScriptSyntacticAndSemanticIssues = true
```

#### `ExtraFileExtensions`
- **Type**: `[]string`
- **Default**: `nil`
- **Description**: Additional file extensions to treat as TypeScript beyond `.ts`, `.tsx`, `.mts`, `.cts`. Each extension must start with a dot.

```go
opts.ExtraFileExtensions = []string{".vue", ".svelte"}
```

#### `PreserveNodeMaps`
- **Type**: `*bool`
- **Default**: `true`
- **Description**: Preserves TypeScript's internal AST node maps for language services. Required for certain type-checking operations.

```go
preserve := true
opts.PreserveNodeMaps = &preserve
```

#### `Project`
- **Type**: `[]string`
- **Default**: `nil`
- **Description**: Paths to TypeScript configuration files (`tsconfig.json`) or directories containing them. Supports glob patterns. Enables type-aware parsing.

```go
opts.Project = []string{
    "./tsconfig.json",
    "./packages/*/tsconfig.json",
}
```

#### `ProjectFolderIgnoreList`
- **Type**: `[]string`
- **Default**: `["**/node_modules/**"]`
- **Description**: Folder patterns to ignore when searching for project files.

```go
opts.ProjectFolderIgnoreList = []string{
    "**/node_modules/**",
    "**/dist/**",
    "**/.cache/**",
}
```

#### `ProjectService`
- **Type**: `bool`
- **Default**: `false`
- **Description**: Enables TypeScript's project service for managing multiple projects with shared state.

```go
opts.ProjectService = true
```

#### `TSConfigRootDir`
- **Type**: `string`
- **Default**: Current working directory
- **Description**: Root directory for resolving relative tsconfig paths specified in `Project`.

```go
opts.TSConfigRootDir = "/path/to/project"
```

#### `Programs`
- **Type**: `[]interface{}`
- **Default**: `nil`
- **Description**: Pre-created TypeScript Program instances. Advanced option for performance optimization.

```go
opts.Programs = []interface{}{myTSProgram}
```

#### `WarnOnUnsupportedTypeScriptVersion`
- **Type**: `*bool`
- **Default**: `true`
- **Description**: Controls whether to warn when using an unsupported TypeScript version.

```go
warn := false
opts.WarnOnUnsupportedTypeScriptVersion = &warn
```

## Builder Pattern

The package provides builder types for fluent configuration:

### ParseOptionsBuilder

```go
opts, err := typescriptestree.NewBuilder().
    WithSourceType(typescriptestree.SourceTypeModule).
    WithFilePath("src/app.tsx").
    WithLoc(true).
    WithRange(true).
    WithComment(true).
    WithTokens(true).
    Build()

if err != nil {
    // handle validation error
}
```

### ParseAndGenerateServicesOptionsBuilder

```go
opts, err := typescriptestree.NewServicesBuilder().
    WithProject("./tsconfig.json").
    WithTSConfigRootDir(".").
    WithGlobCacheLifetime(300).
    WithExtraFileExtensions(".vue").
    Build()

if err != nil {
    // handle validation error
}
```

## Examples

### Basic Parsing

```go
package main

import (
    "fmt"
    "github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"
)

func main() {
    // Create options
    opts := typescriptestree.NewBuilder().
        WithSourceType(typescriptestree.SourceTypeModule).
        WithLoc(true).
        WithRange(true).
        MustBuild()

    // Parse TypeScript code
    source := "const x: number = 42;"
    result, err := typescriptestree.Parse(source, opts)
    if err != nil {
        panic(err)
    }

    // Use the AST
    fmt.Printf("Parsed successfully: %v\n", result.AST != nil)
}
```

### Type-Aware Parsing

```go
package main

import (
    "github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"
)

func main() {
    // Configure for type-aware parsing
    opts := typescriptestree.NewServicesBuilder().
        WithProject("./tsconfig.json", "./tsconfig.test.json").
        WithTSConfigRootDir(".").
        WithErrorOnTypeScriptSyntacticAndSemanticIssues(true).
        MustBuild()

    source := "const x: string = 42;" // Type error
    result, err := typescriptestree.ParseAndGenerateServices(source, opts)
    if err != nil {
        // Will error due to type mismatch
        panic(err)
    }

    // Use result.AST and result.Services
}
```

### JSX/TSX Parsing

```go
opts := typescriptestree.NewBuilder().
    WithFilePath("components/App.tsx").  // Auto-enables JSX
    WithSourceType(typescriptestree.SourceTypeModule).
    MustBuild()

source := `
export function App() {
    return <div>Hello World</div>;
}
`
result, err := typescriptestree.Parse(source, opts)
```

### Custom Logging

```go
var logs []string

opts := typescriptestree.NewBuilder().
    WithLoggerFn(func(msg string) {
        logs = append(logs, msg)
    }).
    WithDebugLevel("parser", "lexer").
    MustBuild()
```

### Lenient Parsing

```go
// Parse potentially malformed code
opts := typescriptestree.NewBuilder().
    WithAllowInvalidAST(true).
    WithErrorOnUnknownASTType(false).
    MustBuild()
```

### Monorepo Configuration

```go
opts := typescriptestree.NewServicesBuilder().
    WithProject(
        "./packages/*/tsconfig.json",
        "./apps/*/tsconfig.json",
    ).
    WithProjectFolderIgnoreList(
        "**/node_modules/**",
        "**/dist/**",
        "**/.cache/**",
        "**/build/**",
    ).
    WithGlobCacheLifetime(600).  // 10 minutes
    MustBuild()
```

### Framework Integration (Vue/Svelte)

```go
opts := typescriptestree.NewServicesBuilder().
    WithExtraFileExtensions(".vue", ".svelte").
    WithProject("./tsconfig.json").
    MustBuild()
```

## Validation

All options are validated when using the builder's `Build()` method:

```go
opts, err := typescriptestree.NewBuilder().
    WithSourceType("invalid"). // Invalid!
    Build()

if err != nil {
    // err: "invalid sourceType: must be 'script' or 'module', got \"invalid\""
}
```

Use `MustBuild()` to panic on validation errors:

```go
opts := typescriptestree.NewBuilder().
    WithSourceType(typescriptestree.SourceTypeModule).
    MustBuild() // Panics if validation fails
```

## Default Values Summary

| Option | Default Value |
|--------|---------------|
| SourceType | "script" |
| AllowInvalidAST | false |
| Comment | false |
| DebugLevel | nil |
| ErrorOnUnknownASTType | false |
| FilePath | "" |
| JSDocParsingMode | "all" |
| JSX | false (true for .tsx) |
| Loc | false |
| Range | false |
| Tokens | false |
| SuppressDeprecatedPropertyWarnings | false |
| CacheLifetime | nil |
| DisallowAutomaticSingleRunInference | false |
| ErrorOnTypeScriptSyntacticAndSemanticIssues | false |
| ExtraFileExtensions | nil |
| PreserveNodeMaps | true |
| Project | nil |
| ProjectFolderIgnoreList | ["**/node_modules/**"] |
| ProjectService | false |
| TSConfigRootDir | "." |
| Programs | nil |
| WarnOnUnsupportedTypeScriptVersion | true |

## Compatibility

These options match the API of `@typescript-eslint/typescript-estree` version 8.x. For the most up-to-date reference, see:
- [TypeScript ESTree Documentation](https://typescript-eslint.io/packages/typescript-estree/)
- [Parser Options Source](https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/typescript-estree/src/parser-options.ts)
