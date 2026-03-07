package main

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
	"HtmxBlog/services"
	"time"
)

func CheckAndInstall() error {
	_, ok := config.Cfg.Settings["init_at"]
	if ok {
		return nil
	}

	if err := createDeafultPages(); err != nil {
		return err
	}

	if err := createDefaultSettings(); err != nil {
		return err
	}

	// set init_at on finish
	err := services.CreateSetting(&model.Setting{
		Key:   "init_at",
		Value: time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return err
	}

	return nil
}

func createDeafultPages() error {
	// create home page
	err := services.CreatePage(&model.Page{
		Name:     "Home",
		Route:    "/",
		Template: "index",
		Sort:     1,
	})
	if err != nil {
		return err
	}

	return nil
}

func createDefaultSettings() error {
	var settings []model.Setting

	settings = append(settings, model.Setting{
		Key:   "site_name",
		Value: "HtmxBlog",
	})

	settings = append(settings, model.Setting{
		Key:   "site_slogan",
		Value: "Hello World!",
	})

	for _, setting := range settings {
		err := services.CreateSetting(&setting)
		if err != nil {
			return err
		}
	}

	return nil
}
