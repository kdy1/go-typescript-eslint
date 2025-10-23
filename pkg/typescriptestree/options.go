package typescriptestree

import (
	"fmt"
	"path/filepath"
	"strings"
)

// SourceType specifies the type of source code being parsed.
type SourceType string

const (
	// SourceTypeScript indicates the code should be parsed as a script (no imports/exports).
	SourceTypeScript SourceType = "script"

	// SourceTypeModule indicates the code should be parsed as a module (with imports/exports).
	SourceTypeModule SourceType = "module"
)

// JSDocParsingMode controls how JSDoc comments are parsed.
type JSDocParsingMode string

const (
	// JSDocParsingModeAll parses all JSDoc comments.
	JSDocParsingModeAll JSDocParsingMode = "all"

	// JSDocParsingModeNone skips JSDoc parsing entirely.
	JSDocParsingModeNone JSDocParsingMode = "none"

	// JSDocParsingModeTypeInfo only parses JSDoc comments needed for type information.
	JSDocParsingModeTypeInfo JSDocParsingMode = "type-info"
)

// DebugLevel specifies which modules to enable debugging for.
type DebugLevel []string

// LoggerFn is a function that receives log messages.
// If set to nil, logging is disabled.
type LoggerFn func(message string)

// CacheDurationSeconds specifies cache lifetime in seconds.
type CacheDurationSeconds int

// CacheLifetime configures cache behavior for the parser.
type CacheLifetime struct {
	// Glob is the cache lifetime for glob searches in seconds.
	Glob *CacheDurationSeconds `json:"glob,omitempty"`
}

// ParseOptions configures the behavior of the TypeScript parser.
// It matches the options available in @typescript-eslint/typescript-estree.
type ParseOptions struct {
	// SourceType specifies whether the code is a "script" or "module".
	// Default: "script"
	SourceType SourceType `json:"sourceType,omitempty"`

	// AllowInvalidAST prevents the parser from throwing an error if it receives
	// an invalid AST from TypeScript. This can be useful for parsing malformed code.
	// Default: false
	AllowInvalidAST bool `json:"allowInvalidAST,omitempty"`

	// Comment indicates whether to create a top-level comments array containing
	// all comments in the source code.
	// Default: false
	Comment bool `json:"comment,omitempty"`

	// DebugLevel enables debugging for specific modules. Pass module names to enable
	// detailed logging for those modules.
	// Default: empty (no debugging)
	DebugLevel DebugLevel `json:"debugLevel,omitempty"`

	// ErrorOnUnknownASTType causes the parser to throw an error if it encounters
	// an unknown AST node type (useful for catching unsupported syntax).
	// Default: false
	ErrorOnUnknownASTType bool `json:"errorOnUnknownASTType,omitempty"`

	// FilePath is the absolute (or relative to cwd) path to the file being parsed.
	// This is used for error messages and determining file type (.ts vs .tsx).
	// Default: ""
	FilePath string `json:"filePath,omitempty"`

	// JSDocParsingMode controls how JSDoc comments are parsed and included in the AST.
	// Options: "all", "none", "type-info"
	// Default: "all"
	JSDocParsingMode JSDocParsingMode `json:"jsDocParsingMode,omitempty"`

	// JSX enables parsing of JSX syntax. This is automatically enabled for .tsx files.
	// Default: false (true for .tsx files)
	JSX bool `json:"jsx,omitempty"`

	// Loc indicates whether to add line/column location information to AST nodes.
	// When enabled, nodes will have a `loc` property with start/end positions.
	// Default: false
	Loc bool `json:"loc,omitempty"`

	// LoggerFn allows overriding the logging function used by the parser.
	// Set to nil to disable logging.
	// Default: logs to stderr
	LoggerFn LoggerFn `json:"-"`

	// Range indicates whether to add [start, end] range information to AST nodes.
	// When enabled, nodes will have a `range` property with character offsets.
	// Default: false
	Range bool `json:"range,omitempty"`

	// Tokens indicates whether to create a top-level array containing all tokens
	// found in the source code.
	// Default: false
	Tokens bool `json:"tokens,omitempty"`

	// SuppressDeprecatedPropertyWarnings prevents warnings about deprecated AST
	// properties from being logged.
	// Default: false
	SuppressDeprecatedPropertyWarnings bool `json:"suppressDeprecatedPropertyWarnings,omitempty"`
}

