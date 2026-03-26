package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"

func generatePassword(length int) string {
	password := ""
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, charsetLength)
		password += string(charset[num.Int64()])
	}

	return password
}

func main() {
	var length int

	fmt.Println("==== GO PASSWORD GENERATOR ====")
	fmt.Print("Enter password length: ")
	fmt.Scanln(&length)

	if length <= 0 {
		fmt.Println("Invalid length")
		return
	}

	password := generatePassword(length)
	fmt.Println("Generated Password:", password)
}