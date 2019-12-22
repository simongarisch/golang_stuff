package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// create the Animal interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

// define three types Cow, Bird, and Snake
type Cow struct{}
type Bird struct{}
type Snake struct{}

// For each of these three types,
// define methods Eat(), Move(), and Speak()
// first for the cow
func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

// bird
func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

// snake
func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func NewAnimal(animalType string) (Animal, error) {
	animalType = strings.ToLower(animalType)
	var animal Animal
	switch animalType {
	case "cow":
		return Cow{}, nil
	case "bird":
		return Bird{}, nil
	case "snake":
		return Snake{}, nil
	default:
		return animal, fmt.Errorf("Animal type %q is not valid", animalType)
	}
}

func AnimalQuery(animal Animal, info string) error {
	info = strings.ToLower(info)
	switch info {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		return fmt.Errorf("Info %q is not valid", info)
	}
	return nil
}

func getUserSelection() ([]string, error) {
	var inputStrings []string

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return inputStrings, err
	}

	inputStrings = strings.Fields(input)
	if len(inputStrings) != 3 {
		return inputStrings, errors.New("expected user to input exactly 3 strings separated by a space")
	}

	return inputStrings, nil
}

func main() {
	description := `
	Every command from the user must be either a 'newanimal' command or a 'query' command.
	Each 'newanimal' command must be a single line containing three strings: newanimal name type ('Cow', 'Bird', 'Snake')
	Each 'query' command must be a single line containing 3 strings: query name info ('eat', 'move', 'speak')
	`
	fmt.Println(description)
	animals := map[string]Animal{}

	for {
		fmt.Printf("> ")
		inputStrings, err := getUserSelection()
		if err != nil {
			fmt.Println("Error: ", err.Error())
			continue
		}

		command := inputStrings[0]
		name := inputStrings[1]
		typeOrInfo := inputStrings[2]

		switch command {
		case "newanimal":
			animal, err := NewAnimal(typeOrInfo)
			if err != nil {
				fmt.Println("Error: ", err.Error())
				continue
			}
			animals[name] = animal
			fmt.Println("Created it!")
			//fmt.Println("animals", animals)
		case "query":
			animal, ok := animals[name]
			if !ok {
				fmt.Printf("Animal name %q is not valid. Please use 'newanimal' to create %q first\n", name, name)
				continue
			}
			err := AnimalQuery(animal, typeOrInfo)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
		default:
			fmt.Printf("Command %q is not valid. Options are 'newanimal' or 'query'\n", command)
			continue
		}
	}
}
