package template

import "HtmxBlog/model"

type App struct {
	PageTitle  string
	Navigation []NavigationItem
	Posts      []model.Post
}

type NavigationItem struct {
	Name string
	Url  string
}
