package main

import (
	"crypto/sha256"
	"fmt"
	"syscall"

	"golang.org/x/term"
)

/*
TODOS
- Program flow
- Add storage for passwords

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

/*

POSTGRES DB SCHEMA

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS passwords (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    website VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

Database schema ^


*/
