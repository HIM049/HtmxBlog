package services

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
)

func CreateTag(tag *model.Tag) error {
	return config.DB.Create(tag).Error
}

func ReadTag(tag *model.Tag) error {
	err := config.DB.First(tag).Error
	return err
}

func ReadAllTags() ([]model.Tag, error) {
	var tags []model.Tag
	err := config.DB.Find(&tags).Error
	return tags, err
}

func UpdateTag(tag *model.Tag) error {
	err := config.DB.Save(tag).Error
	return err
}

func DeleteTag(tag *model.Tag) error {
	err := config.DB.Delete(tag).Error
	return err
}

func GetTag(name string) (*model.Tag, error) {
	var tag model.Tag
	err := config.DB.Where(model.Tag{Name: name}).FirstOrCreate(&tag).Error
	return &tag, err
}
