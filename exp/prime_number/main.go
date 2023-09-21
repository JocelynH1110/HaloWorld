package main

import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var min, max int
	fmt.Print("Minimum: ")
	fmt.Scanf("%d", &min)

	fmt.Print("Maximum: ")
	fmt.Scanf("%d", &max)

	for i := min; i <= max; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}
