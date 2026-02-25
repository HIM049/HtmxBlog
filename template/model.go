package template

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
	"HtmxBlog/services"
)

var currentState App

type App struct {
	PageTitle  string
	Navigation []model.Page
	Categories []model.ViewCategory
	Posts      []model.ViewPost
	Settings   map[string]string
}

// InitBaseApp initializes the base application data
// It should be called after config.Init(), config.InitDB(), services.UpdateConfig()
func InitBaseApp() {
	UpdateSettings()
	UpdateNavigation()
	UpdateCategories()
}

// GetBaseApp returns the base application data
func GetBaseApp() App {
	return currentState
}

// UpdateNavigation updates the navigation data
func UpdateNavigation() error {
	pages, err := services.ReadNavPages()
	if err != nil {
		return err
	}
	currentState.Navigation = pages
	return nil
}

// UpdateCategories updates the categories data
func UpdateCategories() error {
	categories, err := services.ReadViewCategories()
	if err != nil {
		return err
	}
	currentState.Categories = categories
	return nil
}

func UpdateSettings() {
	currentState.Settings = config.Cfg.Settings
}
