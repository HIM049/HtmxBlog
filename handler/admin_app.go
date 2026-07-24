package handler

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
	"HtmxBlog/services"
	"HtmxBlog/state"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/charmbracelet/log"
)

type AdminApp struct {
	Req       *http.Request
	I18n      *model.I18n
	PageTitle string
}

// InitI18n initializes the i18n model, should run after the config is initialized.
func InitI18n() {
	i18n, err := loadI18n(nil)
	if err != nil {
		log.Errorf("failed to load i18n: %v, trying fallback to en_us", err)
		lang := "en_us"
		i18n, err = loadI18n(&lang)
		if err != nil {
			log.Fatalf("failed to load i18n: %v", err)
		}
	}
	state.I18n = i18n
}

func TryUpdateI18n() bool {
	i18n, err := loadI18n(nil)
	if err != nil {
		return false
	}
	state.I18n = i18n
	return true
}

func loadI18n(forceLang *string) (*model.I18n, error) {
	var lang string
	if forceLang == nil {
		lang = config.Cfg.Settings["language"]
		if lang == "" {
			log.Warnf("language not set, using default: en_us")
			lang = "en_us"
		}
	} else {
		lang = *forceLang
	}

	file, err := os.Open("templates/admin/i18n/" + lang + ".json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var i18n model.I18n
	err = json.Unmarshal(data, &i18n)
	if err != nil {
		return nil, err
	}
	return &i18n, nil
}

func NewAdminApp(r *http.Request, name string) *AdminApp {
	title := "Dashboard"
	if name != "" {
		title += " - " + name
	}
	return &AdminApp{
		Req:       r,
		I18n:      state.I18n,
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
