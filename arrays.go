package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 2
	fmt.Println(x) // [2 0 0 0 0]

	// array literal
	y := [5]int{10, 20, 30, 40, 50}
	fmt.Println(y) // [10 20 30 40 50]

	// use of '...' -> will infer size
	z := [...]int{1, 2, 3, 4}
	fmt.Println(z) // [1 2 3 4]

	// iteration
	for index, value := range z {
		fmt.Println(index, value)
	}
	fmt.Println("-----")

}
