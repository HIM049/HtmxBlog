package database

import "HtmxBlog/model"

func CreatePost(post *model.Post) error {
	return db.Create(post).Error
}

func ReadPost(id uint) (*model.Post, error) {
	var post model.Post
	if err := db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func UpdatePost(post *model.Post) error {
	return db.Save(post).Error
}

func DeletePost(post *model.Post) error {
	return db.Delete(post).Error
}

func DeletePostById(id uint) error {
	return db.Delete(&model.Post{}, id).Error
}

func ReadPosts(num int, offset int) ([]model.Post, error) {
	var posts []model.Post
	err := db.Limit(num).Offset(offset).Find(&posts).Error
	return posts, err
}
