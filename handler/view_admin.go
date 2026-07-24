package handler

import (
	"HtmxBlog/state"
	"bytes"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
)

func GenericAdminView(w http.ResponseWriter, r *http.Request) {
	pageName := chi.URLParam(r, "name")
	if pageName == "" {
		pageName = "admin"
	}

	app := NewAdminApp(r, pageName)

	var buf bytes.Buffer
	err := state.AdminTmpl.ExecuteTemplate(&buf, pageName, app)
	if err != nil {
		http.Error(w, "Page not found or error loading data", http.StatusNotFound)
		log.Errorf("[AdminTmpl] Page (%s) error: %v", pageName, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}
