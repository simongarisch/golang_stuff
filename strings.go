package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	book := "The color of magic"
	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	fmt.Println(book[4:11])
	fmt.Println(book[4:])
	fmt.Println(book[:4])
	fmt.Println("------")

	x, y := "astring", "astring"
	fmt.Println(strings.Compare(x, y))                 // 0 when strings are equal
	fmt.Println(strings.Compare("astring", "bstring")) // -1 if <
	fmt.Println(strings.Compare("bstring", "astring")) // +1 if >
	fmt.Println("------")

	fmt.Println(strings.Contains("this is a string", "is")) // true
	fmt.Println(strings.HasPrefix("thisIsThe", "this"))     // true
	fmt.Println(strings.HasSuffix("test.csv", ".csv"))      // true
	fmt.Println("------")

	fmt.Println(strings.Index("this is a string", "is"))  // 2
	fmt.Println(strings.Index("this is a string", "zzz")) // -1
	fmt.Println("------")

	// strings are immutable, but modified strings can be returned
	fmt.Println(strings.Replace("this is a string", "is", "**", 1)) // 'th** is a string', so just replacing the first instance
	fmt.Println(strings.ToLower("willGoLower"))                     // willgolower
	fmt.Println(strings.ToUpper("willGoUpper"))                     // WILLGOUPPER
	fmt.Println(strings.TrimSpace(" someSpace "))                   // 'someSpace'
	fmt.Println("------")

	fmt.Println("Now for string conversion")
	sint := "42"
	aint, err := strconv.Atoi(sint)
	if err != nil {
		panic(err)
	}
	fmt.Println(aint) // 42 ... reverse is Itoa
	// we also have FormatFloat and ParseFloat for floats...
}
