package handler

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HandleRedirectCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	sourcePath := r.FormValue("source_path")
	targetPath := r.FormValue("target_path")
	statusCodeStr := r.FormValue("status_code")

	if sourcePath == "" || targetPath == "" {
		http.Error(w, "Source and target paths are required", http.StatusBadRequest)
		return
	}

	statusCode := 301
	if statusCodeStr != "" {
		statusCode, err = strconv.Atoi(statusCodeStr)
		if err != nil || (statusCode != 301 && statusCode != 302) {
			http.Error(w, "Status code must be 301 or 302", http.StatusBadRequest)
			return
		}
	}

	err = services.CreateRedirect(sourcePath, targetPath, statusCode)
	if err != nil {
		http.Error(w, "Failed to create redirect", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "redirectChanged")
	w.WriteHeader(http.StatusCreated)
}

func HandleRedirectDelete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Redirect ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Redirect ID", http.StatusBadRequest)
		return
	}

	if err := services.DeleteRedirect(uint(id)); err != nil {
		http.Error(w, "Failed to delete redirect", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "redirectChanged")
	w.WriteHeader(http.StatusOK)
}

func HandleRedirectUpdate(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Redirect ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Redirect ID", http.StatusBadRequest)
		return
	}

	redirects, err := services.ReadAllRedirects()
	if err != nil {
		http.Error(w, "Failed to read redirects", http.StatusInternalServerError)
		return
	}

	var target *model.Redirect
	for i := range redirects {
		if redirects[i].ID == uint(id) {
			target = &redirects[i]
			break
		}
	}
	if target == nil {
		http.Error(w, "Redirect not found", http.StatusNotFound)
		return
	}

	sourcePath := r.FormValue("source_path")
	targetPath := r.FormValue("target_path")
	statusCodeStr := r.FormValue("status_code")

	if sourcePath != "" {
		target.SourcePath = sourcePath
	}
	if targetPath != "" {
		target.TargetPath = targetPath
	}
	if statusCodeStr != "" {
		statusCode, err := strconv.Atoi(statusCodeStr)
		if err == nil && (statusCode == 301 || statusCode == 302) {
			target.StatusCode = statusCode
		}
	}

	err = services.UpdateRedirect(uint(id), target.SourcePath, target.TargetPath, target.StatusCode)
	if err != nil {
		http.Error(w, "Failed to update redirect", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "redirectChanged")
	w.WriteHeader(http.StatusOK)
}
