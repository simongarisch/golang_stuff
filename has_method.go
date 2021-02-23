// go run has_method.go
package main

import (
	"fmt"
	"reflect"
)

type myStruct struct {
	a int
}

type myInterface interface {
	SetA(int)
}

func (s *myStruct) SetA(a int) {
	s.a = a
}

func main() {
	s := myStruct{}
	s.SetA(1)
	fmt.Println(s) // {1}

	st := reflect.ValueOf(&s)
	fmt.Println(st) // main.myStruct
	m1 := st.MethodByName("SetA")
	m2 := st.MethodByName("SetB")
	fmt.Println(m1)           // 0x487860
	fmt.Println(m2)           // <invalid reflect.Value>
	fmt.Println(m2.IsValid()) // false

	if m1.IsValid() {
		args := []int{7} // just one arg
		inputs := make([]reflect.Value, len(args))
		for i := range args {
			inputs[i] = reflect.ValueOf(args[i])
		}
		m1.Call(inputs)
	}

	fmt.Println(s) // {7}
}
