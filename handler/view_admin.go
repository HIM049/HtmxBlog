package handler

import (
	"HtmxBlog/services"
	"HtmxBlog/state"
	"net/http"
)

func AdminView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	state.AdminTmpl.ExecuteTemplate(w, "admin", nil)
}

func StatisticsView(w http.ResponseWriter, r *http.Request) {
	stats, err := services.GetStats()
	if err != nil {
		http.Error(w, "Failed to get statistics: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	state.AdminTmpl.ExecuteTemplate(w, "statistics", map[string]interface{}{
		"PageTitle": "Visit Statistics - Admin Dashboard",
		"Stats":     stats,
	})
}
