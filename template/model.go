package template

import (
	"HtmxBlog/database"
	"HtmxBlog/model"
)

var currentState App

type App struct {
	PageTitle  string
	Navigation []model.Page
	Posts      []ViewPost
}

type ViewPost struct {
	model.Post
	Content string
}

func GetBaseApp() App {
	return currentState
}

// UpdateNavigation updates the navigation data
func UpdateNavigation() error {
	pages, err := database.ReadAllPages()
	if err != nil {
		return err
	}
	currentState.Navigation = pages
	return nil
}
