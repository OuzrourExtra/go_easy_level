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

func vowels(str string) int {
	counter := 0
	vowelMaps := map[rune]bool{'a':true,'e':true,'u':true,'i':true,'o':true}
	for _,char := range(str){
		if vowelMaps[char]{
			counter++
		}
	}
	return counter
}
func main(){
	str , err:= inputString()
	if err != nil{
		log.Fatal(err)
	}
	lowerCase := strings.ToLower(str)
	fmt.Printf("The Number of Vowels in '%s' is : %d",str,vowels(lowerCase))
}