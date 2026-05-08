package maintain

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

func CheckAndInstall() error {
	config.Init()
	config.InitDB()

	_, ok := config.Cfg.Settings["init_at"]
	if ok {
		return errors.New("database already installed")
	}

	return config.DB.Transaction(func(tx *gorm.DB) error {
		if err := createDefaultPages(tx); err != nil {
			return err
		}

		if err := createDefaultSettings(tx); err != nil {
			return err
		}

		// set init_at on finish (use FirstOrCreate to handle edge case)
		return tx.Where(model.Setting{Key: "init_at"}).
			Assign(model.Setting{Value: time.Now().Format("2006-01-02 15:04:05")}).
			FirstOrCreate(&model.Setting{}).Error
	})
}

func createDefaultPages(tx *gorm.DB) error {
	// create home page (skip if already exists)
	page := model.Page{
		Name:     "Home",
		Route:    "/",
		Template: "index",
		Sort:     1,
	}
	return tx.Where(model.Page{Name: "Home"}).FirstOrCreate(&page).Error
}

func createDefaultSettings(tx *gorm.DB) error {
	defaults := []model.Setting{
		{Key: "site_name", Value: "HtmxBlog"},
		{Key: "site_slogan", Value: "Hello World!"},
	}

	for _, s := range defaults {
		if err := tx.Where(model.Setting{Key: s.Key}).FirstOrCreate(&s).Error; err != nil {
			return err
		}
	}

	return nil
}
