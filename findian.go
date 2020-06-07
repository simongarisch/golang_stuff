package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter a string...")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	/*
		for index, char := range input {
			fmt.Printf("Character %q at index %d\n", char, index)
		}
		fmt.Println("-----")
	*/

	// remove the characters '\r' and '\n', make lowercase
	input = strings.TrimRight(input, "\r\n")
	input = strings.ToLower(input)

	// run checks
	startsI := strings.HasPrefix(input, "i")
	endsN := strings.HasSuffix(input, "n")
	containsA := strings.Contains(input, "a")
	//fmt.Println(input, startsI, endsN, containsA)

	if startsI && endsN && containsA {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
