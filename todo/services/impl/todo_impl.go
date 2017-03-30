package impl

import (
	"database/sql"
	"todoGolang/common/constants"
	"todoGolang/common/databases/sqliteDb"
	"todoGolang/common/exceptions"
	"todoGolang/todo/model"
	"todoGolang/todo/services"
)

var sqlite3Db *sql.DB

func init() {
	sqlite3Db = sqliteDb.CreateDatabase(constants.SqliteFileInfo)
	sqliteDb.CreateTable(sqlite3Db)
}

type TodoDao struct {
	services.TodoService
}

//GetTodoList : get TodoDataList from tb_todo table
func (todoDao *TodoDao) GetTodoList() model.TodoDataList {
	return getTodoList(sqlite3Db)
}

func (todoDao *TodoDao) GetTodoListJson() model.TodoJsonList{
	return getTodoListJson(sqlite3Db)
}

//AddTodoData : insert TodoData into tb_todo table
func (todoDao *TodoDao) AddTodoData(todoParam *model.TodoData) bool {
	return addTodoData(sqlite3Db, todoParam)
}

//UpdateTodoData : update TodoData's isdone on tb_todo table
func (todoDao *TodoDao) UpdateTodoData(todoParam *model.TodoData) bool {
	return updateTodo(sqlite3Db, todoParam)
}

//DeleteTodoData : delete TodoData from tb_todo table
func (todoDao *TodoDao) DeleteTodoData(todoParam *model.TodoData) bool {
	return deleteTodo(sqlite3Db, todoParam)
}

func getTodoList(db *sql.DB) (todoList model.TodoDataList) {
	getSQL := `
		SELECT
			id
			, title
			, doneYN
			, createDt
			, updateDt
		FROM tb_todo order by id desc;
	`

	rows, err := db.Query(getSQL)
	exceptions.ProcError(err)
	defer rows.Close()

	var idx string
	var todoTitle string
	var doneYN bool
	var createDt string
	var updateDt string

	for rows.Next() {
		rowsErr := rows.Scan(&idx, &todoTitle, &doneYN, &createDt, &updateDt)
		exceptions.ProcError(rowsErr)

		todoResul := model.ResultTodo(idx, todoTitle, doneYN, createDt, updateDt)

		todoList = append(todoList, *todoResul)
	}
	return todoList
}

func getTodoListJson(db *sql.DB) (todoJsonList model.TodoJsonList) {
	getSQL := `
		SELECT
			id
			, title
			, doneYN
			, createDt
			, updateDt
		FROM tb_todo order by id desc;
	`

	rows, err := db.Query(getSQL)
	exceptions.ProcError(err)
	defer rows.Close()

	for rows.Next() {
		todoJson := new(model.TodoJson)
		rowsErr := rows.Scan(&todoJson.TodoIdx, &todoJson.TodoTitle, &todoJson.IsDone, &todoJson.CreateDt, &todoJson.UpdateDt)
		exceptions.ProcError(rowsErr)

		todoJsonList = append(todoJsonList, *todoJson)
	}
	return todoJsonList
}

func addTodoData(db *sql.DB, addParam *model.TodoData) bool {
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
	exceptions.ProcError(prepareErr)

	res, execErr := stmt.Exec(addParam.TodoTitle(), addParam.TodoIsDone())
	defer stmt.Close()

	exceptions.ProcError(execErr)

	id, insertResultErr := res.LastInsertId()
	exceptions.ProcError(insertResultErr)

	if id != 0 {
		bResult = true
	} else {
		bResult = false
	}

	return bResult
}

func updateTodo(db *sql.DB, todoParam *model.TodoData) bool {
	var bResult bool

	updateSQL := "UPDATE tb_todo SET doneYN = ? where id = ?;"

	stmt, prepareErr := db.Prepare(updateSQL)
	exceptions.ProcError(prepareErr)

	res, execErr := stmt.Exec(todoParam.TodoIsDone(), todoParam.TodoIdx())
	exceptions.ProcError(execErr)

	defer stmt.Close()

	affect, affErr := res.RowsAffected()
	exceptions.ProcError(affErr)

	if affect != 0 {
		bResult = true
	} else {
		bResult = false
	}
	return bResult
}

func deleteTodo(db *sql.DB, todoParam *model.TodoData) bool {
	var bResult bool
	deleteSQL := "DELETE FROM tb_todo where id = ?;"

	stmt, prepareErr := db.Prepare(deleteSQL)
	exceptions.ProcError(prepareErr)

	res, execErr := stmt.Exec(todoParam.TodoIdx())
	exceptions.ProcError(execErr)

	defer stmt.Close()

	affect, affErr := res.RowsAffected()
	exceptions.ProcError(affErr)

	if affect != 0 {
		bResult = true
	} else {
		bResult = false
	}
	return bResult
}
