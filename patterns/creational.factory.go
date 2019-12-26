package main

import (
	"errors"
	"fmt"
)

const (
	TypeRural = iota
	TypeCity
)

type Person interface {
	GetName() string
}

type RuralPerson struct{}

type CityPerson struct{}

func (p RuralPerson) GetName() string {
	return "Cletus"
}

func (p CityPerson) GetName() string {
	return "Dan the man"
}

func PersonFactory(person_type int) (Person, error) {
	if person_type == TypeRural {
		return RuralPerson{}, nil
	} else if person_type == TypeCity {
		return CityPerson{}, nil
	} else {
		return nil, errors.New("Unsupported person type")
	}
}

func main() {
	var p Person
	var err error

	// create a rural person
	p, err = PersonFactory(TypeRural)
	if err != nil {
		panic(err)
	}
	fmt.Println(p, p.GetName()) // {} Cletus

	// create a city person
	p, err = PersonFactory(TypeCity)
	if err != nil {
		panic(err)
	}
	fmt.Println(p, p.GetName()) // {} Dan the man
}
