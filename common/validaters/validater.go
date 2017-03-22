package validaters

import (
	"errors"
	"todoGolang/todo/model"
)

//CheckParam : validate parameter
func CheckParam(todoParam *model.TodoData, methodType string) error {
	if methodType == "A" {
		if todoParam.TodoTitle() == "" {
			return errors.New("Title is nil")
		}
	} else if methodType == "U" || methodType == "D" {
		if todoParam.TodoIdx() == "" {
			return errors.New("Idx is nil")
		}
	} else {
		return errors.New("methodType is nil")
	}

	return nil
}
