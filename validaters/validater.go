package validaters

import (
	"errors"
	"todoGolang/databases/sqlite3"
)

//CheckParam : validate parameter
func CheckParam(todoParam *sqlite3.TodoData, methodType string) error {
	if methodType == "A" {
		if todoParam.TodoTitle() == "" {
			return errors.New("Title is nil")
		}
	} else if methodType == "U" {
		if todoParam.TodoIdx() == "" {
			return errors.New("Idx is nil")
		}
	} else if methodType == "D" {
		if todoParam.TodoIdx() == "" {
			return errors.New("Idx is nil")
		}
	} else {
		return errors.New("methodType is nil")
	}

	return nil
}
