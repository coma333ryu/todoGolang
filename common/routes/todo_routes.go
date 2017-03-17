package routes

import (
	"net/http"
	"todoGolang/todo/controller"

	"github.com/gorilla/mux"
)

type TodoRoute struct {
	method      string
	path        string
	handlerFunc http.HandlerFunc
}

type TodoRoutes []TodoRoute

var routeInfo = TodoRoutes{
	TodoRoute{
		method:      "GET",
		path:        "/",
		handlerFunc: controller.Index,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routeInfo {
		router.Methods(route.method).
			Path(route.path).
			HandlerFunc(route.handlerFunc)
	}
	return router
}
