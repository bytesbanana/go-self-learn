package main

import "fmt"

func factorail(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorail(n-1)
}

func main() {
	n := 7

	fmt.Printf("Factorial of %d is %d\n", n, factorail(n))

	var fibonanci func(n int) int

	fibonanci = func(n int) int {
		if n < 2 {
			return n
		}
		return fibonanci(n-1) + fibonanci(n-2)
	}

	fmt.Printf("Fibonanci of %d is %d", n, fibonanci(n))
}
