package main

import "fmt"

func main() {
	q := make(chan string, 2)
	q <- "one"
	q <- "two"
	close(q)

	// Can use range when channel "q" is closed
	for elem := range q {
		fmt.Println(elem)
	}
}
