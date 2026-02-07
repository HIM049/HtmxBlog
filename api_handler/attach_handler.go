package api_handler

import (
	"HtmxBlog/database"
	"HtmxBlog/model"
	"HtmxBlog/services"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// LoadAttachHandler handle browser request to load attach file
func LoadAttachHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	attach, err := database.ReadAttachById(uint(id))
	if err != nil {
		http.Error(w, "Attach not found", http.StatusNotFound)
		return
	}

	if attach.Permission == model.PermissionPrivate {
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
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	attach, err := services.UploadAttach(&file, header.Filename, header.Header.Get("Content-Type"))
	if err != nil {
		http.Error(w, "Failed to upload attach", http.StatusInternalServerError)
		return
	}

	result := fmt.Sprintf("<div>Attach id %d, name %s, Hash %s, Uid %s, Mime %s</div>", attach.ID, attach.Name, attach.Hash, attach.Uid, attach.Mime)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(result))
}

// TODO file reference system
