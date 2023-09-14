package main

import "fmt"

func getHeight() int {
	var height int

	for { //等同於Ｃ裡的while(true)
		fmt.Print("Height: ")
		fmt.Scanf("%d", &height) //Go的scanf可以返回多個值，所以要寫個＆讓他準確抓到位置

		if height >= 1 && height <= 8 {
			return height
		}

	}
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
