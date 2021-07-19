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
	posts = []Post{{Id: 1, Title: "Titulo 1", Text: "Texto 1"}}
}

func getPost(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error": "Error marshaling the posts array"}`))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(result)
}

func addPost(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	var post Post
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
