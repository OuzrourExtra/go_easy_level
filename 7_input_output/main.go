package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

func main(){
	fmt.Println("================================")
	fmt.Printf("Enter Something : ")
	reader := bufio.NewReader(os.Stdin)
	str , _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	fmt.Println("================================")
	fmt.Printf("---> Your input is : '%s' ",str)
	fmt.Println("\n ================================")
}