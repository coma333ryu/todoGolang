package services

import "todoGolang/todo/model"

type TodoServicer interface {
	GetTodoList() model.TodoDataList
	AddTodoData(addParam *model.TodoData) bool
	UpdateTodoData(todoParam *model.TodoData) bool
	DeleteTodoData(todoParam *model.TodoData) bool
}
