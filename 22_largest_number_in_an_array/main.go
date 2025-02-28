package main

import "fmt"

func largestInArray(tab [7]int) int{
	var largest int = -1
	for _ , value := range tab{
		if value >= largest {
			largest = value
		}
	}
	return largest
}

func main(){
	tab := [7]int{789,234,320,492,394,29,29}
	fmt.Printf("The Largest number of %v is : %d", tab , largestInArray(tab))
}