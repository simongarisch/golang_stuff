// go run golinq.go
// go get github.com/ahmetb/go-linq/v3
package main

import (
	"fmt"

	. "github.com/ahmetb/go-linq"
)

type Car struct {
	year         int
	owner, model string
}

func main() {

	var cars []Car
	cars = append(cars, Car{year: 2020, owner: "Simon", model: "jeep"})
	cars = append(cars, Car{year: 2015, owner: "Joe", model: "BMW"})
	cars = append(cars, Car{year: 2010, owner: "Jim", model: "Toyota"})

	var owners []string
	From(cars).Where(func(c interface{}) bool {
		return c.(Car).year >= 2015
	}).Select(func(c interface{}) interface{} {
		return c.(Car).owner
	}).ToSlice(&owners)

	fmt.Println(owners) // [Simon Joe]
}
