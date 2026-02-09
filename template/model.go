package template

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"fmt"
	"os"
	"sort"
	"strings"
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
	content, err := os.ReadFile(vp.ContentPath())
	if err != nil {
		return err
	}
	vp.Content = string(content)
	return nil
}

func (p *ViewPost) TagsToString() string {
	return strings.Join(p.Tags, ", ")
}

func (p *ViewPost) CustomVarsToString() string {
	var lines []string
	for k, v := range p.CustomVars {
		lines = append(lines, fmt.Sprintf("%s: %v", k, v))
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n")
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
	pages, err := services.ReadNavPages()
	if err != nil {
		return err
	}
	currentState.Navigation = pages
	return nil
}
