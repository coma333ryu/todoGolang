package controller

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"todoGolang/common/exceptions"
	"todoGolang/common/logger"
	"todoGolang/common/validaters"
	"todoGolang/todo/model"
	"todoGolang/todo/services/impl"
	"encoding/json"
	"fmt"
)

var (
	todoService   *impl.TodoDao
	controllerLog *log.Logger
)

func init() {
	todoService = new(impl.TodoDao)
	controllerLog = logger.NewLogger()
}

func GetTodoList(res http.ResponseWriter, req *http.Request) {
	view := path.Join("templates", "todo.html")

	result := todoService.GetTodoList()

	todoTemplates := template.Must(template.ParseFiles(view))
	todoTemplates.Execute(res, result)
}

func GetTodoListJson(res http.ResponseWriter, req *http.Request) {
	result := todoService.GetTodoList()
	fmt.Println("result ========> ",result)

	//res.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(res).Encode(result)
	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

func AddTodoData(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	//form value 받아내는 법
	//req.PostFormValue("todoTitle")
	//req.Form["todoTitle"] <== 배열이 들어옴. 그래서 아래와 같이 변경해야 함.
	//strings.Join(req.Form["todoTitle"], "")

	todoAddParam := model.NewTodoData("", req.PostFormValue("todoTitle"), false)

	validErr := validaters.CheckParam(todoAddParam, "A")

	if validErr != nil {
		exceptions.ProcError(validErr)
	}

	bAddResult := todoService.AddTodoData(todoAddParam)

	if bAddResult {
		http.Redirect(res, req, "http://"+req.Host+"/", http.StatusSeeOther)
	} else {
		addErr := errors.New("Todo can not insert into tb_todo")
		exceptions.ProcError(addErr)
	}
}

func UpdateTodoData(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	isDone, _ := strconv.ParseBool(req.PostFormValue("isDone"))

	todoUpdateParam := model.NewTodoData(req.PostFormValue("todoIdx"), "", isDone)

	validErr := validaters.CheckParam(todoUpdateParam, "U")

	if validErr != nil {
		exceptions.ProcError(validErr)
	}

	bUpdateResult := todoService.UpdateTodoData(todoUpdateParam)

	if bUpdateResult {
		http.Redirect(res, req, "http://"+req.Host+"/", http.StatusSeeOther)
	} else {
		addErr := errors.New("Todo can not update into tb_todo")
		exceptions.ProcError(addErr)
	}
}

func DeleteTodoData(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	todoDeleteParam := model.NewTodoData(req.PostFormValue("todoIdx"), "", false)

	validErr := validaters.CheckParam(todoDeleteParam, "D")

	if validErr != nil {
		exceptions.ProcError(validErr)
	}

	bDeleteResult := todoService.DeleteTodoData(todoDeleteParam)

	if bDeleteResult {
		http.Redirect(res, req, "http://"+req.Host+"/", http.StatusSeeOther)
	} else {
		addErr := errors.New("Todo can not delete into tb_todo")
		exceptions.ProcError(addErr)
	}
}
