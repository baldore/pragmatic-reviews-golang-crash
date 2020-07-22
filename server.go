package main

import (
	"fmt"
	"log"

	"github.com/baldore/pragmatic-reviews-golang/controller"
	"github.com/baldore/pragmatic-reviews-golang/repository"
	"github.com/baldore/pragmatic-reviews-golang/router"
	"github.com/baldore/pragmatic-reviews-golang/service"
)

const (
	port string = ":9090"
)

var (
	r              router.Router             = router.NewGorillaMuxRouter()
	posts          repository.PostRepository = repository.NewPostRepository()
	postService    service.PostService       = service.NewPostService(posts)
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	r.Get("/posts", postController.GetPosts)
	r.Post("/posts", postController.AddPost)

	fmt.Printf("Listening on port %s\n", port)
	err := r.ListenAndServe(port)
	if err != nil {
		log.Fatal(err)
	}
}
