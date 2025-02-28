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
func inputNumber() (nb int64, err error) {
	fmt.Println("=======================")
	fmt.Printf("Enter Your number: ")
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

func appendToTheSlice(tab *[]int){
	i:= true
	for i{
		value , err := inputNumber()
		if err != nil{
			log.Fatal(err)
		}
		*tab = append(*tab, int(value))
		fmt.Printf("Do You want to stop ? y/n : ")
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		if str == "y"{
			i=false
		}
	}
}

func main(){
	tab := make([]int , 5)
	appendToTheSlice(&tab)
	fmt.Printf(" The Result : %v",tab)
}