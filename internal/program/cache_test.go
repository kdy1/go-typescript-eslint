package program

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestProgramCache(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a tsconfig
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")
	if err := os.WriteFile(tsconfigPath, []byte(`{"compilerOptions":{}}`), 0600); err != nil {
		t.Fatalf("Failed to write tsconfig: %v", err)
	}

	// Create cache
	cache := NewProgramCache(5 * time.Minute)

	// Initially empty
	if cached := cache.Get(tsconfigPath); cached != nil {
		t.Error("Expected cache to be empty initially")
	}

	// Create and cache a program
	opts := &ProgramOptions{
		TSConfigPath: tsconfigPath,
		RootDir:      tmpDir,
	}

	program, err := CreateProgram(opts)
	if err != nil {
		t.Fatalf("Failed to create program: %v", err)
	}

	cache.Set(tsconfigPath, program)

	// Should now retrieve it
	cached := cache.Get(tsconfigPath)
	if cached == nil {
		t.Error("Expected to retrieve cached program")
	}

	if cached != program {
		t.Error("Retrieved program does not match cached program")
	}

	// Check size
	if cache.Size() != 1 {
		t.Errorf("Expected cache size 1, got %d", cache.Size())
	}
}

func TestProgramCacheExpiration(t *testing.T) {
	tmpDir := t.TempDir()
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")
	if err := os.WriteFile(tsconfigPath, []byte(`{"compilerOptions":{}}`), 0600); err != nil {
		t.Fatalf("Failed to write tsconfig: %v", err)
	}

	// Create cache with very short expiration
	cache := NewProgramCache(100 * time.Millisecond)

	opts := &ProgramOptions{
		TSConfigPath: tsconfigPath,
		RootDir:      tmpDir,
	}

	program, err := CreateProgram(opts)
	if err != nil {
		t.Fatalf("Failed to create program: %v", err)
	}

	cache.Set(tsconfigPath, program)

	// Should retrieve immediately
	if cached := cache.Get(tsconfigPath); cached == nil {
		t.Error("Expected to retrieve cached program immediately")
	}

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)

	// Should now be expired
	if cached := cache.Get(tsconfigPath); cached != nil {
		t.Error("Expected cached program to be expired")
	}
}

func TestProgramCacheClear(t *testing.T) {
	tmpDir := t.TempDir()
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")
	if err := os.WriteFile(tsconfigPath, []byte(`{"compilerOptions":{}}`), 0600); err != nil {
		t.Fatalf("Failed to write tsconfig: %v", err)
	}

	cache := NewProgramCache(5 * time.Minute)

	opts := &ProgramOptions{
		TSConfigPath: tsconfigPath,
		RootDir:      tmpDir,
	}

	program, err := CreateProgram(opts)
	if err != nil {
		t.Fatalf("Failed to create program: %v", err)
	}

	cache.Set(tsconfigPath, program)

	if cache.Size() != 1 {
		t.Error("Expected cache to have 1 entry")
	}

	// Clear cache
	cache.Clear()

	if cache.Size() != 0 {
		t.Error("Expected cache to be empty after clear")
	}

	if cached := cache.Get(tsconfigPath); cached != nil {
		t.Error("Expected no cached program after clear")
	}
}

func TestProgramCacheGetOrCreate(t *testing.T) {
	tmpDir := t.TempDir()
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")
	if err := os.WriteFile(tsconfigPath, []byte(`{"compilerOptions":{}}`), 0600); err != nil {
		t.Fatalf("Failed to write tsconfig: %v", err)
	}

	cache := NewProgramCache(5 * time.Minute)

	opts := &ProgramOptions{
		TSConfigPath: tsconfigPath,
		RootDir:      tmpDir,
	}

	// First call should create
	program1, err := cache.GetOrCreate(opts)
	if err != nil {
		t.Fatalf("Failed to get or create program: %v", err)
	}

	if program1 == nil {
		t.Fatal("Expected non-nil program")
	}

	// Second call should retrieve from cache
	program2, err := cache.GetOrCreate(opts)
	if err != nil {
		t.Fatalf("Failed to get or create program: %v", err)
	}

	if program1 != program2 {
		t.Error("Expected to retrieve same program instance from cache")
	}
}

func TestCleanExpired(t *testing.T) {
	tmpDir := t.TempDir()
	tsconfigPath := filepath.Join(tmpDir, "tsconfig.json")
	if err := os.WriteFile(tsconfigPath, []byte(`{"compilerOptions":{}}`), 0600); err != nil {
		t.Fatalf("Failed to write tsconfig: %v", err)
	}

	cache := NewProgramCache(100 * time.Millisecond)

	opts := &ProgramOptions{
		TSConfigPath: tsconfigPath,
		RootDir:      tmpDir,
	}

	program, err := CreateProgram(opts)
	if err != nil {
		t.Fatalf("Failed to create program: %v", err)
	}

	cache.Set(tsconfigPath, program)

	if cache.Size() != 1 {
		t.Error("Expected cache to have 1 entry")
	}

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)

	// Clean expired
	cache.CleanExpired()

	if cache.Size() != 0 {
		t.Error("Expected cache to be empty after cleaning expired entries")
	}
}
