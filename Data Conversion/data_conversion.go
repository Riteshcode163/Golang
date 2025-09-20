package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 34
	fmt.Println("The number is:", num)
	fmt.Printf("The data type is: %T\n", num)

	/////Int to float

	var data float64 = float64(num)
	fmt.Printf("The data type of num after conversion is : %T\n", data)

	/////Int to string

	str := strconv.Itoa(num)
	fmt.Println("The number is:", str)
	fmt.Printf("The data type is: %T\n", str)

	/////String to int
	number_string := "1204"

	number_int, _ := strconv.Atoi(number_string) // underscore because it return 2 value int and erroe but we dont want error in place of that underscore ignore it

	fmt.Printf("The data type is: %T\n", number_int)

	/////string to float

	num_string := "85.40"

	num_float, _ := strconv.ParseFloat(num_string, 64)

	fmt.Printf("The data type is: %T\n", num_float)

}
