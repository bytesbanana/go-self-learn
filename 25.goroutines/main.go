package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Goroutine = light weight thread of execution

	f("direct") //run func synchronously

	go f("goroutine") // run goroutine -> will execute concurrenyly

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Wait them to finish (Don't have this line won't print any thing in concurrent task)

	time.Sleep(time.Second)
	fmt.Println("done")
}
