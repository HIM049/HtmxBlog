package api_handler

import (
	"HtmxBlog/config"
	"net/http"
)

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

	if passwd == config.Cfg.Service.AdminPasswd {
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    config.Cfg.Service.AdminPasswd,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   3600 * 24, // 24 hours
		})
		w.Header().Set("HX-Redirect", "/admin")
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(`<div class="text-red-500">密码错误</div>`))
}
