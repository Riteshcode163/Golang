package main

import "fmt"

func changenum(num *int) {
	*num = 5

}

func main() {

	num := 10

	changenum(&num)

	fmt.Println(num)

}
