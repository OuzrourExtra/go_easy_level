package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
inputNumber is a function for inputing a number
*/
func inputNumber(order int64) (nb int64, err error) {
	fmt.Println("=======================")
	fmt.Printf("Enter Your number %d : ", order)
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

/*
biggestNumber is a function for finding the biggest number of 3 numbers
*/
func biggestNumber(nb1 int64, nb2 int64, nb3 int64) (nb int64, err error) {
	var result int64 = 0
	if nb1 >= nb2 && nb1 >= nb3 {
		result = nb1
	} else if nb2 >= nb1 && nb2 >= nb3 {
		result = nb2
	} else {
		result = nb3
	}

	if result == 0 {
		return 0, errors.New("all numbers are 0")
	} else {
		return result, nil
	}
}

func main() {
	nb1, _ := inputNumber(1)
	nb2, _ := inputNumber(2)
	nb3, _ := inputNumber(3)
	fmt.Println("=======================")
	result, err := biggestNumber(nb1, nb2, nb3)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("The biggest number of (%d,%d,%d) is : %d", nb1, nb2, nb3, result)
	}
}
