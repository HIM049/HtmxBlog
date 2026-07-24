package handler

import (
	"HtmxBlog/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// // ManageCategoriesView renders the category management page skeleton.
func HandleCategoryCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		HtmxError(w, "Invalid form data")
		return
	}

	name := r.FormValue("name")
	color := r.FormValue("color")

	if name == "" || color == "" {
		HtmxError(w, "Name and color are required")
		return
	}

	_, err = services.CreateCategory(name, color)
	if err != nil {
		HtmxError(w, "Failed to create category")
		return
	}

	go services.UpdateCategories()

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("HX-Trigger", "categoryChanged")
	HtmxSuccess(w, "Category created successfully!")
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

	go services.UpdateCategories()

	w.Header().Set("HX-Trigger", "categoryChanged")
	w.WriteHeader(http.StatusOK)
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

	go services.UpdateCategories()

	w.Header().Set("HX-Trigger", "categoryChanged")
	w.WriteHeader(http.StatusOK)
}
