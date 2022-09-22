package main

import "fmt"

func initSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

func main() {

	//* Go support "anonymous functions" which can form closures
	// Ex 1. simple closures
	addFunc := func(a, b int) int {
		return a + b
	}
	a := 10
	b := 5
	fmt.Printf("%d + %d = %d\n", a, b, addFunc(a, b))

	/**
	Ex2. Function(1) return another function(2)
	this function(2) capture it own `i` value which updated each time
	we call nextInt
	*/
	nextInt := initSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := initSeq()
	fmt.Println(newInts())

}
