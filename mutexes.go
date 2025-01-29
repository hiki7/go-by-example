package main

import (
	"fmt"
	"sync"
)

//mutex is a locker that allows only one goroutine to access some resource/variable at a time

type Container struct {
	mu       sync.Mutex //we use mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()         //locks the mutex before modifying counters
	defer c.mu.Unlock() //using defer to ensure that Unlock is always used
	c.counters[name]++  //generally this piece of code prevents from race conditions when multiple goroutines increment the same key
}

func main() {
	//creaing container and waitgroup
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup //this ensure that main() waits for all goroutines to finish before printing counters

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
