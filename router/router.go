package router

import "net/http"

func Init() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<h1>Hello World</h1>"))
	})

	r.HandleFunc("GET /posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<h1>Posts</h1>"))
	})

	return r
}
