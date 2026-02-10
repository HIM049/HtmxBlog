package view_handler

import (
	"HtmxBlog/services"
	"HtmxBlog/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func AdminView(w http.ResponseWriter, r *http.Request) {
	// pages, _ := services.ReadAllPages()
	// posts, _ := services.ReadPosts(100, 0)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.Tmpl.ExecuteTemplate(w, "admin", map[string]interface{}{
		// "Pages": pages,
		// "Posts": posts,
	})
}

func ManagePostsView(w http.ResponseWriter, r *http.Request) {
	posts, _ := services.ReadPosts(100, 0)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.Tmpl.ExecuteTemplate(w, "post_manage", map[string]interface{}{
		"Posts": posts,
	})
}

func EditView(w http.ResponseWriter, r *http.Request) {
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

	vp := &template.ViewPost{Post: *post}
	vp.LoadContent()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.Tmpl.ExecuteTemplate(w, "update_post", map[string]interface{}{
		"Post": vp,
	})
}
