package main

import (
	"fmt"
	"os"
)

func replace(word string) string {
	chars := []rune(word)
	for i, chr := range chars {
		switch chr {
		case 'a':
			chars[i] = '6'
		case 'e':
			chars[i] = '3'
		case 'i':
			chars[i] = '1'
		case 'o':
			chars[i] = '0'
		}
	}
	return string(chars)
}

func main() {

	for i := 1; i < len(os.Args); i++ {
		fmt.Print(replace(os.Args[i]))
		fmt.Print(" ")
	}
	fmt.Print("\n")
}
