package main

import("fmt")


func main(){
	var tableStrings = [7]string {"monday","tuesday","wednesday","thursday","friday","saturday","sunday"}
	tableNumbers := [6]int{1,2,3,4,5,6}
	fmt.Println("The Values of the table 1 : ")
	for _ , value := range tableStrings {
		fmt.Printf("\n %s",value)
	}
	fmt.Println("The Values of the table 2 : ")
	for _ , value := range tableNumbers {
		fmt.Printf("\n %d",value)
	}
}