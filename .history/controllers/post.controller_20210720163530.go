package controllers

import (
	"encoding/json"
	"net/http"
	"server-with-firebase-go/entities"
	"server-with-firebase-go/errors"
	"server-with-firebase-go/services"
)

type controller struct{}

var (
	postService = services.NewPostService()
)

type PostController interface {
	GetPost(rw http.ResponseWriter, r *http.Request)
	CreatePost(rw http.ResponseWriter, r *http.Request)
}

func NewPostController() PostController {
	return &controller{}
}

func (*controller) GetPost(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(errors.ServiceError{Message: "Error al obtener la lista de publicaciones"})
		return
	}
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(posts)
}

func (*controller) CreatePost(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	var post entities.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(errors.ServiceError{Message: "Error al obtener al decodificar la petición"})
		return
	}

	err = postService.Validate(&post)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(errors.ServiceError{Message: err.Error()})
		return
	}
	createdPost, err := postService.Create(&post)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(errors.ServiceError{Message: "Error al crear la publicación"})
		return
	}
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(createdPost)
}
