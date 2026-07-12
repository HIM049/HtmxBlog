package handler

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"HtmxBlog/state"
	"net/http"
	"sort"

	"github.com/go-chi/chi/v5"
)

var RefreshRoutes = func() {}

// ManagePagesView renders the pages management page skeleton.
func ManagePagesView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	state.AdminTmpl.ExecuteTemplate(w, "page_manage", nil)
}

// PageListComponent renders the pages list fragment (with sorting logic).
func PageListComponent(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "text/html")
	state.AdminTmpl.ExecuteTemplate(w, "manage_pages", map[string]interface{}{
		"SortedPages": sortedPages,
		"HiddenPages": hiddenPages,
	})
}

func HandlePageCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	route := r.FormValue("route")
	template := r.FormValue("template")

	if name == "" || route == "" || template == "" {
		http.Error(w, "Name and route are required", http.StatusBadRequest)
		return
	}

	err = services.CreatePage(&model.Page{
		Name:     name,
		Route:    route,
		Template: template,
	})
	if err != nil {
		http.Error(w, "Failed to create page", http.StatusInternalServerError)
		return
	}

	go func() {
		RefreshRoutes()
		services.UpdateNavigation()
	}()

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("HX-Trigger", "pageChanged")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("<div>Page created successfully!</div>"))
}

func HandlePageDelete(w http.ResponseWriter, r *http.Request) {
	pageId := chi.URLParam(r, "id")
	if pageId == "" {
		http.Error(w, "Page ID is required", http.StatusBadRequest)
		return
	}

	if err := services.DeletePage(pageId); err != nil {
		http.Error(w, "Failed to delete page", http.StatusInternalServerError)
		return
	}

	go func() {
		RefreshRoutes()
		services.UpdateNavigation()
	}()

	w.Header().Set("HX-Trigger", "pageChanged")
	w.WriteHeader(http.StatusOK)
}

func HandlePageReorder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	names := r.Form["pages"]
	if len(names) == 0 {
		http.Error(w, "No pages provided", http.StatusBadRequest)
		return
	}

	err = services.ReorderPages(names)
	if err != nil {
		http.Error(w, "Failed to reorder pages", http.StatusInternalServerError)
		return
	}

	go func() {
		RefreshRoutes()
		services.UpdateNavigation()
	}()

	w.Header().Set("HX-Trigger", "pageChanged")
	w.WriteHeader(http.StatusOK)
}

func HandlePageUnsort(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	page := r.FormValue("page")
	if page == "" {
		http.Error(w, "Page name is required", http.StatusBadRequest)
		return
	}

	err = services.UnsortPage(page)
	if err != nil {
		http.Error(w, "Failed to unsort page", http.StatusInternalServerError)
		return
	}

	go func() {
		RefreshRoutes()
		services.UpdateNavigation()
	}()

	w.Header().Set("HX-Trigger", "pageChanged")
	w.WriteHeader(http.StatusOK)
}
