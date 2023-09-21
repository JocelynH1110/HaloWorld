package main

import "fmt"

func getNonNegativeInt() int { //（）外是反回值
	var num int

	for {
		fmt.Print("Change owed: ")
		_, err := fmt.Scanf("%d", &num)

		if num >= 0 && err == nil {
			return num
		}

	}

}

func main() {
	change := getNonNegativeInt()
	fmt.Println(change)
}
