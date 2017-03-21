package logger

import (
	"log"
	"os"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "INFO: ", log.LstdFlags)
}
