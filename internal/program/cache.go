package program

import (
	"path/filepath"
	"sync"
	"time"
)

// ProgramCache provides caching of TypeScript programs for performance optimization.
type ProgramCache struct {
	// programs maps tsconfig paths to cached programs
	programs map[string]*CachedProgram

	// mu protects concurrent access
	mu sync.RWMutex

	// maxAge is the maximum age of cached programs before they expire
	maxAge time.Duration
}

// CachedProgram wraps a program with cache metadata.
type CachedProgram struct {
	Program   *Program
	CachedAt  time.Time
	TSConfigPath string
}

// NewProgramCache creates a new program cache with the specified max age.
// If maxAge is 0, programs are cached indefinitely.
func NewProgramCache(maxAge time.Duration) *ProgramCache {
	return &ProgramCache{
		programs: make(map[string]*CachedProgram),
		maxAge:   maxAge,
	}
}

// Get retrieves a cached program for the given tsconfig path.
// Returns nil if the program is not in cache or has expired.
func (c *ProgramCache) Get(tsconfigPath string) *Program {
	c.mu.RLock()
	defer c.mu.RUnlock()

	absPath, err := filepath.Abs(tsconfigPath)
	if err != nil {
		return nil
	}

	cached, ok := c.programs[absPath]
	if !ok {
		return nil
	}

	// Check if expired
	if c.maxAge > 0 && time.Since(cached.CachedAt) > c.maxAge {
		// Expired, will be cleaned up later
		return nil
	}

	return cached.Program
}

// Set caches a program for the given tsconfig path.
func (c *ProgramCache) Set(tsconfigPath string, program *Program) {
	c.mu.Lock()
	defer c.mu.Unlock()

	absPath, err := filepath.Abs(tsconfigPath)
	if err != nil {
		return
	}

	c.programs[absPath] = &CachedProgram{
		Program:      program,
		CachedAt:     time.Now(),
		TSConfigPath: absPath,
	}
}

// Clear removes all cached programs.
func (c *ProgramCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.programs = make(map[string]*CachedProgram)
}

// CleanExpired removes expired programs from the cache.
func (c *ProgramCache) CleanExpired() {
	if c.maxAge == 0 {
		// No expiration
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for path, cached := range c.programs {
		if now.Sub(cached.CachedAt) > c.maxAge {
			delete(c.programs, path)
		}
	}
}

// Size returns the number of cached programs.
func (c *ProgramCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.programs)
}

// GetOrCreate retrieves a cached program or creates a new one if not found.
func (c *ProgramCache) GetOrCreate(opts *ProgramOptions) (*Program, error) {
	// Try to get from cache first
	if opts.TSConfigPath != "" {
		if cached := c.Get(opts.TSConfigPath); cached != nil {
			return cached, nil
		}
	}

	// Create new program
	program, err := CreateProgram(opts)
	if err != nil {
		return nil, err
	}

	// Cache it
	if opts.TSConfigPath != "" {
		c.Set(opts.TSConfigPath, program)
	}

	return program, nil
}

// GlobalCache is the default global program cache with 5 minute expiration.
var GlobalCache = NewProgramCache(5 * time.Minute)
