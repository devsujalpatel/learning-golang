package main

import (
	"fmt"
	"time"
)

// Order Struct

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // Nanosecond precision
}

func main() {
	myOrder := Order{
		id:        "123",
		amount:    100.0,
		status:    "received",
		createdAt: time.Now(),
	}
	fmt.Println("Order struct", myOrder)
}
