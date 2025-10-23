# Program Package

This package provides TypeScript program creation and management utilities for type-aware parsing and linting.

## Overview

The `program` package implements:

- **TSConfig Parsing**: Parse and resolve `tsconfig.json` files with inheritance support
- **Program Creation**: Create TypeScript programs from configuration files
- **Program Caching**: Efficient caching of program instances for performance
- **Type Information**: Foundation for accessing TypeScript type checker (future)

## Components

### TSConfig Management

Parse and resolve TypeScript configuration files:

```go
// Parse a single tsconfig.json
config, err := program.ParseTSConfig("./tsconfig.json")

// Resolve with inheritance (handles "extends")
config, err := program.ResolveTSConfig("./tsconfig.json")

// Find config for a specific file
configPath, err := program.FindConfigForFile("./src/app.ts")
```

### TSConfig Structure

The `TSConfig` type represents a complete TypeScript configuration:

```go
type TSConfig struct {
    Extends         string
    CompilerOptions CompilerOptions
    Files           []string
    Include         []string
    Exclude         []string
    References      []ProjectReference
    CompileOnSave   *bool
}
```

Supports all major compiler options including:
- Target and module settings
- Strict mode flags
- Module resolution options
- Path mappings
- Type roots and type acquisition
- Decorators and experimental features

### Program Creation

Create TypeScript programs for type-aware operations:

```go
opts := &program.ProgramOptions{
    TSConfigPath: "./tsconfig.json",
    RootDir:      ".",
    SourceFiles:  []string{"src/app.ts", "src/utils.ts"},
}

prog, err := program.CreateProgram(opts)
if err != nil {
    log.Fatal(err)
}

// Access compiler options
compilerOpts := prog.GetCompilerOptions()
fmt.Println("Target:", compilerOpts.Target)

// Add parsed source files
prog.AddSourceFile("src/app.ts", astNode)

// Retrieve source files
ast, ok := prog.GetSourceFile("src/app.ts")
```

### Program Caching

Efficient caching for improved performance:

```go
// Create cache with 5-minute expiration
cache := program.NewProgramCache(5 * time.Minute)

// Set a program
cache.Set("./tsconfig.json", prog)

// Get from cache
if cached := cache.Get("./tsconfig.json"); cached != nil {
    // Use cached program
}

// Get or create
prog, err := cache.GetOrCreate(&program.ProgramOptions{
    TSConfigPath: "./tsconfig.json",
})

// Clean expired entries
cache.CleanExpired()

// Clear all
cache.Clear()
```

### Global Cache

A default global cache is available:

```go
// Use the global cache (5-minute expiration)
prog, err := program.GlobalCache.GetOrCreate(opts)
```

## TSConfig Inheritance

The package fully supports TypeScript's configuration inheritance via the `extends` field:

```json
// tsconfig.base.json
{
  "compilerOptions": {
    "strict": true,
    "target": "ES2015"
  }
}

// tsconfig.json
{
  "extends": "./tsconfig.base.json",
  "compilerOptions": {
    "target": "ES2020"  // Overrides base
  }
}
```

When resolved, the child configuration inherits from the parent and can override specific fields:

```go
config, err := program.ResolveTSConfig("./tsconfig.json")
// config.CompilerOptions.Target == "ES2020" (child override)
// config.CompilerOptions.Strict == true (inherited from base)
```

## Integration with ParserServices

Programs are used by ParserServices to provide type information:

```go
// In pkg/typescriptestree
services := typescriptestree.NewParserServices(prog)

// Access compiler options
opts := services.GetCompilerOptions()

// Future: Access type checker
typeChecker, err := services.GetTypeChecker()
```

## Thread Safety

The `Program` type includes mutex protection for concurrent access:

```go
// Safe for concurrent reads/writes
go func() {
    prog.AddSourceFile("file1.ts", ast1)
}()

go func() {
    ast, ok := prog.GetSourceFile("file2.ts")
}()
```

## Testing

The package includes comprehensive tests:

```bash
go test ./internal/program/... -v
```

Test coverage includes:
- TSConfig parsing and inheritance
- Program creation and management
- Cache operations and expiration
- Concurrent access patterns
- Config file discovery

## Future Enhancements

Planned features:

1. **Type Checker Integration**: Full integration with TypeScript's type checker or microsoft/typescript-go
2. **Symbol Resolution**: Symbol table management and lookup
3. **Incremental Compilation**: Support for incremental program updates
4. **Project References**: Full support for composite projects
5. **Diagnostic Reporting**: TypeScript compiler diagnostics
6. **Module Resolution**: Complete module resolution implementation

## References

- [TypeScript Compiler API](https://github.com/microsoft/TypeScript/wiki/Using-the-Compiler-API)
- [typescript-eslint createProgram](https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/typescript-estree/src/create-program/)
- [microsoft/typescript-go](https://github.com/microsoft/typescript-go)
- [TSConfig Reference](https://www.typescriptlang.org/tsconfig)
