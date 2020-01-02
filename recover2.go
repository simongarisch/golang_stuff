package main

import "fmt"

// Recover is useful only when called inside deferred functions.
// Executing a call to recover inside a deferred function stops
// the panicking sequence by restoring normal execution and retrieves
// the error value passed to the call of panic.

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverName()

	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}

	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

// recovered from  runtime error: last name cannot be nil
// returned normally from main
// deferred call in main
