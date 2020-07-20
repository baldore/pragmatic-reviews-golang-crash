package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/baldore/pragmatic-reviews-golang/entity"
	"github.com/baldore/pragmatic-reviews-golang/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	posts, err := repo.FindAll()
	if err != nil {
		http.Error(w, "Error getting posts", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, "Error encoding posts", http.StatusInternalServerError)
	}
}

func addPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Error marshalling json", http.StatusInternalServerError)
		return
	}

	post.ID = rand.Int63()
	_, err = repo.Save(&post)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error saving post: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}
