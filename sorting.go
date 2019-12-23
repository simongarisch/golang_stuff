package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func getUserInput() ([]int, error) {
	var input string
	var inputInts []int

	fmt.Printf("Please provide at least 4 integers to sort separated by a space: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return inputInts, err
	}

	inputStrings := strings.Fields(input)
	if len(inputStrings) < 4 {
		return inputInts, errors.New("At least 4 values required to sort")
	}

	for _, str := range inputStrings {
		intValue, err := strconv.Atoi(str)
		if err != nil {
			return inputInts, err
		}
		inputInts = append(inputInts, intValue)
	}
	return inputInts, nil

}

func sortSlice(inputInts []int, wg *sync.WaitGroup, name string) {
	defer wg.Done()
	msg := fmt.Sprintf("goroutine name=%q sorting slice", name)
	fmt.Println(msg, inputInts)
	sort.Sort(sort.IntSlice(inputInts))
}

func main() {
	inputInts, err := getUserInput()
	if err != nil {
		panic(err)
	}
	//fmt.Println(inputInts)

	partitionSize := int(len(inputInts) / 4)
	//fmt.Println("partitionSize", partitionSize)

	parts := make(map[string][]int)
	parts["go1"] = inputInts[:partitionSize]
	parts["go2"] = inputInts[partitionSize : 2*partitionSize]
	parts["go3"] = inputInts[2*partitionSize : 3*partitionSize]
	parts["go4"] = inputInts[3*partitionSize:]
	fmt.Println("----------")

	var wg sync.WaitGroup
	for name, part := range parts {
		wg.Add(1)
		go sortSlice(part, &wg, name)
	}

	wg.Wait()
	fmt.Println("----------")
	fmt.Println("After goroutines are finished sorting their 'chunks'...")
	fmt.Println("The entine slice is not sorted yet")
	fmt.Println(inputInts)

	fmt.Println("----------")
	fmt.Println("Here is the full sorted slice...")
	sort.Sort(sort.IntSlice(inputInts))
	fmt.Println(inputInts)
}
