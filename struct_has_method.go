// go run struct_has_method.go
package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct{}

func (m MyStruct) DoThis() {
	fmt.Println("doing this...")
}

func (m MyStruct) DoThat() {
	fmt.Println("doing that...")
}

type ParentStruct struct {
	MyStruct
}

func (p ParentStruct) HasMethod(methodName string) bool {
	t := reflect.TypeOf(p)
	_, ok := t.MethodByName(methodName)
	return ok
}

func main() {
	p := ParentStruct{}
	p.DoThis() // doing this...
	p.DoThat() // doing that...

	t := reflect.TypeOf(p)
	fmt.Println(t)                        // main.ParentStruct
	fmt.Println(t.MethodByName("DoThis")) // {DoThis  func(main.ParentStruct) <func(main.ParentStruct) Value> 1} true

	fmt.Println(p.HasMethod("DoThis")) // true
	fmt.Println(p.HasMethod("DoThat")) // true
	fmt.Println(p.HasMethod("DoThen")) // false
}
