package template

import "HtmxBlog/model"

type App struct {
	PageTitle  string
	Navigation []model.Page
	Posts      []model.Post
}
