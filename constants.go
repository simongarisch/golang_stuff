package main

import "fmt"

const (
	x = 1
	y = "this"
	z = true
)

type Weekday int

const (
	Sunday    Weekday = 0
	Monday    Weekday = 1
	Tuesday   Weekday = 2
	Wednesday Weekday = 3
	Thursday  Weekday = 4
	Friday    Weekday = 5
	Saturday  Weekday = 6
)

func Weekend(day Weekday) bool {
	switch day {
	case Sunday, Saturday:
		return true
	default:
		return false
	}
}

const s string = "constant"

// We can use iota to simulate C’s enum or #define constant.
type Season uint8

const (
	Spring = Season(iota)
	Summer
	Autumn
	Winter
)

func main() {
	fmt.Println(x, y, z)
	fmt.Println(Sunday)            // 0
	fmt.Println(Saturday)          // 6
	fmt.Println(Weekend(Saturday)) // true
	fmt.Println(Weekend(Tuesday))  // false
	fmt.Println("--------------")

	fmt.Println(Spring) // 0
	fmt.Println(Summer) // 1
	fmt.Println(Autumn) // 2
	fmt.Println(Winter) // 3
	fmt.Println("--------------")

	s := Summer
	fmt.Println(s) // 1
	s = Season(9)
	fmt.Println(s) // 9
}
