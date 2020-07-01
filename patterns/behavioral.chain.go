package main

import "fmt"

type cleaner interface {
	Clean()
	SetNext(cleaner)
}

type Wash struct {
	next cleaner
}
type Polish struct {
	next cleaner
}
type Buff struct {
	next cleaner
}

func (c Wash) Clean() {
	fmt.Println("Washing...")
	if c.next != nil {
		c.next.Clean()
	}
}

func (c Polish) Clean() {
	fmt.Println("Polishing...")
	if c.next != nil {
		c.next.Clean()
	}
}

func (c Buff) Clean() {
	fmt.Println("Buffing...")
	if c.next != nil {
		c.next.Clean()
	}
}

func (c *Wash) SetNext(next cleaner) {
	c.next = next
}

func (c *Polish) SetNext(next cleaner) {
	c.next = next
}

func (c *Buff) SetNext(next cleaner) {
	c.next = next
}

func makeCleaningChain() cleaner {
	wash := &Wash{}
	polish := &Polish{}
	buff := &Buff{}

	wash.SetNext(polish)
	polish.SetNext(buff)
	return wash
}

func main() {
	carWash := makeCleaningChain()
	carWash.Clean()
	// Washing...
	// Polishing...
	// Buffing...
}
