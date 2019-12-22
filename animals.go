package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func getUserSelection() ([]string, error) {
	var inputStrings []string

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return inputStrings, err
	}

	inputStrings = strings.Fields(input)
	if len(inputStrings) != 2 {
		return inputStrings, errors.New("expected user to input exactly 2 strings separated by a space")
	}

	return inputStrings, nil
}

func main() {
	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	//fmt.Println(animals)

	description := `
	Your request should be an animal followed by a space and animal action...
	Available animals are 'cow', 'bird', or 'snake'.
	Avaliable actions are 'eat', 'move', or 'speak'. 
	`
	fmt.Println(description)
	for {
		fmt.Printf("> ")
		inputStrings, err := getUserSelection()
		if err != nil {
			fmt.Println("Error: ", err.Error())
			continue
		}

		animalName := inputStrings[0]
		animalAction := inputStrings[1]

		animal, found := animals[animalName]
		if !found {
			fmt.Println("Animal selection %q is not valid", animalName)
			continue
		}

		switch animalAction {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Animal action selection is invalid")
		}
	}
}
