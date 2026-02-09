package router

import (
	"HtmxBlog/services"
	"HtmxBlog/view_handler"

	"github.com/go-chi/chi/v5"
)

// RegisterPagesRouter registers all pages as routes.
// It will panic if any error occurs.
func RegisterPagesRouter(r chi.Router) error {
	pages, err := services.ReadAllPages()
	if err != nil {
		return err
	}

	for _, page := range pages {
		r.Get(page.Route, view_handler.GenericViewLoader(page.Template))
	}
	return nil
}
