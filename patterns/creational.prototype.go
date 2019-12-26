package main

// WIP - look into go deepcopy, otherwise you'll have bugs
// perhaps json marshall and unmarshall...

import "fmt"

type Dog struct {
	nickname string
	sound    string
}

type DogPrototypes struct {
	prototypes map[string]Dog
}

func (dp DogPrototypes) Register(name string, instance Dog) {
	db.prototypes[name] = instance
}

func (dp DogPrototypes) Unregister(name string) {
	delete(dp.prototypes, name)
}

func (dp DogPrototypes) Clone(name string) {
	delete(dp.prototypes, name)
}

func main() {
	dog := Dog{"Fido", "Woof"}
	fmt.Println(dog)
}
