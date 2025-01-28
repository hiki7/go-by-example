package main

import (
	"fmt"
	"time"
)

func main() {
	//Generally, timers are useful for scheduling delayed execution
	timer1 := time.NewTimer(2 * time.Second) //creating a timer that will fire after 2 seconds
	//Firing a timer means that timer reaches a zero and sends a signal through timer.C
	//this allows goroutine to unblock and continue execution

	//It's important to note that - timer is not a CHANNEL, it just provides channel

	<-timer1.C //blocks the execution until the timer fires
	fmt.Println("Timer 1 fired")

	//if we wanted just to wait, we could've use Sleep
	// but one of the main features of timers is that it can be stopped, as it is shown in the example below
	timer2 := time.NewTimer(time.Second)

	//starting a goroutine that waits for timer2 to fire
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	//this code will execute faster, since it is in the main() function
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)

	//it's important to note that: main() runs synchronously(sequentially), whereas goroutines run asynchronously(concurrently)
}
