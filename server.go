package main

import (
	"fmt"
	"net/http"

	"github.com/baldore/pragmatic-reviews-golang/controller"
	"github.com/baldore/pragmatic-reviews-golang/repository"
	"github.com/baldore/pragmatic-reviews-golang/service"
	"github.com/gorilla/mux"
)

var (
	port string = ":9090"

	posts          repository.PostRepository = repository.NewPostRepository()
	postService    service.PostService       = service.NewPostService(posts)
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/posts", postController.GetPosts).Methods(http.MethodGet)
	r.HandleFunc("/posts", postController.AddPost).Methods(http.MethodPost)

	fmt.Printf("Listening on port %s\n", port)
	http.ListenAndServe(port, r)
}
