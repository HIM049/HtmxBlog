package api_handler

import (
	"HtmxBlog/services"
	"fmt"
	"net/http"
)

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
