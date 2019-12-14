package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		var isDiv3, isDiv5 bool
		isDiv3 = (i%3 == 0)
		isDiv5 = (i%5 == 0)
		if isDiv3 && isDiv5 {
			fmt.Println("fizzbuzz")
		} else if isDiv3 {
			fmt.Println("fizz")
		} else if isDiv5 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
