package main

import (
	"fmt"
	"todoGolang/common/constants"
	"todoGolang/common/databases/sqlite3"
	"todoGolang/common/validaters"

	"todoGolang/todo/model"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqlite3Db := sqlite3.CreateDatabase(constants.SqliteFileInfo)
	defer sqlite3Db.Close()
	sqlite3.CreateTable(sqlite3Db)

	todoData := new(model.TodoData)
	todoParam := model.NewTodoData("", "new new new", false)

	paramErr := validaters.CheckParam(todoParam, "A")

	if paramErr == nil {
		bAddResult := todoData.AddTodoData(sqlite3Db, todoParam)

		if bAddResult {
			fmt.Println("Insert success")
		} else {
			fmt.Println("Insert fail")
		}
	} else {
		fmt.Println(paramErr)
	}

	resultListTodo := todoData.GetTodoList(sqlite3Db)
	fmt.Println("insert result list  =========> ", resultListTodo)

	todoParam = model.NewTodoData("4", "", true)
	paramErr = validaters.CheckParam(todoParam, "U")

	if paramErr == nil {
		bAddResult := todoData.UpdateTodoData(sqlite3Db, todoParam)

		if bAddResult {
			fmt.Println("Update success")
		} else {
			fmt.Println("Update fail")
		}
	} else {
		fmt.Println(paramErr)
	}

	resultListTodo = todoData.GetTodoList(sqlite3Db)
	fmt.Println("update result list  =========> ", resultListTodo)

	todoParam = model.NewTodoData("4", "", true)
	paramErr = validaters.CheckParam(todoParam, "D")

	if paramErr == nil {
		bAddResult := todoData.DeleteTodoData(sqlite3Db, todoParam)

		if bAddResult {
			fmt.Println("Delete success")
		} else {
			fmt.Println("Delete fail")
		}
	} else {
		fmt.Println(paramErr)
	}

	resultListTodo = todoData.GetTodoList(sqlite3Db)
	fmt.Println("Delete result list  =========> ", resultListTodo)
}
