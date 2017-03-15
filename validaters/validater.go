package validaters

import (
	"errors"
	"todoGolang/databases/sqlite3"
)

//CheckParam : validate parameter
func CheckParam(todoParam *sqlite3.TodoData, methodType string) error {
	if methodType == "A" {
		if todoParam.GetTodoTitle() == "" {
			return errors.New("Title is nil")
		}
	} else if methodType == "U" {
		if todoParam.GetTodoIdx() == "" {
			return errors.New("Idx is nil")
		}
	} else if methodType == "D" {
		if todoParam.GetTodoIdx() == "" {
			return errors.New("Idx is nil")
		}
	} else {
		return errors.New("methodType is nil")
	}

	return nil
}
