package router

import (
	"HtmxBlog/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		template.Tmpl.ExecuteTemplate(w, "index", template.App{PageTitle: "Hello World", Navigation: []template.NavigationItem{{Name: "Home", Url: "/"}, {Name: "About", Url: "/about"}}})
	})

	return r
}
