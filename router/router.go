package router

import (
	"HtmxBlog/handler"
	viewhandler "HtmxBlog/view_handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	r.Get("/", viewhandler.IndexView)
	r.Get("/manage", viewhandler.ManageView)

	r.Route("/api", func(r chi.Router) {
		r.Route("/admin", func(r chi.Router) {

			r.Post("/post", handler.HandlePostCreate)
			r.Post("/page", handler.HandlePageCreate)
			r.Delete("/page/{id}", handler.HandlePageDelete)

		})
	})

	return r
}
