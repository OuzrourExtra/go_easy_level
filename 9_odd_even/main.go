package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
	"log"
)

func main(){
	fmt.Println("=======================")
	fmt.Printf("Enter Your number : " )
	reader := bufio.NewReader(os.Stdin)
	str , _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	nb , err := strconv.ParseInt(str,10,64)
	fmt.Println("=======================")
	if err != nil {
		log.Fatal(err)
	}
	if nb % 2 == 0 {
		fmt.Printf("The number '%d' is EVEN !",nb)
	} else {
		fmt.Printf("The number '%d' is ODD ! ",nb)
	}

}