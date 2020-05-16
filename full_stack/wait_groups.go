package main

import (
	"fmt"
	"sync"
)

// Create a global waitgroup:
var wg = &sync.WaitGroup{}

func main() {
	myChannel := make(chan int)
	//Increment the wait group internal counter by 2
	wg.Add(2)
	go runLoopSend(10, myChannel)
	go runLoopReceive(myChannel)
	//Wait till the wait group counter is 0
	wg.Wait()
}

func runLoopSend(n int, ch chan int) {
	//Ensure that the wait group counter decrements by one after our function exits
	defer wg.Done()
	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch)
}

func runLoopReceive(ch chan int) {
	//Ensure that the wait group counter decrements after our function exits
	defer wg.Done()
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Received value:", i)
	}
}
