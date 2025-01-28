package main

import (
	"fmt"
	"sync"
	"time"
)

//to wait for multiple goroutines to finish, we can use a Wait group

// this particular function is going to be used in every goroutine that we launch
func worker(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
