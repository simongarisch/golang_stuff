package main

import (
	"fmt"
	"reflect"
)

type singleton struct {
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()
	fmt.Println(s1, reflect.TypeOf(s1)) // &{} *main.singleton
	fmt.Println(s2, reflect.TypeOf(s2)) // &{} *main.singleton

	// check these pointers are the same
	fmt.Println(s1 == s2) // true
}
