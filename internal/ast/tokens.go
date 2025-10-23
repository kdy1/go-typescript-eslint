package ast

import "sort"

// Token type constants
const (
	TokenTypeIdentifier = "Identifier"
	TokenTypeKeyword    = "Keyword"
	TokenTypePunctuator = "Punctuator"
	TokenTypeString     = "String"
	TokenTypeNumeric    = "Numeric"
)

// TokenUtils provides utilities for working with tokens in the AST.

// GetTokensInRange returns all tokens within a source range.
func GetTokensInRange(tokens []*Token, start, end int) []*Token {
	var result []*Token
	for _, token := range tokens {
		if token.Range != nil && (*token.Range)[0] >= start && (*token.Range)[1] <= end {
			result = append(result, token)
		}
	}
	return result
}

// GetTokensBefore returns all tokens that appear before a position.
func GetTokensBefore(tokens []*Token, pos int) []*Token {
	var result []*Token
	for _, token := range tokens {
		if token.Range != nil && (*token.Range)[1] <= pos {
			result = append(result, token)
		}
	}
	return result
}

// GetTokensAfter returns all tokens that appear after a position.
func GetTokensAfter(tokens []*Token, pos int) []*Token {
	var result []*Token
	for _, token := range tokens {
		if token.Range != nil && (*token.Range)[0] >= pos {
			result = append(result, token)
		}
	}
	return result
}

// GetTokenAtPosition returns the token at a specific position.
func GetTokenAtPosition(tokens []*Token, pos int) *Token {
	for _, token := range tokens {
		if token.Range != nil && (*token.Range)[0] <= pos && pos < (*token.Range)[1] {
			return token
		}
	}
	return nil
}

// GetFirstToken returns the first token in the list.
func GetFirstToken(tokens []*Token) *Token {
	if len(tokens) == 0 {
		return nil
	}
	return tokens[0]
}

// GetLastToken returns the last token in the list.
func GetLastToken(tokens []*Token) *Token {
	if len(tokens) == 0 {
		return nil
	}
	return tokens[len(tokens)-1]
}

// SortTokens sorts tokens by their start position.
func SortTokens(tokens []*Token) {
	sort.Slice(tokens, func(i, j int) bool {
		if tokens[i].Range == nil || tokens[j].Range == nil {
			return false
		}
		return (*tokens[i].Range)[0] < (*tokens[j].Range)[0]
	})
}

// TokenSpan returns the length of a token in characters.
func TokenSpan(token *Token) int {
	if token == nil || token.Range == nil {
		return 0
	}
	return (*token.Range)[1] - (*token.Range)[0]
}

// GetTokensByType returns all tokens of a specific type.
func GetTokensByType(tokens []*Token, tokenType string) []*Token {
	var result []*Token
	for _, token := range tokens {
		if token.Type == tokenType {
			result = append(result, token)
		}
	}
	return result
}

// GetTokensByValue returns all tokens with a specific value.
func GetTokensByValue(tokens []*Token, value string) []*Token {
	var result []*Token
	for _, token := range tokens {
		if token.Value == value {
			result = append(result, token)
		}
	}
	return result
}

// FindTokenIndex returns the index of a token in a slice, or -1 if not found.
func FindTokenIndex(tokens []*Token, target *Token) int {
	for i, token := range tokens {
		if token == target {
			return i
		}
	}
	return -1
}

// GetNextToken returns the token immediately following the given token.
func GetNextToken(tokens []*Token, current *Token) *Token {
	idx := FindTokenIndex(tokens, current)
	if idx >= 0 && idx < len(tokens)-1 {
		return tokens[idx+1]
	}
	return nil
}

// GetPreviousToken returns the token immediately preceding the given token.
func GetPreviousToken(tokens []*Token, current *Token) *Token {
	idx := FindTokenIndex(tokens, current)
	if idx > 0 {
		return tokens[idx-1]
	}
	return nil
}

// GetTokensForNode returns all tokens that belong to a specific node.
func GetTokensForNode(tokens []*Token, node Node) []*Token {
	if node == nil {
		return nil
	}
	return GetTokensInRange(tokens, node.Pos(), node.End())
}

// GetFirstTokenOfNode returns the first token of a node.
func GetFirstTokenOfNode(tokens []*Token, node Node) *Token {
	if node == nil {
		return nil
	}
	nodeTokens := GetTokensForNode(tokens, node)
	return GetFirstToken(nodeTokens)
}

// GetLastTokenOfNode returns the last token of a node.
func GetLastTokenOfNode(tokens []*Token, node Node) *Token {
	if node == nil {
		return nil
	}
	nodeTokens := GetTokensForNode(tokens, node)
	return GetLastToken(nodeTokens)
}

// IsTokenBefore checks if token 'a' appears before token 'b'.
func IsTokenBefore(a, b *Token) bool {
	if a == nil || b == nil || a.Range == nil || b.Range == nil {
		return false
	}
	return (*a.Range)[1] <= (*b.Range)[0]
}

// IsTokenAfter checks if token 'a' appears after token 'b'.
func IsTokenAfter(a, b *Token) bool {
	if a == nil || b == nil || a.Range == nil || b.Range == nil {
		return false
	}
	return (*a.Range)[0] >= (*b.Range)[1]
}

