package middleware

import (
	"HtmxBlog/services"
	"net/http"
)

// RedirectInterceptor checks the request path against the redirect rules in the database.
// If a matching rule is found, it issues the configured redirect (301/302).
// Otherwise, it passes the request to the next handler.
func RedirectInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		redirect, err := services.FindRedirectBySource(r.URL.Path)
		if err == nil {
			http.Redirect(w, r, redirect.TargetPath, redirect.StatusCode)
			return
		}
		next.ServeHTTP(w, r)
	})
}
