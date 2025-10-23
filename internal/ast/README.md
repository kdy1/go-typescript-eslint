# AST Node Utilities

This package provides comprehensive utilities for manipulating, traversing, and analyzing Abstract Syntax Tree (AST) nodes in the go-typescript-eslint project.

## Overview

The AST utilities are organized into several key areas:

- **Visitor Keys**: Defines which child properties should be traversed for each node type
- **Type Guards**: Safe type checking and assertions for nodes
- **Traversal**: Walking and searching the AST with visitor patterns
- **Node Utilities**: Common operations on nodes (comparison, positioning, etc.)
- **Comment Utilities**: Attaching and managing comments
- **Token Utilities**: Working with lexical tokens

## Visitor Keys (`visitor_keys.go`)

Visitor keys define the child properties that should be traversed for each AST node type. These are essential for AST traversal operations.

### Usage

```go
// Get visitor keys for a node type
keys := ast.GetVisitorKeys("BinaryExpression")
// Returns: ["left", "right"]

// Check if a node type has visitor keys
if ast.HasVisitorKeys("Identifier") {
    // ...
}
```

### Key Features

- Complete coverage of all 177+ ESTree and TypeScript node types
- Keys ordered by source code appearance (not alphabetically)
- Based on TypeScript ESLint visitor keys specification

## Type Guards (`guards.go`)

Type guards provide safe type checking and assertions for AST nodes.

### Interface Guards

```go
// Check if a node implements an interface
if ast.IsExpression(node) {
    // node is an Expression
}

if ast.IsStatement(node) {
    // node is a Statement
}

if ast.IsPattern(node) {
    // node is a Pattern (destructuring)
}

if ast.IsDeclaration(node) {
    // node is a Declaration
}

if ast.IsTSNode(node) {
    // node is a TypeScript-specific node
}
```

### Type-Specific Guards

```go
// Check specific node types
if ast.IsIdentifier(node) { }
if ast.IsLiteral(node) { }
if ast.IsMemberExpression(node) { }
if ast.IsCallExpression(node) { }
if ast.IsFunctionExpression(node) { }
if ast.IsClassDeclaration(node) { }
```

### Category Guards

```go
// Check node categories
if ast.IsFunction(node) {
    // FunctionExpression, ArrowFunctionExpression, or FunctionDeclaration
}

if ast.IsClass(node) {
    // ClassExpression or ClassDeclaration
}

if ast.IsLoop(node) {
    // ForStatement, ForInStatement, ForOfStatement, WhileStatement, DoWhileStatement
}

if ast.IsJSX(node) {
    // Any JSX* node
}

if ast.IsTypeScript(node) {
    // Any TS* node
}

if ast.IsImport(node) {
    // Any import-related node
}

if ast.IsExport(node) {
    // Any export-related node
}
```

## AST Traversal (`traverse.go`)

Powerful utilities for walking and searching the AST.

### Basic Traversal

```go
// Simple traversal with a callback
ast.Traverse(rootNode, func(node ast.Node) bool {
    fmt.Println(node.Type())
    return true // return false to skip children
})
```

### Visitor Pattern

```go
// Implement the Visitor interface
type MyVisitor struct{}

func (v *MyVisitor) Visit(node ast.Node) bool {
    // Process the node
    return true
}

ast.Walk(rootNode, &MyVisitor{})

// Or use a visitor function
ast.Walk(rootNode, ast.VisitorFunc(func(node ast.Node) bool {
    // Process the node
    return true
}))
```

### Context-Aware Traversal

```go
// Traverse with parent and ancestor information
ast.TraverseWithContext(rootNode, func(node ast.Node, ctx *ast.TraverseContext) bool {
    fmt.Printf("Node: %s\n", node.Type())
    fmt.Printf("Parent: %v\n", ctx.Parent)
    fmt.Printf("Ancestors: %v\n", ctx.Ancestors)
    fmt.Printf("Property: %s\n", ctx.Key)
    if ctx.Index != nil {
        fmt.Printf("Array index: %d\n", *ctx.Index)
    }
    return true
})
```

### Finding Nodes

```go
// Find the first node matching a predicate
firstId := ast.FindFirst(rootNode, func(node ast.Node) bool {
    return ast.IsIdentifier(node)
})

// Find all nodes matching a predicate
allIds := ast.FindAll(rootNode, func(node ast.Node) bool {
    return ast.IsIdentifier(node)
})

// Find all nodes of a specific type
allBinary := ast.FindByType(rootNode, "BinaryExpression")
```

### Tree Relationships

