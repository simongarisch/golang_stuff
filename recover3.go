package main

// Recover works only when it is called from the same goroutine.
// It's not possible to recover from a panic that has happened in a different goroutine.
// The panic below will not be recovered.

import (
	"fmt"
	"time"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	go b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
}

func main() {
	a()
	fmt.Println("normally returned from main")
}

// Inside A
// Inside Bpanic: oh! B panicked
// ...
// exit status 2
