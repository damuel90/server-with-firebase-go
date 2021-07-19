package controllers

import (
	"encoding/json"
	"net/http"
	"server-with-firebase-go/entities"
	"server-with-firebase-go/repositories"
)

var (
	posts []entities.Post
	repo  = repositories.NewPostRepository()
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
	post.Id = len(posts) + 1
	posts = append(posts, post)
	rw.WriteHeader(http.StatusCreated)
	result, _ := json.Marshal(post)
	rw.Write(result)
}
