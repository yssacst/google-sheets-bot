package logger

import (
	"log"
	"os"
	"strings"
)

type Logger struct {
	info  *log.Logger
	error *log.Logger
	debug *log.Logger
	level string
}

func New() *Logger {
	level := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if level == "" {
		level = "info"
	}

	return &Logger{
		level: level,
		info:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime),
		error: log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
		debug: log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime),
	}
}

func (l *Logger) Info(v ...any) {
	l.info.Println(v...)
}

func (l *Logger) Error(v ...any) {
	l.error.Println(v...)
}

func (l *Logger) Debug(v ...any) {
	if l.level != "debug" {
		return
	}
	l.debug.Println(v...)
}