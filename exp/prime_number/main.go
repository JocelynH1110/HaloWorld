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

// func getNumber(num, num2 int) []int {
// 	primes := []int{}
// 	for i := num; i <= num2; i++ {
// 		if isPrime(i) {
// 			primes = append(primes, i)

// 		}

// 	}
// 	return primes
// }

func main() {
	for i:=1;i<=10;i++{
		if isPrime(i){
			fmt.Println(i)
		}
	}
	// for {
	// 	fmt.Print("Minimun: ")
	// 	fmt.Scanf("%d", &num)

	// 	fmt.Print("Maximun: ")
	// 	fmt.Scanf("%d", &num2)

	// 	if num > num2 {
	// 		return
	// 	} else if num < 1 {
	// 		return
	// 	}

	// }

	//fmt.Printf()

}
