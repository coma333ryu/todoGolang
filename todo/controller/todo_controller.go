package controller

import (
	"html/template"
	"net/http"
	"path"
)

func GetTodoList(res http.ResponseWriter, req *http.Request) {
	view := path.Join("templates", "todo.html")

	todoTemplates := template.Must(template.ParseFiles(view))
	todoTemplates.Execute(res, "")
}

func AddTodoData(res http.ResponseWriter, req *http.Request) {
	view := path.Join("templates", "todo.html")

	todoTemplates := template.Must(template.ParseFiles(view))
	todoTemplates.Execute(res, "")
}

func UpdateTodoData(res http.ResponseWriter, req *http.Request) {
	view := path.Join("templates", "todo.html")

	todoTemplates := template.Must(template.ParseFiles(view))
	todoTemplates.Execute(res, "")
}

func DeleteTodoData(res http.ResponseWriter, req *http.Request) {
	view := path.Join("templates", "todo.html")

	todoTemplates := template.Must(template.ParseFiles(view))
	todoTemplates.Execute(res, "")
}
