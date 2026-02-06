package handler

import (
	"HtmxBlog/database"
	"HtmxBlog/model"
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

	if name == "" || route == "" {
		http.Error(w, "Name and route are required", http.StatusBadRequest)
		return
	}

	database.CreatePage(&model.Page{
		Name:  name,
		Route: route,
	})

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

	database.DeletePage(pageId)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
