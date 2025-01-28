package main

import (
	"fmt"
	"time"
)

// timeout is a way to limit time a goroutine waits for an operation (waiting for a channel or a function to return)
// reasons to use timeouts:
//   - prevents blocking indefinitely when waiting for responses
//
// - ensures fast failure when dealing with network requests, databases or long computations
func main() {

	//creating a buffered channel
	c1 := make(chan string, 1)
	//executing a call that returns its result on a channel c1 after 2s
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	//this select is implementing timeout
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second): //since the operation of sending the result to c1 takes 2 s, this particular case will be executed
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	//c2 channel takes 2 seconds to send the result
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	//in this select statement the first case will be executed, since timeout is 3s and operation time is 2s
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

	//generally this program shows the first operation timing out and the second operation succeeding
}
