package main

import "fmt"

func sums(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := sums(input...)

	fmt.Printf("Sums of %+v is %d\n", input, result)
}
