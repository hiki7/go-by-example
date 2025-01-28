package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

// the main function waits for worker to complete before exiting
func main() {
	done := make(chan bool, 1) //buffered channel size 1, this channel is used to signal main function when the worker function is done
	go worker(done)            //running worker in a separate goroutine. This starts worker in parallel, allowing main() to keep running

	<-done
}
