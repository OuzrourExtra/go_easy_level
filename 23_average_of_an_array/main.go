package main

import "fmt"

func averageTable(tab [10]int) (int){
	sum := 0
	for _ , value := range tab {
		sum += value
	}
	return sum
}

func main(){
	tab := [10]int{1,2,3,4,55,6,7,39,20,30}
	fmt.Printf("The average is : %.2f" , float64(averageTable(tab))/10)
}