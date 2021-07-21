package services

import (
	"errors"
	"math/rand"
	"server-with-firebase-go/entities"
	"server-with-firebase-go/repositories"
	"strconv"
)

type PostService interface {
	Validate(post *entities.Post) error
	Create(post *entities.Post) (*entities.Post, error)
	FindAll() ([]entities.Post, error)
}

type service struct{}

var (
	repo = repositories.NewFirestoreRepository()
)

func NewPostService() PostService {
	return &service{}
}

func (*service) Validate(post *entities.Post) error {
	if post == nil {
		err := errors.New("La publicación esta vacia")
		return err
	}
	if post.Title == "" {
		err := errors.New("El campo title es obligatorio")
		return err
	}
	return nil
}

func (*service) Create(post *entities.Post) (*entities.Post, error) {
	post.Id = strconv.Itoa(rand.Int())
	return repo.Save(post)
}

func (*service) FindAll() ([]entities.Post, error) {
	return repo.FindAll()
}