// ParseAndGenerateServicesOptions extends ParseOptions with additional options
// for generating TypeScript language services (required for type-aware linting).
type ParseAndGenerateServicesOptions struct {
	ParseOptions

	// CacheLifetime controls the internal cache expiry for various operations.
	// This can improve performance for repeated parsing operations.
	CacheLifetime *CacheLifetime `json:"cacheLifetime,omitempty"`

	// DisallowAutomaticSingleRunInference disables the performance heuristic that
	// infers whether the parser is being used for a single run or multiple runs.
	// Default: false
	DisallowAutomaticSingleRunInference bool `json:"disallowAutomaticSingleRunInference,omitempty"`

	// ErrorOnTypeScriptSyntacticAndSemanticIssues causes the parser to throw an
	// error if TypeScript reports any syntactic or semantic issues with the code.
	// Default: false
	ErrorOnTypeScriptSyntacticAndSemanticIssues bool `json:"errorOnTypeScriptSyntacticAndSemanticIssues,omitempty"`

	// ExtraFileExtensions specifies file extensions that should be treated as
	// TypeScript files beyond the standard .ts, .tsx, .mts, .cts extensions.
	// Each extension should include the leading dot (e.g., ".vue").
	// Default: empty
	ExtraFileExtensions []string `json:"extraFileExtensions,omitempty"`

	// PreserveNodeMaps controls whether to preserve TypeScript's internal AST node
	// maps for use in language services. Required for certain type-checking operations.
	// Default: true
	PreserveNodeMaps *bool `json:"preserveNodeMaps,omitempty"`

	// Project specifies one or more paths to TypeScript configuration files (tsconfig.json)
	// or directories containing them. This enables type-aware parsing.
	// Supports glob patterns.
	// Default: empty
	Project []string `json:"project,omitempty"`

	// ProjectFolderIgnoreList specifies folder names or glob patterns to ignore when
	// searching for project files using the `project` option.
	// Default: ["**/node_modules/**"]
	ProjectFolderIgnoreList []string `json:"projectFolderIgnoreList,omitempty"`

	// ProjectService enables use of TypeScript's project service for managing
	// multiple projects and shared state across parses.
	// Default: false
	ProjectService bool `json:"projectService,omitempty"`

	// TSConfigRootDir specifies the root directory for relative tsconfig paths.
	// When set, paths in the `project` option are resolved relative to this directory.
	// Default: current working directory
	TSConfigRootDir string `json:"tsconfigRootDir,omitempty"`

	// Programs provides pre-created TypeScript Program instances to use instead of
	// creating new ones. This is an advanced option for performance optimization.
	// Default: nil
	Programs []interface{} `json:"-"`

	// WarnOnUnsupportedTypeScriptVersion controls whether to warn when using an
	// unsupported TypeScript version.
	// Default: true
	WarnOnUnsupportedTypeScriptVersion *bool `json:"warnOnUnsupportedTypeScriptVersion,omitempty"`
}

// NewParseOptions creates a new ParseOptions with sensible defaults matching
// @typescript-eslint/typescript-estree behavior.
func NewParseOptions() *ParseOptions {
	return &ParseOptions{
		SourceType:                         SourceTypeScript,
		AllowInvalidAST:                    false,
		Comment:                            false,
		DebugLevel:                         nil,
		ErrorOnUnknownASTType:              false,
		FilePath:                           "",
		JSDocParsingMode:                   JSDocParsingModeAll,
		JSX:                                false,
		Loc:                                false,
		LoggerFn:                           defaultLogger,
		Range:                              false,
		Tokens:                             false,
		SuppressDeprecatedPropertyWarnings: false,
	}
}

// NewParseAndGenerateServicesOptions creates a new ParseAndGenerateServicesOptions
// with sensible defaults.
func NewParseAndGenerateServicesOptions() *ParseAndGenerateServicesOptions {
	preserveNodeMaps := true
	warnOnUnsupported := true

	return &ParseAndGenerateServicesOptions{
		ParseOptions:                                *NewParseOptions(),
		CacheLifetime:                               nil,
		DisallowAutomaticSingleRunInference:         false,
		ErrorOnTypeScriptSyntacticAndSemanticIssues: false,
		ExtraFileExtensions:                         nil,
		PreserveNodeMaps:                            &preserveNodeMaps,
		Project:                                     nil,
		ProjectFolderIgnoreList:                     []string{"**/node_modules/**"},
		ProjectService:                              false,
		TSConfigRootDir:                             "",
		Programs:                                    nil,
		WarnOnUnsupportedTypeScriptVersion:          &warnOnUnsupported,
	}
}

