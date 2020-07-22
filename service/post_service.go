package service

import (
	"math/rand"

	"github.com/baldore/pragmatic-reviews-golang/entity"
	"github.com/baldore/pragmatic-reviews-golang/repository"
)

type PostService interface {
	FindAll() ([]entity.Post, error)
	AddPost(post *entity.Post) error
}

type service struct {
	posts repository.PostRepository
}

func NewPostService(posts repository.PostRepository) PostService {
	return &service{
		posts: posts,
	}
}

func (s *service) FindAll() ([]entity.Post, error) {
	return s.posts.FindAll()
}

func (s *service) AddPost(post *entity.Post) error {
	post.ID = rand.Int63()
	_, err := s.posts.Save(post)
	return err
}
