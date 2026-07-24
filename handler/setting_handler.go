package handler

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HandleSettingCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		HtmxError(w, "Invalid form data")
		return
	}

	key := r.FormValue("key")
	value := r.FormValue("value")

	if key == "" {
		HtmxError(w, "Key is required")
		return
	}

	err = services.CreateSetting(&model.Setting{
		Key:   key,
		Value: value,
	})
	if err != nil {
		HtmxError(w, "Failed to create setting")
		return
	}

	go func() {
		services.UpdateConfig()
		services.UpdateSettings()
	}()

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("HX-Trigger", "settingChanged")
	HtmxSuccess(w, "Setting created successfully!")
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

	go func() {
		services.UpdateConfig()
		services.UpdateSettings()
	}()

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

	go func() {
		services.UpdateConfig()
		services.UpdateSettings()
	}()

	w.Header().Set("HX-Trigger", "settingChanged")
	w.WriteHeader(http.StatusOK)
}
