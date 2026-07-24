package state

import (
	"HtmxBlog/model"
)

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
	Comments   []*model.CommentNode
}

// GetBaseApp returns the base application data
func GetBaseApp() App {
	return CurrentState
}

const PageSize = 5
