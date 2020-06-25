/*
The Composite design pattern favors composition (commonly defined as a has a
relationship) over inheritance (an is a relationship).

With Go, we can use two types of composition–the direct composition
and the embedding composition.
*/
package main

type Athlete struct{}

func (a *Athlete) Train() {
	println("Training")
}

func Swim() {
	println("Swimming!")
}

type CompositeSwimmerA struct { // direct composition in this case
	MyAthlete Athlete // when new struct created initialized as zero value for it's type
	MySwim    *func()
}

//******************************
type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

//******************************
type Animal struct{}

func (r *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal // embedding composition
	Swim   func()
}

//******************************
type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

//******************************
type Parent struct {
	SomeField int
}

type Son struct {
	P Parent
}

func GetParentField(p Parent) int {
	return p.SomeField
}
