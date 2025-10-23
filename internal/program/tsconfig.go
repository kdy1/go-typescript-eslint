package program

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CompilerOptions represents TypeScript compiler options from tsconfig.json.
type CompilerOptions struct {
	Target                             string              `json:"target,omitempty"`
	Module                             string              `json:"module,omitempty"`
	Lib                                []string            `json:"lib,omitempty"`
	AllowJs                            *bool               `json:"allowJs,omitempty"`
	CheckJs                            *bool               `json:"checkJs,omitempty"`
	JSX                                string              `json:"jsx,omitempty"`
	Declaration                        *bool               `json:"declaration,omitempty"`
	DeclarationMap                     *bool               `json:"declarationMap,omitempty"`
	SourceMap                          *bool               `json:"sourceMap,omitempty"`
	OutFile                            string              `json:"outFile,omitempty"`
	OutDir                             string              `json:"outDir,omitempty"`
	RootDir                            string              `json:"rootDir,omitempty"`
	RemoveComments                     *bool               `json:"removeComments,omitempty"`
	NoEmit                             *bool               `json:"noEmit,omitempty"`
	Strict                             *bool               `json:"strict,omitempty"`
	NoImplicitAny                      *bool               `json:"noImplicitAny,omitempty"`
	StrictNullChecks                   *bool               `json:"strictNullChecks,omitempty"`
	StrictFunctionTypes                *bool               `json:"strictFunctionTypes,omitempty"`
	StrictBindCallApply                *bool               `json:"strictBindCallApply,omitempty"`
	StrictPropertyInitialization       *bool               `json:"strictPropertyInitialization,omitempty"`
	NoImplicitThis                     *bool               `json:"noImplicitThis,omitempty"`
	AlwaysStrict                       *bool               `json:"alwaysStrict,omitempty"`
	NoUnusedLocals                     *bool               `json:"noUnusedLocals,omitempty"`
	NoUnusedParameters                 *bool               `json:"noUnusedParameters,omitempty"`
	NoImplicitReturns                  *bool               `json:"noImplicitReturns,omitempty"`
	NoFallthroughCasesInSwitch         *bool               `json:"noFallthroughCasesInSwitch,omitempty"`
	NoUncheckedIndexedAccess           *bool               `json:"noUncheckedIndexedAccess,omitempty"`
	NoImplicitOverride                 *bool               `json:"noImplicitOverride,omitempty"`
	NoPropertyAccessFromIndexSignature *bool               `json:"noPropertyAccessFromIndexSignature,omitempty"`
	ModuleResolution                   string              `json:"moduleResolution,omitempty"`
	BaseUrl                            string              `json:"baseUrl,omitempty"`
	Paths                              map[string][]string `json:"paths,omitempty"`
	RootDirs                           []string            `json:"rootDirs,omitempty"`
	TypeRoots                          []string            `json:"typeRoots,omitempty"`
	Types                              []string            `json:"types,omitempty"`
	AllowSyntheticDefaultImports       *bool               `json:"allowSyntheticDefaultImports,omitempty"`
	EsModuleInterop                    *bool               `json:"esModuleInterop,omitempty"`
	PreserveSymlinks                   *bool               `json:"preserveSymlinks,omitempty"`
	ForceConsistentCasingInFileNames   *bool               `json:"forceConsistentCasingInFileNames,omitempty"`
	SkipLibCheck                       *bool               `json:"skipLibCheck,omitempty"`
	ResolveJsonModule                  *bool               `json:"resolveJsonModule,omitempty"`
	IsolatedModules                    *bool               `json:"isolatedModules,omitempty"`
	ExperimentalDecorators             *bool               `json:"experimentalDecorators,omitempty"`
	EmitDecoratorMetadata              *bool               `json:"emitDecoratorMetadata,omitempty"`
	Incremental                        *bool               `json:"incremental,omitempty"`
	TsBuildInfoFile                    string              `json:"tsBuildInfoFile,omitempty"`
}

// ProjectReference represents a TypeScript project reference.
type ProjectReference struct {
	Path      string `json:"path"`
	Prepend   *bool  `json:"prepend,omitempty"`
	Composite *bool  `json:"composite,omitempty"`
}

