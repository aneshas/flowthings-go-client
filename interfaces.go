package main

// Logger interface
type ILogger interface {
	Init()
	Info(...interface{})    // stdout info messages
	Log(...interface{})     // main file logger
	Warning(...interface{}) // stdout error
	Error(...interface{})   // stderr
}
