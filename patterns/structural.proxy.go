package main

import (
	"fmt"
	"strings"
)

type Passwords interface {
	GetMasterPassword() string
}

type SimonsPasswords struct{}
type PasswordsProxy struct {
	user          string
	userPasswords SimonsPasswords
}

func NewPasswordsProxy(user string) PasswordsProxy {
	proxy := PasswordsProxy{
		user,
		SimonsPasswords{},
	}
	return proxy
}

func (passwords SimonsPasswords) GetMasterPassword() string {
	return "Wubba lubba dub dub"
}

func (proxy PasswordsProxy) GetMasterPassword() string {
	password := proxy.userPasswords.GetMasterPassword()
	if strings.ToLower(proxy.user) == "simon" {
		return password
	}
	return strings.Repeat("*", len(password))
}

func main() {
	var proxy Passwords
	proxy = NewPasswordsProxy("Bob")
	fmt.Println(proxy.GetMasterPassword()) // *******************

	proxy = NewPasswordsProxy("Simon")
	fmt.Println(proxy.GetMasterPassword()) // Wubba lubba dub dub
}
