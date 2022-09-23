package main

import "fmt"

// * "person" struct type
type person struct {
	name string
	age  int
}

// * "newPerson" constructs a new person with given name
func newPerson(name string) *person {
	p := person{
		name: name,
	}

	p.age = 42

	/*
	* Can safely return pointer to local variable
	* As a local variable will SURVIVE the scope of the function
	 */
	fmt.Printf("New person %s %p \n", name, &p)
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})

	john := newPerson("John")

	//* Pointer survive but different address
	fmt.Println(john)
	fmt.Printf("john ptr: %p \n", &john)

	sean := person{name: "Sean", age: 50}
	fmt.Println(sean)

	seanPtr := &sean
	fmt.Println(sean.age)

	seanPtr.age = 51
	fmt.Printf("%p\n", &sean)
	fmt.Printf("%p\n", seanPtr)
	fmt.Println(seanPtr.age)

}
