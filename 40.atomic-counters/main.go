package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//Mechainsm for managing state in Go is communication over channels

	var ops uint64 // Always positive counter
	var ops2 uint64 = 0

	var wg sync.WaitGroup

	// start 50 goroutines that each increment the counter exactly 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func(i int) {
			fmt.Printf("start goroutines#%d\n", i)
			for c := 0; c < 1000; c++ {
				//* Atomically increment the counter by giving it the mem address
				atomic.AddUint64(&ops, 1)
				// * general incremtn the counter
				ops2 = ops2 + 1
			}
			fmt.Printf("stop #%d\n", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("blocking for result done!")
	fmt.Println("ops:", ops)
	fmt.Println("ops2:", ops2)
}
