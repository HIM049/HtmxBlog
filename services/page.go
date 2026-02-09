package services

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
)

var onPageChange func()

// CreatePage creates a page.
func CreatePage(item *model.Page) error {
	err := config.DB.Create(item).Error
	if err != nil {
		return err
	}
	onPageChange()
	return nil
}

// ReadPage reads a page by its name.
func ReadPage(name string) (*model.Page, error) {
	var item model.Page
	err := config.DB.Where("name = ?", name).First(&item).Error
	return &item, err
}

// UpdatePage updates a page.
func UpdatePage(item *model.Page) error {
	err := config.DB.Save(item).Error
	if err != nil {
		return err
	}
	onPageChange()
	return nil
}

// DeletePage deletes a page.
func DeletePage(name string) error {
	err := config.DB.Where("name = ?", name).Delete(model.Page{}).Error
	if err != nil {
		return err
	}
	onPageChange()
	return nil
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

// RegisterOnPageChange registers a callback that called when page changed.
func RegisterOnPageChange(f func()) {
	onPageChange = f
}
