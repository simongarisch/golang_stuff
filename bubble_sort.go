package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ProcessInputString(input string) ([]int, error) {
	// https://golang.org/pkg/strings/#Fields
	inputStrings := strings.Fields(input)
	if len(inputStrings) == 0 {
		return nil, errors.New("At least one value required to sort")
	}
	var inputInts []int
	for _, str := range inputStrings {
		intValue, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		inputInts = append(inputInts, intValue)
	}
	return inputInts, nil
}

func Swap(intslice []int, i int) {
	intslice[i], intslice[i+1] = intslice[i+1], intslice[i]
}

func BubbleSort(intslice []int) {
	numElements := len(intslice)
	if numElements < 2 {
		return
	}
	for a := 0; a < numElements; a++ {
		for b := numElements - 1; b >= a+1; b-- {
			if intslice[b] < intslice[b-1] {
				Swap(intslice, b-1)
			}
		}
	}

}

func main() {
	var input string
	//input = "10 12 11 5 22 18"
	fmt.Printf("Please provide integers to sort separated by a space: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	//fmt.Println("input", input)

	intslice, err := ProcessInputString(input)
	if err != nil {
		panic(err)
	}
	//fmt.Println("intslice", intslice)

	BubbleSort(intslice)
	fmt.Println(intslice)
}
