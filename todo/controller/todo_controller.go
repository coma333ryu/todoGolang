package controller

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func Index(w http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadFile("public/index.html")
	//data, err := ioutil.ReadFile("public/todo.html")

	if err != nil {
		fmt.Errorf("error reading index file : %s", err)
	}
	w.Header().Add("Content Type", " text/html")
	w.WriteHeader(200)
	w.Write(data)
}
