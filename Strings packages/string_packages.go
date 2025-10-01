package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "Apple,orange,banana,kiwi"
	//i want them in comma seperated slice
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	////counting

	str := "One two three two four two  five"
	count := strings.Count(str, "two")
	fmt.Println(count)

	///trimming

	str = "            hello, Go!   "
	trimmed := strings.TrimSpace(str)
	fmt.Println(trimmed)

	///join string

	str1 := "Ritesh"
	str2 := "Patil"
	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println(result)

}
