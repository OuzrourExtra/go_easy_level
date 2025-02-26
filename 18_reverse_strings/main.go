package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)



/*
inputNumber is a function for inputing a number
*/
func inputString() (str string, err error) {
	fmt.Println("=======================")
	fmt.Printf("Enter Your String : ")
	reader := bufio.NewReader(os.Stdin)
	result, err := reader.ReadString('\n')
	return result , err
}

func main(){
	str , err := inputString()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n You Printed : %s \n The Result is : ",str)
	for i:=len(str)-1 ; i>=0 ; i--{
		fmt.Printf("%s",string(str[i]))
	}
}