package router

import (
	"HtmxBlog/api_handler"
	app_middleware "HtmxBlog/middleware"
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

	r.Group(func(r chi.Router) {
		r.Use(app_middleware.NotFoundInterceptor)

		err := RegisterPagesRouter(r)
		if err != nil {
			panic(err)
		}
		r.Get("/p/{id}", view_handler.PostView)

		r.Route("/admin", func(r chi.Router) {
			r.Get("/auth", view_handler.AuthView)
			r.Group(func(r chi.Router) {
				r.Use(app_middleware.AccessControlMiddleware)

				r.Get("/", view_handler.AdminView)
				r.Get("/pages", view_handler.ManagePagesView)
				r.Get("/posts", view_handler.ManagePostsView)
				r.Get("/categories", view_handler.ManageCategoriesView)
				r.Get("/settings", view_handler.ManageSettingsView)
				r.Get("/post/{id}/edit", view_handler.EditView)
			})
		})
	})

	r.Route("/api", func(r chi.Router) {

		r.Route("/admin", func(r chi.Router) {
			r.Post("/auth", api_handler.AuthHandler)

			r.Group(func(r chi.Router) {
				r.Use(app_middleware.AccessControlMiddleware)

				r.Route("/post", func(r chi.Router) {
					r.Post("/create", api_handler.HandlePostCreate)
					r.Patch("/{id}", api_handler.HandlePostUpdate)
					r.Delete("/{id}", api_handler.HandlePostDelete)

					r.Post("/{id}/attach", api_handler.UploadAttachHandler)
				})

				r.Route("/page", func(r chi.Router) {
					r.Post("/create", api_handler.HandlePageCreate)
					r.Post("/reorder", api_handler.HandlePageReorder)
					r.Post("/unsort", api_handler.HandlePageUnsort)
					r.Delete("/{id}", api_handler.HandlePageDelete)
				})

				r.Route("/category", func(r chi.Router) {
					r.Post("/", api_handler.HandleCategoryCreate)
					r.Delete("/{id}", api_handler.HandleCategoryDelete)
					r.Patch("/{id}", api_handler.HandleCategoryUpdate)
				})

				r.Route("/setting", func(r chi.Router) {
					r.Post("/", api_handler.HandleSettingCreate)
					r.Delete("/{id}", api_handler.HandleSettingDelete)
					r.Patch("/{id}", api_handler.HandleSettingUpdate)
				})
			})
		})
	})

	// TODO share link system. use static url to route share link
	return r
}
