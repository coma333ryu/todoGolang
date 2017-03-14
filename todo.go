package main

import (
	"fmt"
	"todoGolang/constants"
	"todoGolang/databases/sqlite3"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := sqlite3.CreateDatabase(constants.SqliteFileInfo)
	defer db.Close()

	sqlite3.CreateTable(db)

	todoData := new(sqlite3.TodoData)

	result := todoData.GetTodoList(db)

	sqlite3.TodoData{id: 1}
	fmt.Println("resultresultresult", result)
}
