package main

import "fmt"

func greet(name string) {
	if name == "" {
		fmt.Println("HaloWorld")
	} else {
		fmt.Printf("Haloooo,%s\n", name)
	}
}

func main() {
	var name string
	fmt.Println("whats your name?")
	fmt.Scanln(&name)
	greet(name)
}
