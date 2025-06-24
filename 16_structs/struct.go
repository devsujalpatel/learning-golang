package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

// Order Struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // Nanosecond precision
	customer
}

// func newOrder(id string, amount float32, status string) *order {
// 	// initial setup of struct
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
// }

// // receiver type
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func (o order) getAmount() float32 {
// 	return o.amount
// }

func main() {
	newCustomer := customer{
		name:  "John Doe",
		phone: "1234567890",
	}

	newOrder := order{
		id:        "1",
		amount:    100.0,
		status:    "received",
		createdAt: time.Now(),
		customer: newCustomer,
	}
	
fmt.Println(newOrder)
	println(newOrder.customer.name)
	println(newOrder.customer.phone)

	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{
	// 	name:   "Go",
	// 	isGood: true,
	// }

	// println(language.name)
	// println(language.isGood)

	// myOrder := newOrder("1", 100.0, "received")

	// fmt.Println(myOrder)

	// if you don't set any filed, default value is zero value
	// int => 0, string => "", bool => false, other => nil
	// myOrder := order{
	// 	id: "123",
	// 	// amount:    100.0,
	// 	status:    "received",
	// 	createdAt: time.Now(),
	// }

	// myOrder.changeStatus("shipped")

	// fmt.Println(myOrder.getAmount())

	// // fmt.Println(myOrder.status)

	// myOrder2 := order{
	// 	id:        "2",
	// 	amount:    50.0,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }
	// myOrder.amount = 200

	// fmt.Println("Order struct 1", myOrder)
	// fmt.Println("Order struct 1",myOrder2)

}
