package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height

}

func main() {
	var r rect
	r = rect{20, 30}
	fmt.Println("Area: ", r.area())
	fmt.Println("perim: ", r.perim())

	var rp *rect
	rp = &r

	fmt.Println("Area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	fmt.Printf("r -> %p \n", &r)
	fmt.Printf("rp -> %p \n", rp)

}
