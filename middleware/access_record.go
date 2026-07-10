package middleware

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
)

func AccessRecordMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Real-IP")
		if ip == "" {
			ip = r.RemoteAddr
		}

		ua := r.UserAgent()
		referer := r.Referer()
		method := r.Method
		path := r.URL.Path
		query := r.URL.RawQuery
		createdAt := time.Now()

		go func() {
			record := model.AccessRecord{
				CreatedAt: createdAt,
				IP:        ip,
				UserAgent: ua,
				Referer:   referer,
				Method:    method,
				Path:      path,
				Query:     query,
			}

			err := config.DB.Create(&record).Error
			if err != nil {
				log.Errorf("Failed to create record: %v", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
