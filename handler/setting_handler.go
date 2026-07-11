package handler

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"HtmxBlog/state"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// ManageSettingsView renders the settings management page skeleton.
func ManageSettingsView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	state.AdminTmpl.ExecuteTemplate(w, "setting_manage", nil)
}

// SettingListComponent renders the settings list fragment.
func SettingListComponent(w http.ResponseWriter, r *http.Request) {
	settings, _ := services.ReadAllSettings()
	w.Header().Set("Content-Type", "text/html")
	state.AdminTmpl.ExecuteTemplate(w, "setting_list", map[string]interface{}{
		"Settings": settings,
	})
}

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

	err = services.CreateSetting(&model.Setting{
		Key:   key,
		Value: value,
	})
	if err != nil {
		http.Error(w, "Failed to create setting", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("HX-Trigger", "settingChanged")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`<div class="text-green-600 font-bold p-4 bg-green-50 rounded shadow-md border border-green-200">Setting created successfully!</div>`))
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

	if err := services.DeleteSetting(uint(id)); err != nil {
		http.Error(w, "Failed to delete setting", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "settingChanged")
	w.WriteHeader(http.StatusOK)
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

	err = services.UpdateSetting(setting)
	if err != nil {
		http.Error(w, "Failed to update setting", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "settingChanged")
	w.WriteHeader(http.StatusOK)
}
