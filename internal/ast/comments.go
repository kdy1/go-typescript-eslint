package ast

import "sort"

// CommentAttachment represents a comment attached to a node.
type CommentAttachment struct {
	Comment *Comment
	Node    Node
	Type    CommentAttachmentType
}

// CommentAttachmentType indicates how a comment is attached to a node.
type CommentAttachmentType int

const (
	// CommentLeading indicates the comment appears before the node.
	CommentLeading CommentAttachmentType = iota
	// CommentTrailing indicates the comment appears after the node on the same line.
	CommentTrailing
	// CommentInner indicates the comment is inside the node.
	CommentInner
)

// AttachComments attaches comments to their associated nodes based on position.
// This is typically called after parsing to associate comments with AST nodes.
func AttachComments(root Node, comments []*Comment) []CommentAttachment {
	if root == nil || len(comments) == 0 {
		return nil
	}

	// Sort comments by position
	sortedComments := make([]*Comment, len(comments))
	copy(sortedComments, comments)
	sort.Slice(sortedComments, func(i, j int) bool {
		if sortedComments[i].Range == nil || sortedComments[j].Range == nil {
			return false
		}
		return (*sortedComments[i].Range)[0] < (*sortedComments[j].Range)[0]
	})

	// Collect all nodes
	var nodes []Node
	Traverse(root, func(node Node) bool {
		nodes = append(nodes, node)
		return true
	})

	// Sort nodes by position
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Pos() < nodes[j].Pos()
	})

	var attachments []CommentAttachment
	commentIdx := 0

	for _, node := range nodes {
		nodeStart := node.Pos()
		nodeEnd := node.End()

		// Attach leading comments
		for commentIdx < len(sortedComments) {
			comment := sortedComments[commentIdx]
			if comment.Range != nil && (*comment.Range)[1] <= nodeStart {
				attachments = append(attachments, CommentAttachment{
					Comment: comment,
					Node:    node,
					Type:    CommentLeading,
				})
				commentIdx++
			} else {
				break
			}
		}

		// Attach inner comments (comments inside the node's range)
		for commentIdx < len(sortedComments) {
			comment := sortedComments[commentIdx]
			if comment.Range != nil && (*comment.Range)[0] >= nodeStart && (*comment.Range)[1] <= nodeEnd {
				attachments = append(attachments, CommentAttachment{
					Comment: comment,
					Node:    node,
					Type:    CommentInner,
				})
				commentIdx++
			} else {
				break
			}
		}
	}

	// Handle remaining comments as trailing comments of the last node
	if len(nodes) > 0 && commentIdx < len(sortedComments) {
		lastNode := nodes[len(nodes)-1]
		for ; commentIdx < len(sortedComments); commentIdx++ {
			attachments = append(attachments, CommentAttachment{
				Comment: sortedComments[commentIdx],
				Node:    lastNode,
				Type:    CommentTrailing,
			})
		}
	}

	return attachments
}

// GetLeadingComments returns all leading comments for a node.
func GetLeadingComments(attachments []CommentAttachment, node Node) []*Comment {
	var comments []*Comment
	for _, att := range attachments {
		if att.Node == node && att.Type == CommentLeading {
			comments = append(comments, att.Comment)
		}
	}
	return comments
}

// GetTrailingComments returns all trailing comments for a node.
func GetTrailingComments(attachments []CommentAttachment, node Node) []*Comment {
	var comments []*Comment
	for _, att := range attachments {
		if att.Node == node && att.Type == CommentTrailing {
			comments = append(comments, att.Comment)
		}
	}
	return comments
}

// GetInnerComments returns all inner comments for a node.
func GetInnerComments(attachments []CommentAttachment, node Node) []*Comment {
	var comments []*Comment
	for _, att := range attachments {
		if att.Node == node && att.Type == CommentInner {
			comments = append(comments, att.Comment)
		}
	}
	return comments
}

// GetAllComments returns all comments associated with a node.
func GetAllComments(attachments []CommentAttachment, node Node) []*Comment {
	var comments []*Comment
	for _, att := range attachments {
		if att.Node == node {
			comments = append(comments, att.Comment)
		}
	}
	return comments
}

