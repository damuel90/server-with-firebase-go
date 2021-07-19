package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	port := ":8080"
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Hola mundo!")
	})
	router.HandleFunc("/posts", getPost).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")

	fmt.Println("Server running in port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
