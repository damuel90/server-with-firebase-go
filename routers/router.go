package routers

import "net/http"

type Router interface {
	Get(uri string, callback func(rw http.ResponseWriter, r *http.Request))
	Post(uri string, callback func(rw http.ResponseWriter, r *http.Request))
	RunServer(port string) error
}
