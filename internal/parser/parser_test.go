package parser

import (
	"testing"

	"github.com/kdy1/go-typescript-eslint/internal/ast"
)

func TestParserBasic(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "empty program",
			input:   "",
			wantErr: false,
		},
		{
			name:    "simple variable declaration",
			input:   "const x = 42;",
			wantErr: false,
		},
		{
			name:    "simple function declaration",
			input:   "function foo() { return 42; }",
			wantErr: false,
		},
		{
			name:    "simple class declaration",
			input:   "class Foo { constructor() {} }",
			wantErr: false,
		},
		{
			name:    "arrow function",
			input:   "const foo = (x) => x + 1;",
			wantErr: false,
		},
		{
			name:    "if statement",
			input:   "if (x > 0) { console.log('positive'); }",
			wantErr: false,
		},
		{
			name:    "for loop",
			input:   "for (let i = 0; i < 10; i++) { console.log(i); }",
			wantErr: false,
		},
		{
			name:    "while loop",
			input:   "while (x > 0) { x--; }",
			wantErr: false,
		},
		{
			name:    "switch statement",
			input:   "switch (x) { case 1: break; default: break; }",
			wantErr: false,
		},
		{
			name:    "try-catch",
			input:   "try { foo(); } catch (e) { console.error(e); }",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := New(tt.input)
			node, err := parser.Parse()

			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && node == nil {
				t.Errorf("Parse() returned nil node")
				return
			}

			if !tt.wantErr {
				program, ok := node.(*ast.Program)
				if !ok {
					t.Errorf("Parse() did not return a Program node")
					return
				}

				if program.NodeType != ast.NodeTypeProgram {
					t.Errorf("Program node type = %v, want %v", program.NodeType, ast.NodeTypeProgram)
				}
			}
		})
	}
}

func TestParserExpressions(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "binary expression",
			input:   "x + y;",
			wantErr: false,
		},
		{
			name:    "complex binary expression",
			input:   "x + y * z - a / b;",
			wantErr: false,
		},
		{
			name:    "logical expression",
			input:   "x && y || z;",
			wantErr: false,
		},
		{
			name:    "conditional expression",
			input:   "x > 0 ? 'positive' : 'negative';",
			wantErr: false,
		},
		{
			name:    "unary expression",
			input:   "!x;",
			wantErr: false,
		},
		{
			name:    "update expression",
			input:   "x++;",
			wantErr: false,
		},
		{
			name:    "member expression",
			input:   "obj.prop;",
			wantErr: false,
		},
		{
			name:    "computed member expression",
			input:   "obj[key];",
			wantErr: false,
		},
		{
			name:    "call expression",
			input:   "foo(1, 2, 3);",
			wantErr: false,
		},
		{
			name:    "new expression",
			input:   "new Foo(x, y);",
			wantErr: false,
		},
		{
			name:    "array expression",
			input:   "[1, 2, 3];",
			wantErr: false,
		},
		{
			name:    "object expression",
			input:   "{ a: 1, b: 2 };",
			wantErr: false,
		},
		{
			name:    "template literal",
			input:   "`hello ${world}`;",
			wantErr: false,
		},
		{
			name:    "spread operator",
			input:   "[...arr];",
			wantErr: false,
		},
		{
			name:    "optional chaining",
			input:   "obj?.prop?.method();",
			wantErr: false,
		},
		{
			name:    "nullish coalescing",
			input:   "x ?? y;",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := New(tt.input)
			node, err := parser.Parse()

			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && node == nil {
				t.Errorf("Parse() returned nil node")
			}
		})
	}
}

