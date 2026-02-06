package handler

import (
	"HtmxBlog/database"
	"HtmxBlog/model"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
	w.Write([]byte("<div>Post created successfully!</div>"))
}
