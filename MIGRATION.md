# Migration Guide: From typescript-estree to go-typescript-eslint

This guide helps you migrate from [@typescript-eslint/typescript-estree](https://typescript-eslint.io/packages/typescript-estree/) (JavaScript/TypeScript) to go-typescript-eslint (Go).

## Table of Contents

- [Overview](#overview)
- [Quick Reference](#quick-reference)
- [API Mapping](#api-mapping)
- [Configuration Options](#configuration-options)
- [Code Examples](#code-examples)
- [Common Patterns](#common-patterns)
- [Feature Compatibility](#feature-compatibility)
- [Performance Comparison](#performance-comparison)
- [Common Gotchas](#common-gotchas)
- [Best Practices](#best-practices)

## Overview

go-typescript-eslint is a faithful Go port of typescript-estree that maintains API compatibility where possible while following Go idioms and best practices.

### Key Differences

| Aspect | typescript-estree | go-typescript-eslint |
|--------|------------------|---------------------|
| Language | JavaScript/TypeScript | Go |
| Package Manager | npm | Go modules |
| Import | `import { parse } from '@typescript-eslint/typescript-estree'` | `import "github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"` |
| Configuration | Object literals | Builder pattern or structs |
| Naming | camelCase | PascalCase (exported) |
| Error Handling | Exceptions | Error returns |
| Type Safety | TypeScript types | Go types |

## Quick Reference

### Installation

**typescript-estree:**
```bash
npm install @typescript-eslint/typescript-estree
```

**go-typescript-eslint:**
```bash
go get github.com/kdy1/go-typescript-eslint
```

### Basic Usage

**typescript-estree:**
```typescript
import { parse } from '@typescript-eslint/typescript-estree';

const result = parse(code, {
  sourceType: 'module',
  loc: true,
  range: true
});
```

**go-typescript-eslint:**
```go
import "github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"

opts := typescriptestree.NewBuilder().
    WithSourceType(typescriptestree.SourceTypeModule).
    WithLoc(true).
    WithRange(true).
    MustBuild()

result, err := typescriptestree.Parse(code, opts)
if err != nil {
    // Handle error
}
```

### Type-Aware Parsing

**typescript-estree:**
```typescript
import { parseAndGenerateServices } from '@typescript-eslint/typescript-estree';

const result = parseAndGenerateServices(code, {
  project: './tsconfig.json',
  tsconfigRootDir: '.'
});
```

**go-typescript-eslint:**
```go
opts := typescriptestree.NewServicesBuilder().
    WithProject("./tsconfig.json").
    WithTSConfigRootDir(".").
    Build()

result, err := typescriptestree.ParseAndGenerateServices(code, opts)
if err != nil {
    // Handle error
}
```

## API Mapping

### Functions

| typescript-estree | go-typescript-eslint | Notes |
|------------------|---------------------|-------|
| `parse(code, options)` | `Parse(source, opts)` | Same functionality |
| `parseAndGenerateServices(code, options)` | `ParseAndGenerateServices(source, opts)` | Same functionality |
| `createProgram(configFile)` | Handled internally by program cache | Not exposed in public API |
| `clearCaches()` | `ClearProgramCache()` | Clears program cache |
| N/A | `ClearDefaultProjectMatchedFiles()` | Compatibility stub |

### Types

| typescript-estree | go-typescript-eslint | Notes |
|------------------|---------------------|-------|
| `TSESTree.Program` | `*ast.Program` | Root AST node |
| `TSESTree.Node` | `ast.Node` | Base node interface |
| `ParserServices` | `*Services` or `*ParserServices` | Type information |
| `ParseOptions` | `*ParseOptions` | Basic parsing options |
| `ParseAndGenerateServicesOptions` | `*ParseAndGenerateServicesOptions` | Type-aware options |

### Constants

| typescript-estree | go-typescript-eslint | Notes |
|------------------|---------------------|-------|
| `AST_NODE_TYPES.Identifier` | `AST_NODE_TYPES.Identifier` | Exact same |
| `AST_TOKEN_TYPES.Keyword` | `AST_TOKEN_TYPES.Keyword` | Exact same |

## Configuration Options

### ParseOptions Mapping

| typescript-estree | go-typescript-eslint | Type | Notes |
|------------------|---------------------|------|-------|
| `sourceType` | `SourceType` | `SourceType` (enum) | Use `SourceTypeScript` or `SourceTypeModule` |
| `jsx` | `JSX` | `bool` | Same |
| `comment` | `Comment` | `bool` | Same |
| `tokens` | `Tokens` | `bool` | Same |
| `loc` | `Loc` | `bool` | Same |
| `range` | `Range` | `bool` | Same |
| `filePath` | `FilePath` | `string` | Same |
| `allowInvalidAST` | `AllowInvalidAST` | `bool` | Same |
| `jsDocParsingMode` | `JSDocParsingMode` | `JSDocParsingMode` (enum) | Use `JSDocParsingModeAll`, `JSDocParsingModeNone`, or `JSDocParsingModeTypeInfo` |
| `errorOnUnknownASTType` | `ErrorOnUnknownASTType` | `bool` | Same |
| `debugLevel` | `DebugLevel` | `[]string` | Same |
| `loggerFn` | `LoggerFn` | `func(string)` | Same concept, different signature |
| `suppressDeprecatedPropertyWarnings` | `SuppressDeprecatedPropertyWarnings` | `bool` | Same |

### ParseAndGenerateServicesOptions Mapping

| typescript-estree | go-typescript-eslint | Type | Notes |
|------------------|---------------------|------|-------|
| All ParseOptions | All ParseOptions | | Inherits all basic options |
| `project` | `Project` | `[]string` | Same |
| `tsconfigRootDir` | `TSConfigRootDir` | `string` | Same |
| `projectService` | `ProjectService` | `bool` | Same (not yet implemented) |
| `preserveNodeMaps` | `PreserveNodeMaps` | `*bool` | Pointer for optional value |
| `programs` | `Programs` | `[]*program.Program` | Different type |
| `extraFileExtensions` | `ExtraFileExtensions` | `[]string` | Same |
| `projectFolderIgnoreList` | `ProjectFolderIgnoreList` | `[]string` | Same |
| `cacheLifetime` | `CacheLifetime` | `*CacheLifetime` | Same structure |
| `errorOnTypeScriptSyntacticAndSemanticIssues` | `ErrorOnTypeScriptSyntacticAndSemanticIssues` | `bool` | Same |
| `disallowAutomaticSingleRunInference` | `DisallowAutomaticSingleRunInference` | `bool` | Same |
| `warnOnUnsupportedTypeScriptVersion` | `WarnOnUnsupportedTypeScriptVersion` | `*bool` | Pointer for optional value |

## Code Examples

### Example 1: Basic Parsing

**typescript-estree:**
```typescript
import { parse, AST_NODE_TYPES } from '@typescript-eslint/typescript-estree';

const code = 'const x: number = 42;';
const ast = parse(code, {
  sourceType: 'module',
  loc: true,
  range: true
});

console.log(ast.type); // 'Program'

// Check node type
if (ast.body[0].type === AST_NODE_TYPES.VariableDeclaration) {
  console.log('Found variable declaration');
}
```

**go-typescript-eslint:**
```go
package main

import (
    "fmt"
    "log"
    "github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"
)

func main() {
    code := "const x: number = 42;"
    opts := typescriptestree.NewBuilder().
        WithSourceType(typescriptestree.SourceTypeModule).
        WithLoc(true).
        WithRange(true).
        MustBuild()

    result, err := typescriptestree.Parse(code, opts)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(result.AST.Type()) // "Program"

    // Check node type
    if len(result.AST.Body) > 0 {
        if result.AST.Body[0].Type() == typescriptestree.AST_NODE_TYPES.VariableDeclaration {
            fmt.Println("Found variable declaration")
        }
    }
}
```

### Example 2: Type-Aware Parsing

**typescript-estree:**
```typescript
import { parseAndGenerateServices } from '@typescript-eslint/typescript-estree';

const result = parseAndGenerateServices(code, {
  filePath: './src/index.ts',
  project: './tsconfig.json',
  tsconfigRootDir: process.cwd()
});

// Access type information
const checker = result.services.program.getTypeChecker();
const sourceFile = result.services.program.getSourceFile('./src/index.ts');
```

**go-typescript-eslint:**
```go
package main

import (
    "log"
    "os"
    "github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"
)

func main() {
    code := "/* TypeScript code */"
    cwd, _ := os.Getwd()

    opts := typescriptestree.NewServicesBuilder().
        WithFilePath("./src/index.ts").
        WithProject("./tsconfig.json").
        WithTSConfigRootDir(cwd).
        Build()

    result, err := typescriptestree.ParseAndGenerateServices(code, opts)
    if err != nil {
        log.Fatal(err)
    }

    // Access compiler options
    if result.Services != nil {
        compilerOpts := result.Services.GetCompilerOptions()
        // Use compiler options...
    }
}
```

### Example 3: JSX/TSX Parsing

**typescript-estree:**
```typescript
import { parse } from '@typescript-eslint/typescript-estree';

const jsxCode = 'const App = () => <div>Hello</div>;';
const ast = parse(jsxCode, {
  jsx: true,
  sourceType: 'module'
});

// Or auto-detect from file path
const ast2 = parse(jsxCode, {
  filePath: 'Component.tsx'  // Automatically enables JSX
});
```

**go-typescript-eslint:**
```go
package main

import (
    "log"
    "github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"
)

func main() {
    jsxCode := "const App = () => <div>Hello</div>;"

    // Explicit JSX option
    opts1 := typescriptestree.NewBuilder().
        WithJSX(true).
        WithSourceType(typescriptestree.SourceTypeModule).
        MustBuild()

    result1, err := typescriptestree.Parse(jsxCode, opts1)
    if err != nil {
        log.Fatal(err)
    }

    // Or auto-detect from file path
    opts2 := typescriptestree.NewBuilder().
        WithFilePath("Component.tsx").  // Automatically enables JSX
        MustBuild()

    result2, err := typescriptestree.Parse(jsxCode, opts2)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Example 4: Error Handling

**typescript-estree:**
```typescript
import { parse } from '@typescript-eslint/typescript-estree';

try {
  const ast = parse(invalidCode, {
    allowInvalidAST: false
  });
} catch (error) {
  console.error('Parse error:', error.message);
}

// Or allow invalid AST
const ast = parse(invalidCode, {
  allowInvalidAST: true
});
// ast may be partial but won't throw
```

**go-typescript-eslint:**
```go
package main

import (
    "fmt"
    "log"
    "github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"
)

func main() {
    invalidCode := "const x = "  // Incomplete

    // Strict parsing (return error)
    opts1 := typescriptestree.NewBuilder().
        WithAllowInvalidAST(false).
        MustBuild()

    result1, err := typescriptestree.Parse(invalidCode, opts1)
    if err != nil {
        fmt.Printf("Parse error: %v\n", err)
        // Handle error
    }

    // Allow invalid AST (partial result)
    opts2 := typescriptestree.NewBuilder().
        WithAllowInvalidAST(true).
        MustBuild()

    result2, err := typescriptestree.Parse(invalidCode, opts2)
    // result2.AST may contain partial AST
    // err may still contain error information
    if err != nil {
        fmt.Printf("Parse warning: %v\n", err)
    }
    if result2 != nil && result2.AST != nil {
        fmt.Println("Partial AST available")
    }
}
```

### Example 5: Node Mappings

**typescript-estree:**
```typescript
import { parseAndGenerateServices } from '@typescript-eslint/typescript-estree';

const result = parseAndGenerateServices(code, {
  project: './tsconfig.json',
  preserveNodeMaps: true
});

// Get TypeScript node from ESTree node
const tsNode = result.services.esTreeNodeToTSNodeMap.get(esTreeNode);

// Get ESTree node from TypeScript node
const esTreeNode = result.services.tsNodeToESTreeNodeMap.get(tsNode);
```

**go-typescript-eslint:**
```go
package main

import (
    "log"
    "github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"
)

func main() {
    code := "/* code */"

    opts := typescriptestree.NewServicesBuilder().
        WithProject("./tsconfig.json").
        WithPreserveNodeMaps(true).
        Build()

    result, err := typescriptestree.ParseAndGenerateServices(code, opts)
    if err != nil {
        log.Fatal(err)
    }

    // Get TypeScript node from ESTree node
    if result.Services != nil {
        tsNode, ok := result.Services.GetTSNodeForESTreeNode(esTreeNode)
        if ok {
            // Use tsNode...
        }

        // Get ESTree node from TypeScript node
        esTreeNode, ok := result.Services.GetESTreeNodeForTSNode(tsNode)
        if ok {
            // Use esTreeNode...
        }
    }
}
```

## Common Patterns

### Pattern 1: Builder Pattern vs. Object Literals

**typescript-estree:**
```typescript
const options = {
  sourceType: 'module',
  jsx: true,
  loc: true,
  range: true,
  comment: true,
  tokens: true
};

const ast = parse(code, options);
```

**go-typescript-eslint:**
```go
// Option 1: Builder pattern (recommended)
opts := typescriptestree.NewBuilder().
    WithSourceType(typescriptestree.SourceTypeModule).
    WithJSX(true).
    WithLoc(true).
    WithRange(true).
    WithComment(true).
    WithTokens(true).
    MustBuild()

result, err := typescriptestree.Parse(code, opts)

// Option 2: Direct struct construction
opts := &typescriptestree.ParseOptions{
    SourceType: typescriptestree.SourceTypeModule,
    JSX:        true,
    Loc:        true,
    Range:      true,
    Comment:    true,
    Tokens:     true,
}

result, err := typescriptestree.Parse(code, opts)
```

### Pattern 2: Default Options

**typescript-estree:**
```typescript
// Use defaults
const ast = parse(code);

// Or explicit defaults
const ast = parse(code, {});
```

**go-typescript-eslint:**
```go
// Pass nil for defaults
result, err := typescriptestree.Parse(code, nil)

// Or explicit defaults
opts := typescriptestree.NewBuilder().MustBuild()
result, err := typescriptestree.Parse(code, opts)
```

### Pattern 3: AST Traversal

**typescript-estree:**
```typescript
import { AST_NODE_TYPES, simpleTraverse } from '@typescript-eslint/typescript-estree';

simpleTraverse(ast, {
  enter(node) {
    if (node.type === AST_NODE_TYPES.Identifier) {
      console.log('Found identifier:', node.name);
    }
  }
});
```

**go-typescript-eslint:**
```go
import (
    "fmt"
    "github.com/kdy1/go-typescript-eslint/internal/ast"
    "github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"
)

// Use ast.Walk for traversal
ast.Walk(result.AST, ast.VisitorFunc(func(node ast.Node) bool {
    if node.Type() == typescriptestree.AST_NODE_TYPES.Identifier {
        if id, ok := node.(*ast.Identifier); ok {
            fmt.Printf("Found identifier: %s\n", id.Name)
        }
    }
    return true  // Continue traversal
}))
```

## Feature Compatibility

### Fully Compatible

- ✅ Basic parsing (`parse()` / `Parse()`)
- ✅ Type-aware parsing (`parseAndGenerateServices()` / `ParseAndGenerateServices()`)
- ✅ All parse options
- ✅ JSX/TSX support
- ✅ Node type constants (`AST_NODE_TYPES`)
- ✅ Token type constants (`AST_TOKEN_TYPES`)
- ✅ Comments and tokens collection
- ✅ Location and range information
- ✅ Program caching
- ✅ tsconfig.json support
- ✅ Node mappings

### Partially Compatible

- 🟡 **Project service**: API exists but not fully implemented
- 🟡 **Type checker**: Basic access available, full integration planned
- 🟡 **Debugging**: Debug levels supported, format differs

### Not Yet Implemented

- ❌ **Incremental parsing**: Not yet available
- ❌ **createProgram()**: Handled internally, not exposed
- ❌ **simpleTraverse()**: Use `ast.Walk()` instead

## Performance Comparison

### Memory Usage

- **typescript-estree**: Node.js runtime overhead + AST
- **go-typescript-eslint**: Go runtime (lower baseline) + AST

**Expected**: go-typescript-eslint uses ~30-40% less memory

### Parse Speed

- **typescript-estree**: V8 JIT compilation benefits
- **go-typescript-eslint**: Native code, no JIT warmup

**Expected**: Similar performance, go-typescript-eslint faster for cold starts

### Program Caching

Both implementations cache TypeScript programs. Performance is comparable.

## Common Gotchas

### 1. Error Handling

**typescript-estree** uses exceptions:
```typescript
try {
  const ast = parse(code);
} catch (error) {
  // Handle error
}
```

**go-typescript-eslint** uses error returns:
```go
result, err := typescriptestree.Parse(code, opts)
if err != nil {
    // Handle error
}
```

### 2. Nil Checks

In Go, always check for nil:

```go
result, err := typescriptestree.Parse(code, nil)
if err != nil {
    return err
}

// Check result
if result != nil && result.AST != nil {
    // Use AST
}

// Check services
if result.Services != nil {
    // Use services
}
```

### 3. Builder Validation

**MustBuild()** panics on invalid options:

```go
// This panics if options are invalid
opts := typescriptestree.NewBuilder().
    WithSourceType("invalid").  // Invalid value
    MustBuild()  // PANICS!

// Use Build() for error handling
opts, err := typescriptestree.NewBuilder().
    WithSourceType("invalid").
    Build()  // Returns error
```

### 4. Node Type Assertions

In Go, use type assertions carefully:

```go
if node.Type() == typescriptestree.AST_NODE_TYPES.Identifier {
    // Type assertion with check
    if id, ok := node.(*ast.Identifier); ok {
        fmt.Println(id.Name)
    }
}
```

### 5. Concurrent Access

Go makes concurrency explicit. If parsing multiple files:

```go
// Sequential
for _, file := range files {
    result, err := typescriptestree.Parse(file, opts)
    // Process result
}

// Concurrent (with goroutines)
var wg sync.WaitGroup
for _, file := range files {
    wg.Add(1)
    go func(f string) {
        defer wg.Done()
        result, err := typescriptestree.Parse(f, opts)
        // Process result
    }(file)
}
wg.Wait()
```

### 6. Import Paths

Go uses full import paths:

```go
// Correct
import "github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"

// Incorrect
import "typescriptestree"  // Won't work
```

## Best Practices

### 1. Reuse Options

Create options once and reuse them:

```go
// Good
opts := typescriptestree.NewBuilder().
    WithSourceType(typescriptestree.SourceTypeModule).
    WithLoc(true).
    MustBuild()

for _, code := range sources {
    result, err := typescriptestree.Parse(code, opts)
    // Process result
}
```

### 2. Clear Cache in Tests

Always clear the program cache in tests:

```go
func TestMyParser(t *testing.T) {
    defer typescriptestree.ClearProgramCache()

    // Test code
}
```

### 3. Check Errors

Always check error returns:

```go
// Good
result, err := typescriptestree.Parse(code, opts)
if err != nil {
    return fmt.Errorf("parse failed: %w", err)
}

// Bad
result, _ := typescriptestree.Parse(code, opts)  // Ignoring errors
```

### 4. Use Type-Safe Constants

Use exported constants instead of strings:

```go
// Good
if node.Type() == typescriptestree.AST_NODE_TYPES.Identifier {
    // ...
}

// Bad
if node.Type() == "Identifier" {  // Typo-prone
    // ...
}
```

### 5. Handle Nil Properly

Check for nil before dereferencing:

```go
// Good
if result != nil && result.Services != nil {
    opts := result.Services.GetCompilerOptions()
    if opts != nil {
        // Use opts
    }
}

// Bad
opts := result.Services.GetCompilerOptions()  // May panic if Services is nil
```

## Migration Checklist

When migrating from typescript-estree:

- [ ] Install go-typescript-eslint: `go get github.com/kdy1/go-typescript-eslint`
- [ ] Replace import statements
- [ ] Convert function calls: `parse()` → `Parse()`
- [ ] Convert configuration: object literals → builder pattern or structs
- [ ] Add error handling: exceptions → error returns
- [ ] Update constant references: keep the same
- [ ] Convert AST traversal: `simpleTraverse` → `ast.Walk`
- [ ] Add nil checks where needed
- [ ] Update tests to use Go testing package
- [ ] Clear cache in tests: `ClearProgramCache()`
- [ ] Review performance and adjust as needed

## Getting Help

If you encounter migration issues:

1. Check this guide for common patterns
2. Review the [examples/](examples/) directory
3. Read the [API documentation](https://pkg.go.dev/github.com/kdy1/go-typescript-eslint)
4. Check existing [GitHub Issues](https://github.com/kdy1/go-typescript-eslint/issues)
5. Ask in [GitHub Discussions](https://github.com/kdy1/go-typescript-eslint/discussions)
6. Refer to [ARCHITECTURE.md](ARCHITECTURE.md) for internals

## Additional Resources

- [typescript-estree Documentation](https://typescript-eslint.io/packages/typescript-estree/)
- [Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Testing](https://go.dev/doc/tutorial/add-a-test)

---

For questions or feedback about this migration guide, please open a [GitHub Issue](https://github.com/kdy1/go-typescript-eslint/issues).
