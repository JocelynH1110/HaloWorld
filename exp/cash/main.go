package main

import (
	"fmt"
	"strconv"
)

func getNonNegativeInt() int { //（）外是反回值
	for {
		var input string
		fmt.Print("Change owed: ")
		fmt.Scanln(&input)

		num, err := strconv.Atoi(input) //字串轉換成int
		//Ａtoi函數的返回值為(輸入的值,error)

		if num >= 0 && err == nil { //err == nil 指沒有錯誤的話
			return num
		}

	}

}

func main() {
	denominations := []int{25, 10, 5, 1}
	totalCoins := 0

	change := getNonNegativeInt()

	for _, coinValue := range denominations {
		numCoins := change / coinValue
		totalCoins += numCoins
		change -= numCoins * coinValue
	}

	fmt.Println(totalCoins)
}