// TSConfig represents a parsed TypeScript configuration file.
type TSConfig struct {
	// Extends specifies the base configuration file to inherit from.
	Extends string `json:"extends,omitempty"`

	// CompilerOptions contains TypeScript compiler options.
	CompilerOptions CompilerOptions `json:"compilerOptions,omitempty"`

	// Files is a list of files to include in the program.
	Files []string `json:"files,omitempty"`

	// Include specifies patterns of files to include.
	Include []string `json:"include,omitempty"`

	// Exclude specifies patterns of files to exclude.
	Exclude []string `json:"exclude,omitempty"`

	// References lists project references for composite projects.
	References []ProjectReference `json:"references,omitempty"`

	// CompileOnSave indicates whether to compile on save in IDEs.
	CompileOnSave *bool `json:"compileOnSave,omitempty"`

	// TypeAcquisition configures automatic type acquisition.
	TypeAcquisition map[string]interface{} `json:"typeAcquisition,omitempty"`

	// path is the absolute path to this tsconfig file (internal use).
	path string
}

// ParseTSConfig parses a tsconfig.json file from the given path.
func ParseTSConfig(path string) (*TSConfig, error) {
	// Normalize path to absolute
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve absolute path for %s: %w", path, err)
	}

	// Read the file
	data, err := os.ReadFile(absPath) // #nosec G304 -- absPath is validated as an absolute path
	if err != nil {
		return nil, fmt.Errorf("failed to read tsconfig file %s: %w", absPath, err)
	}

	// Parse JSON (supports comments through lenient parsing)
	var config TSConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse tsconfig file %s: %w", absPath, err)
	}

	config.path = absPath

	return &config, nil
}

// ResolveTSConfig resolves a tsconfig.json file with inheritance.
// It handles the "extends" field and merges configurations.
func ResolveTSConfig(path string) (*TSConfig, error) {
	config, err := ParseTSConfig(path)
	if err != nil {
		return nil, err
	}

	// If no extends, return as-is
	if config.Extends == "" {
		return config, nil
	}

	// Resolve the base configuration
	baseConfigPath := resolveExtendsPath(config.path, config.Extends)
	baseConfig, err := ResolveTSConfig(baseConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve extended config %s: %w", config.Extends, err)
	}

	// Merge configurations (child overrides parent)
	merged := mergeConfigs(baseConfig, config)
	return merged, nil
}

// resolveExtendsPath resolves the path to an extended configuration file.
func resolveExtendsPath(currentConfigPath, extendsPath string) string {
	// If it's a package name (not starting with . or /), it's from node_modules
	if !strings.HasPrefix(extendsPath, ".") && !strings.HasPrefix(extendsPath, "/") {
		// For simplicity, assume it's in node_modules relative to current config
		dir := filepath.Dir(currentConfigPath)
		return filepath.Join(dir, "node_modules", extendsPath, "tsconfig.json")
	}

	// Relative path
	dir := filepath.Dir(currentConfigPath)
	resolved := filepath.Join(dir, extendsPath)

	// If no extension, try .json
	if filepath.Ext(resolved) == "" {
		if _, err := os.Stat(resolved + ".json"); err == nil {
			return resolved + ".json"
		}
		// Try as directory with tsconfig.json
		if stat, err := os.Stat(resolved); err == nil && stat.IsDir() {
			return filepath.Join(resolved, "tsconfig.json")
		}
	}

	return resolved
}

// mergeConfigs merges two TSConfig objects, with child taking precedence.
func mergeConfigs(parent, child *TSConfig) *TSConfig {
	merged := &TSConfig{
		path: child.path,
	}

	// Merge compiler options
	merged.CompilerOptions = mergeCompilerOptions(&parent.CompilerOptions, &child.CompilerOptions)

	// Files: child overrides if present
	if len(child.Files) > 0 {
		merged.Files = child.Files
	} else {
		merged.Files = parent.Files
	}

	// Include: child overrides if present
	if len(child.Include) > 0 {
		merged.Include = child.Include
	} else {
		merged.Include = parent.Include
	}

	// Exclude: child overrides if present
	if len(child.Exclude) > 0 {
		merged.Exclude = child.Exclude
	} else {
		merged.Exclude = parent.Exclude
	}

	// References: combine
	merged.References = append(parent.References, child.References...)

	// CompileOnSave: child overrides if present
	if child.CompileOnSave != nil {
		merged.CompileOnSave = child.CompileOnSave
	} else {
		merged.CompileOnSave = parent.CompileOnSave
	}

	return merged
}

