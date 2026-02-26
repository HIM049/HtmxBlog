package api_handler

import (
	"HtmxBlog/config"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var CurrentToken string
var CreateTime time.Time

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	passwd := r.FormValue("password")
	if passwd == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`<div class="text-red-500">密码不能为空</div>`))
		return
	}

	newToken := uuid.New().String()

	if passwd == config.Cfg.Service.AdminPasswd {
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    newToken,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   3600 * config.Cfg.Service.ValidTime,
		})

		CurrentToken = newToken
		CreateTime = time.Now()

		w.Header().Set("HX-Redirect", "/admin")
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(`<div class="text-red-500">密码错误</div>`))
}
