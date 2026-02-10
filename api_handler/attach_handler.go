package api_handler

import (
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

	url := fmt.Sprintf("/attach/%s", attach.Uid)
	markdown := fmt.Sprintf("![%s](%s)", attach.Name, url)

	result := fmt.Sprintf(`
		<div class="flex items-center justify-between p-4 bg-gray-50 rounded border border-gray-200">
			<div>
				<div class="font-bold text-gray-700">%s</div>
				<div class="text-xs text-gray-500">%s</div>
			</div>
			<button onclick="navigator.clipboard.writeText('%s'); this.innerText = 'Copied!'; setTimeout(() => this.innerText = 'Copy Link', 2000)"
				class="bg-blue-600 hover:bg-blue-700 text-white text-sm font-bold py-1 px-3 rounded focus:outline-none focus:shadow-outline transition-colors">
				Copy Link
			</button>
		</div>
	`, attach.Name, attach.Mime, markdown)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(result))
}

// TODO file reference system
