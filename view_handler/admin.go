package view_handler

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"HtmxBlog/template"
	"net/http"
	"sort"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func AdminView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.Tmpl.ExecuteTemplate(w, "admin", nil)
}

func ManagePagesView(w http.ResponseWriter, r *http.Request) {
	pages, _ := services.ReadAllPages()

	var sortedPages []model.Page
	var hiddenPages []model.Page

	for _, page := range pages {
		if page.Sort > 0 {
			sortedPages = append(sortedPages, page)
		} else {
			hiddenPages = append(hiddenPages, page)
		}
	}

	sort.Slice(sortedPages, func(i, j int) bool {
		return sortedPages[i].Sort < sortedPages[j].Sort
	})

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.Tmpl.ExecuteTemplate(w, "page_manage", map[string]interface{}{
		"SortedPages": sortedPages,
		"HiddenPages": hiddenPages,
	})
}

func ManageCategoriesView(w http.ResponseWriter, r *http.Request) {
	categories, _ := services.ReadCategories()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.Tmpl.ExecuteTemplate(w, "category_manage", map[string]interface{}{
		"Categories": categories,
	})
}

func ManagePostsView(w http.ResponseWriter, r *http.Request) {
	posts, _ := services.ReadPosts(100, 0)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.Tmpl.ExecuteTemplate(w, "post_manage", map[string]interface{}{
		"Posts": posts,
	})
}

func ManageSettingsView(w http.ResponseWriter, r *http.Request) {
	settings, _ := services.ReadAllSettings()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.Tmpl.ExecuteTemplate(w, "manage_settings", map[string]interface{}{
		"Settings": settings,
	})
}

func EditView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	post, err := services.ReadPost(uint(id))
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	vp := &model.ViewPost{Post: *post}
	vp.LoadContent()
	categories, _ := services.ReadCategories()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.Tmpl.ExecuteTemplate(w, "update_post", map[string]interface{}{
		"Post":       vp,
		"Categories": categories,
	})
}
