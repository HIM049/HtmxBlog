package handler

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
	"HtmxBlog/services"
	"HtmxBlog/state"
	"net/http"
	"strconv"
)

func GenericViewLoader(pageModel model.Page) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		categoryID := r.URL.Query().Get("category")
		pageStr := r.URL.Query().Get("page")
		tag := r.URL.Query().Get("tag")

		page, _ := strconv.Atoi(pageStr)
		if page < 1 {
			page = 1
		}

		offset := (page - 1) * state.PageSize
		posts, err := services.ReadPostsWithConditions(state.PageSize, offset, model.VisibilityPublic, "", model.StateRelease, categoryID, tag, pageModel.FilterMode, pageModel.FilterCategoryIDs)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		totalPosts, err := services.CountPostsWithConditions(model.VisibilityPublic, "", model.StateRelease, categoryID, tag, pageModel.FilterMode, pageModel.FilterCategoryIDs)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		totalPages := int((totalPosts + int64(state.PageSize) - 1) / int64(state.PageSize))

		base := state.GetBaseApp()
		if pageModel.Name != "" && pageModel.Name != "Home" {
			base.PageTitle = pageModel.Name + " - " + config.Cfg.Settings["site_name"]
		} else {
			base.PageTitle = config.Cfg.Settings["site_name"]
		}

		// Filter sidebar categories based on the page's filter configuration
		if pageModel.FilterMode == model.FilterInclude && len(pageModel.FilterCategoryIDs) > 0 {
			idSet := make(map[uint]bool, len(pageModel.FilterCategoryIDs))
			for _, id := range pageModel.FilterCategoryIDs {
				idSet[id] = true
			}
			var filteredCats []model.ViewCategory
			for _, cat := range base.Categories {
				if idSet[cat.ID] {
					filteredCats = append(filteredCats, cat)
				}
			}
			base.Categories = filteredCats
		} else if pageModel.FilterMode == model.FilterExclude && len(pageModel.FilterCategoryIDs) > 0 {
			idSet := make(map[uint]bool, len(pageModel.FilterCategoryIDs))
			for _, id := range pageModel.FilterCategoryIDs {
				idSet[id] = true
			}
			var filteredCats []model.ViewCategory
			for _, cat := range base.Categories {
				if !idSet[cat.ID] {
					filteredCats = append(filteredCats, cat)
				}
			}
			base.Categories = filteredCats
		}
		for _, post := range posts {
			base.Posts = append(base.Posts, model.ViewPost{
				Post: post,
			})
		}

		// Pagination logic
		base.Pagination = state.Pagination{
			CurrentPage: page,
			TotalPages:  totalPages,
			TotalPosts:  totalPosts,
			HasPrev:     page > 1,
			HasNext:     page < totalPages,
			PrevPage:    page - 1,
			NextPage:    page + 1,
			CategoryID:  categoryID,
			Tag:         tag,
		}

		// Generate page numbers
		if totalPages <= 9 {
			for i := 1; i <= totalPages; i++ {
				base.Pagination.PageNumbers = append(base.Pagination.PageNumbers, i)
			}
		} else {
			if page <= 5 {
				// Show previous 6 pages
				for i := 1; i <= 6; i++ {
					base.Pagination.PageNumbers = append(base.Pagination.PageNumbers, i)
				}
				base.Pagination.PageNumbers = append(base.Pagination.PageNumbers, 0, totalPages-1, totalPages)
			} else if page >= totalPages-4 {
				// Show end 6 pages
				base.Pagination.PageNumbers = append(base.Pagination.PageNumbers, 1, 2, 0)
				for i := totalPages - 5; i <= totalPages; i++ {
					base.Pagination.PageNumbers = append(base.Pagination.PageNumbers, i)
				}
			} else {
				// Show middle pages
				base.Pagination.PageNumbers = append(base.Pagination.PageNumbers, 1, 2, 0, page-1, page, page+1, 0, totalPages-1, totalPages)
			}
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		state.Tmpl.ExecuteTemplate(w, pageModel.Template, base)
	}
}
