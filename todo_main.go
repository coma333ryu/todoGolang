package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"todoGolang/common/routes"
)

var rootDir, _ = os.Getwd()
var templatePath = filepath.Join(rootDir, "/templates/")
var todoTemplates = template.Must(template.ParseFiles(templatePath + "/todo.html"))

func serve(res http.ResponseWriter, req *http.Request) {

	// todoTemplates.Execute(res, myVars)

}

func main() {
	/*
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
	*/

	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))
	todoRouter := routes.NewRouter()
	// http.ListenAndServe(":3000", router)
	http.HandleFunc("/", serve)
	http.ListenAndServe(":8888", todoRouter)
}