// TokensOverlap checks if two tokens overlap in source code.
func TokensOverlap(a, b *Token) bool {
	if a == nil || b == nil || a.Range == nil || b.Range == nil {
		return false
	}
	aStart, aEnd := (*a.Range)[0], (*a.Range)[1]
	bStart, bEnd := (*b.Range)[0], (*b.Range)[1]
	return (aStart <= bStart && bStart < aEnd) || (bStart <= aStart && aStart < bEnd)
}

// GetTokenText returns the text content of a token.
func GetTokenText(token *Token) string {
	if token == nil {
		return ""
	}
	return token.Value
}

// IsKeywordToken checks if a token is a keyword.
func IsKeywordToken(token *Token) bool {
	if token == nil {
		return false
	}
	return token.Type == TokenTypeKeyword
}

// IsIdentifierToken checks if a token is an identifier.
func IsIdentifierToken(token *Token) bool {
	if token == nil {
		return false
	}
	return token.Type == TokenTypeIdentifier
}

// IsPunctuatorToken checks if a token is a punctuator.
func IsPunctuatorToken(token *Token) bool {
	if token == nil {
		return false
	}
	return token.Type == TokenTypePunctuator
}

// IsStringToken checks if a token is a string literal.
func IsStringToken(token *Token) bool {
	if token == nil {
		return false
	}
	return token.Type == TokenTypeString
}

// IsNumericToken checks if a token is a numeric literal.
func IsNumericToken(token *Token) bool {
	if token == nil {
		return false
	}
	return token.Type == TokenTypeNumeric
}

// IsOperatorToken checks if a token is an operator.
func IsOperatorToken(token *Token) bool {
	if token == nil {
		return false
	}
	// Common operator punctuators
	operators := map[string]bool{
		"+": true, "-": true, "*": true, "/": true, "%": true,
		"++": true, "--": true, "**": true,
		"&": true, "|": true, "^": true, "~": true,
		"<<": true, ">>": true, ">>>": true,
		"&&": true, "||": true, "??": true,
		"!":  true,
		"==": true, "!=": true, "===": true, "!==": true,
		"<": true, "<=": true, ">": true, ">=": true,
		"=": true, "+=": true, "-=": true, "*=": true, "/=": true,
		"%=": true, "&=": true, "|=": true, "^=": true,
		"<<=": true, ">>=": true, ">>>=": true, "**=": true,
		"?.": true, "??=": true, "||=": true, "&&=": true,
	}
	return IsPunctuatorToken(token) && operators[token.Value]
}

// IsBinaryOperator checks if a token is a binary operator.
func IsBinaryOperator(token *Token) bool {
	if token == nil {
		return false
	}
	binaryOps := map[string]bool{
		"+": true, "-": true, "*": true, "/": true, "%": true, "**": true,
		"&": true, "|": true, "^": true,
		"<<": true, ">>": true, ">>>": true,
		"&&": true, "||": true, "??": true,
		"==": true, "!=": true, "===": true, "!==": true,
		"<": true, "<=": true, ">": true, ">=": true,
		"in": true, "instanceof": true,
	}
	return (IsPunctuatorToken(token) || IsKeywordToken(token)) && binaryOps[token.Value]
}

// IsUnaryOperator checks if a token is a unary operator.
func IsUnaryOperator(token *Token) bool {
	if token == nil {
		return false
	}
	unaryOps := map[string]bool{
		"+": true, "-": true, "!": true, "~": true,
		"++": true, "--": true,
		"typeof": true, "void": true, "delete": true,
	}
	return (IsPunctuatorToken(token) || IsKeywordToken(token)) && unaryOps[token.Value]
}

// IsAssignmentOperator checks if a token is an assignment operator.
func IsAssignmentOperator(token *Token) bool {
	if token == nil {
		return false
	}
	assignOps := map[string]bool{
		"=": true, "+=": true, "-=": true, "*=": true, "/=": true,
		"%=": true, "&=": true, "|=": true, "^=": true,
		"<<=": true, ">>=": true, ">>>=": true, "**=": true,
		"??=": true, "||=": true, "&&=": true,
	}
	return IsPunctuatorToken(token) && assignOps[token.Value]
}

// GetTokensOnLine returns all tokens on a specific line.
func GetTokensOnLine(tokens []*Token, line int) []*Token {
	var result []*Token
	for _, token := range tokens {
		if token.Loc != nil && token.Loc.Start.Line == line {
			result = append(result, token)
		}
	}
	return result
}

// CountTokens returns the number of tokens in a range.
func CountTokens(tokens []*Token, start, end int) int {
	count := 0
	for _, token := range tokens {
		if token.Range != nil && (*token.Range)[0] >= start && (*token.Range)[1] <= end {
			count++
		}
	}
	return count
}

// GetWhitespaceAfter returns the amount of whitespace after a token.
// This requires comparing with the next token or end position.
func GetWhitespaceAfter(tokens []*Token, token *Token) int {
	if token == nil || token.Range == nil {
		return 0
	}
	next := GetNextToken(tokens, token)
	if next == nil || next.Range == nil {
		return 0
	}
	return (*next.Range)[0] - (*token.Range)[1]
}

// GetWhitespaceBefore returns the amount of whitespace before a token.
// This requires comparing with the previous token or start position.
func GetWhitespaceBefore(tokens []*Token, token *Token) int {
	if token == nil || token.Range == nil {
		return 0
	}
	prev := GetPreviousToken(tokens, token)
	if prev == nil || prev.Range == nil {
		return 0
	}
	return (*token.Range)[0] - (*prev.Range)[1]
}
