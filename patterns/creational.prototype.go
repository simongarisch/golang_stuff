package main

// WIP - look into go deepcopy, otherwise you'll have bugs
// perhaps json marshall and unmarshall...

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Dog struct {
	Nickname string
	Sound    string
}

type DogPrototypes struct {
	prototypes map[string][]byte
}

func NewDogPrototypes() DogPrototypes {
	dogPrototypes := DogPrototypes{}
	dogPrototypes.prototypes = map[string][]byte{}
	return dogPrototypes
}

func (dp DogPrototypes) Register(name string, instance Dog) {
	bytes, err := json.Marshal(instance)
	if err != nil {
		panic(err)
	}
	dp.prototypes[name] = bytes
}

func (dp DogPrototypes) Unregister(name string) {
	delete(dp.prototypes, name)
}

func (dp DogPrototypes) Clone(name string) (Dog, error) {
	var dog Dog
	bytes, ok := dp.prototypes[name]
	if !ok {
		return dog, errors.New("No prototype available for " + name)
	}
	json.Unmarshal(bytes, &dog)
	return dog, nil
}

func main() {
	dog1 := Dog{"Fido", "Woof"}
	fmt.Println(dog1) // {Fido Woof}

	prototypes := NewDogPrototypes()
	prototypes.Register("Fido", dog1)
	fmt.Println(prototypes) // {map[Fido:[123 125 ...]]}

	dog2, err := prototypes.Clone("Fido")
	if err != nil {
		panic(err)
	}
	fmt.Println(dog2) // {Fido Woof}

	fmt.Println(&dog1 == &dog2) // false
	dog2.Nickname = "Catdog"
	dog2.Sound = "Meow"
	fmt.Println(dog1, dog2) // {Fido Woof} {Catdog Meow}
}
