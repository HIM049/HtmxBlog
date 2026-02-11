package services

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
)

func CreateCategory(name, color string) (*model.Category, error) {
	category := &model.Category{
		Name:  name,
		Color: color,
	}
	err := config.DB.Create(category).Error
	if err != nil {
		return nil, err
	}
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

func UpdateCategory(category *model.Category) error {
	return config.DB.Save(category).Error
}

func DeleteCategory(id uint) error {
	return config.DB.Delete(&model.Category{}, id).Error
}
