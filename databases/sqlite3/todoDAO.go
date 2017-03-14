package sqlite3

import (
	"database/sql"
	"fmt"
)

//TodoData : Todo List Data struct
type TodoData struct {
	id       int
	Title    string
	DoneYN   bool
	createDt string
	updateDt string
}

//TodoDataList : TodoData's slice
type TodoDataList []TodoData

//CreateDatabase : create Sqlite3 Db file about filePath
func CreateDatabase(filePath string) *sql.DB {
	sqliteDB, err := sql.Open("sqlite3", filePath)
	procError(err)
	return sqliteDB
}

//CreateTable : create table about tb_todo
func CreateTable(db *sql.DB) {
	tableSQL := `
	CREATE TABLE IF NOT EXISTS tb_todo (
        id integer primary key autoincrement, 
        title text not null,
        doneYN boolean not null,
        createDt varchar(14) not null,
        updateDt varchar(14) not null
    );
	`
	_, err := db.Exec(tableSQL)
	procError(err)
}

func GetTodoList(db *sql.DB) (todoList TodoDataList) {
	getSQL := `
		SELECT
			id
			, title
			, doneYN
			, createDt
			, updateDt
		FROM tb_todo;
	`

	rows, err := db.Query(getSQL)
	procError(err)
	defer rows.Close()

	todoData := new(TodoData)
	for rows.Next() {
		rowsErr := rows.Scan(&todoData.id, &todoData.Title, &todoData.DoneYN, &todoData.createDt, &todoData.updateDt)
		procError(rowsErr)

		todoList = append(todoList, *todoData)
	}
	return todoList
}

func AddTodoData(db *sql.DB, addParam *TodoData) {
	addSQL := `
		INSERT INTO tb_todo (
			title
			, doneYN
			, createDt
			, updateDt
		) VALUES (
			?
			, ?
			, strftime('%Y%m%d%H%M%S','now')
			, strftime('%Y%m%d%H%M%S','now')
		);
	`

	stmt, prepareErr := db.Prepare(addSQL)
	procError(prepareErr)
	res, execErr := stmt.Exec(addParam.Title, addParam.DoneYN)
	defer stmt.Close()
	procError(execErr)
	id, insertResultErr := res.LastInsertId()
	procError(insertResultErr)

	fmt.Println("ididididid", id)
}

func UpdateTodoData() {

}

func DeleteTodoData() {

}

func procError(err error) {
	if err != nil {
		panic(err)
	}
}
