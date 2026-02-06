package database

import "HtmxBlog/model"

func CreatePage(item *model.Page) error {
	return db.Create(item).Error
}

func ReadPage(name string) (*model.Page, error) {
	var item model.Page
	err := db.Where("name = ?", name).First(&item).Error
	return &item, err
}

func UpdatePage(item *model.Page) error {
	return db.Save(item).Error
}

func DeletePage(name string) error {
	return db.Where("name = ?", name).Delete(model.Page{}).Error
}

func ReadAllPages() ([]model.Page, error) {
	var pages []model.Page
	err := db.Find(&pages).Error
	return pages, err
}
