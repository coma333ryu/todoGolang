package model

//TodoData : Todo List Data struct
type TodoData struct {
	todoIdx   string
	todoTitle string
	isDone    bool
	createDt  string
	updateDt  string
}

//TodoDataList : TodoData's slice
type TodoDataList []TodoData

//NewTodoData : create TodoData struct
func NewTodoData(idx string, title string, done bool) *TodoData {
	return &TodoData{
		todoIdx:   idx,
		todoTitle: title,
		isDone:    done,
	}
}

//ResultTodo : set TodoData for query result
func ResultTodo(idx string, title string, done bool, createDt string, updateDt string) *TodoData {
	return &TodoData{
		todoIdx:   idx,
		todoTitle: title,
		isDone:    done,
		createDt:  createDt,
		updateDt:  updateDt,
	}
}

//TodoIdx : get todoIdx
func (todo *TodoData) TodoIdx() string {
	return todo.todoIdx
}

//TodoTitle : get totoTitle
func (todo *TodoData) TodoTitle() string {
	return todo.todoTitle
}

//TodoIsDone : get isDone
func (todo *TodoData) TodoIsDone() bool {
	return todo.isDone
}

//TodoCreateDt : get createDt
func (todo *TodoData) TodoCreateDt() string {
	return todo.createDt
}

//TodoUpdateDt : get updateDt
func (todo *TodoData) TodoUpdateDt() string {
	return todo.updateDt
}
