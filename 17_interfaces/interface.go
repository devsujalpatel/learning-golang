package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, accountNum string)
}

type payment struct {
	gateway paymenter
}

// Open close principle is violating here
func (p payment) makePayment(amount float32) {
	// // razorpayPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	// // razorpayPaymentGw.pay(amount)
	// stripePaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}
type paypal struct{}

func (p paypal) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using paypal", amount)
}

func (p paypal) refund(amount float32, accountNum string) {
	// logic to make refund
	fmt.Println("making refund using paypal", amount, accountNum)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using stripe", amount)
}

func main() {
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}
	paypalPaymentGw := paypal{}
	// newPayment := payment{
	// 	gateway: razorpayPaymentGw,
	// }
	// newPayment2 := payment{
	// 	gateway: stripePaymentGw,
	// }
	newPayment3 := payment{
		gateway: paypalPaymentGw,

	}
	// newPayment.makePayment(100.0)
	// newPayment2.makePayment(50.0)
	newPayment3.makePayment(150.0)

	newPayment3.gateway.refund(100.0, "1234")
}
