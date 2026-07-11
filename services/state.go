package services

import (
	"HtmxBlog/config"
	"HtmxBlog/state"
	"sort"
)

// InitBaseApp initializes the base application data
// It should be called after config.Init(), config.InitDB(), services.UpdateConfig()
func InitBaseApp() {
	UpdateSettings()
	UpdateNavigation()
	UpdateCategories()
	UpdateTags()
}

func UpdateSettings() {
	state.CurrentState.Settings = config.Cfg.Settings
}

// UpdateNavigation updates the navigation data
func UpdateNavigation() error {
	pages, err := ReadNavPages()
	if err != nil {
		return err
	}
	state.CurrentState.Navigation = pages
	return nil
}

// UpdateCategories updates the categories data
func UpdateCategories() error {
	categories, err := ReadViewCategories()
	if err != nil {
		return err
	}
	state.CurrentState.Categories = categories
	return nil
}

// UpdateTags updates the tags data
func UpdateTags() error {
	tags, err := ReadAllTags()
	if err != nil {
		return err
	}

	sort.Slice(tags, func(i, j int) bool {
		return len(tags[i].Posts) > len(tags[j].Posts)
	})

	state.CurrentState.Tags = tags
	return nil
}
