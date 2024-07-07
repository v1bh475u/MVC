package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: hashpassword <password>")
		os.Exit(1)
	}
	password := os.Args[1]
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		os.Exit(2)
	}
	fmt.Println(string(hashedPassword))
}
