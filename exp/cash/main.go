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
	coinValues := map[int]int{
		25: 0,
		10: 0,
		5:  0,
		1:  0,
	}
	totalCoins := 0

	change := getNonNegativeInt()
	fmt.Println(change)

	fmt.Println("----------")
	for coinValue := range coinValues {
		numCoins := change / coinValue
		coinValues[coinValue] = numCoins
		totalCoins += numCoins
		change -= numCoins * coinValue
	}

	fmt.Println("投入數量：")
	for coinValue, numCoins := range coinValues {
		if numCoins > 0 {
			fmt.Printf("%d元 :%d枚\n", coinValue, numCoins)
		}

	}
	fmt.Println("----------")
	fmt.Printf("%d枚\n", totalCoins)
}