```go
// Get parent of a node
parent := ast.GetParent(rootNode, targetNode)

// Get all ancestors
ancestors := ast.GetAncestors(rootNode, targetNode)

// Get siblings
siblings := ast.GetSiblings(rootNode, targetNode)

// Check if tree contains a node
if ast.Contains(rootNode, targetNode) {
    // ...
}

// Get depth of a node
depth := ast.GetDepth(rootNode, targetNode)

// Count total nodes in tree
count := ast.CountNodes(rootNode)
```

## Node Utilities (`utils.go`)

Common operations for working with individual nodes.

### Position and Range

```go
// Get node range as [start, end]
rng := ast.GetNodeRange(node)

// Get node span (length)
span := ast.NodeSpan(node)

// Check if position is in node's range
if ast.IsInRange(node, position) {
    // ...
}

// Find deepest node at a position
node := ast.GetNodeAtPosition(rootNode, position)

// Get all nodes in a range
nodes := ast.GetNodesInRange(rootNode, start, end)
```

### Node Relationships

```go
// Check if nodes overlap
if ast.NodesOverlap(nodeA, nodeB) {
    // ...
}

// Check if node A contains node B
if ast.NodeContains(nodeA, nodeB) {
    // ...
}

// Check ordering
if ast.IsBefore(nodeA, nodeB) {
    // nodeA appears before nodeB
}

if ast.IsAfter(nodeA, nodeB) {
    // nodeA appears after nodeB
}
```

### Source Text

```go
// Get source text of a node
source := "const foo = 42;"
text := ast.GetNodeText(node, source)
```

### Node Comparison

```go
// Deep equality check
if ast.NodeEquals(nodeA, nodeB) {
    // nodes are structurally equal
}
```

### Literal Checks

```go
if ast.IsNullLiteral(node) { }
if ast.IsBooleanLiteral(node) { }
if ast.IsStringLiteral(node) { }
if ast.IsNumberLiteral(node) { }
if ast.IsRegExpLiteral(node) { }
if ast.IsBigIntLiteral(node) { }
```

### Member Expression Checks

```go
// Check if static (obj.prop) or computed (obj[prop])
if ast.IsStaticMemberExpression(node) { }
if ast.IsComputedMemberExpression(node) { }

// Check if method call (obj.method())
if ast.IsMethodCall(node) { }
```

### Function Checks

```go
if ast.IsAsyncFunction(node) { }
if ast.IsGeneratorFunction(node) { }
if ast.IsArrowFunction(node) { }

// Check if function tree contains await/yield
if ast.HasAwait(node) { }
if ast.HasYield(node) { }

// Get function name
name := ast.GetFunctionName(funcNode)
```

### Class Utilities

```go
// Get class name
name := ast.GetClassName(classNode)
```

### Identifier Utilities

```go
// Get all identifiers in tree
identifiers := ast.GetAllIdentifiers(rootNode)

// Get unique identifier names
names := ast.GetIdentifierNames(rootNode)
```

### Export Checks

```go
if ast.IsExported(node) { }
if ast.IsDefaultExport(node) { }
if ast.IsNamedExport(node) { }
```

## Comment Utilities (`comments.go`)

Utilities for attaching and working with comments in the AST.

### Comment Attachment

```go
// Attach comments to nodes
attachments := ast.AttachComments(rootNode, comments)

// Get comments for a node
leading := ast.GetLeadingComments(attachments, node)
trailing := ast.GetTrailingComments(attachments, node)
inner := ast.GetInnerComments(attachments, node)
all := ast.GetAllComments(attachments, node)

// Check for comments
if ast.HasLeadingComment(attachments, node) {
    // ...
}
```

### Comment Filtering

```go
// Filter by type
lineComments := ast.GetLineComments(comments)
blockComments := ast.GetBlockComments(comments)
docComments := ast.GetDocComments(comments)

// Check if doc comment (JSDoc-style)
if ast.IsDocComment(comment) {
    // comment starts with /**
}
```

### Comment Position

```go
// Get comments in range
rangeComments := ast.GetCommentsInRange(comments, start, end)

// Get comments before/after position
before := ast.GetCommentsBefore(comments, position)
after := ast.GetCommentsAfter(comments, position)

// Get comments on specific line
lineComments := ast.GetCommentsOnLine(comments, lineNumber)
```

### Comment Utilities

```go
// Sort comments by position
ast.SortComments(comments)

// Get comment span
span := ast.CommentSpan(comment)

// Get comment text
text := ast.GetCommentText(comment)

// Check if comment is on same line as a location
if ast.IsCommentOnSameLine(comment, loc) {
    // ...
}
```

## Token Utilities (`tokens.go`)

Utilities for working with lexical tokens.

### Token Position

