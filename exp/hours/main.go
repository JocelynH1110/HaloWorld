package main

import (
	"fmt"
)

func caculateHour() float64 {
	panic("unimplemented")
}

func main() {
	var weeks int
	fmt.Print("Number of weeks taking CS50:")
	fmt.Scanf("%d", &weeks)

	hours := make([]int, weeks)
	for i := 0; i < weeks; i++ {
		fmt.Printf("Week %d Hours: ", i)
		fmt.Scanf("%d", &hours[i])
	}

	var answer rune //rune=char
	fmt.Print("Enter T for total hours, A for average hours per week: ")
	fmt.Scanf("%c", &answer)
	fmt.Println(string(answer))
}
