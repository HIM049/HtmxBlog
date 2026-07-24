package services

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
)

func CreateSetting(item *model.Setting) error {
	return config.DB.Create(item).Error
}

func ReadSetting(id uint) (*model.Setting, error) {
	var item model.Setting
	err := config.DB.First(&item, id).Error
	return &item, err
}

func UpdateSetting(item *model.Setting) error {
	return config.DB.Save(item).Error
}

func DeleteSetting(id uint) error {
	return config.DB.Delete(&model.Setting{}, id).Error
}

func ReadAllSettings() ([]model.Setting, error) {
	var items []model.Setting
	err := config.DB.Find(&items).Error
	return items, err
}

func UpdateConfig() {
	settings, err := ReadAllSettings()
	if err != nil {
		return
	}
	config.Cfg.Settings = make(map[string]string)
	for _, setting := range settings {
		config.Cfg.Settings[setting.Key] = setting.Value
	}
}
