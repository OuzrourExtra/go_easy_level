package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
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


func tableMultiplication(n int) {
	fmt.Println("-------------------------------")
	fmt.Printf(" The Table of Multiplication of %d \n",n)
	fmt.Println("-------------------------------")
	for i:=1 ; i<=10 ; i++{
		fmt.Printf("==========>  %d x %d = %d \n" , n,i,n*i)
		fmt.Println("-------------------------------")
	}
}

func main(){
	number , err := inputNumber()
	if err != nil {
		log.Fatal(err)
	}else{
		tableMultiplication(int(number))
	}
}