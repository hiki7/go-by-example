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

func main() {
	done := make(chan bool, 1) //buffered channel size 1
	go worker(done)            //running worker in a separate goroutine. This starts worker in parallel, allowing main() to keep running

	<-done
}
