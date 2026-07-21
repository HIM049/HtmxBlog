package handler

import (
	"HtmxBlog/config"
	"HtmxBlog/state"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func AuthView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	state.AdminTmpl.ExecuteTemplate(w, "auth", nil)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	passwd := r.FormValue("password")
	if passwd == "" {
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

		state.CurrentToken = newToken
		state.CreateTime = time.Now()

		w.Header().Set("HX-Redirect", "/admin")
		return
	}

	w.Write([]byte(`<div class="text-red-500">密码错误</div>`))
}
