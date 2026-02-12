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
	err := config.DB.Where("sort > ?", 0).Order("sort asc").Find(&pages).Error
	return pages, err
}

// ReorderPages reorders pages by their names.
func ReorderPages(names []string) error {
	tx := config.DB.Begin()
	for i, name := range names {
		// sort index starts from 1
		// 0 is default value, mean no sort
		if err := tx.Model(&model.Page{}).Where("name = ?", name).Update("sort", i+1).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	onPageChange()
	return nil
}

// UnsortPage sets the sort value of a page to 0.
func UnsortPage(name string) error {
	err := config.DB.Model(&model.Page{}).Where("name = ?", name).Update("sort", 0).Error
	if err != nil {
		return err
	}
	onPageChange()
	return nil
}

// RegisterOnPageChange registers a callback that called when page changed.
func RegisterOnPageChange(f func()) {
	onPageChange = f
}
