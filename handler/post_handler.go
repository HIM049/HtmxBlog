package handler

import (
	"HtmxBlog/database"
	"HtmxBlog/model"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
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
	content := r.FormValue("content")

	if title == "" || content == "" {
		http.Error(w, "Title and content are required", http.StatusBadRequest)
		return
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
		Color: "#0099cc", // TODO set color
	}
	post := &model.Post{
		Title:       title,
		Category:    category,
		ContentPath: filePath,
	}

	if err := database.CreatePost(post); err != nil {
		http.Error(w, "Failed to save post to database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("<div>Post created successfully!</div>"))
}
