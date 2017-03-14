package sqlite3

import "database/sql"

//TodoData : Todo List Data struct
type TodoData struct {
	id       int
	title    string
	doneYN   bool
	createDt string
	updateDt string
}

//TodoDataList : TodoData's slice
type TodoDataList []TodoData

func NewTodo() {

}

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

func (todo *TodoData) GetTodoList(db *sql.DB) (todoList TodoDataList) {
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

	for rows.Next() {
		rowsErr := rows.Scan(&todo.id, &todo.title, &todo.doneYN, &todo.createDt, &todo.updateDt)
		procError(rowsErr)

		todoList = append(todoList, *todo)
	}
	return todoList
}

func (todo *TodoData) AddTodoData(db *sql.DB) {

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
