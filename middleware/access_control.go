package middleware

import (
	"HtmxBlog/config"
	"net/http"
)

func AccessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil || cookie.Value != config.Cfg.Service.AdminPasswd {
			http.Redirect(w, r, "/admin/auth", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
