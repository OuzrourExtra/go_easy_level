package main

import (
	"fmt"
	"log"
	"bufio"
	"strings"
	"strconv"
	"os"
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

// leapYear a function for detecting if the year is leap or not
func leapYear(year int) bool {
	switch{
	case year%400 == 0 :
		return true
	case year%100 == 0:
		return false
	case year%4 == 0 :
		return true
	default :
		return false
	}
}


func main(){
	year , err := inputNumber()
	if err != nil {
		log.Fatal(err)
	}
	leapBool := leapYear(int(year))
	if leapBool {
		fmt.Printf("The Year %d is Leap !",year)
	}else{
		fmt.Printf("The Year %d is not Leap ! ",year)
	}
}