package main

import (
	"fmt"
)

func main() {

	/*
	* By default "channel are unbuffered"
	* "Bufferd channels" is limited number of values without a receiver for those value
	* because by default channel must have sent and reciever
	 */
	messages := make(chan string, 2)
	messages <- "Hello"
	messages <- "World"
	go func() {
		messages <- "!"
	}()

	fmt.Print(<-messages)
	fmt.Print(<-messages)
	fmt.Print(<-messages)

}
