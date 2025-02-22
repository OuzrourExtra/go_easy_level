/*
Project created for training

By:Ouzrour Ilyas
Date: 23/2/25
*/

package main

import (
	"fmt"
	//"Log"
	"bufio"
	"os"
	"log"
	"strconv"
	"strings"
)

func swap(nb1 int64 , nb2 int64) (int64,int64) {
	return nb2,nb1
}

func main(){
	fmt.Println("Enter the first number ( must be an integer ) : ")
	reader := bufio.NewReader(os.Stdin)
	a , err := reader.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Enter the second number ( must be an integer ) : ")
	b ,err2 := bufio.NewReader(os.Stdin).ReadString('\n')
	if err2 != nil {
		log.Fatal(err)
	}
	// for removing the \n from the input
	a= strings.TrimSpace(a)
	b= strings.TrimSpace(b)
	// for converting values to int to show them when using %d instead
	// of being just strings
	c , _ := strconv.ParseInt(a,10,64)
	d , _ := strconv.ParseInt(b,10,64)

	fmt.Printf("\n You have entrered : %d %d ",c,d)
	c , d = swap(c,d)
	fmt.Printf("\n Swaped values : %d %d ",c,d)
}