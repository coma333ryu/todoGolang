package main

import (
	"log"
	"net/http"
	"todoGolang/common/logger"
	"todoGolang/common/routes"
)

var mainLog *log.Logger

func main() {
	mainLog = logger.NewLogger()
	todoRouter := routes.NewRouter()
	mainLog.Println("============ Started Web Server ============")

	http.ListenAndServe(":8888", todoRouter)
}
