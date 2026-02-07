package api_handler

import (
	"HtmxBlog/database"
	"HtmxBlog/model"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

const POSTS_DIR = "./app_data/posts"

func HandlePostCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	cName := r.FormValue("category")
	cColor := r.FormValue("cat_color")
	content := r.FormValue("content")
	customVarsRaw := r.FormValue("custom_vars")

	if title == "" || content == "" {
		http.Error(w, "Title and content are required", http.StatusBadRequest)
		return
	}

	// Parse Custom Vars
	customVars := make(map[string]interface{})
	if customVarsRaw != "" {
		lines := strings.Split(customVarsRaw, "\n")
		for _, line := range lines {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				if key != "" {
					customVars[key] = value
				}
			}
		}
	}

	// Save content to file
	// TODO: uuid as file name
	filename := fmt.Sprintf("%s.md", title)
	filePath := filepath.Join(POSTS_DIR, filename)
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		http.Error(w, "Failed to save post content", http.StatusInternalServerError)
		return
	}

	// Create post record in database
	category := model.Category{
		Name:  cName,
		Color: cColor,
	}
	post := &model.Post{
		Permission:  model.PermissionPublic,
		Title:       title,
		Category:    category,
		ContentPath: filePath,
		CustomVars:  customVars,
	}

	if err := database.CreatePost(post); err != nil {
		http.Error(w, "Failed to save post to database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`<div class="text-green-600 font-bold p-4 bg-green-50 rounded shadow-md border border-green-200">Post created successfully!</div>`))
}

func HandlePostDelete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Post ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Post ID", http.StatusBadRequest)
		return
	}

	if err := database.DeletePostById(uint(id)); err != nil {
		http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("HX-Trigger", `{"showMessage": "Post deleted successfully!"}`)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
