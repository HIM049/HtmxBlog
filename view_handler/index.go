package viewhandler

import (
	"HtmxBlog/database"
	"HtmxBlog/model"
	"HtmxBlog/template"
	"net/http"
)

func IndexView(w http.ResponseWriter, r *http.Request) {
	pages, err := database.ReadAllPages()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.Tmpl.ExecuteTemplate(w, "index",
		template.App{
			PageTitle:  "Hello World",
			Navigation: pages,
			Posts: []model.Post{
				{
					Title: "Post 1", Category: model.Category{Name: "技术", Color: "#c0efff"},
					Tags: []string{"Tag 1", "Tag 2"}, ContentPath: "/post/1",
					CustomVars: map[string]interface{}{"intro": "This is a test post, and this is a intro of this post."},
				},
				{
					Title: "Post 2", Category: model.Category{Name: "杂谈", Color: "#fba1ff"},
					Tags: []string{"Tag 2", "Tag 3"}, ContentPath: "/post/2",
					CustomVars: map[string]interface{}{"cover": "/assets/background2.png", "intro": "This is a test post, and this is a intro of this post."},
				},
			},
		},
	)

}
