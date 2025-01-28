package main

import "fmt"

//range can also be used to iterate over the values received from the channel

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	//but if we compare to the manual for loop, that we've used in closing-channel.go code
	//range over a channel only returns one value - the received value, whereas manual for loop also provides second value
	for elem := range queue {
		fmt.Println(elem) //since we closed the channel, the iteration terminates after the last value
	}
}
