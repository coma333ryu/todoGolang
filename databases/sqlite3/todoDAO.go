package sqlite3

import "database/sql"

//TodoData : Todo List Data struct
type TodoData struct {
	todoIdx   string
	totoTitle string
	isDone    bool
	createDt  string
	updateDt  string
}

//TodoDataList : TodoData's slice
type TodoDataList []TodoData

type todoServicer interface {
	CreateDatabase(filePath string) *sql.DB
	CreateTable(db *sql.DB)
	GetTodoList(db *sql.DB) TodoDataList
	AddTodoData(db *sql.DB, addParam *TodoData) bool
	UpdateTodoData(db *sql.DB, todoParam *TodoData) bool
	DeleteTodoData(db *sql.DB, todoParam *TodoData) bool
}

//NewTodoData : create TodoData struct
func NewTodoData(idx string, title string, done bool) *TodoData {
	return &TodoData{
		todoIdx:   idx,
		totoTitle: title,
		isDone:    done,
	}
}

//TodoTitle : get Todo Title
func (todo *TodoData) TodoTitle() string {
	return todo.totoTitle
}

//TodoIdx : get Todo Idx
func (todo *TodoData) TodoIdx() string {
	return todo.todoIdx
}

//CreateDatabase : create Sqlite3 Database file about filePath
func (todo *TodoData) CreateDatabase(filePath string) *sql.DB {
	return createDatabase(filePath)
}

//CreateTable : create table about tb_todo table
func (todo *TodoData) CreateTable(db *sql.DB) {
	createTable(db)
}

//GetTodoList : get TodoDataList from tb_todo table
func (todo *TodoData) GetTodoList(db *sql.DB) TodoDataList {
	return getTodoList(db)
}

//AddTodoData : insert TodoData into tb_todo table
func (todo *TodoData) AddTodoData(db *sql.DB, todoParam *TodoData) bool {
	return addTodoData(db, todoParam)
}

//UpdateTodoData : update TodoData's isdone on tb_todo table
func (todo *TodoData) UpdateTodoData(db *sql.DB, todoParam *TodoData) bool {
	return updateTodo(db, todoParam)
}

//DeleteTodoData : delete TodoData from tb_todo table
func (todo *TodoData) DeleteTodoData(db *sql.DB, todoParam *TodoData) bool {
	return deleteTodo(db, todoParam)
}

func createDatabase(filePath string) *sql.DB {
	sqliteDB, err := sql.Open("sqlite3", filePath)
	procError(err)
	return sqliteDB
}

func createTable(db *sql.DB) {
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

func getTodoList(db *sql.DB) (todoList TodoDataList) {
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
		rowsErr := rows.Scan(&todoData.todoIdx, &todoData.totoTitle, &todoData.isDone, &todoData.createDt, &todoData.updateDt)
		procError(rowsErr)

		todoList = append(todoList, *todoData)
	}
	return todoList
}

func addTodoData(db *sql.DB, addParam *TodoData) bool {
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
	var bResult bool

	stmt, prepareErr := db.Prepare(addSQL)
	procError(prepareErr)

	res, execErr := stmt.Exec(addParam.totoTitle, addParam.isDone)
	defer stmt.Close()
	procError(execErr)

	id, insertResultErr := res.LastInsertId()
	procError(insertResultErr)

	if id != 0 {
		bResult = true
	} else {
		bResult = false
	}

	return bResult
}

func updateTodo(db *sql.DB, todoParam *TodoData) bool {
	var bResult bool

	updateSQL := "UPDATE tb_todo SET doneYN = ? where id = ?;"

	stmt, prepareErr := db.Prepare(updateSQL)
	procError(prepareErr)

	res, execErr := stmt.Exec(todoParam.isDone, todoParam.todoIdx)
	procError(execErr)

	defer stmt.Close()

	affect, affErr := res.RowsAffected()
	procError(affErr)

	if affect != 0 {
		bResult = true
	} else {
		bResult = false
	}
	return bResult
}

func deleteTodo(db *sql.DB, todoParam *TodoData) bool {
	var bResult bool
	deleteSQL := "DELETE FROM tb_todo where id = ?;"

	stmt, prepareErr := db.Prepare(deleteSQL)
	procError(prepareErr)

	res, execErr := stmt.Exec(todoParam.todoIdx)
	procError(execErr)

	defer stmt.Close()

	affect, affErr := res.RowsAffected()
	procError(affErr)

	if affect != 0 {
		bResult = true
	} else {
		bResult = false
	}
	return bResult
}

func procError(err error) {
	if err != nil {
		panic(err)
	}
}
