package router

import (
	"HtmxBlog/api_handler"
	"HtmxBlog/view_handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	r.Get("/", view_handler.IndexView)
	r.Get("/p/{id}", view_handler.PostView)

	r.Get("/admin", view_handler.AdminView)

	r.Route("/api", func(r chi.Router) {
		r.Route("/admin", func(r chi.Router) {

			r.Post("/post", api_handler.HandlePostCreate)
			r.Delete("/post/{id}", api_handler.HandlePostDelete)

			r.Post("/page", api_handler.HandlePageCreate)
			r.Delete("/page/{id}", api_handler.HandlePageDelete)

			r.Post("/attach", api_handler.UploadAttachHandler)

		})
	})

	return r
}
