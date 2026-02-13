package services

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
)

var onCategoryChange func()

func CreateCategory(name, color string) (*model.Category, error) {
	category := &model.Category{
		Name:       name,
		Color:      color,
		Visibility: model.VisibilityPublic,
	}
	err := config.DB.Create(category).Error
	if err != nil {
		return nil, err
	}
	onCategoryChange()
	return category, nil
}

func ReadCategory(id uint) (*model.Category, error) {
	var category model.Category
	err := config.DB.First(&category, id).Error
	return &category, err
}

func ReadCategories() ([]model.Category, error) {
	var categories []model.Category
	err := config.DB.Find(&categories).Error
	return categories, err
}

func ReadViewCategories() ([]model.ViewCategory, error) {
	var results []model.ViewCategory
	err := config.DB.Debug().Model(&model.Category{}).
		Select("categories.*, count(posts.id) as count").
		Joins(
			"left join posts on posts.category_id = categories.id AND posts.deleted_at IS NULL AND posts.visibility = ? AND posts.state = ?",
			model.VisibilityPublic,
			model.StateRelease,
		).
		Where("categories.visibility = ?", model.VisibilityPublic).
		Group("categories.id").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func UpdateCategory(category *model.Category) error {
	err := config.DB.Save(category).Error
	if err != nil {
		return err
	}
	onCategoryChange()
	return nil
}

func DeleteCategory(id uint) error {
	err := config.DB.Delete(&model.Category{}, id).Error
	if err != nil {
		return err
	}
	onCategoryChange()
	return nil
}

// SetCategoryVisibility sets the visibility of a category
func SetCategoryVisibility(id uint, visibility string) (*model.Category, error) {
	category, err := ReadCategory(id)
	if err != nil {
		return nil, err
	}

	category.Visibility = visibility

	err = UpdateCategory(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

// RegisterOnCategoryChange registers a callback that called when category changed.
func RegisterOnCategoryChange(f func()) {
	onCategoryChange = f
}
