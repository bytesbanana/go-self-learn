package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("\n------")
	fmt.Println(g)
	fmt.Printf("area: %.3f\n", g.area())
	fmt.Printf("perim: %.3f\n", +g.perim())
}

func main() {
	r := rect{width: 30, height: 20}
	c := circle{radius: 10}

	measure(r)
	measure(c)
}
