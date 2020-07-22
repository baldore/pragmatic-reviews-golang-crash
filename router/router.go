package router

import "net/http"

type Router interface {
	Get(route string, f func(w http.ResponseWriter, r *http.Request))
	Post(route string, f func(w http.ResponseWriter, r *http.Request))
	ListenAndServe(port string) error
}
