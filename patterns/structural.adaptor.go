package main

import "fmt"

type Man struct{}

type Cat struct{}

func (man Man) Talk() {
	fmt.Println("Blah de blah")
}

func (cat Cat) MakeNoise() {
	fmt.Println("Meow de meow")
}

// we want a common interface
type Animal interface {
	Talk()
}

func (cat Cat) Talk() {
	cat.MakeNoise()
}

func main() {
	animals := []Animal{Man{}, Cat{}}
	for _, animal := range animals {
		animal.Talk()
	}
	// Blah de blah
	// Meow de meow
}
