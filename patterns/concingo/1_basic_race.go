package main

import (
	"fmt"
	"time"
)

func main() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data) // 0
	}

	time.Sleep(time.Second)
	fmt.Printf("the value is %v.\n", data) // 1
}
