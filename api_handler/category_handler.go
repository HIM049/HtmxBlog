package api_handler

import (
	"HtmxBlog/services"
	"HtmxBlog/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HandleCategoryCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	color := r.FormValue("color")

	if name == "" || color == "" {
		http.Error(w, "Name and color are required", http.StatusBadRequest)
		return
	}

	_, err = services.CreateCategory(name, color)
	if err != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`<div class="text-green-600 font-bold p-4 bg-green-50 rounded shadow-md border border-green-200">Category created successfully!</div>`))
	w.Header().Set("HX-Trigger", "newCategory")
}

func HandleCategoryDelete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Category ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	if err := services.DeleteCategory(uint(id)); err != nil {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

func HandleCategoryUpdate(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Category ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	category, err := services.ReadCategory(uint(id))
	if err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	if name := r.FormValue("name"); name != "" {
		category.Name = name
	}
	if color := r.FormValue("color"); color != "" {
		category.Color = color
	}
	if visibility := r.FormValue("visibility"); visibility != "" {
		category.Visibility = visibility
	}

	err = services.UpdateCategory(category)
	if err != nil {
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	template.Tmpl.ExecuteTemplate(w, "category_item", category)
}
