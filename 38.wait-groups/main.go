package main

import (
	"fmt"
	"sync"
	"time"
)

// This fuc will run in every goroutine
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // For simulate expensive tasks
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// This WaitGroup is used to wait for all goroutines launched here to finish
	// * Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) //  increment the WaitGroup counter for each.

		i := i //Avoid re-use of the same i value in each goroutine closure

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait() //Block until the WaitGroup counter goes back to 0
}
