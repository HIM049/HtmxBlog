package router

import (
	"HtmxBlog/config"
	"HtmxBlog/handler"
	app_middleware "HtmxBlog/middleware"
	"HtmxBlog/services"
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
	r.Get("/attach/{id}", handler.LoadAttachHandler)

	r.Group(func(r chi.Router) {
		r.Use(app_middleware.NotFoundInterceptor)

		r.Group(func(r chi.Router) {
			r.Use(app_middleware.AccessRecordMiddleware)

			err := RegisterPagesRouter(r)
			if err != nil {
				panic("failed to register router: " + err.Error())
			}

			r.Get("/post/{id}", handler.PostView)
		})

		r.Route("/admin", func(r chi.Router) {
			r.Get("/auth", handler.AuthView)
			r.Group(func(r chi.Router) {
				r.Use(app_middleware.AccessControlMiddleware)

				r.Get("/", handler.AdminView)
				r.Get("/pages", handler.ManagePagesView)
				r.Get("/posts", handler.ManagePostsView)
				r.Get("/categories", handler.ManageCategoriesView)
				r.Get("/settings", handler.ManageSettingsView)
				r.Get("/comments", handler.ManageCommentsView)
				r.Get("/redirects", handler.ManageRedirectsView)
				r.Get("/statistics", handler.StatisticsView)
				r.Get("/post/{id}/edit", handler.EditView)

				r.Route("/fragment", func(r chi.Router) {
					r.Get("/category-list", handler.CategoryListComponent)
					r.Get("/setting-list", handler.SettingListComponent)
					r.Get("/redirect-list", handler.RedirectListComponent)
					r.Get("/comment-list", handler.CommentListComponent)
					r.Get("/post-list", handler.PostListComponent)
					r.Get("/page-list", handler.PageListComponent)
				})
			})
		})
	})

	r.Route("/api", func(r chi.Router) {
		r.Get("/version", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(config.APP_VERSION))
		})

		r.Post("/comment", handler.HandleCommentCreate)

		r.Route("/admin", func(r chi.Router) {
			r.Post("/auth", handler.AuthHandler)

			r.Group(func(r chi.Router) {
				r.Use(app_middleware.AccessControlMiddleware)

				r.Delete("/comment/{id}", handler.HandleCommentDelete)
				r.Post("/comment/{id}/approve", handler.HandleCommentApprove)

				r.Route("/post", func(r chi.Router) {
					r.Post("/create", handler.HandlePostCreate)
					r.Patch("/{id}", handler.HandlePostUpdate)
					r.Post("/{id}/publish", handler.HandlePostPublish)
					r.Delete("/{id}", handler.HandlePostDelete)

					r.Post("/{id}/attach", handler.UploadAttachHandler)
				})

				r.Route("/page", func(r chi.Router) {
					r.Post("/create", handler.HandlePageCreate)
					r.Post("/reorder", handler.HandlePageReorder)
					r.Post("/unsort", handler.HandlePageUnsort)
					r.Delete("/{id}", handler.HandlePageDelete)
				})

				r.Route("/category", func(r chi.Router) {
					r.Post("/", handler.HandleCategoryCreate)
					r.Delete("/{id}", handler.HandleCategoryDelete)
					r.Patch("/{id}", handler.HandleCategoryUpdate)
				})

				r.Route("/setting", func(r chi.Router) {
					r.Post("/", handler.HandleSettingCreate)
					r.Delete("/{id}", handler.HandleSettingDelete)
					r.Patch("/{id}", handler.HandleSettingUpdate)
				})

				r.Route("/redirect", func(r chi.Router) {
					r.Post("/", handler.HandleRedirectCreate)
					r.Delete("/{id}", handler.HandleRedirectDelete)
					r.Patch("/{id}", handler.HandleRedirectUpdate)
				})
			})
		})
	})

	// TODO share link system. use static url to route share link

	// Fallback: check redirect rules first, otherwise redirect to home
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		redirect, err := services.FindRedirectBySource(r.URL.Path)
		if err == nil {
			http.Redirect(w, r, redirect.TargetPath, redirect.StatusCode)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	})

	return r
}
