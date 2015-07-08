package main

import (
	"log"
	"os"
)

type DefaultLogger struct {
	AllToFile bool // if true all logging is redirected to file
	info      *log.Logger
	log       *log.Logger
	warning   *log.Logger
	error     *log.Logger
}

func (l *DefaultLogger) Init() {
	l.info = log.New(
		os.Stdout,
		"INFO: ", log.Ldate|log.Ltime)

	l.error = log.New(
		os.Stdout,
		"ERROR: ", log.Ldate|log.Ltime)

	l.warning = log.New(
		os.Stdout,
		"WARNING: ", log.Ldate|log.Ltime)
}

func (l *DefaultLogger) Info(message ...interface{}) {
	go l.info.Println(message)
}

func (l *DefaultLogger) Log(message ...interface{}) {
	go l.info.Println(message)
}

func (l *DefaultLogger) Warning(message ...interface{}) {
	go l.warning.Println(message)
}

func (l *DefaultLogger) Error(message ...interface{}) {
	go l.error.Println(message)
}
