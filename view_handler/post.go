package view_handler

import (
	"HtmxBlog/database"
	"HtmxBlog/template"
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

	post, err := database.ReadPost(uint(id))
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	vp := template.ViewPost{Post: *post}
	if err = vp.LoadContent(); err != nil {
		http.Error(w, "Failed to load post content", http.StatusInternalServerError)
		return
	}
	base := template.GetBaseApp()
	base.Posts = []template.ViewPost{vp}

	template.Tmpl.ExecuteTemplate(w, "post", base)

}
