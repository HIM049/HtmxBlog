package view_handler

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"HtmxBlog/template"
	"net/http"
)

func GenericViewLoader(tmpl string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		categoryID := r.URL.Query().Get("category")
		posts, err := services.ReadPostsWithConditions(10, 0, model.VisibilityPublic, "", model.StateRelease, categoryID)
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
		template.Tmpl.ExecuteTemplate(w, tmpl, base)
	}
}
