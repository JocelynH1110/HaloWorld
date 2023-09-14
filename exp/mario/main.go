package main

import "fmt"

func getHeight() int {
	var height int
	fmt.Print("Height: ")
	fmt.Scanf("%d", &height)
	return height
}
func main() {
	var height int = getHeight()

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
