package main

import (
	"fmt"
	"time"
)

/*

* Channels are the pipes that connect concurrent goroutines.
* You can send values into channels from one goroutine and receive those values into another goroutine.

 */

func main() {

	/*
	* Create a new channel with make(chan val-type).
	* Channels are typed by the values they convey.
	 */
	fmt.Println("Start:", time.Now())
	messages := make(chan string)

	go func() {
		//* Send a value into a channel using the channel <- syntax
		time.Sleep(time.Second * 2)
		fmt.Println("Excute:", time.Now())
		messages <- "ping"
	}()

	fmt.Println("Recieved:", time.Now())
	msg := <-messages
	// fmt.Println(<-messages)
	fmt.Println("Done:", time.Now())
	fmt.Println(msg)
	// !By default sends and receives block until both the sender and receiver are ready
	// !This property allowed us to "wait at the end of our program" for the "ping" message without having to use any other synchronizatio

}
