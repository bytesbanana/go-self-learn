package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"ac", "ab", "aa"}
	sort.Strings(strs)
	fmt.Println("strings", strs)

	ints := []int{3, 1, 2}
	sort.Ints(ints)
	fmt.Println("ints", ints)

	fmt.Println("Is strings are sorted? ", sort.StringsAreSorted((strs)))
}
