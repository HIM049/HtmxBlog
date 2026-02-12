package template

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
)

var currentState App

type App struct {
	PageTitle  string
	Navigation []model.Page
	Categories []model.ViewCategory
	Posts      []model.ViewPost
}

func InitBaseApp() {
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
