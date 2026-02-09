package services

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
)

// CreatePage creates a page.
func CreatePage(item *model.Page) error {
	return config.DB.Create(item).Error
}

// ReadPage reads a page by its name.
func ReadPage(name string) (*model.Page, error) {
	var item model.Page
	err := config.DB.Where("name = ?", name).First(&item).Error
	return &item, err
}

// UpdatePage updates a page.
func UpdatePage(item *model.Page) error {
	return config.DB.Save(item).Error
}

// DeletePage deletes a page.
func DeletePage(name string) error {
	return config.DB.Where("name = ?", name).Delete(model.Page{}).Error
}

// ReadAllPages reads all pages.
func ReadAllPages() ([]model.Page, error) {
	var pages []model.Page
	err := config.DB.Find(&pages).Error
	return pages, err
}

// ReadNavPages reads all pages that are shown in the navigation.
func ReadNavPages() ([]model.Page, error) {
	var pages []model.Page
	err := config.DB.Where("show_in_nav = ?", true).Find(&pages).Error
	return pages, err
}
