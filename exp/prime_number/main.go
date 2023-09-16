package main

import "fmt"

func main(){

	var num,num2 int
	fmt.Print("Minimun: ")
	fmt.Scanf("%d",&num)

	fmt.Print("Maximun: ")
	fmt.Scanf("%d",&num2)

	for i:=1;i<=num2;i++{
		if num2%i:=0{

			sum=sum+i
			fmt.Println(sum)
		}
	}
}