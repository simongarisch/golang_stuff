package main

import "fmt"

type ReadingStrategy interface {
	Read()
}

type ReadAloud struct{}
type SpeedRead struct{}
type SkimRead struct{}

func (r ReadAloud) Read() {
	fmt.Println("Blah blah blah")
}

func (r SpeedRead) Read() {
	fmt.Println("Faster than a speeding bullet ...")
}

func (r SkimRead) Read() {
	fmt.Println("Faster ... speeding ...")
}

type Kindle struct {
	readingStrategy ReadingStrategy
}

func (k Kindle) Read() {
	k.readingStrategy.Read()
}

func main() {
	var k Kindle

	k = Kindle{ReadAloud{}}
	k.Read() // Blah blah blah

	k = Kindle{SpeedRead{}}
	k.Read() // Faster than a speeding bullet ...

	k = Kindle{SkimRead{}}
	k.Read() // Faster ... speeding ...
}
