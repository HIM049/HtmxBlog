package services

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
)

func CreatePage(item *model.Page) error {
	return config.DB.Create(item).Error
}

func ReadPage(name string) (*model.Page, error) {
	var item model.Page
	err := config.DB.Where("name = ?", name).First(&item).Error
	return &item, err
}

func UpdatePage(item *model.Page) error {
	return config.DB.Save(item).Error
}

func DeletePage(name string) error {
	return config.DB.Where("name = ?", name).Delete(model.Page{}).Error
}

func ReadAllPages() ([]model.Page, error) {
	var pages []model.Page
	err := config.DB.Find(&pages).Error
	return pages, err
}
