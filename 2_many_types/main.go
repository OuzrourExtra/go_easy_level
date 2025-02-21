package main

import "fmt"

func main(){
	// int
	var int8Var int8 = 127
	var int16Var int16 = 32767
	var int32Var int32 = 2147483647
	var int64Var int64 = 9223372036854775807
	var intVar int = 9223372036854775807

	// unsigned int
	var uint8Var uint8 = 255
	var uint16Var uint16 = 65535
	var uint32Var uint32 = 4294967295
	var uint64Var uint64 = 18446744073709551615
	// Float

	var float32Var float32 = 92233333368547758079223372036854775807.3
	var float64Var float64 = 92233720368547758079223372036854775807922337203685477580792233720368547758079223372036854775807922337203685477580792233720368547758079223372036854775807922337203685477580792233720368547758079223372036854775807922337203685477580792233720368549999775807922337203685477580792233720368547758079223372036854775807.3
	// Complex types
	var complex64Var complex64 = complex(1, 2)
	var complex128Var complex128 = complex(1, 2)


	// String
	var stringVar string = "string :)"

	//bool
	var boolVar bool = false

	// runes
	var runeVar rune = '#'

	// bytes
	var byteVar byte = 'l'

	/** 
	for printing the values 
	*/
	
	// int
	fmt.Printf("int : %d \n",intVar)
	fmt.Printf("int8 : %d \n",int8Var)
	fmt.Printf("int16 : %d \n",int16Var)
	fmt.Printf("int32 : %d \n",int32Var)
	fmt.Printf("int64 : %d \n",int64Var)
	fmt.Printf("uint8 : %d \n",uint8Var)
	// uint
	fmt.Printf("uint16 : %d \n",uint16Var)
	fmt.Printf("uint32 : %d \n",uint32Var)
	fmt.Printf("uint64 : %d \n",uint64Var)
	// float
	fmt.Printf("float32 : %f \n",float32Var)
	fmt.Printf("float64 : %f \n",float64Var)
	// complex
	fmt.Printf("complex64: %v\ncomplex128: %v\n\n", complex64Var, complex128Var)
	// string
	fmt.Printf("stringVar : %s \n",stringVar)
	// boolean
	fmt.Printf("boolVar : %t \n",boolVar)
	// rune
	fmt.Printf("runeVar : %c %d \n",runeVar , runeVar)
	// byte
	fmt.Printf("byteVar : %c %d \n",byteVar,byteVar)
}
