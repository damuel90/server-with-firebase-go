package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Router interface {
	GET(uri string, callback func(rw http.ResponseWriter, r *http.Request))
	POST(uri string, callback func(rw http.ResponseWriter, r *http.Request))
	SERVER(port string)
}

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, callback func(rw http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, callback).Methods("GET")
}

func (*muxRouter) POST(uri string, callback func(rw http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, callback).Methods("POST")
}

func (*muxRouter) SERVER(port string) {
	fmt.Printf("Server running in port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}
