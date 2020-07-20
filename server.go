package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port = ":9090"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "yes")
	})

	r.HandleFunc("/posts", getPosts).Methods(http.MethodGet)
	r.HandleFunc("/posts", addPost).Methods(http.MethodPost)

	fmt.Printf("Listening on port %s", port)
	http.ListenAndServe(port, r)
}
