package main

import (
	"fmt"
)

// send
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second)
// 	}
// }

// func sum(result chan int, num1 int, num2 int) {
// 	total := num1 + num2

// 	result <- total
// }

// goroutine synchronizer
// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("Processing...")
// 	done <- true
// }

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("Sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "golang"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}

	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i:= 0; i < 15; i++{
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("All emails sent")

	// // this is important
	// close(emailChan)

	// <- done

	// done := make(chan bool)

	// go task(done)

	// do := <-done // block

	// fmt.Println(do)

	// result := make(chan int)

	// go sum(result, 3, 4)

	// res := <-result
	// fmt.Println(res)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// numChan <- 5

	// messageChan := make(chan string)

	// messageChan <- "ping" // blocking

	// msg := <-messageChan

	// fmt.Println(msg)
}
