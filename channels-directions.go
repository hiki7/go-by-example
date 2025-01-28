package main

import "fmt"

//when using channels as function parameters, we can specify that the channel is meant to be sending or receiving
// this specifity increases type-safety

// this ping function used to only send values
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// the pong function accepts one channel for receiving(pings) and another channel for sending(pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
