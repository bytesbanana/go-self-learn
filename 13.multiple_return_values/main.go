package main

import "fmt"

func addTenToBoth(a, b int) (int, int) {
	return a + 10, b + 10
}

func main() {
	a := 10
	b := 12
	resA, resB := addTenToBoth(a, b)

	fmt.Printf("Add 10 to %d , %d equals %d, %d", a, b, resA, resB)
}
