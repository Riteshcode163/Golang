package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])

	}
	/////////

	var num = make([]int, 0, 5)
	num = append(num, 1)
	num = append(num, 2)
	num = append(num, 3)
	num = append(num, 4)
	num = append(num, 5)
	num = append(num, 6)

	for i := 0; i < len(num); i++ {
		fmt.Print(num[i])

	}

	// fmt.Println(num)
	fmt.Println(cap(num))
	fmt.Println(len(num))
}
