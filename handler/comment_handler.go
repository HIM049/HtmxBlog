package handler

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"HtmxBlog/state"
	"HtmxBlog/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// ManageCommentsView renders the comments management page skeleton.
func ManageCommentsView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	state.AdminTmpl.ExecuteTemplate(w, "comment_manage", nil)
}

// CommentListComponent renders the comments list fragment.
func CommentListComponent(w http.ResponseWriter, r *http.Request) {
	comments, _ := services.ReadAllComments()
	w.Header().Set("Content-Type", "text/html")
	state.AdminTmpl.ExecuteTemplate(w, "manage_comments", map[string]interface{}{
		"Comments": comments,
	})
}

// HandleCommentCreate is a handler for creating a new comment.
func HandleCommentCreate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		HtmxError(w, "Invalid form data")
		return
	}

	postIDStr := r.FormValue("post_id")
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		HtmxError(w, "Invalid post ID")
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
		HtmxError(w, "Name, Email, and Content are required")
		return
	}

	comment := &model.Comment{
		PostID:    uint(postID),
		Parent:    parent,
		Name:      name,
		Email:     email,
		Url:       url,
		Content:   content,
		IP:        utils.GetRealIP(r),
		UserAgent: r.UserAgent(),
		State:     model.StatePending,
	}

	if err := services.CreateComment(comment); err != nil {
		HtmxError(w, "Failed to create comment")
		return
	}

	// For HTMX, we can return a success message or the new comment partial.
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("HX-Trigger", "commentAdded")
	HtmxSuccess(w, "Comment submitted successfully!")
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

	w.Header().Set("HX-Trigger", "commentChanged")
	w.WriteHeader(http.StatusOK)
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

	w.Header().Set("HX-Trigger", "commentChanged")
	w.WriteHeader(http.StatusOK)
}
