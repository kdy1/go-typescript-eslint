package program

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseTSConfig(t *testing.T) {
	// Create temporary directory for test files
	tmpDir := t.TempDir()

	// Create a simple tsconfig.json
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")
	tsconfigContent := `{
		"compilerOptions": {
			"target": "ES2020",
			"module": "commonjs",
			"strict": true,
			"esModuleInterop": true
		},
		"include": ["src/**/*"],
		"exclude": ["node_modules", "dist"]
	}`

	if err := os.WriteFile(tsconfigPath, []byte(tsconfigContent), 0600); err != nil {
		t.Fatalf("Failed to write test tsconfig: %v", err)
	}

	// Test parsing
	config, err := ParseTSConfig(tsconfigPath)
	if err != nil {
		t.Fatalf("Failed to parse tsconfig: %v", err)
	}

	// Verify parsed values
	if config.CompilerOptions.Target != "ES2020" {
		t.Errorf("Expected target ES2020, got %s", config.CompilerOptions.Target)
	}

	if config.CompilerOptions.Module != "commonjs" {
		t.Errorf("Expected module commonjs, got %s", config.CompilerOptions.Module)
	}

	if config.CompilerOptions.Strict == nil || !*config.CompilerOptions.Strict {
		t.Error("Expected strict to be true")
	}

	if len(config.Include) != 1 || config.Include[0] != "src/**/*" {
		t.Errorf("Expected include [src/**/*], got %v", config.Include)
	}

	if len(config.Exclude) != 2 {
		t.Errorf("Expected 2 exclude patterns, got %d", len(config.Exclude))
	}
}

func TestResolveTSConfigWithExtends(t *testing.T) {
	// Create temporary directory for test files
	tmpDir := t.TempDir()

	// Create base config
	baseConfigPath := filepath.Join(tmpDir, "tsconfig.base.json")
	baseConfigContent := `{
		"compilerOptions": {
			"target": "ES2015",
			"strict": true,
			"moduleResolution": "node"
		}
	}`

	if err := os.WriteFile(baseConfigPath, []byte(baseConfigContent), 0600); err != nil {
		t.Fatalf("Failed to write base config: %v", err)
	}

	// Create child config that extends base
	childConfigPath := filepath.Join(tmpDir, "tsconfig.json")
	childConfigContent := `{
		"extends": "./tsconfig.base.json",
		"compilerOptions": {
			"target": "ES2020",
			"module": "esnext"
		},
		"include": ["src/**/*"]
	}`

	if err := os.WriteFile(childConfigPath, []byte(childConfigContent), 0600); err != nil {
		t.Fatalf("Failed to write child config: %v", err)
	}

	// Resolve with inheritance
	config, err := ResolveTSConfig(childConfigPath)
	if err != nil {
		t.Fatalf("Failed to resolve tsconfig: %v", err)
	}

	// Child should override target
	if config.CompilerOptions.Target != "ES2020" {
		t.Errorf("Expected target ES2020 (child override), got %s", config.CompilerOptions.Target)
	}

	// Child should override module
	if config.CompilerOptions.Module != "esnext" {
		t.Errorf("Expected module esnext (child override), got %s", config.CompilerOptions.Module)
	}

	// Child should inherit strict from base
	if config.CompilerOptions.Strict == nil || !*config.CompilerOptions.Strict {
		t.Error("Expected strict true (inherited from base)")
	}

	// Child should inherit moduleResolution from base
	if config.CompilerOptions.ModuleResolution != "node" {
		t.Errorf("Expected moduleResolution node (inherited), got %s", config.CompilerOptions.ModuleResolution)
	}

	// Child should have its own include
	if len(config.Include) != 1 || config.Include[0] != "src/**/*" {
		t.Errorf("Expected include [src/**/*], got %v", config.Include)
	}
}

func TestGetConfigDir(t *testing.T) {
	tmpDir := t.TempDir()
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")

	config := &TSConfig{
		path: tsconfigPath,
	}

	dir := config.GetConfigDir()
	if dir != tmpDir {
		t.Errorf("Expected directory %s, got %s", tmpDir, dir)
	}
}

func TestGetPath(t *testing.T) {
	tmpDir := t.TempDir()
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")

	config := &TSConfig{
		path: tsconfigPath,
	}

	path := config.GetPath()
	if path != tsconfigPath {
		t.Errorf("Expected path %s, got %s", tsconfigPath, path)
	}
}
