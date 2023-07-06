package main

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/term"
)

/*
TODOS
- Program flow
- Add storage for passwords
- Connect to DB

*/

func main() {

	db, err := sql.Open("postgres", "postgres://aaron:pleasechangeme!@localhost/gopass?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (username, password_hash, created_at) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	fmt.Println("Please enter your password, it will not show up when typing")
	var userString string
	bytePass, _ := term.ReadPassword(int(syscall.Stdin))
	username := "Plato"
	passwordHash := string(bytePass)
	createdAt := time.Now()

	//write to DB

	_, err = stmt.Exec(username, passwordHash, createdAt)
	if err != nil {
		log.Fatal(err)
	}

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
