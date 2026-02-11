package services

import (
	"HtmxBlog/config"
	"HtmxBlog/model"

	"github.com/google/uuid"
)

// CreateDefaultPost creates a post with default values
func CreateDefaultPost() (*model.Post, error) {
	uid := uuid.New().String()
	post := &model.Post{
		Uid:        uid,
		Visibility: model.VisibilityPublic, // TODO config default permission
		State:      model.StateDraft,
	}
	err := config.DB.Create(post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func ReadPost(id uint) (*model.Post, error) {
	var post model.Post
	if err := config.DB.Preload("Category").First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func ReadPosts(num int, offset int) ([]model.Post, error) {
	var posts []model.Post
	err := config.DB.Preload("Category").Limit(num).Offset(offset).Find(&posts).Error
	return posts, err
}

func ReadPostsWithConditions(num, offset int, visibility, protect, state string) ([]model.Post, error) {
	var posts []model.Post
	query := config.DB.Model(&model.Post{})

	if visibility != "" {
		query.Where("visibility = ?", visibility)
	}

	if protect != "" {
		query.Where("protect = ?", protect)
	}

	if state != "" {
		query.Where("state = ?", state)
	}

	query.Preload("Category").Limit(num).Offset(offset).Find(&posts)
	return posts, query.Error
}

func UpdatePost(post *model.Post) error {
	return config.DB.Save(post).Error
}

func DeletePost(id uint) error {
	return config.DB.Delete(&model.Post{}, id).Error
}
