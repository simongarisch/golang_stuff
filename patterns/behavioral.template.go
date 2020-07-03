// https://stackoverflow.com/questions/37635769/elegant-way-to-implement-template-method-pattern-in-golang
package main

import (
	"fmt"
)

type Runner struct {
	run func() // 1. this has to get assigned the actual implementation
}

func NewRunner(i func()) *Runner {
	return &Runner{i}
}

func (r *Runner) Start() {
	r.run()
}

type Logger struct {
	Runner
}

func NewLogger() *Logger {
	l := Logger{}
	l.run = l.loggerRun // 2. the actual version is assigned
	return &l
}

func (l *Logger) loggerRun() {
	fmt.Println("Logger is running...")
}

func main() {
	l := NewLogger() // 3. constructor should be used, to get the assignment working
	l.Start()
}
