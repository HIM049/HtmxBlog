package view_handler

import (
	"HtmxBlog/services"
	"HtmxBlog/template"
	"net/http"
)

func IndexView(w http.ResponseWriter, r *http.Request) {
	posts, err := services.ReadPosts(10, 0)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	base := template.GetBaseApp()
	base.PageTitle = "HIMs Blog"
	for _, post := range posts {
		base.Posts = append(base.Posts, template.ViewPost{
			Post:    post,
			Content: post.ContentPath(),
		})
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.Tmpl.ExecuteTemplate(w, "index", base)

}
