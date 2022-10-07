package main

import (
	"fmt"
	"sort"
)

type byLength []string // Just alias of []string type

// ************ Implement sort.Interface -> Len():int, Swap(i, j), less(i,j):bool
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// ********** End Implement sort.Interface

func main() {
	// Sort by custom function need to implement sort.Interface
	fruits := []string{"peach", "banana", "papaya", "apple", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

}
