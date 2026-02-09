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
