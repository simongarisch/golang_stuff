package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter some numbers ('X' to quit)...")

	numbers := make([]int, 0, 3)
	for {
		// collect user input as a string
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("You entered %q\n", input)

		if strings.Compare(input, "X") == 0 {
			break
		}

		// try to convert this to an integer
		intInput, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Cannot convert %q to int! Please try again...\n", input)
			continue
		}
		//fmt.Printf("Int equivalent = %d\n", intInput)

		numbers = append(numbers, intInput)
		sort.Slice(numbers, func(i, j int) bool {
			return numbers[i] < numbers[j]
		})
		fmt.Println("So far you've entered", numbers)
	}
}
