package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// this will block
	/*
		<-ch
		fmt.Println("Here")
	*/

	go func() {
		// send number to the channel
		ch <- 353
	}()

	// receive from the channel
	val := <-ch
	fmt.Printf("got %d\n", val)
	fmt.Println("-----")

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	// close the channel when we are done
	fmt.Println("-----")
	fmt.Println("Closing a channel...")
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("received %d\n", i)
	}
}
