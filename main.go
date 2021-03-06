package main

import (
	"fmt"
	"log"
	"net/http"
	"server-with-firebase-go/controllers"
	"server-with-firebase-go/repositories"
	"server-with-firebase-go/routers"
	"server-with-firebase-go/services"
)

var (
	muxRouter      = routers.NewMuxRouter()
	postRepository = repositories.NewFirestoreRepository()
	postService    = services.NewPostService(postRepository)
	postController = controllers.NewPostController(postService)
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
