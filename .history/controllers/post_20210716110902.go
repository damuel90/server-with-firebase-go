package controllers

import (
	"encoding/json"
	"net/http"
	"server-with-firebase-go/entities"
	"server-with-firebase-go/repositories"
)

var (
	posts []entities.Post
	repo  repositories.PostRepository = repositories.NewPostRepository()
)

func GetPost(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	//posts, err := repo.Fi
	/*result, err := json.Marshal(posts)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error": "Error marshaling the posts array"}`))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(result)*/
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
