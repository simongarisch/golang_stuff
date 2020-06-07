// hasarace.go
// go run -race hasarace.go
package main

import "fmt"

var x int

func main() {
	for i := 1; i <= 1000; i++ {
		go func() {
			x++
		}()
	}

	fmt.Println(x)
}
