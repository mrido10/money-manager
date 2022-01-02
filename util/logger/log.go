package logger

import (
	"log"
	"os"
)

type Log interface {
	Info()  *log.Logger
	Error() *log.Logger
}

func Info() *log.Logger {
	return log.New(os.Stdout, "[INFO] \t", log.Lshortfile | log.LstdFlags)
}

func Error() *log.Logger {
	return log.New(os.Stdout, "[ERROR]\t", log.LstdFlags)
}