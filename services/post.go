package services

import (
	"HtmxBlog/database"
	"HtmxBlog/model"

	"github.com/google/uuid"
)

// CreateDefaultPost creates a post with default values
func CreateDefaultPost() (*model.Post, error) {
	uid := uuid.New().String()
	post := &model.Post{
		Uid:        uid,
		Permission: model.PermissionPublic, // TODO config default permission
		State:      model.StateDraft,
	}
	err := database.CreatePost(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}
