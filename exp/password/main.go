package main

import (
	"fmt"
	"unicode"
)

func check(word string) bool {
	var (
		lowerCase bool
		upperCase bool
		number    bool
		symbol    bool
	)
	for _, char := range word {
		if char >= 'a' && char <= 'z' {
			lowerCase = true
			continue
		}
		if char >= 'A' && char <= 'Z' {
			upperCase = true
			continue
		}
		if char >= '0' && char <= '9' {
			number = true
			continue
		}
		if char < 128 && unicode.IsPrint(char) {
			symbol = true
		}
	}
	return lowerCase && upperCase && number && symbol
}

func main() {
	var password string
	fmt.Print("Enter your password: ")
	fmt.Scanf("%s", &password)

	if check(password) {
		fmt.Println("Your password is valid!")
	} else {
		fmt.Println("Your password needs at least one uppercase letter, lowercase letter, number ,and symbol!")
	}
}
