package logger

import (
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	Info *log.Logger
	Warn *log.Logger
	Err  *log.Logger
}

func InitLoger() *Logger {
	warn := log.New(os.Stderr, color.HiYellowString("[ WARN  ]"), log.Ltime|log.Lshortfile)
	info := log.New(os.Stderr, color.HiGreenString("[ INFO  ]"), log.Ltime|log.Lshortfile)
	err := log.New(os.Stderr, color.HiRedString("[ ERROR ]"), log.Ltime|log.Lshortfile)

	return &Logger{
		Info: info,
		Warn: warn,
		Err:  err,
	}
}
