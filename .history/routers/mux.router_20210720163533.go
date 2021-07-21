package routers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Router interface {
	Get(uri string, callback func(rw http.ResponseWriter, r *http.Request))
	Post(uri string, callback func(rw http.ResponseWriter, r *http.Request))
	RunServer(port string) error
}

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) Get(uri string, callback func(rw http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, callback).Methods("GET")
}

func (*muxRouter) Post(uri string, callback func(rw http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, callback).Methods("POST")
}

func (*muxRouter) RunServer(port string) error {
	fmt.Printf("Server running in port %v", port)
	return http.ListenAndServe(port, muxDispatcher)
}