```go
// Get tokens in range
tokens := ast.GetTokensInRange(allTokens, start, end)

// Get tokens before/after position
before := ast.GetTokensBefore(allTokens, position)
after := ast.GetTokensAfter(allTokens, position)

// Get token at specific position
token := ast.GetTokenAtPosition(allTokens, position)
```

### Token Navigation

```go
// Get first/last token
first := ast.GetFirstToken(tokens)
last := ast.GetLastToken(tokens)

// Get next/previous token
next := ast.GetNextToken(tokens, currentToken)
prev := ast.GetPreviousToken(tokens, currentToken)
```

### Tokens for Nodes

```go
// Get all tokens for a node
nodeTokens := ast.GetTokensForNode(allTokens, node)

// Get first/last token of a node
first := ast.GetFirstTokenOfNode(allTokens, node)
last := ast.GetLastTokenOfNode(allTokens, node)
```

### Token Filtering

```go
// Get tokens by type or value
keywordTokens := ast.GetTokensByType(tokens, "Keyword")
plusTokens := ast.GetTokensByValue(tokens, "+")

// Get tokens on specific line
lineTokens := ast.GetTokensOnLine(tokens, lineNumber)

// Count tokens in range
count := ast.CountTokens(tokens, start, end)
```

### Token Type Checks

```go
if ast.IsKeywordToken(token) { }
if ast.IsIdentifierToken(token) { }
if ast.IsPunctuatorToken(token) { }
if ast.IsStringToken(token) { }
if ast.IsNumericToken(token) { }
if ast.IsOperatorToken(token) { }
if ast.IsBinaryOperator(token) { }
if ast.IsUnaryOperator(token) { }
if ast.IsAssignmentOperator(token) { }
```

### Token Relationships

```go
// Check token ordering
if ast.IsTokenBefore(tokenA, tokenB) { }
if ast.IsTokenAfter(tokenA, tokenB) { }
if ast.TokensOverlap(tokenA, tokenB) { }
```

### Token Utilities

```go
// Sort tokens by position
ast.SortTokens(tokens)

// Get token span
span := ast.TokenSpan(token)

// Get token text
text := ast.GetTokenText(token)

// Get whitespace around token
before := ast.GetWhitespaceBefore(tokens, token)
after := ast.GetWhitespaceAfter(tokens, token)
```

## Best Practices

### 1. Use Type Guards for Safety

```go
// Good
if ast.IsIdentifier(node) {
    id := node.(*ast.Identifier)
    // Work with id safely
}

// Also good - with assertion helper
if id, ok := ast.AsExpression(node); ok {
    // Work with expression
}
```

### 2. Leverage Context-Aware Traversal

```go
// When you need parent information
ast.TraverseWithContext(root, func(node ast.Node, ctx *ast.TraverseContext) bool {
    if ast.IsReturnStatement(node) && ast.IsArrowFunction(ctx.Parent) {
        // Check returns in arrow functions
    }
    return true
})
```

### 3. Use Specific Finders

```go
// Instead of traversing manually
// Good
identifiers := ast.FindByType(root, "Identifier")

// Instead of implementing custom search
// Good
firstAsync := ast.FindFirst(root, ast.IsAsyncFunction)
```

### 4. Batch Operations

```go
// Get all needed information in one traversal
type Collector struct {
    functions []ast.Node
    classes   []ast.Node
    imports   []ast.Node
}

collector := &Collector{}
ast.Traverse(root, func(node ast.Node) bool {
    if ast.IsFunction(node) {
        collector.functions = append(collector.functions, node)
    }
    if ast.IsClass(node) {
        collector.classes = append(collector.classes, node)
    }
    if ast.IsImport(node) {
        collector.imports = append(collector.imports, node)
    }
    return true
})
```

## Performance Considerations

1. **Visitor Keys**: O(1) map lookups for visitor key access
2. **Traversal**: Single-pass traversal algorithms
3. **Finding**: Early termination with `FindFirst`
4. **Caching**: Consider caching frequently accessed relationships (parent, ancestors)

## References

- [ESTree Specification](https://github.com/estree/estree)
- [TypeScript ESLint AST Spec](https://typescript-eslint.io/packages/typescript-estree/ast-spec/)
- [TypeScript ESLint Visitor Keys](https://github.com/typescript-eslint/typescript-eslint/tree/main/packages/visitor-keys)
- [ESRecurse](https://github.com/estools/esrecurse)

## Contributing

When adding new utilities:

1. Follow existing naming conventions
2. Add comprehensive tests
3. Document with examples
4. Update this README
5. Ensure null safety (check for nil nodes)
6. Use visitor keys for traversal
