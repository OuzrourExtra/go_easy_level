package main

import(
	"fmt"
	"strings"
	"strconv"
	"log"
	"bufio"
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

func number_to_day(number int) string{
	switch number{
	case 1 :
		return "Monday"
	case 2 : 
		return "Tuesday"
	case 3 :
		return "Wednesday"
	case 4 :
		return "Thursday"
	case 5 : 
		return "Friday"
	case 6 : 
		return "Saturday"
	case 7 :
		return "Sunday"
	default :
		return "Your number isn't in the range (1-7) . Sorry!"
	}
}

func main(){
	fmt.Println("==============================")
	fmt.Println("Select a number between 1-7 ")
	number , err := inputNumber()
	if err != nil{
		log.Fatal(err)
	}
	if number > 7 || number < 1{
		fmt.Printf("%s",number_to_day(int(number)))
	}else{
		fmt.Printf("The Day is : %s",number_to_day(int(number)))
	}
}