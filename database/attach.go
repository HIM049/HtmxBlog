package database

import "HtmxBlog/model"

func CreateAttach(attach *model.Attach) error {
	return db.Create(attach).Error
}

func ReadAttachById(id uint) (*model.Attach, error) {
	var attach model.Attach
	err := db.First(&attach, id).Error
	return &attach, err
}

func ReadAttachByHash(hash string) (*model.Attach, error) {
	var attach model.Attach
	err := db.Where("hash = ?", hash).First(&attach).Error
	return &attach, err
}

func ReadAttachList(limit, offset int) ([]model.Attach, error) {
	var attaches []model.Attach
	err := db.Limit(limit).Offset(offset).Find(&attaches).Error
	return attaches, err
}

func UpdateAttach(attach *model.Attach) error {
	return db.Save(attach).Error
}

func DeleteAttach(id uint) error {
	return db.Delete(&model.Attach{}, id).Error
}
