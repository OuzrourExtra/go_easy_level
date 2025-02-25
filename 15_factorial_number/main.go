package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	"bufio"
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

func factorial(n int) int{
	fact:=1
	for i:=1 ; i<=n ; i++{
		fact *= i
	}
	return fact
}

func main(){
	number , err := inputNumber()
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Printf("\n------------------------------- \n")
	fmt.Printf("            %d! = %d", number , factorial(int(number)))
	fmt.Printf("\n------------------------------- \n")

}