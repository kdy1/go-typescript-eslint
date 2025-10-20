// Package main provides the command-line interface for go-typescript-eslint.
//
// This is the main entry point for the go-typescript-eslint tool, which
// can be used to parse TypeScript files and output their AST representation.
//
// # Usage
//
//	go-typescript-eslint [options] <file>
//
// # Options
//
//	-format string
//	    Output format: json, pretty (default "json")
//	-tokens
//	    Include tokens in output
//	-comments
//	    Include comments in output
//	-loc
//	    Include location information
//	-range
//	    Include range information
//
// # Examples
//
//	# Parse a TypeScript file and output JSON
//	go-typescript-eslint file.ts
//
//	# Parse with all metadata
//	go-typescript-eslint -tokens -comments -loc -range file.ts
//
//	# Pretty print the AST
//	go-typescript-eslint -format pretty file.ts
package main
