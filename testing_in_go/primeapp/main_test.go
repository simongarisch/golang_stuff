package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	n := 0
	result, msg := isPrime(n)
	if result {
		t.Errorf("with %d as test parameter, got true but expected false", n)
	}

	if msg != "0: zero and one are not prime" {
		t.Errorf("unexpected message: %q", msg)
	}
}

func Test_isPrimeWithTable(t *testing.T) {
	primeTests := []struct {
		name    string
		testNum int
		result  bool
		msg     string
	}{ // 2, 3, 5, 7, 11, ...
		{"-3 not prime", -3, false, "-3: negative numbers are not prime"},
		{"-2 not prime", -2, false, "-2: negative numbers are not prime"},
		{"-1 not prime", -1, false, "-1: negative numbers are not prime"},
		{"0 not prime", 0, false, "0: zero and one are not prime"},
		{"1 not prime", 1, false, "1: zero and one are not prime"},
		{"2 prime", 2, true, "2 is a prime number"},
		{"3 prime", 3, true, "3 is a prime number"},
		{"4 not prime", 4, false, "4 is not a prime number because it's divisible by 2"},
		{"4 not prime", 4, false, "4 is not a prime number because it's divisible by 2"},
		{"5 prime", 5, true, "5 is a prime number"},
		{"6 not prime", 6, false, "6 is not a prime number because it's divisible by 2"},
		{"7 prime", 7, true, "7 is a prime number"},
		{"8 not prime", 8, false, "8 is not a prime number because it's divisible by 2"},
		{"9 not prime", 9, false, "9 is not a prime number because it's divisible by 3"},
		{"10 not prime", 10, false, "10 is not a prime number because it's divisible by 2"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.result != result {
			t.Errorf("%s: expected %t but got %t", e.name, e.result, result)
		}
		if e.msg != msg {
			t.Errorf("%s: expected %q but got %q", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	expected := "-> "
	stringOut := string(out)
	if stringOut != expected {
		t.Errorf("incorrect prompt: expected %q, got %q", expected, stringOut)
	}
}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	expected := "Enter a whole number"
	stringOut := string(out)
	if !strings.Contains(stringOut, expected) {
		t.Errorf("intro text not correct, got %q", stringOut)
	}
}

func Test_checkNumber(t *testing.T) {
	input := strings.NewReader("7")
	reader := bufio.NewScanner(input)

	res, _ := checkNumber(reader)
	if !strings.EqualFold(res, "7 is a prime number") {
		t.Errorf("incorrect value returned: %q", res)
	}
}

func Test_checkNumberWithTable(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "\"\" is not int. Please enter a whole number..."},
		{name: "zero", input: "0", expected: "0: zero and one are not prime"},
		{name: "one", input: "1", expected: "1: zero and one are not prime"},
		{name: "two", input: "2", expected: "2 is a prime number"},
		{name: "three", input: "3", expected: "3 is a prime number"},
		{name: "negative", input: "-1", expected: "-1: negative numbers are not prime"},
		{name: "typed", input: "three", expected: "\"three\" is not int. Please enter a whole number..."},
		{name: "decimal", input: "1.1", expected: "\"1.1\" is not int. Please enter a whole number..."},
		{name: "quit", input: "q", expected: ""},
		{name: "QUIT", input: "Q", expected: ""},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)

		res, _ := checkNumber(reader)
		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: incorrect value returned: %q, wanted %q", e.name, res, e.expected)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	// to test this function we need a channel
	// and an instance of io.Reader
	doneChan := make(chan bool)

	// create a reference to a bytes.Buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)

	<-doneChan
	close(doneChan)
}
