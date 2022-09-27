package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	* 1. Sends and recieves on channels are blocking
	* 2. We can use select with a "default" clause to implement non-blocking sends, receievs and even non-blocking multi-way selects
	 */
	messages := make(chan string)
	signals := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		messages <- "hello"
	}()

	// Non-blocking recieve
	// If value is available on "messages" then select will take the <-messages case with that value
	// otherwise immediately take the default case
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	// Non-blocking sent works like above
	// "msg" cannot be sent to messages channel
	// because the channel has no buffer and there is no receiver
	// Therefore the default case i selected
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Use multiple cases above the "default" clause
	// to implement a "multi-way non-blocking select"
	// Here we attempt non-blocking receives on both message and signals
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	fmt.Println("End of program")
}
