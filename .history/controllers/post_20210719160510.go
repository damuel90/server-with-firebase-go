package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"server-with-firebase-go/entities"
	"server-with-firebase-go/repositories"
)

var (
	repo = repositories.NewPostRepository()
)

func GetPost(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error": "Error al obtener la lista de publicaciones"}`))
		return
	}
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(posts)
}

func CreatePost(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	var post entities.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error": "Error unmarshaling the request"}`))
		return
	}

	post.Id = rand.Int()
	repo.Save(&post)
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(post)
}
