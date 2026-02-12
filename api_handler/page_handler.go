package api_handler

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

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

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("<div>Page created successfully!</div>"))
}

func HandlePageDelete(w http.ResponseWriter, r *http.Request) {
	pageId := chi.URLParam(r, "id")
	if pageId == "" {
		http.Error(w, "Page ID is required", http.StatusBadRequest)
		return
	}

	services.DeletePage(pageId)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

func HandlePageReorder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	names := r.Form["pages"]
	if len(names) == 0 {
		// Try to parse as JSON if form data is empty
		// This is just a fallback, usually not needed if frontend is updated correctly
		// But let's keep it simple and just rely on form data first.
		// Actually, let's just stick to form data as requested by HTMX approach.
		http.Error(w, "No pages provided", http.StatusBadRequest)
		return
	}

	err = services.ReorderPages(names)
	if err != nil {
		http.Error(w, "Failed to reorder pages", http.StatusInternalServerError)
		return
	}

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

	w.WriteHeader(http.StatusOK)
}
