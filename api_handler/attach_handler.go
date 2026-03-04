package api_handler

import (
	"HtmxBlog/model"
	"HtmxBlog/services"
	"HtmxBlog/template"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// LoadAttachHandler handle browser request to load attach file
func LoadAttachHandler(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")

	attach, err := services.ReadAttachByUid(uid)
	if err != nil {
		http.Error(w, "Attach not found", http.StatusNotFound)
		return
	}

	if attach.Permission == model.VisibilityPrivate {
		http.Error(w, "Access Forbidden", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", attach.Mime)

	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", attach.Name))

	filePath := filepath.Join(services.ATTACHES_DIR, attach.Uid)
	http.ServeFile(w, r, filePath)
}

// UploadAttachHandler handles the upload of an attach file.
func UploadAttachHandler(w http.ResponseWriter, r *http.Request) {
	// Get id and get post
	postId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Post ID", http.StatusBadRequest)
		return
	}

	// Get file from form
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	attach, err := services.CreateAttach(&file, header.Filename, header.Header.Get("Content-Type"), uint(postId))
	if err != nil {
		http.Error(w, "Failed to upload attach", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusCreated)
	template.Tmpl.ExecuteTemplate(w, "attach_item", attach)
}

// TODO file reference system
