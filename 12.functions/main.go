package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	a := 10
	b := 12
	res := add(a, b)

	fmt.Printf("Result of %d + %d = %d", a, b, res)
}
