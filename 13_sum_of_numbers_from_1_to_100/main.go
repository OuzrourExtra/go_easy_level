package main

import "fmt"
/*
 sum is a function that find the sum of numbers from 1 to a given number.
 */
 func sum(n int) int {
	result := 0
	for i:=1 ; i<=n ; i++{
		result += i
	}
	return result
}

func main(){
	fmt.Printf("The sum of numbers from 1 to 100 is : %d",sum(100))
}