// defaultLogger is the default logging function that writes to stderr.
func defaultLogger(message string) {
	// Intentionally ignore error - logging to stderr is best-effort
	//nolint:errcheck // Logging to stderr is best-effort, errors are intentionally ignored
	_, _ = fmt.Fprintln(getStderr(), message)
}

// getStderr returns the standard error writer (extracted for testability).
var getStderr = func() interface{ Write([]byte) (int, error) } {
	// This will be implemented to return os.Stderr
	return &nullWriter{}
}

type nullWriter struct{}

func (*nullWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

// Validate checks if the options are valid and returns an error if not.
func (o *ParseOptions) Validate() error {
	// Validate SourceType
	if o.SourceType != "" && o.SourceType != SourceTypeScript && o.SourceType != SourceTypeModule {
		return fmt.Errorf("invalid sourceType: must be 'script' or 'module', got %q", o.SourceType)
	}

	// Validate JSDocParsingMode
	if o.JSDocParsingMode != "" &&
		o.JSDocParsingMode != JSDocParsingModeAll &&
		o.JSDocParsingMode != JSDocParsingModeNone &&
		o.JSDocParsingMode != JSDocParsingModeTypeInfo {
		return fmt.Errorf("invalid jsDocParsingMode: must be 'all', 'none', or 'type-info', got %q", o.JSDocParsingMode)
	}

	return nil
}

// Validate checks if the service options are valid and returns an error if not.
func (o *ParseAndGenerateServicesOptions) Validate() error {
	// First validate base ParseOptions
	if err := o.ParseOptions.Validate(); err != nil {
		return err
	}

	// Validate ExtraFileExtensions
	for _, ext := range o.ExtraFileExtensions {
		if !strings.HasPrefix(ext, ".") {
			return fmt.Errorf("invalid extra file extension %q: must start with a dot", ext)
		}
	}

	// Validate that Project and ProjectService aren't both used
	if len(o.Project) > 0 && o.ProjectService {
		return fmt.Errorf("cannot use both 'project' and 'projectService' options")
	}

	return nil
}

// InferJSXFromFilePath automatically enables JSX parsing for .tsx files.
// This should be called after setting FilePath to match typescript-estree behavior.
func (o *ParseOptions) InferJSXFromFilePath() {
	if o.FilePath != "" {
		ext := strings.ToLower(filepath.Ext(o.FilePath))
		if ext == ".tsx" || ext == ".jsx" {
			o.JSX = true
		}
	}
}

// ParseOptionsBuilder provides a fluent API for constructing ParseOptions.
type ParseOptionsBuilder struct {
	opts *ParseOptions
}

// NewBuilder creates a new ParseOptionsBuilder with default values.
func NewBuilder() *ParseOptionsBuilder {
	return &ParseOptionsBuilder{
		opts: NewParseOptions(),
	}
}

// WithSourceType sets the source type (script or module).
func (b *ParseOptionsBuilder) WithSourceType(sourceType SourceType) *ParseOptionsBuilder {
	b.opts.SourceType = sourceType
	return b
}

// WithAllowInvalidAST enables or disables parsing of invalid ASTs.
func (b *ParseOptionsBuilder) WithAllowInvalidAST(allow bool) *ParseOptionsBuilder {
	b.opts.AllowInvalidAST = allow
	return b
}

// WithComment enables or disables comment collection.
func (b *ParseOptionsBuilder) WithComment(comment bool) *ParseOptionsBuilder {
	b.opts.Comment = comment
	return b
}

// WithDebugLevel sets the debug level for specific modules.
func (b *ParseOptionsBuilder) WithDebugLevel(modules ...string) *ParseOptionsBuilder {
	b.opts.DebugLevel = modules
	return b
}

// WithErrorOnUnknownASTType enables or disables errors on unknown AST types.
func (b *ParseOptionsBuilder) WithErrorOnUnknownASTType(errorOn bool) *ParseOptionsBuilder {
	b.opts.ErrorOnUnknownASTType = errorOn
	return b
}

// WithFilePath sets the file path and automatically infers JSX if needed.
func (b *ParseOptionsBuilder) WithFilePath(path string) *ParseOptionsBuilder {
	b.opts.FilePath = path
	b.opts.InferJSXFromFilePath()
	return b
}

// WithJSDocParsingMode sets the JSDoc parsing mode.
func (b *ParseOptionsBuilder) WithJSDocParsingMode(mode JSDocParsingMode) *ParseOptionsBuilder {
	b.opts.JSDocParsingMode = mode
	return b
}

// WithJSX enables or disables JSX parsing.
func (b *ParseOptionsBuilder) WithJSX(jsx bool) *ParseOptionsBuilder {
	b.opts.JSX = jsx
	return b
}

// WithLoc enables or disables location information.
func (b *ParseOptionsBuilder) WithLoc(loc bool) *ParseOptionsBuilder {
	b.opts.Loc = loc
	return b
}

// WithLoggerFn sets the logger function.
func (b *ParseOptionsBuilder) WithLoggerFn(fn LoggerFn) *ParseOptionsBuilder {
	b.opts.LoggerFn = fn
	return b
}

// WithRange enables or disables range information.
func (b *ParseOptionsBuilder) WithRange(rang bool) *ParseOptionsBuilder {
	b.opts.Range = rang
	return b
}

// WithTokens enables or disables token collection.
func (b *ParseOptionsBuilder) WithTokens(tokens bool) *ParseOptionsBuilder {
	b.opts.Tokens = tokens
	return b
}

// WithSuppressDeprecatedPropertyWarnings enables or disables deprecated property warnings.
func (b *ParseOptionsBuilder) WithSuppressDeprecatedPropertyWarnings(suppress bool) *ParseOptionsBuilder {
	b.opts.SuppressDeprecatedPropertyWarnings = suppress
	return b
}

// Build returns the constructed ParseOptions after validation.
func (b *ParseOptionsBuilder) Build() (*ParseOptions, error) {
	if err := b.opts.Validate(); err != nil {
		return nil, err
	}
	return b.opts, nil
}

// MustBuild returns the constructed ParseOptions or panics if validation fails.
func (b *ParseOptionsBuilder) MustBuild() *ParseOptions {
	opts, err := b.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to build ParseOptions: %v", err))
	}
	return opts
}

