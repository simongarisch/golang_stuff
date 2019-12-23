package main

import (
	"fmt"
	"sync"
)

func withRaceCondition() int {
	x := 0

	var wg sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			x++
			wg.Done()
		}()
	}

	wg.Wait()
	return x
}

func withoutRaceCondition() int {
	x := 0

	for i := 1; i <= 1000; i++ {
		func() {
			x++
		}()
	}

	return x
}

func main() {

	fmt.Println("Under this race condition it is not clear what value x will have...")
	for i := 1; i <= 5; i++ {
		x := withRaceCondition()
		fmt.Println("x =", x)
	}

	fmt.Println("----------")
	fmt.Println("If we didn't have a race condition...")
	for i := 1; i <= 5; i++ {
		x := withoutRaceCondition()
		fmt.Println("x =", x)
	}

}
