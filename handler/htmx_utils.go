package handler

import "net/http"

// HtmxError writes an error message with a 200 OK status code so that HTMX
// will properly render the error message in the target element.
func HtmxError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<div class="text-red-600 font-bold p-4 bg-red-50 rounded shadow-md border border-red-200">` + msg + `</div>`))
}

// HtmxSuccess writes a success message with a 200 OK status code.
func HtmxSuccess(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<div class="text-green-600 font-bold p-4 bg-green-50 rounded shadow-md border border-green-200">` + msg + `</div>`))
}
