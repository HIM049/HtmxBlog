package handler

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"HtmxBlog/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
)

// HandlePostCreate is a handler for creating a new post.
// It creates a default post and redirects to the editor page.
func HandlePostCreate(w http.ResponseWriter, r *http.Request) {
	post, err := services.CreateDefaultPost()
	if err != nil {
		HtmxError(w, "Failed to create post: "+err.Error())
		return
	}

	go func() {
		services.UpdateCategories()
		services.UpdateTags()
	}()

	w.Header().Set("HX-Redirect", fmt.Sprintf("/admin/editor?id=%d", post.ID))
}

func HandlePostUpdate(w http.ResponseWriter, r *http.Request) {
	vp, err := parseRequest(r)
	if err != nil {
		http.Error(w, "Failed to parse request", http.StatusBadRequest)
		return
	}

	// Update saves to draft only when released
	if vp.Post.State == model.StateDraft {
		err = services.UpdatePostWithContent(vp)
	} else {
		err = services.SaveDraft(vp.Post.Uid, vp)
	}
	if err != nil {
		http.Error(w, "Failed to save update", http.StatusInternalServerError)
		return
	}

	go func() {
		services.UpdateCategories()
		services.UpdateTags()
	}()

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Saved to draft"))
}

func HandlePostPublish(w http.ResponseWriter, r *http.Request) {
	vp, err := parseRequest(r)
	if err != nil {
		http.Error(w, "Failed to parse request", http.StatusBadRequest)
		return
	}

	draftFile := false
	if vp.State == model.StateDraft {
		// Set state to release
		vp.State = model.StateRelease
	} else {
		draftFile = true
	}

	// Perform actual update (promote draft/current state to published)
	err = services.UpdatePostWithContent(vp)
	if err != nil {
		http.Error(w, "Failed to publish post", http.StatusInternalServerError)
		return
	}

	if draftFile {
		services.DeleteDraft(vp.Post.Uid)
	}

	go func() {
		services.UpdateCategories()
		services.UpdateTags()
	}()

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

	go func() {
		services.UpdateCategories()
		services.UpdateTags()
	}()

	w.Header().Set("HX-Trigger", "postChanged")
	w.WriteHeader(http.StatusOK)
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
			for tagName := range strings.SplitSeq(tagsVal, ",") {
				trimmed := strings.TrimSpace(tagName)
				if trimmed == "" {
					continue
				}
				tag, err := services.GetTag(trimmed)
				if err != nil {
					log.Errorf("Failed to get tag %v", err)
					continue
				}
				postTags = append(postTags, *tag)
			}
		}
		post.Tags = postTags
	}

	vars := make(map[string]any)
	if keys, ok := r.Form["custom_var_keys"]; ok {
		values := r.Form["custom_var_values"]
		for i, key := range keys {
			key = strings.TrimSpace(key)
			if key != "" && i < len(values) {
				vars[key] = strings.TrimSpace(values[i])
			}
		}
	}
	post.CustomVars = vars

	if createdAt := r.FormValue("created_at"); createdAt != "" {
		t, err := utils.ParseDateTimeLocal(createdAt)
		if err == nil {
			post.CreatedAt = t
		}
	}
	return nil
}

func parseRequest(r *http.Request) (*model.ViewPost, error) {
	// get id & read base struct
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return nil, err
	}
	post, err := services.ReadPost(uint(id))
	if err != nil {
		return nil, err
	}

	// patch post struct
	if err := parsePostForm(post, r); err != nil {
		return nil, err
	}

	// get new content or fallback to current
	// TODO: fix: fallback to current draft/release
	var vp *model.ViewPost
	if content := r.FormValue("content"); content != "" {
		vp = &model.ViewPost{
			Post:    *post,
			Content: content,
		}
	} else {
		if post.State == model.StateRelease {
			vp, err = services.GetDraft(uint(id))
			if err == nil {
				return vp, nil
			}
		}

		vp = &model.ViewPost{Post: *post}
		vp.LoadContent()
	}

	return vp, nil
}
