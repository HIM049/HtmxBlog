package middleware

import (
	"HtmxBlog/api_handler"
	"net/http"
	"time"
)

func AccessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil || cookie.Value != api_handler.CurrentToken || time.Since(api_handler.CreateTime) > 24*time.Hour {
			http.Redirect(w, r, "/admin/auth", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
