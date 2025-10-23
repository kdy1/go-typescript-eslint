package typescriptestree

import (
	"github.com/kdy1/go-typescript-eslint/internal/program"
)

// ClearProgramCache clears all cached TypeScript programs.
// This is intended exclusively for test isolation between lint operations.
//
// Example usage:
//
//	// Before running tests
//	typescriptestree.ClearProgramCache()
//
// In production code, you typically don't need to call this function.
// The cache automatically expires entries based on the configured lifetime.
func ClearProgramCache() {
	program.GlobalCache.Clear()
}

// ClearDefaultProjectMatchedFiles clears the tracked project-matched file records.
// This is intended for test isolation and cleanup.
//
// Example usage:
//
//	// Clean up after tests
//	typescriptestree.ClearDefaultProjectMatchedFiles()
//
// This is primarily for internal use and testing. In most cases, you don't
// need to call this function manually.
func ClearDefaultProjectMatchedFiles() {
	// This is a placeholder for future implementation
	// In typescript-estree, this tracks which files were matched by default projects
	// For now, we don't have this tracking, but we provide the API for compatibility
}
