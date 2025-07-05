package db

import "HtmxBlog/models"

func ListPostTable() ([]models.Post, error) {
	var posts []models.Post
	err := db.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func CreatePostTable(post *models.Post) error {
	return db.Create(&post).Error
}

func ReadPostTable(uid string) (*models.Post, error) {
	var post models.Post
	err := db.Where("uid = ?", uid).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func UpdatePostTable(post *models.Post) error {
	return db.Save(&post).Error
}

func DeletePostTable(uid string) error {
	var post models.Post
	err := db.Where("uid = ?", uid).First(&post).Error
	if err != nil {
		return err
	}
	return db.Delete(&post).Error
}
