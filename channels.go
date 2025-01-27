package main

import "fmt"

//channels are the pipes that connect concurrent goroutines.
// you can send values from one goroutine and receive these values in another goroutine

func main() {
	messages := make(chan string)

	// channel <- value - this is the syntax for sending the values into the channel
	go func() {
		messages <- "ping"
	}()

	// <-channel - this is the syntax for receiving the values from the channel
	msg := <-messages

	fmt.Println(msg)
}
