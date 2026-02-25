package router

import (
	"HtmxBlog/services"
	"HtmxBlog/view_handler"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

type HotRouter struct {
	mux  *chi.Mux
	lock sync.RWMutex
}

func (h *HotRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.lock.RLock()
	defer h.lock.RUnlock()

	h.mux.ServeHTTP(w, r)
}

func NewHotRouter() *HotRouter {
	return &HotRouter{}
}

func (h *HotRouter) Update(newMux *chi.Mux) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.mux = newMux
}

// RegisterPagesRouter registers all pages as routes.
// It will panic if any error occurs.
func RegisterPagesRouter(r chi.Router) error {
	pages, err := services.ReadAllPages()
	if err != nil {
		return err
	}

	for _, page := range pages {
		r.Get(page.Route, view_handler.GenericViewLoader(page.Template))
	}
	return nil
}