// ParseAndGenerateServicesOptionsBuilder provides a fluent API for constructing
// ParseAndGenerateServicesOptions.
type ParseAndGenerateServicesOptionsBuilder struct {
	opts *ParseAndGenerateServicesOptions
}

// NewServicesBuilder creates a new ParseAndGenerateServicesOptionsBuilder with default values.
func NewServicesBuilder() *ParseAndGenerateServicesOptionsBuilder {
	return &ParseAndGenerateServicesOptionsBuilder{
		opts: NewParseAndGenerateServicesOptions(),
	}
}

// WithParseOptions sets the base parse options.
func (b *ParseAndGenerateServicesOptionsBuilder) WithParseOptions(
	opts *ParseOptions,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.ParseOptions = *opts
	return b
}

// Base ParseOptions builder methods for convenience

// WithSourceType sets the source type (script or module).
func (b *ParseAndGenerateServicesOptionsBuilder) WithSourceType(
	sourceType SourceType,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.SourceType = sourceType
	return b
}

// WithAllowInvalidAST enables or disables parsing of invalid ASTs.
func (b *ParseAndGenerateServicesOptionsBuilder) WithAllowInvalidAST(
	allow bool,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.AllowInvalidAST = allow
	return b
}

// WithComment enables or disables comment collection.
func (b *ParseAndGenerateServicesOptionsBuilder) WithComment(comment bool) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.Comment = comment
	return b
}

// WithDebugLevel sets the debug level for specific modules.
func (b *ParseAndGenerateServicesOptionsBuilder) WithDebugLevel(
	modules ...string,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.DebugLevel = modules
	return b
}

// WithErrorOnUnknownASTType enables or disables errors on unknown AST types.
func (b *ParseAndGenerateServicesOptionsBuilder) WithErrorOnUnknownASTType(
	errorOn bool,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.ErrorOnUnknownASTType = errorOn
	return b
}

// WithFilePath sets the file path and automatically infers JSX if needed.
func (b *ParseAndGenerateServicesOptionsBuilder) WithFilePath(path string) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.FilePath = path
	b.opts.InferJSXFromFilePath()
	return b
}

// WithJSDocParsingMode sets the JSDoc parsing mode.
func (b *ParseAndGenerateServicesOptionsBuilder) WithJSDocParsingMode(
	mode JSDocParsingMode,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.JSDocParsingMode = mode
	return b
}

// WithJSX enables or disables JSX parsing.
func (b *ParseAndGenerateServicesOptionsBuilder) WithJSX(jsx bool) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.JSX = jsx
	return b
}

// WithLoc enables or disables location information.
func (b *ParseAndGenerateServicesOptionsBuilder) WithLoc(loc bool) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.Loc = loc
	return b
}

