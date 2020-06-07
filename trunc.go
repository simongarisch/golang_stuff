package main

import "fmt"

func main() {
	var userInputFloat float64
	var userInputInt int64

	fmt.Println("Enter a number")

	_, err := fmt.Scan(&userInputFloat)
	if err != nil {
		panic(err)
	}

	userInputInt = int64(userInputFloat)
	fmt.Println(userInputInt)
}
