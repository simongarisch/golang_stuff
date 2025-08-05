package main

import "fmt"

func main() {
	n := 2

	_, msg := isPrime(n)
	fmt.Println(msg)
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
