package main

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
	"log"
)

/*
inputNumber is a function for inputing a number
*/
func inputNumber(sentence string) (nb int64, err error) {
	fmt.Println("=======================")
	fmt.Printf("%s: ",sentence)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	nb, err1 := strconv.ParseInt(str, 10, 64)
	if err1 != nil {
		return 0, err1
	} else {
		return nb, nil
	}
}

func appendToTheSlice(tab *[]int , cap int){
	
	for i:= 0 ; i<cap ; i++{
		value , err := inputNumber("Enter Your Number")
		if err != nil{
			log.Fatal(err)
		}
		*tab = append(*tab, int(value))
	}
}

func sumSlice(slice []int) int{
	sum := 0
	for _ , value := range slice {
		sum += value
	}
	return sum
}

func main(){
	numberOfElements , err :=inputNumber("Enter The Number of Elements ")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("===================================")
	tab := make([]int , 0 , numberOfElements )
	appendToTheSlice(&tab , int(numberOfElements))
	fmt.Printf(" The Slice : %v \n",tab)
	fmt.Printf(" The Sum of the Elements %d",sumSlice(tab))
}
