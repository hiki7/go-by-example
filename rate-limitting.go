package main

import (
	"fmt"
	"time"
)

func main() {
	//sending requests to the channel
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	//closing a channel is necessary when you want to signal that no more values will be sent on it
	//especially it is important when we use range over the channel, so that it will know when to stop waiting for the values
	close(requests) //if we don't close the channel and try to range over it, we will get deadlock, since the receiver keeps waiting for new values

	limiter := time.Tick(200 * time.Millisecond) //this limiter channel will receive value every 200ms

	for req := range requests {
		<-limiter //by blocking on a receive from the limiter channel before serving each request, we limit ourselves
		//to 1 request every 200ms

		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3) //this limiter can handle 3 requests instantly, since it is a buffered channel

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	} //fill up the channel to represent allowed bursting

	//every 200ms we'll try to add a new value to burstyLimiter
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	//Now simulate 5 more incomming requests, the first 3 will benefit from the burst capability of burstyLimiter
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyRequests
		fmt.Println("request", req, time.Now())
	}
}
