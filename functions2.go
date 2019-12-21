package main

import (
	"errors"
	"fmt"
)

// functions are first class values
// we can treat them like other types
var funcVar func(int) int

func incFn(x int) int {
	return x + 1
}

func applyIt(afunct func(int) int, val int) int {
	// we are passing functions as args to another function
	return afunct(val)
}

// functions can also return other functions
func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func getOperatingFunction(operation string) (func(int, int) int, error) {
	switch operation {
	case "+":
		return add, nil
	case "-":
		return subtract, nil
	default:
		return nil, errors.New("Operation not available")
	}

}

// variable number of args to functions
func getMax(vals ...int) int {
	maxValue := vals[0]
	for _, value := range vals {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func main() {
	funcVar = incFn
	fmt.Println(funcVar(1))        // 2
	fmt.Println(applyIt(incFn, 2)) // 3

	// anonymous functions
	z := applyIt(func(x int) int { return x + 1 }, 2)
	fmt.Println(z) // 3

	operatingFunction, err := getOperatingFunction("-")
	if err != nil {
		panic(err)
	}
	fmt.Println(operatingFunction(10, 2)) // 8
	fmt.Println("----------")

	fmt.Println(getMax(1, 11, 2, 7)) // 11

	myslice := []int{1, 11, 2, 7}
	fmt.Println(getMax(myslice...)) // 11
}
