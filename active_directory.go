// go run active_directory.go
// go get github.com/contester/runlib/win32
package main

import (
	"fmt"

	"github.com/contester/runlib/win32"
	"golang.org/x/sys/windows"
)

var domain = "thedomain"

func auth(username string, password string) bool {
	pusername, _ := windows.UTF16PtrFromString(username)
	pdomain, _ := windows.UTF16PtrFromString(domain)
	ppassword, _ := windows.UTF16PtrFromString(password)

	_, err := win32.LogonUser(
		pusername,
		pdomain,
		ppassword,
		3,
		0,
	)

	fmt.Println(err) // LogonUser: The user name or password is incorrect.
	if err == nil {
		return true
	}
	return false
}

func main() {
	username := "simon"
	password := "notMyPassword"

	fmt.Println(auth(username, password)) // false
}
