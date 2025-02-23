/*
Project created for training

By:Ouzrour Ilyas
Date: 23/2/25
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// converting float64t to string using package "strconv"
func float64_to_string(number float64) string {
	var fmt byte = 'f' // no exponent
	precision := -1    // for precision , i want just the minimum numbers
	bitSize := 64      // for float64
	return strconv.FormatFloat(number, fmt, precision, bitSize)
}

// converting string to float64 using package "strconv"
func string_to_float64(str string) float64 {
	bitSize := 64 // for float64
	number, _ := strconv.ParseFloat(str, bitSize)
	return number
}

func main() {
	for {
		fmt.Println("1. String to Float64 ")
		fmt.Println("2. Float64 to String ")
		fmt.Printf("What do you want to do ? I want :  ")
		choice, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		choice = strings.TrimSpace(choice)
		choice_to_int, _ := strconv.ParseInt(choice, 10, 32)
		fmt.Println(choice_to_int)
		if choice_to_int != 1 && choice_to_int != 2 {
			fmt.Println("========================")
			fmt.Println("Wrong choice ! ")
			fmt.Println("========================")
		} else {
			if choice_to_int == 1 {
				fmt.Printf("\nEnter your string :")
				str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				str = strings.TrimSpace(str)
				fmt.Printf("\n String : '%s' ==> Float64 :%f", str, string_to_float64(str))
			} else {
				fmt.Printf("\nEnter your number :")
				float, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				float = strings.TrimSpace(float)
				floatVar := string_to_float64(float)
				fmt.Printf("\n Float64 : %f ==> String :'%s'\n", floatVar, float64_to_string(floatVar))
			}
			wait, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			fmt.Println(wait)
			fmt.Println("========================")
			fmt.Println("succesfully done !")
			fmt.Println("========================")

		}
	}
}
