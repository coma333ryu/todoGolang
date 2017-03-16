package model

import (
	"database/sql"
	"todoGolang/common/exceptions"
)

//TodoDataList : TodoData's slice
type TodoDataList []TodoData

type todoServicer interface {
	GetTodoList(db *sql.DB) TodoDataList
	AddTodoData(db *sql.DB, addParam *TodoData) bool
	UpdateTodoData(db *sql.DB, todoParam *TodoData) bool
	DeleteTodoData(db *sql.DB, todoParam *TodoData) bool
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

		todoResul := ResultTodo(idx, todoTitle, doneYN, createDt, updateDt)

		todoList = append(todoList, *todoResul)
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

func updateTodo(db *sql.DB, todoParam *TodoData) bool {
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

func deleteTodo(db *sql.DB, todoParam *TodoData) bool {
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
