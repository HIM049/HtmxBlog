package handler

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"HtmxBlog/state"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func PostView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := services.ReadPost(uint(id))
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	// Check premission
	if post.State != model.StateRelease || post.Visibility == model.VisibilityPrivate {
		cookie, err := r.Cookie("token")
		if err != nil || cookie.Value != state.CurrentToken || services.IsTokenExpired() {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}
	}

	vp := model.ViewPost{Post: *post}
	if err = vp.LoadContent(); err != nil {
		http.Error(w, "Failed to load post content", http.StatusInternalServerError)
		return
	}

	base := state.GetBaseApp()
	base.Posts = []model.ViewPost{vp}

	comments, err := services.ReadCommentsByPostID(uint(id))
	if err == nil {
		base.Comments = services.BuildCommentTree(comments)
	}

	state.Tmpl.ExecuteTemplate(w, "post", base)

}
