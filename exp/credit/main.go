package main

import "fmt"
import "log"
import "errors"

const CARD_NUMBER string = "4003600000000014"

func stringToDigits(str string) ([]int, error) {

	digits := make([]int, len(str))

	for i, chr := range str {
		if chr < 48 || chr > 57 {
			err := errors.New("不正確的數字")
			return []int{}, err
		}
		digit := int(chr) - 48
		digits[i] = digit

	}

	return digits, nil
}
func main() {
	digits, err := stringToDigits(CARD_NUMBER)

	if err != nil {
		log.Fatalln("輸入錯誤")
	}

	for _, digit := range digits {

		fmt.Println(digit)
	}
}
