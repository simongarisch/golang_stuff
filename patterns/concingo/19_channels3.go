package main

import "fmt"

func add(ch chan interface{}, value interface{}) {
	ch <- value
}

func main() {
	ch := make(chan interface{}, 3) // channel is type agnostic, deadlock if you don't buffer
	add(ch, 2)
	add(ch, "this is a string")
	add(ch, true)

	close(ch)
	for value := range ch {
		fmt.Println(value)
	}
	// 2
	// this is a string
	// true
}
