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

func calculateChecksumFromDigits(digits []int) int {
	checksum := 0
	for i := len(digits) - 2; i >= 0; i -= 2 {
		digit := digits[i]
		doubled := digit * 2
		if doubled > 9 {
			checksum += doubled - 9
		} else {
			checksum += doubled
		}
	}

	for i := len(digits) - 1; i >= 0; i -= 2 {
		checksum += digits[i]
	}

	return checksum
}

func validateCardNumber(cardNumber string) (bool, error) {
	digits, err := stringToDigits(cardNumber)
	if err != nil {
		return false, err
	}
	checksum := calculateChecksumFromDigits(digits)
	return checksum%10 == 0, nil
}

func main() {
	valid, err := validateCardNumber(CARD_NUMBER)

	if err != nil {
		log.Fatalln("輸入錯誤")
	}

	fmt.Println(valid)

}
