package viewhandler

import (
	"HtmxBlog/database"
	"HtmxBlog/template"
	"net/http"
)

func IndexView(w http.ResponseWriter, r *http.Request) {
	pages, err := database.ReadAllPages()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	posts, err := database.ReadPosts(10, 0)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.Tmpl.ExecuteTemplate(w, "index",
		template.App{
			PageTitle:  "Hello World",
			Navigation: pages,
			Posts:      posts,
		},
	)

}
