package template

import (
	"HtmxBlog/database"
	"HtmxBlog/model"
	"os"
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

// LoadContent loads the content of the post
func (vp *ViewPost) LoadContent() error {
	content, err := os.ReadFile(vp.ContentPath)
	if err != nil {
		return err
	}
	vp.Content = string(content)
	return nil
}

// GetBaseApp returns the base application data
func GetBaseApp() App {
	if currentState.Navigation == nil {
		UpdateNavigation()
	}
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
