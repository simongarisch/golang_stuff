package main

import (
	"encoding/json"
	"fmt"
)

// capitalize attributes for the marshalling to work
type Person struct {
	Name  string
	Addr  string
	Phone string
}

func main() {
	p1 := Person{"Joe", "a st.", "123"}
	fmt.Println(p1) // {Joe a st. 123}

	byteArr, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	fmt.Println(byteArr) // [123 34 78 97 109 ...

	// convert this byte array back to a person struct
	var p2 Person
	err = json.Unmarshal(byteArr, &p2)
	if err != nil {
		panic(err)
	}
	fmt.Println(p2) // {Joe a st. 123}
}
