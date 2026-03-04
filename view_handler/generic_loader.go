package view_handler

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
	"HtmxBlog/services"
	"HtmxBlog/template"
	"net/http"
	"strconv"
)

func GenericViewLoader(tmpl string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		categoryID := r.URL.Query().Get("category")
		pageStr := r.URL.Query().Get("page")
		page, _ := strconv.Atoi(pageStr)
		if page < 1 {
			page = 1
		}

		offset := (page - 1) * template.PageSize
		posts, err := services.ReadPostsWithConditions(template.PageSize, offset, model.VisibilityPublic, "", model.StateRelease, categoryID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		totalPosts, err := services.CountPostsWithConditions(model.VisibilityPublic, "", model.StateRelease, categoryID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		totalPages := int((totalPosts + int64(template.PageSize) - 1) / int64(template.PageSize))

		base := template.GetBaseApp()
		base.PageTitle = config.Cfg.Settings["site_name"]
		for _, post := range posts {
			base.Posts = append(base.Posts, model.ViewPost{
				Post:    post,
				Content: post.ContentPath(),
			})
		}

		// Pagination logic
		base.Pagination = template.Pagination{
			CurrentPage: page,
			TotalPages:  totalPages,
			TotalPosts:  totalPosts,
			HasPrev:     page > 1,
			HasNext:     page < totalPages,
			PrevPage:    page - 1,
			NextPage:    page + 1,
			CategoryID:  categoryID,
		}

		// Generate page numbers (simple version: all pages)
		for i := 1; i <= totalPages; i++ {
			base.Pagination.PageNumbers = append(base.Pagination.PageNumbers, i)
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		template.Tmpl.ExecuteTemplate(w, tmpl, base)
	}
}
