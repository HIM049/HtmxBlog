package viewhandler

import (
	"HtmxBlog/database"
	"HtmxBlog/template"
	"net/http"
	"os"
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

	content, err := os.ReadFile(post.ContentPath)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	base := template.GetBaseApp()
	base.Posts = []template.ViewPost{
		{
			Post:    *post,
			Content: string(content),
		},
	}

	template.Tmpl.ExecuteTemplate(w, "post", base)

}
