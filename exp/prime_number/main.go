package main

import "fmt"

func isPrime(num3 int)bool{
	for{fmt.Print("Minimun: ")
	fmt.Scanf("%d",&num)
	
	fmt.Print("Maximun: ")
	fmt.Scanf("%d",&num2)

	var num,num2 int
	
	if num3<=1{
		return false
	}
	for i:=num;i<num2;i++{
		if num2%i==0{
			return false
		}
	}

	return true
}
}
func main(){
	var num,num2 int=isPrime(int)

	for i:=num;i<=num2;i++{
		if isPrime(i){
			fmt.Println(i)

		}
		}

	}
	