package converter

import (
	"testing"

	"github.com/kdy1/go-typescript-eslint/internal/ast"
)

// TestNewConverter tests creating a new converter.
func TestNewConverter(t *testing.T) {
	source := "const x: number = 42;"
	opts := &Options{
		PreserveNodeMaps: true,
	}

	converter := NewConverter(source, opts)

	if converter == nil {
		t.Fatal("NewConverter returned nil")
	}

	if converter.source != source {
		t.Errorf("Expected source %q, got %q", source, converter.source)
	}

	if !converter.options.PreserveNodeMaps {
		t.Error("Expected PreserveNodeMaps to be true")
	}

	if converter.esTreeNodeToTSNodeMap == nil {
		t.Error("Expected esTreeNodeToTSNodeMap to be initialized")
	}

	if converter.tsNodeToESTreeNodeMap == nil {
		t.Error("Expected tsNodeToESTreeNodeMap to be initialized")
	}
}

// TestConvertProgram tests converting a Program node.
func TestConvertProgram(t *testing.T) {
	source := "const x = 1;"
	converter := NewConverter(source, nil)

	// Create a simple program
	originalProgram := &ast.Program{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeProgram.String(),
		},
		SourceType: "module",
		Body:       []ast.Statement{},
		Comments:   []ast.Comment{},
		Tokens:     []ast.Token{},
	}

	result := converter.ConvertProgram(originalProgram)

	if result == nil {
		t.Fatal("ConvertProgram returned nil")
	}

	if result.NodeType != ast.NodeTypeProgram.String() {
		t.Errorf("Expected node type %q, got %q", ast.NodeTypeProgram.String(), result.NodeType)
	}

	if result.SourceType != "module" {
		t.Errorf("Expected source type 'module', got %q", result.SourceType)
	}
}

// TestConvertIdentifier tests converting an Identifier node.
func TestConvertIdentifier(t *testing.T) {
	source := "x"
	converter := NewConverter(source, nil)

	original := &ast.Identifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeIdentifier.String(),
		},
		Name: "x",
	}

	result := converter.convertIdentifier(original)

	if result == nil {
		t.Fatal("convertIdentifier returned nil")
	}

	if result.Name != "x" {
		t.Errorf("Expected name 'x', got %q", result.Name)
	}

	if result.NodeType != ast.NodeTypeIdentifier.String() {
		t.Errorf("Expected node type %q, got %q", ast.NodeTypeIdentifier.String(), result.NodeType)
	}
}

// TestConvertLiteral tests converting a Literal node.
func TestConvertLiteral(t *testing.T) {
	source := "42"
	converter := NewConverter(source, nil)

	original := &ast.Literal{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeLiteral.String(),
		},
		Value: 42,
		Raw:   "42",
	}

	result := converter.convertLiteral(original)

	if result == nil {
		t.Fatal("convertLiteral returned nil")
	}

	if result.Value != 42 {
		t.Errorf("Expected value 42, got %v", result.Value)
	}

	if result.Raw != "42" {
		t.Errorf("Expected raw '42', got %q", result.Raw)
	}
}

// TestConvertBinaryExpression tests converting a BinaryExpression node.
func TestConvertBinaryExpression(t *testing.T) {
	source := "1 + 2"
	converter := NewConverter(source, nil)

	left := &ast.Literal{
		BaseNode: ast.BaseNode{NodeType: ast.NodeTypeLiteral.String()},
		Value:    1,
		Raw:      "1",
	}

	right := &ast.Literal{
		BaseNode: ast.BaseNode{NodeType: ast.NodeTypeLiteral.String()},
		Value:    2,
		Raw:      "2",
	}

	original := &ast.BinaryExpression{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeBinaryExpression.String(),
		},
		Operator: "+",
		Left:     left,
		Right:    right,
	}

	result := converter.convertBinaryExpression(original)

	if result == nil {
		t.Fatal("convertBinaryExpression returned nil")
	}

	if result.Operator != "+" {
		t.Errorf("Expected operator '+', got %q", result.Operator)
	}

	if result.Left == nil {
		t.Error("Expected left operand to be converted")
	}

	if result.Right == nil {
		t.Error("Expected right operand to be converted")
	}
}

// TestConvertFunctionDeclaration tests converting a FunctionDeclaration node.
func TestConvertFunctionDeclaration(t *testing.T) {
	source := "function foo() {}"
	converter := NewConverter(source, nil)

	original := &ast.FunctionDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeFunctionDeclaration.String(),
		},
		ID: &ast.Identifier{
			BaseNode: ast.BaseNode{NodeType: ast.NodeTypeIdentifier.String()},
			Name:     "foo",
		},
		Params: []ast.Pattern{},
		Body: &ast.BlockStatement{
			BaseNode: ast.BaseNode{NodeType: ast.NodeTypeBlockStatement.String()},
			Body:     []ast.Statement{},
		},
	}

	result := converter.convertFunctionDeclaration(original)

	if result == nil {
		t.Fatal("convertFunctionDeclaration returned nil")
	}

	if result.ID == nil || result.ID.Name != "foo" {
		t.Error("Expected function id to be 'foo'")
	}

	if result.Body == nil {
		t.Error("Expected function body to be converted")
	}
}

// TestConvertVariableDeclaration tests converting a VariableDeclaration node.
func TestConvertVariableDeclaration(t *testing.T) {
	source := "const x = 1;"
	converter := NewConverter(source, nil)

	original := &ast.VariableDeclaration{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeVariableDeclaration.String(),
		},
		Kind:         "const",
		Declarations: []ast.VariableDeclarator{},
	}

	result := converter.convertVariableDeclaration(original)

	if result == nil {
		t.Fatal("convertVariableDeclaration returned nil")
	}

	if result.Kind != "const" {
		t.Errorf("Expected kind 'const', got %q", result.Kind)
	}
}

