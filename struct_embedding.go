// go run struct_embedding.go
// https://eli.thegreenplace.net/2020/embedding-in-go-part-1-structs-in-structs/
package main

import "fmt"

type User struct {
	Name string
}

func (u User) SayHello() {
	fmt.Println("Hello, " + u.Name)
}

type Admin struct {
	User
	Permissions map[string]string
}

func main() {
	admin := Admin{User: User{Name: "Simon"}}
	fmt.Println(admin.Name) // Simon
	admin.SayHello()        // Hello, Simon
}
