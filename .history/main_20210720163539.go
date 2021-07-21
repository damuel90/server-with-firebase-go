package main

import (
	"fmt"
	"log"
	"net/http"
	"server-with-firebase-go/controllers"
	"server-with-firebase-go/routers"
)

var (
	muxRouter      = routers.NewMuxRouter()
	postController = controllers.NewPostController()
)

func main() {
	port := ":8080"
	muxRouter.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Hola mundo!")
	})
	muxRouter.Get("/posts", postController.GetPost)
	muxRouter.Post("/posts", postController.CreatePost)

	log.Fatalln(muxRouter.RunServer(port))
}
