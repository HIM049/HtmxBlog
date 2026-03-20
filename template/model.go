package template

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
	"HtmxBlog/services"
	"sort"
)

var currentState App

const PageSize = 5

type Pagination struct {
	CurrentPage int
	TotalPages  int
	TotalPosts  int64
	HasPrev     bool
	HasNext     bool
	PrevPage    int
	NextPage    int
	PageNumbers []int
	CategoryID  string
	Tag         string
}

type App struct {
	PageTitle  string
	Navigation []model.Page
	Categories []model.ViewCategory
	Tags       []model.Tag
	Posts      []model.ViewPost
	Settings   map[string]string
	Pagination Pagination
	Comments   []*services.CommentNode
}

// InitBaseApp initializes the base application data
// It should be called after config.Init(), config.InitDB(), services.UpdateConfig()
func InitBaseApp() {
	UpdateSettings()
	UpdateNavigation()
	UpdateCategories()
	UpdateTags()
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

// UpdateTags updates the tags data
func UpdateTags() error {
	tags, err := services.ReadAllTags()
	if err != nil {
		return err
	}

	sort.Slice(tags, func(i, j int) bool {
		return len(tags[i].Posts) > len(tags[j].Posts)
	})

	currentState.Tags = tags
	return nil
}

func UpdateSettings() {
	currentState.Settings = config.Cfg.Settings
}