func TestParserTypeScript(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "variable with type annotation",
			input:   "const x: number = 42;",
			wantErr: false,
		},
		{
			name:    "function with type annotation",
			input:   "function foo(x: number): string { return 'hello'; }",
			wantErr: false,
		},
		{
			name:    "interface declaration",
			input:   "interface Foo { x: number; y: string; }",
			wantErr: false,
		},
		{
			name:    "type alias",
			input:   "type Foo = { x: number; y: string; };",
			wantErr: false,
		},
		{
			name:    "enum declaration",
			input:   "enum Color { Red, Green, Blue }",
			wantErr: false,
		},
		{
			name:    "class with type parameters",
			input:   "class Foo<T> { value: T; }",
			wantErr: false,
		},
		{
			name:    "union type",
			input:   "type Foo = string | number;",
			wantErr: false,
		},
		{
			name:    "intersection type",
			input:   "type Foo = A & B;",
			wantErr: false,
		},
		{
			name:    "tuple type",
			input:   "type Foo = [string, number];",
			wantErr: false,
		},
		{
			name:    "generic function",
			input:   "function foo<T>(x: T): T { return x; }",
			wantErr: false,
		},
		{
			name:    "as expression",
			input:   "const x = foo as string;",
			wantErr: false,
		},
		{
			name:    "non-null assertion",
			input:   "const x = foo!;",
			wantErr: false,
		},
		{
			name:    "namespace declaration",
			input:   "namespace Foo { export const x = 42; }",
			wantErr: false,
		},
		{
			name:    "import type",
			input:   "import type { Foo } from './foo';",
			wantErr: false,
		},
		{
			name:    "export type",
			input:   "export type { Foo };",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := New(tt.input)
			node, err := parser.Parse()

			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && node == nil {
				t.Errorf("Parse() returned nil node")
			}
		})
	}
}

func TestParserModules(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "import default",
			input:   "import foo from './foo';",
			wantErr: false,
		},
		{
			name:    "import named",
			input:   "import { foo, bar } from './foo';",
			wantErr: false,
		},
		{
			name:    "import namespace",
			input:   "import * as foo from './foo';",
			wantErr: false,
		},
		{
			name:    "import mixed",
			input:   "import foo, { bar } from './foo';",
			wantErr: false,
		},
		{
			name:    "export default",
			input:   "export default foo;",
			wantErr: false,
		},
		{
			name:    "export named",
			input:   "export { foo, bar };",
			wantErr: false,
		},
		{
			name:    "export declaration",
			input:   "export const foo = 42;",
			wantErr: false,
		},
		{
			name:    "export all",
			input:   "export * from './foo';",
			wantErr: false,
		},
		{
			name:    "export all as",
			input:   "export * as foo from './foo';",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := New(tt.input)
			node, err := parser.Parse()

			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && node == nil {
				t.Errorf("Parse() returned nil node")
			}
		})
	}
}

func TestParserDestructuring(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "array destructuring",
			input:   "const [a, b, c] = arr;",
			wantErr: false,
		},
		{
			name:    "object destructuring",
			input:   "const { a, b, c } = obj;",
			wantErr: false,
		},
		{
			name:    "nested destructuring",
			input:   "const { a: { b, c } } = obj;",
			wantErr: false,
		},
		{
			name:    "rest element",
			input:   "const [a, ...rest] = arr;",
			wantErr: false,
		},
		{
			name:    "default values",
			input:   "const { a = 1, b = 2 } = obj;",
			wantErr: false,
		},
		{
			name:    "function parameter destructuring",
			input:   "function foo({ a, b }) { return a + b; }",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := New(tt.input)
			node, err := parser.Parse()

			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && node == nil {
				t.Errorf("Parse() returned nil node")
			}
		})
	}
}

func TestParserAsync(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "async function",
			input:   "async function foo() { return await bar(); }",
			wantErr: false,
		},
		{
			name:    "async arrow function",
			input:   "const foo = async () => await bar();",
			wantErr: false,
		},
		{
			name:    "for await",
			input:   "for await (const item of items) { console.log(item); }",
			wantErr: false,
		},
		{
			name:    "generator function",
			input:   "function* gen() { yield 1; yield 2; }",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := New(tt.input)
			node, err := parser.Parse()

			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && node == nil {
				t.Errorf("Parse() returned nil node")
			}
		})
	}
}