// mergeCompilerOptions merges compiler options with child taking precedence.
//
//nolint:gocognit // This function is a straightforward merger of many fields
func mergeCompilerOptions(parent, child *CompilerOptions) CompilerOptions {
	merged := *parent

	// Merge each field - child overrides parent if set
	if child.Target != "" {
		merged.Target = child.Target
	}
	if child.Module != "" {
		merged.Module = child.Module
	}
	if len(child.Lib) > 0 {
		merged.Lib = child.Lib
	}
	if child.AllowJs != nil {
		merged.AllowJs = child.AllowJs
	}
	if child.CheckJs != nil {
		merged.CheckJs = child.CheckJs
	}
	if child.JSX != "" {
		merged.JSX = child.JSX
	}
	if child.Declaration != nil {
		merged.Declaration = child.Declaration
	}
	if child.DeclarationMap != nil {
		merged.DeclarationMap = child.DeclarationMap
	}
	if child.SourceMap != nil {
		merged.SourceMap = child.SourceMap
	}
	if child.OutFile != "" {
		merged.OutFile = child.OutFile
	}
	if child.OutDir != "" {
		merged.OutDir = child.OutDir
	}
	if child.RootDir != "" {
		merged.RootDir = child.RootDir
	}
	if child.RemoveComments != nil {
		merged.RemoveComments = child.RemoveComments
	}
	if child.NoEmit != nil {
		merged.NoEmit = child.NoEmit
	}
	if child.Strict != nil {
		merged.Strict = child.Strict
	}
	if child.NoImplicitAny != nil {
		merged.NoImplicitAny = child.NoImplicitAny
	}
	if child.StrictNullChecks != nil {
		merged.StrictNullChecks = child.StrictNullChecks
	}
	if child.StrictFunctionTypes != nil {
		merged.StrictFunctionTypes = child.StrictFunctionTypes
	}
	if child.StrictBindCallApply != nil {
		merged.StrictBindCallApply = child.StrictBindCallApply
	}
	if child.StrictPropertyInitialization != nil {
		merged.StrictPropertyInitialization = child.StrictPropertyInitialization
	}
	if child.NoImplicitThis != nil {
		merged.NoImplicitThis = child.NoImplicitThis
	}
	if child.AlwaysStrict != nil {
		merged.AlwaysStrict = child.AlwaysStrict
	}
	if child.NoUnusedLocals != nil {
		merged.NoUnusedLocals = child.NoUnusedLocals
	}
	if child.NoUnusedParameters != nil {
		merged.NoUnusedParameters = child.NoUnusedParameters
	}
	if child.NoImplicitReturns != nil {
		merged.NoImplicitReturns = child.NoImplicitReturns
	}
	if child.NoFallthroughCasesInSwitch != nil {
		merged.NoFallthroughCasesInSwitch = child.NoFallthroughCasesInSwitch
	}
	if child.ModuleResolution != "" {
		merged.ModuleResolution = child.ModuleResolution
	}
	if child.BaseUrl != "" {
		merged.BaseUrl = child.BaseUrl
	}
	if len(child.Paths) > 0 {
		merged.Paths = child.Paths
	}
	if len(child.RootDirs) > 0 {
		merged.RootDirs = child.RootDirs
	}
	if len(child.TypeRoots) > 0 {
		merged.TypeRoots = child.TypeRoots
	}
	if len(child.Types) > 0 {
		merged.Types = child.Types
	}
	if child.AllowSyntheticDefaultImports != nil {
		merged.AllowSyntheticDefaultImports = child.AllowSyntheticDefaultImports
	}
	if child.EsModuleInterop != nil {
		merged.EsModuleInterop = child.EsModuleInterop
	}
	if child.SkipLibCheck != nil {
		merged.SkipLibCheck = child.SkipLibCheck
	}
	if child.ResolveJsonModule != nil {
		merged.ResolveJsonModule = child.ResolveJsonModule
	}
	if child.IsolatedModules != nil {
		merged.IsolatedModules = child.IsolatedModules
	}
	if child.ExperimentalDecorators != nil {
		merged.ExperimentalDecorators = child.ExperimentalDecorators
	}
	if child.EmitDecoratorMetadata != nil {
		merged.EmitDecoratorMetadata = child.EmitDecoratorMetadata
	}

	return merged
}

// GetConfigDir returns the directory containing this tsconfig file.
func (c *TSConfig) GetConfigDir() string {
	return filepath.Dir(c.path)
}

// GetPath returns the absolute path to this tsconfig file.
func (c *TSConfig) GetPath() string {
	return c.path
}
