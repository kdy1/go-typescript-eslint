//nolint:revive // Package name "types" is appropriate for type definitions
package types

// Type represents a TypeScript type.
type Type interface {
	// String returns a string representation of the type.
	String() string

	// Equals checks if this type equals another type.
	Equals(other Type) bool
}

// PrimitiveType represents a primitive TypeScript type.
type PrimitiveType int

const (
	// StringType represents the string type.
	StringType PrimitiveType = iota
	// NumberType represents the number type.
	NumberType
	// BooleanType represents the boolean type.
	BooleanType
	// NullType represents the null type.
	NullType
	// UndefinedType represents the undefined type.
	UndefinedType
	// SymbolType represents the symbol type.
	SymbolType
	// BigIntType represents the bigint type.
	BigIntType
	// VoidType represents the void type.
	VoidType
	// NeverType represents the never type.
	NeverType
	// AnyType represents the any type.
	AnyType
	// UnknownType represents the unknown type.
	UnknownType
)

var primitiveTypeNames = map[PrimitiveType]string{
	StringType:    "string",
	NumberType:    "number",
	BooleanType:   "boolean",
	NullType:      "null",
	UndefinedType: "undefined",
	SymbolType:    "symbol",
	BigIntType:    "bigint",
	VoidType:      "void",
	NeverType:     "never",
	AnyType:       "any",
	UnknownType:   "unknown",
}

// String returns a string representation of the primitive type.
func (pt PrimitiveType) String() string {
	if name, ok := primitiveTypeNames[pt]; ok {
		return name
	}
	return "unknown"
}

// Equals checks if this primitive type equals another type.
func (pt PrimitiveType) Equals(other Type) bool {
	if otherPT, ok := other.(PrimitiveType); ok {
		return pt == otherPT
	}
	return false
}
