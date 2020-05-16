package main

import (
	"fmt"
	"time"
)

func runLoopSend(n int, ch chan int) {
	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch)
}

func runLoopReceive(ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Received value:", i)
	}
}

func main() {
	myChannel := make(chan int)
	go runLoopSend(10, myChannel)
	go runLoopReceive(myChannel)
	time.Sleep(2 * time.Second)
}
