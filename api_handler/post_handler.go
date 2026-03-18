package api_handler

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"HtmxBlog/utils"
	"fmt"
	"net/http"
	"strconv"

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

func parsePostForm(post *model.Post, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	if title := r.FormValue("title"); title != "" {
		post.Title = title
	}

	if visibility := r.FormValue("visibility"); visibility != "" {
		post.Visibility = visibility
	}

	if protect := r.FormValue("protect"); protect != "" {
		post.Protect = protect
	}

	if state := r.FormValue("state"); state != "" {
		post.State = state
	}

	if categoryID := r.FormValue("category_id"); categoryID != "" {
		id, err := strconv.Atoi(categoryID)
		if err == nil {
			if id == 0 {
				post.CategoryID = nil
			} else {
				val := uint(id)
				post.CategoryID = &val
			}
		}
	}

	if tags, ok := r.Form["tags"]; ok {
		tagsVal := tags[0]
		var postTags []model.Tag
		if tagsVal != "" {
			tagsList := strings.Split(tagsVal, ",")
			for _, tagName := range tagsList {
				trimmed := strings.TrimSpace(tagName)
				if trimmed == "" {
					continue
				}
				tag, err := services.GetTag(trimmed)
				if err != nil {
					fmt.Println("Failed to get tag: ", err)
					continue
				}
				postTags = append(postTags, *tag)
			}
		}
		post.Tags = postTags
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

	if createdAt := r.FormValue("created_at"); createdAt != "" {
		t, err := utils.ParseDateTimeLocal(createdAt)
		if err == nil {
			post.CreatedAt = t
		}
	}
	return nil
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

	if err := parsePostForm(post, r); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	var vp *model.ViewPost
	if content := r.FormValue("content"); content != "" {
		vp = &model.ViewPost{
			Post:    *post,
			Content: content,
		}
	} else {
		vp = &model.ViewPost{Post: *post}
		vp.LoadContent()
	}

	// Update ALWAYS saves to draft
	err = services.SaveDraft(uint(id), vp)
	if err != nil {
		http.Error(w, "Failed to save draft", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Saved to draft"))
}

func HandlePostPublish(w http.ResponseWriter, r *http.Request) {
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

	// Parse current form state (Sync editor to what will be published)
	if err := parsePostForm(post, r); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	var vp *model.ViewPost
	if content := r.FormValue("content"); content != "" {
		vp = &model.ViewPost{
			Post:    *post,
			Content: content,
		}
	} else {
		vp = &model.ViewPost{Post: *post}
		vp.LoadContent()
	}

	// Set state to release
	vp.State = model.StateRelease

	// Perform actual update (promote draft/current state to published)
	err = services.UpdatePostWithContent(vp)
	if err != nil {
		http.Error(w, "Failed to publish post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Published"))
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
