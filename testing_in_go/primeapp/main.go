package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a welcome message
	intro()

	// create a channel to indicate when the user wants to quit.
	doneChan := make(chan bool)

	// start a goroutine to read user intput and run program
	go readUserInput(os.Stdin, doneChan)

	// block until the doneChan gets a value
	<-doneChan

	// close the channel
	close(doneChan)

	// say goodbye
	fmt.Println("Goodbye!")
}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)
	for {
		res, done := checkNumber(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumber(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	// check to see if the user wants to quit
	text := strings.TrimSpace(scanner.Text())
	if strings.EqualFold(text, "q") {
		return "", true
	}

	// try to covert what the user typed into an int
	numToCheck, err := strconv.Atoi(text)
	if err != nil {
		return fmt.Sprintf("%q is not int. Please enter a whole number...", text), false
	}

	_, msg := isPrime(numToCheck)
	return msg, false
}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("------------")
	fmt.Println("Enter a whole number and we'll tell you if it is a prime number or not. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d: zero and one are not prime", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, fmt.Sprintf("%d: negative numbers are not prime", n)
	}

	// use the modulus operator repeatedly
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime
			return false, fmt.Sprintf("%d is not a prime number because it's divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number", n)
}
