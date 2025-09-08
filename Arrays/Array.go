package main

import "fmt"

func main() {
	Num := [5]int{2, 3, 4, 5, 6}

	for i := 0; i < 5; i++ {

		fmt.Println(Num[i])

	}
	/////////////////////////

	name := [3]string{"Ritesh", "Shivkiran", "patil"}
	for i := 0; i < 3; i++ {

		fmt.Print(name[i], " ")

	}
	////////////////////////////

	number := [2][2]int{{2, 3}, {4, 5}}
	fmt.Println(number[0][0])

}
