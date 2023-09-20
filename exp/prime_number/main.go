package main

import "fmt"

func isPrime(num3 int) bool {

	if num3 <= 1 {
		return false
	}
	for i := num; i < num2; i++ {
		if num2%i == 0 {
			return false
		}
	}

	return true
}

func getNumber(num, num2 int) []int {
	primes := []int{}
	for i := num; i <= num2; i++ {
		if isPrime(i) {
			primes = append(primes, i)

		}

	}
	return primes
}
func main() {

	var num, num2 int
	for {
		fmt.Print("Minimun: ")
		fmt.Scanf("%d", &num)

		fmt.Print("Maximun: ")
		fmt.Scanf("%d", &num2)

		if num > num2 {
			return
		} else if num < 1 {
			return
		}

	}
	fmt.Printf("%v", getNumber)

}
