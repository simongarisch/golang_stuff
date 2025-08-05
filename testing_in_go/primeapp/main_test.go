package main

import "testing"

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
