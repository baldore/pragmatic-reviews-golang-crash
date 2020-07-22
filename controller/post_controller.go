package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/baldore/pragmatic-reviews-golang/entity"
	"github.com/baldore/pragmatic-reviews-golang/service"
)

type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	postService service.PostService
}

func NewPostController(ps service.PostService) PostController {
	return &controller{
		postService: ps,
	}
}

func (c *controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	posts, err := c.postService.FindAll()
	if err != nil {
		http.Error(w, "Error getting posts", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, "Error encoding posts", http.StatusInternalServerError)
	}
}

func (c *controller) AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Error marshalling json", http.StatusInternalServerError)
		return
	}

	err = c.postService.AddPost(&post)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error saving post: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}
