package main

// Recover works only when it is called from the same goroutine.
// It's not possible to recover from a panic that has happened in a different goroutine.
// defer recovery() now placed inside the goroutine.
// We've also used debug.PrintStack() to print the stack trace.

import (
	"fmt"
	"time"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
		//debug.PrintStack()
	}
}

func a() {
	fmt.Println("Inside A")
	go b()
	time.Sleep(1 * time.Second)
}

func b() {
	defer recovery()
	fmt.Println("Inside B")
	panic("oh! B panicked")
}

func main() {
	a()
	fmt.Println("normally returned from main")
}

// Inside A
// Inside B
// recovered: oh! B panicked
// normally returned from main
