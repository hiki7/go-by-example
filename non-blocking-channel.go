package main

import "fmt"

func main() {
	//basic sends and receive operations are blocking, which means that

	messages := make(chan string)
	signals := make(chan bool)

	//here is non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println("message is received", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"

	//non-blocking send
	select {
	case messages <- msg:
		fmt.Println("message is sent", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("message is received", msg)
	case signal := <-signals:
		fmt.Println("signal is received", signal)
	default:
		fmt.Println("no activity")
	}
}
