package services

import (
	"HtmxBlog/config"
	"HtmxBlog/model"

	"github.com/google/uuid"
)

var onPostChange func()

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
	onPostChange()
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

func ReadPostsWithConditions(num, offset int, visibility, protect, state, categoryID string) ([]model.Post, error) {
	var posts []model.Post
	query := config.DB.Model(&model.Post{})

	if visibility != "" {
		query = query.Where("visibility = ?", visibility)
	}

	if protect != "" {
		query = query.Where("protect = ?", protect)
	}

	if state != "" {
		query = query.Where("state = ?", state)
	}

	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	err := query.Preload("Category").Limit(num).Offset(offset).Find(&posts).Error
	return posts, err
}

func UpdatePost(post *model.Post) error {
	err := config.DB.Save(post).Error
	if err != nil {
		return err
	}
	onPostChange()
	return nil
}

func DeletePost(id uint) error {
	err := config.DB.Delete(&model.Post{}, id).Error
	if err != nil {
		return err
	}
	onPostChange()
	return nil
}

func RegisterOnPostChange(f func()) {
	onPostChange = f
}
