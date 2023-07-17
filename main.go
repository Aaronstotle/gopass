package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"syscall"

	_ "github.com/lib/pq"
	"golang.org/x/term"
)

/*
STATUS: Postgres can receive queries, need to ask for username. Users should be able to have multiple entries in the password table.
TODO: put DB writes in a function
*/

func main() {

	db, err := sql.Open("postgres", "postgres://aaron:pleasechangeme!@localhost/gopass?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// stmt, err := db.Prepare("INSERT INTO users (username, password_hash, created_at) VALUES ($1, $2, $3)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()

	userExists := false

	username := greeting() // Assign variables to output of functions to use in DB insert later
	fmt.Printf("Your username is %s\n", username)
	if userExists {
		//TODO: Add DB Query to search for user and  passwords
		fmt.Printf("User exists already, lets get your passwords\n")
	} else {
		fmt.Printf("User does not existing, writing to database")
	}
	passwordhash := hashPass()
	fmt.Println(passwordhash)
	exit()

	// _, err = stmt.Exec(username, passwordHash, createdAt)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%x", h.Sum(nil))

}

func greeting() string {

	var username string

	fmt.Println("Please enter your username")
	fmt.Scanf("%s", &username)

	return username
}

func hashPass() string {
	fmt.Println("Please enter your password, it will not show up when typed")
	bytePass, _ := term.ReadPassword(int(syscall.Stdin))
	pass := string(bytePass)

	h := sha256.New()
	h.Write([]byte(pass))
	passwordHash := hex.EncodeToString(h.Sum(nil))
	//fmt.Printf("Your hashed password is, %x\n", h.Sum(nil))
	return passwordHash
}

func exit() int {
	fmt.Println("\nExiting program, thanks for playing!")
	os.Exit(0)
	return 0
}
