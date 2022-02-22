package logger

import (
	"log"
	"os"
)

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
}

func Info(x interface{}) {
	infoLogger.Println(x)
}

func Warning(x interface{}) {
	warningLogger.Println(x)
}

func Error(x interface{}) {
	errorLogger.Println(x)
}
