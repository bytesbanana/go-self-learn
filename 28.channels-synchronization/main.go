package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	/*
	* We can use channels to synchronize execution across goroutines
	 */
	done := make(chan bool)
	go worker(done) // Start goroutines

	<-done // This line will be block unit recieve notification from the worker on the channel

}
