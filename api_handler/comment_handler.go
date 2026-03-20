package api_handler

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// HandleCommentCreate is a handler for creating a new comment.
func HandleCommentCreate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	postIDStr := r.FormValue("post_id")
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	parentStr := r.FormValue("parent")
	var parent uint
	if parentStr != "" {
		p, err := strconv.ParseUint(parentStr, 10, 64)
		if err == nil {
			parent = uint(p)
		}
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	url := r.FormValue("url")
	content := r.FormValue("content")

	if name == "" || email == "" || content == "" {
		http.Error(w, "Name, Email, and Content are required", http.StatusBadRequest)
		return
	}

	comment := &model.Comment{
		PostID:    uint(postID),
		Parent:    parent,
		Name:      name,
		Email:     email,
		Url:       url,
		Content:   content,
		IP:        r.RemoteAddr,
		UserAgent: r.UserAgent(),
		State:     model.StatePending,
	}

	if err := services.CreateComment(comment); err != nil {
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}

	// For HTMX, we can return a success message or the new comment partial.
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("HX-Trigger", "commentAdded")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`<div class="text-green-600 font-bold p-2 bg-green-50 rounded border border-green-200">Comment submitted successfully!</div>`))
}

// HandleCommentApprove handles comment approval for admin.
func HandleCommentApprove(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Comment ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	if err := services.ApproveComment(uint(id)); err != nil {
		http.Error(w, "Failed to approve comment", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Approved"))
}

// HandleCommentDelete handles comment deletion for admin.
func HandleCommentDelete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Comment ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	if err := services.DeleteComment(uint(id)); err != nil {
		http.Error(w, "Failed to delete comment", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted"))
}
