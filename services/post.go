package services

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const DRAFTS_DIR = "./app_data/drafts"

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
	if err := config.DB.Preload("Category").Preload("Tags").Preload("Attachs").First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func ReadPosts(num int, offset int) ([]model.Post, error) {
	var posts []model.Post
	err := config.DB.Preload("Category").Preload("Tags").Limit(num).Offset(offset).Order("id desc").Find(&posts).Error
	return posts, err
}

func ReadPostsWithConditions(num, offset int, visibility, protect, state, categoryID, tag string) ([]model.Post, error) {
	var posts []model.Post
	query := config.DB.Model(&model.Post{})
	query = query.Where(model.Post{Visibility: visibility, Protect: protect, State: state})

	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	if tag != "" {
		query = query.Joins("JOIN post_tags ON post_tags.post_id = posts.id").
			Joins("JOIN tags ON tags.id = post_tags.tag_id").
			Where("tags.name = ?", tag)
	}

	err := query.Preload("Category").Preload("Tags").Limit(num).Offset(offset).Order("posts.created_at desc").Find(&posts).Error
	return posts, err
}

func CountPostsWithConditions(visibility, protect, state, categoryID, tag string) (int64, error) {
	var count int64
	query := config.DB.Model(&model.Post{})
	query = query.Where(model.Post{Visibility: visibility, Protect: protect, State: state})

	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	if tag != "" {
		query = query.Joins("JOIN post_tags ON post_tags.post_id = posts.id").
			Joins("JOIN tags ON tags.id = post_tags.tag_id").
			Where("tags.name = ?", tag)
	}

	err := query.Count(&count).Error
	return count, err
}

func updatePost(post *model.Post) error {
	err := config.DB.Model(post).Select("*").Omit("Category", "Tags", "Attachs").Updates(post).Error
	if err != nil {
		return err
	}
	return config.DB.Model(post).Association("Tags").Replace(post.Tags)
}

func updateContent(p *model.ViewPost) error {
	if err := os.WriteFile(p.ContentPath(), []byte(p.Content), 0644); err != nil {
		return err
	}
	return nil
}

func UpdatePostWithContent(p model.GenericPost) error {
	post := p.GetPost()
	err := updatePost(post)
	if err != nil {
		return err
	}
	if vp, conv := p.GetViewPost(); !conv {
		err = updateContent(vp)
		if err != nil {
			return err
		}
	}

	// Clean up draft if it exists after publishing
	draftPath := filepath.Join(DRAFTS_DIR, post.Uid+".json")
	if _, err := os.Stat(draftPath); err == nil {
		os.Remove(draftPath)
	}

	onPostChange()
	return nil
}

func SaveDraft(id uint, p *model.ViewPost) error {
	post, err := ReadPost(id)
	if err != nil {
		return err
	}

	// Ensure drafts directory exists
	if _, err := os.Stat(DRAFTS_DIR); os.IsNotExist(err) {
		os.MkdirAll(DRAFTS_DIR, 0755)
	}

	draftPath := filepath.Join(DRAFTS_DIR, post.Uid+".json")
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(draftPath, data, 0644)
}

func GetDraft(id uint) (*model.ViewPost, error) {
	post, err := ReadPost(id)
	if err != nil {
		return nil, err
	}

	draftPath := filepath.Join(DRAFTS_DIR, post.Uid+".json")
	if _, err := os.Stat(draftPath); os.IsNotExist(err) {
		return nil, err // No draft found
	}

	data, err := os.ReadFile(draftPath)
	if err != nil {
		return nil, err
	}

	var vp model.ViewPost
	if err := json.Unmarshal(data, &vp); err != nil {
		return nil, err
	}

	// Ensure ID, UID, and Attachs match the current live record
	vp.ID = post.ID
	vp.Uid = post.Uid
	vp.Attachs = post.Attachs

	return &vp, nil
}

func DeletePost(id uint) error {
	err := config.DB.Delete(&model.Post{}, id).Error
	if err != nil {
		return err
	}
	onPostChange()
	return nil
}

func DestroyPost(id uint) error {
	var post model.Post
	if err := config.DB.Unscoped().First(&post, id).Error; err != nil {
		return err
	}
	if err := os.Remove(post.ContentPath()); err != nil {
		return err
	}
	if err := config.DB.Unscoped().Delete(&post).Error; err != nil {
		return err
	}
	onPostChange()
	return nil
}

func RegisterOnPostChange(f func()) {
	onPostChange = f
}
