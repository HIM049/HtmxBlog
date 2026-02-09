package api_handler

import (
	"HtmxBlog/services"
	"fmt"
	"net/http"
	"strconv"

	"os"
	"strings"

	"github.com/go-chi/chi/v5"
)

// HandlePostCreate is a handler for creating a new post.
// It creates a default post and redirects to the editor page.
func HandlePostCreate(w http.ResponseWriter, r *http.Request) {
	post, err := services.CreateDefaultPost()
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`<div class="text-red-600 font-bold p-4 bg-red-50 rounded shadow-md border border-red-200">Failed to create post: %s</div>`, err.Error())))

		return
	}

	w.Header().Set("HX-Redirect", fmt.Sprintf("/admin/post/%d/edit", post.ID))
}

func HandlePostUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Post ID", http.StatusBadRequest)
		return
	}
	post, err := services.ReadPost(uint(id))
	if err != nil {
		http.Error(w, "Failed to read post", http.StatusInternalServerError)
		return
	}

	// Update logic
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	if title := r.FormValue("title"); title != "" {
		post.Title = title
	}

	if category := r.FormValue("category"); category != "" {
		post.Category.Name = category
	}

	if color := r.FormValue("cat_color"); color != "" {
		post.Category.Color = color
	}

	if tags := r.FormValue("tags"); tags != "" {
		tagsList := strings.Split(tags, ",")
		var cleanTags []string
		for _, tag := range tagsList {
			cleanTags = append(cleanTags, strings.TrimSpace(tag))
		}
		post.Tags = cleanTags
	}

	if customVars := r.FormValue("custom_vars"); customVars != "" {
		vars := make(map[string]interface{})
		lines := strings.Split(customVars, "\n")
		for _, line := range lines {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				if key != "" {
					vars[key] = value
				}
			}
		}
		post.CustomVars = vars
	}

	if content := r.FormValue("content"); content != "" {
		// Content might be very large, writing to file directly
		if err := os.WriteFile(post.ContentPath(), []byte(content), 0644); err != nil {
			http.Error(w, "Failed to save content", http.StatusInternalServerError)
			return
		}
	}

	if status := r.FormValue("status"); status == "published" {
		post.State = "release" // model.StateRelease
	}

	err = services.UpdatePost(post)
	if err != nil {
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

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

	if err := services.DeletePost(uint(id)); err != nil {
		http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("HX-Trigger", `{"showMessage": "Post deleted successfully!"}`)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
