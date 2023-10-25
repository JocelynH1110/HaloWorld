package main

import (
	"fmt"
)

func caculateHour(hours []int) float64 {
	//panic("unimplemented") panic以後程式不會啟動
	sum := 0.0
	for _, i := range hours {
		sum += float64(i)
	}
	return sum
}

func main() {
	var weeks int
	fmt.Print("Number of weeks taking CS50: ")
	fmt.Scanf("%d", &weeks)

	hours := make([]int, weeks)
	for i := 0; i < weeks; i++ {
		fmt.Printf("Week %d Hours: ", i)
		fmt.Scanf("%d", &hours[i])
	}

	var answer rune //rune=char
	fmt.Print("Enter T for total hours, A for average hours per week: ")
	fmt.Scanf("%c", &answer)

	total := caculateHour(hours)

	switch answer {
	case 'T', 't':
		fmt.Printf("%.1f hours\n", total)
	case 'A', 'a':
		fmt.Printf("%.1f hours\n", total/float64(weeks))
	}

}
