package services

import "todoGolang/todo/model"

type TodoService interface {
	GetTodoList() model.TodoDataList
	GetTodoListJson() model.TodoJsonList
	AddTodoData(addParam *model.TodoData) bool
	UpdateTodoData(todoParam *model.TodoData) bool
	DeleteTodoData(todoParam *model.TodoData) bool
}
