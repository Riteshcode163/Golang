package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float64
	status    string
	createdat time.Time
}

func main() {

	order := order{
		id:        "Ritesh",
		amount:    344.76,
		status:    "Pending",
		createdat: time.Now(),
	}

	fmt.Print("Id is : ", order.id, " Amount is : ", order.amount, " Status is : ", order.status, "Time is : ", order.createdat)

}
