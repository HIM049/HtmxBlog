package middleware

import (
	"net/http"
)

func NotFoundInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		iw := &interceptorWriter{ResponseWriter: w, status: 200}
		next.ServeHTTP(iw, r)

		if iw.status == http.StatusNotFound && !iw.wroteBody {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	})
}

type interceptorWriter struct {
	http.ResponseWriter
	status    int
	wroteBody bool
}

func (i *interceptorWriter) WriteHeader(code int) {
	i.status = code
	if code != http.StatusNotFound {
		i.ResponseWriter.WriteHeader(code)
	}
}

func (i *interceptorWriter) Write(b []byte) (int, error) {
	if i.status == http.StatusNotFound {
		return len(b), nil
	}
	i.wroteBody = true
	return i.ResponseWriter.Write(b)
}
