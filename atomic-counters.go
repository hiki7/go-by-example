package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//this counter is used for counting the number of operations that have been performed by multiple goroutines
	var ops atomic.Uint64 //atomic interger type
	//why don't we use just simple int
	//because, atomic integer type prevents race conditions
	//race condition happen when multiple goroutines access and modify the same shared variable at the same time
	//which might lead to unpredictable or incorrect results

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops.Load())
}
