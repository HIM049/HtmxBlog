package router

import (
	"HtmxBlog/api_handler"
	"HtmxBlog/view_handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var HRouter *HotRouter

func Init() {
	HRouter = NewHotRouter()
	HRouter.Update(loadRoutes())
}

func RefreshRoutes() {
	HRouter.Update(loadRoutes())
}

func loadRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	r.Get("/attach/{id}", api_handler.LoadAttachHandler)

	err := RegisterPagesRouter(r)
	if err != nil {
		panic(err)
	}
	r.Get("/p/{id}", view_handler.PostView)

	r.Route("/admin", func(r chi.Router) {
		r.Get("/", view_handler.AdminView)
		r.Get("/pages", view_handler.ManagePagesView)
		r.Get("/posts", view_handler.ManagePostsView)
		r.Get("/post/{id}/edit", view_handler.EditView)
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/admin", func(r chi.Router) {

			r.Route("/post", func(r chi.Router) {
				r.Post("/create", api_handler.HandlePostCreate)
				r.Patch("/{id}", api_handler.HandlePostUpdate)
				r.Delete("/{id}", api_handler.HandlePostDelete)

				r.Post("/{id}/attach", api_handler.UploadAttachHandler)
			})

			r.Post("/page", api_handler.HandlePageCreate)
			r.Delete("/page/{id}", api_handler.HandlePageDelete)

		})
	})

	// TODO share link system. use static url to route share link
	return r
}
