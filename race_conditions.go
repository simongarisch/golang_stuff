/*
Here we have created a race condition in the func 'withRaceCondition'
This function uses two goroutines to run another 'increment' function that
takes two arguments: an integer pointer and an amount to increment by.

If we start with x=0 and increment this by 10K twice then we expect x=20K.
However, if we run the 'withRaceCondition' function 10 times we see that
the final value of x is not certain...

Why is this the case?
Suppose that the current value of x = 1000
goroutine one takes x = 1000 and increments by one
before goroutine one writes this new value (1001) to x
goroutine two also takes x=1000 and increments it by one (1001)
If they were run separately they would each increment x by 1 so x=1002
Instead we have a race condition and sometimes x will not be 20K.
*/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func increment(someInt *int, by int) {
	defer wg.Done()
	for i := 1; i <= by; i++ {
		*someInt++
	}
}

func withRaceCondition() int {
	x := 0

	wg.Add(1)
	go increment(&x, 10000)
	wg.Add(1)
	go increment(&x, 10000)

	wg.Wait()
	return x
}

func main() {

	fmt.Println("Under this race condition it is not clear what value x will have...")
	for i := 1; i <= 10; i++ {
		x := withRaceCondition()
		fmt.Println("x =", x)
	}
}
