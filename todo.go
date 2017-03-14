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

	// todoData := new(sqlite3.TodoData)

	todoParam := sqlite3.TodoData{
		Title:  "aaa",
		DoneYN: false,
	}
	sqlite3.AddTodoData(db, &todoParam)

	result := sqlite3.GetTodoList(db)

	fmt.Println("resultresultresult", result)
}
