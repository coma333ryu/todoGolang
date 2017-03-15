package main

import (
	"fmt"
	"todoGolang/constants"
	"todoGolang/databases/sqlite3"
	"todoGolang/validaters"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	todoData := new(sqlite3.TodoData)

	db := todoData.CreateDatabase(constants.SqliteFileInfo)
	defer db.Close()
	todoData.CreateTable(db)

	todoAddParam := sqlite3.NewTodoData("", "", false)

	vErr := validaters.CheckParam(todoAddParam, "A")

	if vErr == nil {
		bAddResult := todoData.AddTodoData(db, todoAddParam)

		if bAddResult {
			fmt.Println("Insert success")
		} else {
			fmt.Println("Insert fail")
		}
	} else {
		fmt.Println(vErr)
	}

	resultTodoList := todoData.GetTodoList(db)

	fmt.Println("resultTodoList =========> ", resultTodoList)

	todoUpdateParma := sqlite3.NewTodoData("3", "", true)
	bUpdateResult := todoData.UpdateTodoData(db, todoUpdateParma)

	if bUpdateResult {
		fmt.Println("Update success")
	} else {
		fmt.Println("Update fail")
	}

	resultTodoList = todoData.GetTodoList(db)

	fmt.Println("resultTodoList =========> ", resultTodoList)

	todoDeleteParma := sqlite3.NewTodoData("3", "", false)

	bDeleteResult := todoData.DeleteTodoData(db, todoDeleteParma)

	if bDeleteResult {
		fmt.Println("Delete success")
	} else {
		fmt.Println("Delete fail")
	}

	resultTodoList = todoData.GetTodoList(db)

	fmt.Println("resultTodoList =========> ", resultTodoList)
}
