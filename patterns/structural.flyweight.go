package main

import (
	"errors"
	"fmt"
)

// our choice of 2 presidents
const (
	PresidentBush = iota
	PresidentClinton
)

type President interface {
	GetType() int
	Speech() string
}

type Bush struct{}
type Clinton struct{}

func (president Bush) GetType() int {
	return PresidentBush
}

func (president Clinton) GetType() int {
	return PresidentClinton
}

func (president Bush) Speech() string {
	return "To the C students... I say, you too can be President of the United States."
}

func (president Clinton) Speech() string {
	return "I did not have sexual relations with that woman"
}

func PresidentFactory(presidentType int) (President, error) {
	switch presidentType {
	case PresidentBush:
		return Bush{}, nil
	case PresidentClinton:
		return Clinton{}, nil
	default:
		return nil, errors.New("presidentType not available.")
	}

}

type FlyweightObjects struct {
	pool map[int]*President
}

func (f *FlyweightObjects) GetFlyweight(presidentType int) (*President, error) {
	if f.pool == nil {
		f.pool = make(map[int]*President)
	}

	_, ok := f.pool[presidentType]
	if !ok {
		fmt.Println("Creating ", presidentType)
		president, err := PresidentFactory(presidentType)
		if err != nil {
			return &president, err
		}
		f.pool[presidentType] = &president
	}

	return f.pool[presidentType], nil
}

func main() {
	flyweights := FlyweightObjects{}

	bush1, _ := flyweights.GetFlyweight(PresidentBush)
	bush2, _ := flyweights.GetFlyweight(PresidentBush)

	clinton1, _ := flyweights.GetFlyweight(PresidentClinton)
	clinton2, _ := flyweights.GetFlyweight(PresidentClinton)

	fmt.Println((*bush1).GetType(), (*bush1).Speech())       // 0 To the C students...
	fmt.Println((*clinton1).GetType(), (*clinton1).Speech()) // 1 I did not have...

	fmt.Println(bush1 == bush1)       // true, pointers to the same address
	fmt.Println(bush1 == bush2)       // true
	fmt.Println(clinton1 == clinton2) // true
}
