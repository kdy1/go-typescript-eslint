# Examples

This directory contains examples demonstrating how to use the go-typescript-eslint library.

## Running Examples

Each example can be run using `go run`:

```bash
go run examples/<example_name>/main.go
```

## Available Examples

### basic_parse
Demonstrates basic parsing of TypeScript code into an AST.

```bash
go run examples/basic_parse/main.go
```

### with_types
Shows how to use type-aware parsing with TypeScript program services.

```bash
go run examples/with_types/main.go
```

### custom_visitor
Demonstrates how to traverse the AST using a custom visitor pattern.

```bash
go run examples/custom_visitor/main.go
```

### error_handling
Shows proper error handling and recovery during parsing.

```bash
go run examples/error_handling/main.go
```

### jsx_tsx
Demonstrates parsing JSX/TSX syntax.

```bash
go run examples/jsx_tsx/main.go
```

## Directory Structure

```
examples/
├── README.md              # This file
├── doc.go                 # Package documentation
├── basic_parse/           # Basic parsing example
├── with_types/            # Type-aware parsing example
├── custom_visitor/        # AST traversal example
├── error_handling/        # Error handling example
└── jsx_tsx/               # JSX/TSX parsing example
```

## Writing Your Own Examples

To add a new example:

1. Create a new directory under `examples/`
2. Add a `main.go` file with your example code
3. Update this README with a description
4. Ensure your example is self-contained and documented
