package main

import "fmt"

// * pings only recieve value
func ping(pings chan<- string, msg string) {

	pings <- msg
}

// * pings only send value
// * pongs only recieve value
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	// * The way to specify if a channel is meant to be "recieve" or "send" values only
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passeed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
