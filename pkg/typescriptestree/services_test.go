package typescriptestree

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/program"
)

func TestNewParserServices(t *testing.T) {
	tmpDir := t.TempDir()
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")
	if err := os.WriteFile(tsconfigPath, []byte(`{"compilerOptions":{}}`), 0644); err != nil {
		t.Fatalf("Failed to write tsconfig: %v", err)
	}

	opts := &program.ProgramOptions{
		TSConfigPath: tsconfigPath,
		RootDir:      tmpDir,
	}

	prog, err := program.CreateProgram(opts)
	if err != nil {
		t.Fatalf("Failed to create program: %v", err)
	}

	services := NewParserServices(prog)

	if services == nil {
		t.Fatal("Expected non-nil services")
	}

	if services.Program != prog {
		t.Error("Expected services to reference the program")
	}

	if services.ESTreeNodeToTSNodeMap == nil {
		t.Error("Expected non-nil ESTreeNodeToTSNodeMap")
	}

	if services.TSNodeToESTreeNodeMap == nil {
		t.Error("Expected non-nil TSNodeToESTreeNodeMap")
	}
}

func TestAddNodeMapping(t *testing.T) {
	services := NewParserServices(nil)

	// Create test nodes
	estreeNode := &ast.Identifier{
		Name: "test",
	}
	tsNode := "typescript-node-placeholder"

	// Add mapping
	services.AddNodeMapping(estreeNode, tsNode)

	// Verify bidirectional mapping
	retrievedTS, ok := services.GetTSNodeForESTreeNode(estreeNode)
	if !ok {
		t.Error("Expected to find TS node for ESTree node")
	}

	if retrievedTS != tsNode {
		t.Error("Retrieved TS node does not match")
	}

	retrievedESTree, ok := services.GetESTreeNodeForTSNode(tsNode)
	if !ok {
		t.Error("Expected to find ESTree node for TS node")
	}

	if retrievedESTree != estreeNode {
		t.Error("Retrieved ESTree node does not match")
	}
}

func TestHasNodeMapping(t *testing.T) {
	services := NewParserServices(nil)

	estreeNode := &ast.Identifier{Name: "test"}

	// Initially no mapping
	if services.HasNodeMapping(estreeNode) {
		t.Error("Expected no mapping initially")
	}

	// Add mapping
	services.AddNodeMapping(estreeNode, "ts-node")

	// Now should have mapping
	if !services.HasNodeMapping(estreeNode) {
		t.Error("Expected to have mapping after adding")
	}
}

func TestClearNodeMappings(t *testing.T) {
	services := NewParserServices(nil)

	// Add some mappings
	node1 := &ast.Identifier{Name: "test1"}
	node2 := &ast.Identifier{Name: "test2"}

	services.AddNodeMapping(node1, "ts1")
	services.AddNodeMapping(node2, "ts2")

	// Verify they exist
	if !services.HasNodeMapping(node1) || !services.HasNodeMapping(node2) {
		t.Error("Expected mappings to exist")
	}

	// Clear mappings
	services.ClearNodeMappings()

	// Verify they're gone
	if services.HasNodeMapping(node1) || services.HasNodeMapping(node2) {
		t.Error("Expected mappings to be cleared")
	}

	if len(services.ESTreeNodeToTSNodeMap) != 0 {
		t.Error("Expected ESTreeNodeToTSNodeMap to be empty")
	}

	if len(services.TSNodeToESTreeNodeMap) != 0 {
		t.Error("Expected TSNodeToESTreeNodeMap to be empty")
	}
}

func TestGetCompilerOptions(t *testing.T) {
	tmpDir := t.TempDir()
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")
	tsconfigContent := `{
		"compilerOptions": {
			"target": "ES2020",
			"strict": true
		}
	}`
	if err := os.WriteFile(tsconfigPath, []byte(tsconfigContent), 0644); err != nil {
		t.Fatalf("Failed to write tsconfig: %v", err)
	}

	opts := &program.ProgramOptions{
		TSConfigPath: tsconfigPath,
		RootDir:      tmpDir,
	}

	prog, err := program.CreateProgram(opts)
	if err != nil {
		t.Fatalf("Failed to create program: %v", err)
	}

	services := NewParserServices(prog)

	compilerOpts := services.GetCompilerOptions()
	if compilerOpts == nil {
		t.Fatal("Expected non-nil compiler options")
	}

	if compilerOpts.Target != "ES2020" {
		t.Errorf("Expected target ES2020, got %s", compilerOpts.Target)
	}

	if compilerOpts.Strict == nil || !*compilerOpts.Strict {
		t.Error("Expected strict to be true")
	}
}

func TestGetCompilerOptionsNilProgram(t *testing.T) {
	services := NewParserServices(nil)

	compilerOpts := services.GetCompilerOptions()
	if compilerOpts != nil {
		t.Error("Expected nil compiler options when program is nil")
	}
}

func TestGetTypeChecker(t *testing.T) {
	tmpDir := t.TempDir()
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")
	if err := os.WriteFile(tsconfigPath, []byte(`{"compilerOptions":{}}`), 0644); err != nil {
		t.Fatalf("Failed to write tsconfig: %v", err)
	}

	opts := &program.ProgramOptions{
		TSConfigPath: tsconfigPath,
		RootDir:      tmpDir,
	}

	prog, err := program.CreateProgram(opts)
	if err != nil {
		t.Fatalf("Failed to create program: %v", err)
	}

	services := NewParserServices(prog)

	_, err = services.GetTypeChecker()
	if err != ErrNotImplemented {
		t.Errorf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestGetTypeAtLocation(t *testing.T) {
	tmpDir := t.TempDir()
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")
	if err := os.WriteFile(tsconfigPath, []byte(`{"compilerOptions":{}}`), 0644); err != nil {
		t.Fatalf("Failed to write tsconfig: %v", err)
	}

	opts := &program.ProgramOptions{
		TSConfigPath: tsconfigPath,
		RootDir:      tmpDir,
	}

	prog, err := program.CreateProgram(opts)
	if err != nil {
		t.Fatalf("Failed to create program: %v", err)
	}

	services := NewParserServices(prog)

	node := &ast.Identifier{Name: "test"}
	_, err = services.GetTypeAtLocation(node)
	if err != ErrNotImplemented {
		t.Errorf("Expected ErrNotImplemented, got %v", err)
	}
}

func TestGetSymbolAtLocation(t *testing.T) {
	tmpDir := t.TempDir()
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")
	if err := os.WriteFile(tsconfigPath, []byte(`{"compilerOptions":{}}`), 0644); err != nil {
		t.Fatalf("Failed to write tsconfig: %v", err)
	}

	opts := &program.ProgramOptions{
		TSConfigPath: tsconfigPath,
		RootDir:      tmpDir,
	}

	prog, err := program.CreateProgram(opts)
	if err != nil {
		t.Fatalf("Failed to create program: %v", err)
	}

	services := NewParserServices(prog)

	node := &ast.Identifier{Name: "test"}
	_, err = services.GetSymbolAtLocation(node)
	if err != ErrNotImplemented {
		t.Errorf("Expected ErrNotImplemented, got %v", err)
	}
}
