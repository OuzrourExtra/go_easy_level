package main


import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strings"
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


func palindrome (s string) bool {
	var inversed strings.Builder
	strOriginal := strings.ReplaceAll(strings.TrimSpace(s)," ","")
	
	for i:=len(strOriginal)-1 ; i>=0 ; i--{
		inversed.WriteByte(strOriginal[i])
	}
	if strOriginal == inversed.String(){
		return true
	} else {
		return false
	}
}


func main(){
	input , err := inputString()
	if err!=nil{
		log.Fatal(err)
	}
	input = strings.ToLower(input)
	palindrome_bool := palindrome(input)
	if palindrome_bool {
		fmt.Printf("\n ==> The String '%s' is a Palindrome !",input)
	}else{
		fmt.Printf("\n ==> The String '%s' is NOT a palindrome ! ",input)
	}
}