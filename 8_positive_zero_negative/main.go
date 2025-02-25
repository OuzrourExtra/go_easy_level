package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
	"log"
)

func main(){
	fmt.Println("================================")
	// Input the number
	fmt.Printf("Enter a number : ")
	reader := bufio.NewReader(os.Stdin)
	str , _ := reader.ReadString('\n')
	fmt.Printf("================================\n")
	// For deleteing the '\n' that are added automatically when you 
	// press enter because it block the strconv.ParseInt function from parsing it !
	str = strings.TrimSpace(str)
	nb , err := strconv.ParseInt(str,10,64)
	if err != nil {
		log.Fatal(err)
	}
	if nb > 0 {
		fmt.Printf("The Number '%d' is POSITIVE !",nb)
	} else if nb == 0 {
		fmt.Printf("The Number '%d' is NULL ! ",nb)
	} else {
		fmt.Printf("The Number '%d' is NEGATIVE !",nb)
	}
}