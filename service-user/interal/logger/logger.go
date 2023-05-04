package logger

import (
	"log"
	"os"
)

type Logger struct {
	Info *log.Logger
	Warn *log.Logger
	Err  *log.Logger
}

func InitLoger() *Logger {
	warn := log.New(os.Stderr, "[ WARN  ]", log.Ltime|log.Lshortfile)
	info := log.New(os.Stderr, "[ INFO  ]", log.Ltime|log.Lshortfile)
	err := log.New(os.Stderr, "[ ERROR ]", log.Ltime|log.Lshortfile)

	return &Logger{
		Info: info,
		Warn: warn,
		Err:  err,
	}
}
