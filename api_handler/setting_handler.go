package api_handler

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"HtmxBlog/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HandleSettingCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	key := r.FormValue("key")
	value := r.FormValue("value")

	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	err = services.CreateSettings(&model.Setting{
		Key:   key,
		Value: value,
	})
	if err != nil {
		http.Error(w, "Failed to create setting", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`<div class="text-green-600 font-bold p-4 bg-green-50 rounded shadow-md border border-green-200">Setting created successfully!</div>`))
	w.Header().Set("HX-Trigger", "newSetting")
}

func HandleSettingDelete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Setting ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Setting ID", http.StatusBadRequest)
		return
	}

	if err := services.DeleteSettings(uint(id)); err != nil {
		http.Error(w, "Failed to delete setting", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

func HandleSettingUpdate(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Setting ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Setting ID", http.StatusBadRequest)
		return
	}

	setting, err := services.ReadSetting(uint(id))
	if err != nil {
		http.Error(w, "Setting not found", http.StatusNotFound)
		return
	}

	if key := r.FormValue("key"); key != "" {
		setting.Key = key
	}
	if value := r.FormValue("value"); value != "" {
		setting.Value = value
	}

	err = services.UpdateSettings(setting)
	if err != nil {
		http.Error(w, "Failed to update setting", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	template.Tmpl.ExecuteTemplate(w, "setting_item", setting)
}
