package main

import "fmt"

func main() {
	/*
		Use "range" to iterates over elements in a variety of data strutures
	*/

	// 1. Use "range" to iterates over the slice ( Arrays also works too)
	// Ex. 1.1
	nums := []int{2, 3, 4}
	sum := 0

	// _ for ignore using index
	for _, n := range nums {
		sum += n
	}
	fmt.Println(sum)
	// Ex. 2.1
	for i, n := range nums {
		fmt.Printf("index %d value is %d\n", i, n)
	}

	// 2. Use "range" to iterates over the keys fo map

	// Ex.2.1 key and value
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// Ex.2.2 only key
	for k := range kvs {
		fmt.Printf("Key: %s\n", k)
	}

	// 3.Use "range" to interates over string and runes
	for i, c := range "go" {
		fmt.Println(i, c, string(c))
	}

}
