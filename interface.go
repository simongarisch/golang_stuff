// go run interface.go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{}

	x = 7
	print(x) // 7

	x = "this"
	print(x) // 'this'

	print(true) // true
	print(7.88) // 7.88

	nums := []interface{}{1, 2, 3, 1.2} // 7
	fmt.Println(sum(nums...))
}

func print(x interface{}) {
	fmt.Println(x)
}

func sum(items ...interface{}) interface{} {
	var i int
	var f float64
	typeInt := reflect.TypeOf(i)
	typeFloat := reflect.TypeOf(f)

	var sum int
	for _, item := range items {
		itemType := reflect.TypeOf(item)
		// fmt.Println(item)
		// fmt.Println(reflect.TypeOf(itemType)) // *reflect.rtype
		// fmt.Println(itemType)                 // int
		// fmt.Println(itemType == typeInt)      // true
		if itemType == typeInt {
			sum += item.(int)
		}
		if itemType == typeFloat {
			sum += int(item.(float64))
		}
	}

	return sum
}