// HasLeadingComment checks if a node has any leading comments.
func HasLeadingComment(attachments []CommentAttachment, node Node) bool {
	for _, att := range attachments {
		if att.Node == node && att.Type == CommentLeading {
			return true
		}
	}
	return false
}

// HasTrailingComment checks if a node has any trailing comments.
func HasTrailingComment(attachments []CommentAttachment, node Node) bool {
	for _, att := range attachments {
		if att.Node == node && att.Type == CommentTrailing {
			return true
		}
	}
	return false
}

// FilterCommentsByType filters comments by their type.
func FilterCommentsByType(comments []*Comment, commentType string) []*Comment {
	var filtered []*Comment
	for _, comment := range comments {
		if comment.Type == commentType {
			filtered = append(filtered, comment)
		}
	}
	return filtered
}

// GetLineComments returns all line comments from a list.
func GetLineComments(comments []*Comment) []*Comment {
	return FilterCommentsByType(comments, "Line")
}

// GetBlockComments returns all block comments from a list.
func GetBlockComments(comments []*Comment) []*Comment {
	return FilterCommentsByType(comments, "Block")
}

// IsDocComment checks if a comment is a documentation comment (JSDoc-style).
// A doc comment is a block comment that starts with /** (not just /*).
func IsDocComment(comment *Comment) bool {
	if comment == nil || comment.Type != "Block" {
		return false
	}
	return len(comment.Value) > 0 && comment.Value[0] == '*'
}

// GetDocComments returns all documentation comments from a list.
func GetDocComments(comments []*Comment) []*Comment {
	var docs []*Comment
	for _, comment := range comments {
		if IsDocComment(comment) {
			docs = append(docs, comment)
		}
	}
	return docs
}

// GetCommentText returns the text content of a comment (without delimiters).
func GetCommentText(comment *Comment) string {
	if comment == nil {
		return ""
	}
	return comment.Value
}

// GetCommentsInRange returns all comments within a source range.
func GetCommentsInRange(comments []*Comment, start, end int) []*Comment {
	var result []*Comment
	for _, comment := range comments {
		if comment.Range != nil && (*comment.Range)[0] >= start && (*comment.Range)[1] <= end {
			result = append(result, comment)
		}
	}
	return result
}

// GetCommentsBefore returns all comments that appear before a position.
func GetCommentsBefore(comments []*Comment, pos int) []*Comment {
	var result []*Comment
	for _, comment := range comments {
		if comment.Range != nil && (*comment.Range)[1] <= pos {
			result = append(result, comment)
		}
	}
	return result
}

// GetCommentsAfter returns all comments that appear after a position.
func GetCommentsAfter(comments []*Comment, pos int) []*Comment {
	var result []*Comment
	for _, comment := range comments {
		if comment.Range != nil && (*comment.Range)[0] >= pos {
			result = append(result, comment)
		}
	}
	return result
}

// SortComments sorts comments by their start position.
func SortComments(comments []*Comment) {
	sort.Slice(comments, func(i, j int) bool {
		if comments[i].Range == nil || comments[j].Range == nil {
			return false
		}
		return (*comments[i].Range)[0] < (*comments[j].Range)[0]
	})
}

// CommentSpan returns the length of a comment in characters.
func CommentSpan(comment *Comment) int {
	if comment == nil || comment.Range == nil {
		return 0
	}
	return (*comment.Range)[1] - (*comment.Range)[0]
}

// IsCommentOnSameLine checks if a comment is on the same line as a position.
// This requires the comment to have location information with line numbers.
func IsCommentOnSameLine(comment *Comment, loc *SourceLocation) bool {
	if comment == nil || comment.Loc == nil || loc == nil {
		return false
	}
	return comment.Loc.Start.Line == loc.Start.Line
}

// GetCommentsOnLine returns all comments on a specific line.
func GetCommentsOnLine(comments []*Comment, line int) []*Comment {
	var result []*Comment
	for _, comment := range comments {
		if comment.Loc != nil && comment.Loc.Start.Line == line {
			result = append(result, comment)
		}
	}
	return result
}
