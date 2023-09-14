package main

import "fmt"

func main() {
	var height int = 3
	for y := 1; y <= height; y++ {

		for x := 1; x <= height-y; x++ {
			fmt.Print(" ")
		}
		for x := 1; x <= y; x++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

}
