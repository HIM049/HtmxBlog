package router

import (
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
