package main

import (
	"todoGolang/database/sqlite3"

	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := sqlite3.CreateDatabase("./todo.db")
	defer db.Close()

	sqlite3.CreateTable(db)

	rows, err := db.Query("SELECT * FROM tb_todo")

	if err != nil {
		panic(err)
	}

	fmt.Println(rows)
}
