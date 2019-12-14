package main

import "fmt"

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) { // int here is passed by value... slices and maps are passed by reference
	n *= 2
}

func doublePtr(n *int) {
	*n *= 2
}

func main() {
	values := []int{1, 2, 3, 4}
	doubleAt(values, 2)
	fmt.Println(values)

	val := 10
	double(val)      // passed by value
	fmt.Println(val) // 10

	doublePtr(&val)  // by reference
	fmt.Println(val) // 20
}
