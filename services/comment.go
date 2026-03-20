package services

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
)

// CreateComment creates a new comment in the database.
func CreateComment(comment *model.Comment) error {
	return config.DB.Create(comment).Error
}

// ReadCommentsByPostID returns all approved comments for a given post ID.
func ReadCommentsByPostID(postID uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := config.DB.Where("post_id = ? AND state = ?", postID, model.StateApproved).Order("created_at asc").Find(&comments).Error
	return comments, err
}

// CommentNode represents a comment in a tree structure.
type CommentNode struct {
	model.Comment
	Children []*CommentNode
}

// ReadAllComments returns all comments in the database.
func ReadAllComments() ([]model.Comment, error) {
	var comments []model.Comment
	err := config.DB.Order("created_at desc").Find(&comments).Error
	return comments, err
}

// UpdateComment updates a comment in the database.
func UpdateComment(comment *model.Comment) error {
	return config.DB.Save(comment).Error
}

// DeleteComment deletes a comment by ID and all its children.
func DeleteComment(id uint) error {
	return config.DB.Delete(&model.Comment{}, id).Error
}

// ApproveComment updates a comment's state to 'approved'.
func ApproveComment(id uint) error {
	return config.DB.Model(&model.Comment{}).Where("id = ?", id).Update("state", model.StateApproved).Error
}

// BuildCommentTree organizes a flat list of comments into a tree structure.
func BuildCommentTree(comments []model.Comment) []*CommentNode {
	commentMap := make(map[uint]*CommentNode)
	var rootNodes []*CommentNode

	// Create nodes and map them by ID
	for _, c := range comments {
		node := &CommentNode{
			Comment:  c,
			Children: []*CommentNode{},
		}
		commentMap[c.ID] = node
	}

	// Link children to parents
	for _, c := range comments {
		node := commentMap[c.ID]
		if c.Parent == 0 {
			rootNodes = append(rootNodes, node)
		} else {
			if parent, ok := commentMap[c.Parent]; ok {
				parent.Children = append(parent.Children, node)
			} else {
				// Parent not found, treat as root or skip.
				// For moderation, if parent is not approved but child is, we might want to skip.
				continue
			}
		}
	}

	return rootNodes
}
