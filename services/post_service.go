package services

import (
	"HtmxBlog/db"
	"HtmxBlog/models"
	"github.com/google/uuid"
)

func ListPosts() ([]models.Post, error) {
	posts, err := db.ListPostTable()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func CreatePost(title, content string) error {
	// Generate a uid for the new post
	uid, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	post := models.Post{
		Uid:     uid.String(),
		Title:   title,
		Tags:    nil,
		Content: content,
	}

	// Save to database
	if err = db.CreatePostTable(&post); err != nil {
		return err
	}

	return nil
}

func ReadPost(uid string) (*models.Post, error) {
	post, err := db.ReadPostTable(uid)
	if err != nil {
		return nil, err
	}
	return post, nil
}
