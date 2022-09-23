package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.head = &element[T]{val: v}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: v}
		list.tail = list.tail.next
	}
}

func (list *List[T]) GetAll() []T {
	var elems []T
	for e := list.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}

	return elems
}

func main() {
	//* Go 1.18 support generics
	var m map[int]string

	m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys: ", m)
	fmt.Println("values 1: ", m[1])

	var r []int
	r = MapKeys[int, string](m)
	fmt.Println(r)
	list := List[string]{}
	list.Push("2")
	list.Push("Three")
	list.Push("Four")

	fmt.Println(list.GetAll())

}
