package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct {
	mux *mux.Router
}

func NewGorillaMuxRouter() Router {
	return &muxRouter{
		mux: mux.NewRouter(),
	}
}

func (m *muxRouter) Get(route string, f func(w http.ResponseWriter, r *http.Request)) {
	m.mux.HandleFunc(route, f).Methods(http.MethodGet)
}

func (m *muxRouter) Post(route string, f func(w http.ResponseWriter, r *http.Request)) {
	m.mux.HandleFunc(route, f).Methods(http.MethodPost)
}

func (m *muxRouter) ListenAndServe(port string) error {
	return http.ListenAndServe(port, m.mux)
}
