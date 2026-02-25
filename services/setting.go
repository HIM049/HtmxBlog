package services

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
)

var onSettingChange func()

func CreateSettings(item *model.Setting) error {
	err := config.DB.Create(item).Error
	if err != nil {
		return err
	}
	onSettingChange()
	return nil
}

func ReadSetting(id uint) (*model.Setting, error) {
	var item model.Setting
	err := config.DB.First(&item, id).Error
	return &item, err
}

func ReadSettings() (*model.Setting, error) {
	var item model.Setting
	err := config.DB.First(&item).Error
	return &item, err
}

func UpdateSettings(item *model.Setting) error {
	err := config.DB.Save(item).Error
	if err != nil {
		return err
	}
	onSettingChange()
	return nil
}

func DeleteSettings(id uint) error {
	err := config.DB.Delete(&model.Setting{}, id).Error
	if err != nil {
		return err
	}
	onSettingChange()
	return nil
}

func ReadAllSettings() ([]model.Setting, error) {
	var items []model.Setting
	err := config.DB.Find(&items).Error
	return items, err
}

func RegisterOnSettingChange(f func()) {
	onSettingChange = f
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
