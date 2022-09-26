package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	// var a chan int
	// if a == nil {
	// 	fmt.Println("Channel a is nil, going to define it")
	// 	a = make(chan int)
	// 	fmt.Printf("Type of a is %T", a)
	// }
	/*
	 * Sending and receiving from a channel
	 * Example
	 * data := <- a // read from channel a
	 * a <- data // write to channel a
	 */

	// * Ex1
	done := make(chan bool)
	go hello(done)
	fmt.Println("before retrieve done")
	<-done
	fmt.Println("main function")

	// * Ex2 communitaion
	number := 59
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)

	// * Ex3 Deadlock
	ch := make(chan int)
	ch <- 5 // Send data to channel but Goroutine receiving data from the channel "ch"

	// * Ex4 Unidirectional channels

}
