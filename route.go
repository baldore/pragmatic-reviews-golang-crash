package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{
		{
			Id:    1,
			Title: "Common form mistakes",
			Text:  "Stop messing around with forms",
		}, {
			Id:    2,
			Title: "Doing a sudoku with Typescript",
			Text:  "This is great",
		},
	}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, "Error marshalling json", http.StatusInternalServerError)
		return
	}
}

func addPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Error marshalling json", http.StatusInternalServerError)
		return
	}

	lp := posts[len(posts)-1]
	post.Id = lp.Id + 1
	posts = append(posts, post)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshalling json"))
	}
}
