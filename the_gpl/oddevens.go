package main

import "fmt"

func collectNumbers(numbers chan<- int) {
	for i := 1; i <= 10; i++ {
		numbers <- i
	}
	close(numbers)
}

func sortOddsAndEvens(numbers <-chan int, odd chan<- int, even chan<- int) {
	for {
		number, more := <-numbers
		if !more {
			break
		}
		if number%2 == 0 {
			even <- number
		} else {
			odd <- number
		}
	}
	close(odd)
	close(even)
}

func printNumbers(even <-chan int, odd <-chan int, done chan<- bool) {
	for {
		select {
		case n, more := <-even:
			if !more {
				even = nil
			} else {
				fmt.Println("even", n)
			}
		case n, more := <-odd:
			if !more {
				odd = nil
			} else {
				fmt.Println("odd", n)
			}
		}
		if even == nil && odd == nil {
			break
		}
	}
	done <- true
}

func main() {
	buff := 10
	numbers := make(chan int, buff)
	odd := make(chan int, buff)
	even := make(chan int, buff)
	done := make(chan bool)

	go collectNumbers(numbers)
	go sortOddsAndEvens(numbers, odd, even)
	go printNumbers(odd, even, done)

	<-done // wait
}
