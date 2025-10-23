package typescriptestree

import (
	"fmt"
	"sync"

	"github.com/kdy1/go-typescript-eslint/internal/ast"
	"github.com/kdy1/go-typescript-eslint/internal/program"
)

// ParserServices provides access to TypeScript type information and node mappings.
// This enables type-aware linting rules that need access to the TypeScript type checker.
type ParserServices struct {
	// Program is the TypeScript program instance providing type information
	Program *program.Program

	// ESTreeNodeToTSNodeMap maps ESTree nodes to TypeScript AST nodes.
	// TypeScript nodes are represented as interface{} until we have
	// a full TypeScript AST implementation.
	ESTreeNodeToTSNodeMap map[ast.Node]interface{}

	// TSNodeToESTreeNodeMap maps TypeScript AST nodes to ESTree nodes
	TSNodeToESTreeNodeMap map[interface{}]ast.Node

	// mu protects concurrent access to the maps
	mu sync.RWMutex
}

// NewParserServices creates a new ParserServices instance.
func NewParserServices(prog *program.Program) *ParserServices {
	return &ParserServices{
		Program:               prog,
		ESTreeNodeToTSNodeMap: make(map[ast.Node]interface{}),
		TSNodeToESTreeNodeMap: make(map[interface{}]ast.Node),
	}
}

// AddNodeMapping adds a bidirectional mapping between an ESTree node and a TypeScript node.
func (s *ParserServices) AddNodeMapping(estreeNode ast.Node, tsNode interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.ESTreeNodeToTSNodeMap[estreeNode] = tsNode
	s.TSNodeToESTreeNodeMap[tsNode] = estreeNode
}

// GetTSNodeForESTreeNode retrieves the TypeScript node for a given ESTree node.
func (s *ParserServices) GetTSNodeForESTreeNode(estreeNode ast.Node) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tsNode, ok := s.ESTreeNodeToTSNodeMap[estreeNode]
	return tsNode, ok
}

// GetESTreeNodeForTSNode retrieves the ESTree node for a given TypeScript node.
func (s *ParserServices) GetESTreeNodeForTSNode(tsNode interface{}) (ast.Node, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	estreeNode, ok := s.TSNodeToESTreeNodeMap[tsNode]
	return estreeNode, ok
}

// HasNodeMapping checks if a mapping exists for the given ESTree node.
func (s *ParserServices) HasNodeMapping(estreeNode ast.Node) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.ESTreeNodeToTSNodeMap[estreeNode]
	return ok
}

// ClearNodeMappings removes all node mappings.
func (s *ParserServices) ClearNodeMappings() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.ESTreeNodeToTSNodeMap = make(map[ast.Node]interface{})
	s.TSNodeToESTreeNodeMap = make(map[interface{}]ast.Node)
}

// GetCompilerOptions returns the TypeScript compiler options from the program.
func (s *ParserServices) GetCompilerOptions() *program.CompilerOptions {
	if s.Program == nil {
		return nil
	}
	return s.Program.GetCompilerOptions()
}

// GetTypeChecker returns a type checker instance for type-aware operations.
// This is a placeholder for future TypeScript type checker integration.
func (s *ParserServices) GetTypeChecker() (interface{}, error) {
	if s.Program == nil {
		return nil, fmt.Errorf("no program available for type checking")
	}
	// TODO: Integrate with TypeScript type checker or microsoft/typescript-go
	return nil, ErrNotImplemented
}

// GetTypeAtLocation gets the type of a node at a specific location.
// This is a placeholder for future type information access.
func (s *ParserServices) GetTypeAtLocation(node ast.Node) (interface{}, error) {
	if s.Program == nil {
		return nil, fmt.Errorf("no program available for type checking")
	}
	// TODO: Implement type information retrieval
	return nil, ErrNotImplemented
}

// GetSymbolAtLocation gets the symbol for a node at a specific location.
// This is a placeholder for future symbol information access.
func (s *ParserServices) GetSymbolAtLocation(node ast.Node) (interface{}, error) {
	if s.Program == nil {
		return nil, fmt.Errorf("no program available for symbol resolution")
	}
	// TODO: Implement symbol information retrieval
	return nil, ErrNotImplemented
}

// Services is the updated name for ParserServices to match typescript-estree.
// Both names are supported for compatibility.
type Services = ParserServices
