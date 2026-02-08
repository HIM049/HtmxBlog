package view_handler

import (
	"HtmxBlog/database"
	"HtmxBlog/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func AdminView(w http.ResponseWriter, r *http.Request) {
	pages, _ := database.ReadAllPages()
	posts, _ := database.ReadPosts(100, 0)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.Tmpl.ExecuteTemplate(w, "admin", map[string]interface{}{
		"Pages": pages,
		"Posts": posts,
	})
}

func EditView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	post, err := database.ReadPost(uint(id))
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