// WithLoggerFn sets the logger function.
func (b *ParseAndGenerateServicesOptionsBuilder) WithLoggerFn(fn LoggerFn) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.LoggerFn = fn
	return b
}

// WithRange enables or disables range information.
func (b *ParseAndGenerateServicesOptionsBuilder) WithRange(rang bool) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.Range = rang
	return b
}

// WithTokens enables or disables token collection.
func (b *ParseAndGenerateServicesOptionsBuilder) WithTokens(tokens bool) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.Tokens = tokens
	return b
}

// WithSuppressDeprecatedPropertyWarnings enables or disables deprecated property warnings.
func (b *ParseAndGenerateServicesOptionsBuilder) WithSuppressDeprecatedPropertyWarnings(
	suppress bool,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.SuppressDeprecatedPropertyWarnings = suppress
	return b
}

// WithCacheLifetime sets the cache lifetime configuration.
func (b *ParseAndGenerateServicesOptionsBuilder) WithCacheLifetime(
	lifetime *CacheLifetime,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.CacheLifetime = lifetime
	return b
}

// WithGlobCacheLifetime sets the glob cache lifetime in seconds.
func (b *ParseAndGenerateServicesOptionsBuilder) WithGlobCacheLifetime(
	seconds int,
) *ParseAndGenerateServicesOptionsBuilder {
	if b.opts.CacheLifetime == nil {
		b.opts.CacheLifetime = &CacheLifetime{}
	}
	duration := CacheDurationSeconds(seconds)
	b.opts.CacheLifetime.Glob = &duration
	return b
}

// WithDisallowAutomaticSingleRunInference disables automatic single run inference.
func (b *ParseAndGenerateServicesOptionsBuilder) WithDisallowAutomaticSingleRunInference(
	disallow bool,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.DisallowAutomaticSingleRunInference = disallow
	return b
}

// WithErrorOnTypeScriptSyntacticAndSemanticIssues enables errors on TypeScript issues.
func (b *ParseAndGenerateServicesOptionsBuilder) WithErrorOnTypeScriptSyntacticAndSemanticIssues(
	errorOn bool,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.ErrorOnTypeScriptSyntacticAndSemanticIssues = errorOn
	return b
}

// WithExtraFileExtensions sets additional file extensions to treat as TypeScript.
func (b *ParseAndGenerateServicesOptionsBuilder) WithExtraFileExtensions(
	extensions ...string,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.ExtraFileExtensions = extensions
	return b
}

// WithPreserveNodeMaps sets whether to preserve TypeScript node maps.
func (b *ParseAndGenerateServicesOptionsBuilder) WithPreserveNodeMaps(
	preserve bool,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.PreserveNodeMaps = &preserve
	return b
}

// WithProject sets the TypeScript project configuration paths.
func (b *ParseAndGenerateServicesOptionsBuilder) WithProject(paths ...string) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.Project = paths
	return b
}

// WithProjectFolderIgnoreList sets folders to ignore when searching for projects.
func (b *ParseAndGenerateServicesOptionsBuilder) WithProjectFolderIgnoreList(
	patterns ...string,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.ProjectFolderIgnoreList = patterns
	return b
}

// WithProjectService enables or disables the TypeScript project service.
func (b *ParseAndGenerateServicesOptionsBuilder) WithProjectService(
	enable bool,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.ProjectService = enable
	return b
}

// WithTSConfigRootDir sets the root directory for tsconfig paths.
func (b *ParseAndGenerateServicesOptionsBuilder) WithTSConfigRootDir(
	dir string,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.TSConfigRootDir = dir
	return b
}

// WithWarnOnUnsupportedTypeScriptVersion sets whether to warn on unsupported TypeScript versions.
func (b *ParseAndGenerateServicesOptionsBuilder) WithWarnOnUnsupportedTypeScriptVersion(
	warn bool,
) *ParseAndGenerateServicesOptionsBuilder {
	b.opts.WarnOnUnsupportedTypeScriptVersion = &warn
	return b
}

// Build returns the constructed ParseAndGenerateServicesOptions after validation.
func (b *ParseAndGenerateServicesOptionsBuilder) Build() (*ParseAndGenerateServicesOptions, error) {
	if err := b.opts.Validate(); err != nil {
		return nil, err
	}
	return b.opts, nil
}

// MustBuild returns the constructed ParseAndGenerateServicesOptions or panics if validation fails.
func (b *ParseAndGenerateServicesOptionsBuilder) MustBuild() *ParseAndGenerateServicesOptions {
	opts, err := b.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to build ParseAndGenerateServicesOptions: %v", err))
	}
	return opts
}
