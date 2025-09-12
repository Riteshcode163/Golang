package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

func main() {
	result := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) //////always store value and then print
	fmt.Println(result)
}
