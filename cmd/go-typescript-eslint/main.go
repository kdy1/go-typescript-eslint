package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/kdy1/go-typescript-eslint/pkg/typescriptestree"
)

var (
	formatFlag   = flag.String("format", "json", "Output format: json, pretty")
	tokensFlag   = flag.Bool("tokens", false, "Include tokens in output")
	commentsFlag = flag.Bool("comments", false, "Include comments in output")
	locFlag      = flag.Bool("loc", false, "Include location information")
	rangeFlag    = flag.Bool("range", false, "Include range information")
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <file>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	filename := flag.Arg(0)
	// #nosec G304 - The filename is provided by the user via CLI argument, this is expected behavior for a parser tool
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	options := typescriptestree.NewBuilder().
		WithSourceType(typescriptestree.SourceTypeModule).
		WithLoc(*locFlag).
		WithRange(*rangeFlag).
		WithComment(*commentsFlag).
		WithTokens(*tokensFlag).
		WithFilePath(filename).
		MustBuild()

	result, err := typescriptestree.Parse(string(content), options)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse error: %v\n", err)
		os.Exit(1)
	}

	var output []byte
	switch *formatFlag {
	case "json":
		output, err = json.Marshal(result)
	case "pretty":
		output, err = json.MarshalIndent(result, "", "  ")
	default:
		fmt.Fprintf(os.Stderr, "Unknown format: %s\n", *formatFlag)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding output: %v\n", err)
		os.Exit(1)
	}

	if _, err := os.Stdout.WriteString(string(output) + "\n"); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}
}
