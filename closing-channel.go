package main

import "fmt"

func main() {
	//creating buffered and unbuffered channels
	jobs := make(chan int, 5)
	done := make(chan bool)

	//this goroutine continuously receives from jobs channel until the channel is closed
	//more is false when the channel is closed
	//it's very important to note that receiving from the channel can return two values: the actual value and the boolean indicating whether the channel is open or closed
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				//when all jobs are received, it sends true to done channel
				done <- true
				return
			}
		}
	}()

	//sending 3 jobs to the channel
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	//closing the channel after all jobs are sent
	close(jobs)
	fmt.Println("sent all jobs")

	//when we write like this without assigning to variable, we are simply waiting for a value to be sent - we need just the signal that smth happened
	// since done is unbuffered channel, this ensure that main() waits until the goroutine sends a value
	<-done

	//trying to receive smth from the closed jobs channel
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok) //ok will be false
}
