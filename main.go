package main

import (
	"crypto/sha256"
	"fmt"
	"syscall"

	"golang.org/x/term"
)

/*
TODOS


*/

func main() {
	fmt.Println("Please enter your password, it will not show up when typing")
	var userString string
	bytePass, _ := term.ReadPassword(int(syscall.Stdin))

	userString = string(bytePass)

	h := sha256.New()
	h.Write([]byte(userString))
	fmt.Printf("%x", h.Sum(nil))

}
