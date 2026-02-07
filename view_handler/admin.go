package viewhandler

import (
	"HtmxBlog/database"
	"HtmxBlog/template"
	"net/http"
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
