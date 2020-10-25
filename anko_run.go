// go run anko_run.go
package main

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

func main() {
	e := env.NewEnv()

	err := e.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	script := `
println("Hello World :)")

a = map[interface]interface{"x": 1}
println(a) // map[x:1]
a.b = 2
a["c"] = 3
println(a["b"]) // 2
println(a.c) // 3
`

	_, err = vm.Execute(e, nil, script)
	if err != nil {
		log.Fatalf("Execute error: %v\n", err)
	}

	// output: Hello World :)
	fmt.Println("------------")
	time.Sleep(2 * time.Second)
	fmt.Println(e)
	fmt.Println("------------")
	fmt.Println(reflect.TypeOf(e)) // *env.Env from github.com/mattn/anko/env
	value, err := e.Get("a")
	fmt.Println(value) // map[b:2 c:3 x:1]
	fmt.Println(err)   // <nil>
}
