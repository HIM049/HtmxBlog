package handler

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"net/http"
	"sort"
	"strconv"
)

type AdminApp struct {
	Req       *http.Request
	PageTitle string
}

func NewAdminApp(r *http.Request, name string) *AdminApp {
	title := "Dashboard"
	if name != "" {
		title += " - " + name
	}
	return &AdminApp{
		Req:       r,
		PageTitle: title,
	}
}

func (a *AdminApp) GetStats() (*services.StatsSummary, error) {
	return services.GetStats()
}

func (a *AdminApp) Stats() (*services.StatsSummary, error) {
	return a.GetStats()
}

func (a *AdminApp) GetPosts() ([]model.Post, error) {
	return services.ReadPosts(100, 0)
}

func (a *AdminApp) Posts() ([]model.Post, error) {
	return a.GetPosts()
}

func (a *AdminApp) GetPages() ([]model.Page, error) {
	return services.ReadAllPages()
}

func (a *AdminApp) Pages() ([]model.Page, error) {
	return a.GetPages()
}

func (a *AdminApp) GetCategories() ([]model.Category, error) {
	return services.ReadCategories()
}

func (a *AdminApp) Categories() ([]model.Category, error) {
	return a.GetCategories()
}

func (a *AdminApp) GetSettings() ([]model.Setting, error) {
	return services.ReadAllSettings()
}

func (a *AdminApp) Settings() ([]model.Setting, error) {
	return a.GetSettings()
}

func (a *AdminApp) GetComments() ([]model.Comment, error) {
	return services.ReadAllComments()
}

func (a *AdminApp) Comments() ([]model.Comment, error) {
	return a.GetComments()
}

func (a *AdminApp) GetRedirects() ([]model.Redirect, error) {
	return services.ReadAllRedirects()
}

func (a *AdminApp) Redirects() ([]model.Redirect, error) {
	return a.GetRedirects()
}

func (a *AdminApp) GetEditPost() (*model.ViewPost, error) {
	if a.Req == nil {
		return nil, nil
	}
	idStr := a.Req.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	vp, err := services.GetDraft(uint(id))
	if err == nil && vp != nil {
		return vp, nil
	}
	post, err := services.ReadPost(uint(id))
	if err != nil {
		return nil, err
	}
	vp = &model.ViewPost{Post: *post}
	vp.LoadContent()
	return vp, nil
}

func (a *AdminApp) EditPost() (*model.ViewPost, error) {
	return a.GetEditPost()
}

func (a *AdminApp) Post() (*model.ViewPost, error) {
	return a.GetEditPost()
}

func (a *AdminApp) SortedPages() ([]model.Page, error) {
	pages, err := services.ReadAllPages()
	if err != nil {
		return nil, err
	}
	var sortedPages []model.Page
	for _, page := range pages {
		if page.Sort > 0 {
			sortedPages = append(sortedPages, page)
		}
	}
	sort.Slice(sortedPages, func(i, j int) bool {
		return sortedPages[i].Sort < sortedPages[j].Sort
	})
	return sortedPages, nil
}

func (a *AdminApp) GetSortedPages() ([]model.Page, error) {
	return a.SortedPages()
}

func (a *AdminApp) HiddenPages() ([]model.Page, error) {
	pages, err := services.ReadAllPages()
	if err != nil {
		return nil, err
	}
	var hiddenPages []model.Page
	for _, page := range pages {
		if page.Sort <= 0 {
			hiddenPages = append(hiddenPages, page)
		}
	}
	return hiddenPages, nil
}

func (a *AdminApp) GetHiddenPages() ([]model.Page, error) {
	return a.HiddenPages()
}