// TestConvertTSTypeAnnotation tests converting a TSTypeAnnotation node.
func TestConvertTSTypeAnnotation(t *testing.T) {
	source := ": number"
	converter := NewConverter(source, nil)

	original := &ast.TSTypeAnnotation{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeTSTypeAnnotation.String(),
		},
		TypeAnnotation: &ast.TSNumberKeyword{
			BaseNode: ast.BaseNode{NodeType: ast.NodeTypeTSNumberKeyword.String()},
		},
	}

	result := converter.convertTSTypeAnnotation(original)

	if result == nil {
		t.Fatal("convertTSTypeAnnotation returned nil")
	}

	if result.TypeAnnotation == nil {
		t.Error("Expected type annotation to be converted")
	}
}

// TestConvertArrayPattern tests converting an ArrayPattern node.
func TestConvertArrayPattern(t *testing.T) {
	source := "[a, b]"
	converter := NewConverter(source, nil)

	original := &ast.ArrayPattern{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeArrayPattern.String(),
		},
		Elements: []ast.Pattern{
			&ast.Identifier{
				BaseNode: ast.BaseNode{NodeType: ast.NodeTypeIdentifier.String()},
				Name:     "a",
			},
			&ast.Identifier{
				BaseNode: ast.BaseNode{NodeType: ast.NodeTypeIdentifier.String()},
				Name:     "b",
			},
		},
	}

	result := converter.convertArrayPattern(original)

	if result == nil {
		t.Fatal("convertArrayPattern returned nil")
	}

	if len(result.Elements) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(result.Elements))
	}
}

// TestNodeMappings tests that bidirectional node mappings are created correctly.
func TestNodeMappings(t *testing.T) {
	source := "x"
	converter := NewConverter(source, &Options{PreserveNodeMaps: true})

	original := &ast.Identifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeIdentifier.String(),
		},
		Name: "x",
	}

	result := converter.convertIdentifier(original)

	// Check forward mapping (TS -> ESTree)
	if mapped, exists := converter.tsNodeToESTreeNodeMap[original]; !exists {
		t.Error("Expected forward mapping to exist")
	} else if mapped != result {
		t.Error("Forward mapping points to wrong node")
	}

	// Check reverse mapping (ESTree -> TS)
	if mapped, exists := converter.esTreeNodeToTSNodeMap[result]; !exists {
		t.Error("Expected reverse mapping to exist")
	} else if mapped != original {
		t.Error("Reverse mapping points to wrong node")
	}
}

// TestGetNodeMaps tests getting node maps from the converter.
func TestGetNodeMaps(t *testing.T) {
	source := "x"
	converter := NewConverter(source, &Options{PreserveNodeMaps: true})

	original := &ast.Identifier{
		BaseNode: ast.BaseNode{
			NodeType: ast.NodeTypeIdentifier.String(),
		},
		Name: "x",
	}

	result := converter.convertIdentifier(original)

	nodeMaps := converter.GetNodeMaps()

	if nodeMaps == nil {
		t.Fatal("GetNodeMaps returned nil")
	}

	if nodeMaps.ESTreeNodeToTSNodeMap == nil {
		t.Error("Expected ESTreeNodeToTSNodeMap to be initialized")
	}

	if nodeMaps.TSNodeToESTreeNodeMap == nil {
		t.Error("Expected TSNodeToESTreeNodeMap to be initialized")
	}

	// Verify mappings are accessible through the returned maps
	if mapped, exists := nodeMaps.TSNodeToESTreeNodeMap[original]; !exists || mapped != result {
		t.Error("Expected node mapping to be accessible through GetNodeMaps")
	}
}

// TestConvertNodeNil tests that converting nil nodes returns nil.
func TestConvertNodeNil(t *testing.T) {
	converter := NewConverter("", nil)

	if result := converter.ConvertNode(nil); result != nil {
		t.Error("Expected ConvertNode(nil) to return nil")
	}
}

// TestConvertExpressionsNil tests that converting nil expression slices returns nil.
func TestConvertExpressionsNil(t *testing.T) {
	converter := NewConverter("", nil)

	if result := converter.convertExpressions(nil); result != nil {
		t.Error("Expected convertExpressions(nil) to return nil")
	}
}

// TestConvertStatementsNil tests that converting nil statement slices returns nil.
func TestConvertStatementsNil(t *testing.T) {
	converter := NewConverter("", nil)

	if result := converter.convertStatements(nil); result != nil {
		t.Error("Expected convertStatements(nil) to return nil")
	}
}

// TestOptionsDefaults tests that default options are set correctly.
func TestOptionsDefaults(t *testing.T) {
	converter := NewConverter("", nil)

	if converter.options == nil {
		t.Error("Expected default options to be set")
	}

	if !converter.options.PreserveNodeMaps {
		t.Error("Expected PreserveNodeMaps to default to true")
	}
}

// TestWithAllowPattern tests the withAllowPattern context manager.
func TestWithAllowPattern(t *testing.T) {
	converter := NewConverter("", nil)

	if converter.allowPattern {
		t.Error("Expected allowPattern to default to false")
	}

	converter.withAllowPattern(true, func() {
		if !converter.allowPattern {
			t.Error("Expected allowPattern to be true inside withAllowPattern")
		}
	})

	if converter.allowPattern {
		t.Error("Expected allowPattern to be restored to false after withAllowPattern")
	}
}
