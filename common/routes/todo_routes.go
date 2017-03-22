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
		handlerFunc: controller.GetTodoList,
	},
	TodoRoute{
		method:      "POST",
		path:        "/add",
		handlerFunc: controller.AddTodoData,
	},
	TodoRoute{
		method:      "POST",
		path:        "/delete",
		handlerFunc: controller.DeleteTodoData,
	},
	TodoRoute{
		method:      "POST",
		path:        "/update",
		handlerFunc: controller.UpdateTodoData,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/web").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))

	for _, route := range routeInfo {
		router.Methods(route.method).Path(route.path).HandlerFunc(route.handlerFunc)
	}
	return router
}
