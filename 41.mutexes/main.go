package main

import (
	"fmt"
	"sync"
	"time"
)

type Container struct {
	mu      sync.Mutex // Add Mutext to synchronize access N
	couters map[string]int
}

// NOTE: Mutext must not be copied pass around it with pointer

func (c *Container) inc(name string) {
	c.mu.Lock()         // Lock the "mutext" before accessing counters
	defer c.mu.Unlock() //Unlock at the end of function using "defer"
	c.couters[name]++

}

func main() {
	// For manage complex state that Atomic counter
	// use "Mutex" to safely access data across multiple goroutines

	c := Container{
		couters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup
	doIncrement := func(name string, n int) {
		fmt.Println(time.Now())
		time.Sleep(500 * time.Millisecond)
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// Run several goroutines concurrently
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait() //wait goroutines to finish
	fmt.Println(c.couters)
}
