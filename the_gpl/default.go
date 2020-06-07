package main

import (
	"fmt"
	"time"
)

func checkChannel(c chan bool) {
	select {
	case value := <-c:
		fmt.Println("received", value)
		return
	default:
		fmt.Println("waiting...")
	}
}

func main() {
	c := make(chan bool, 10)

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
		go checkChannel(c)
		if i == 3 {
			fmt.Println("sending to channel...")
			c <- true
		}
	}
}
