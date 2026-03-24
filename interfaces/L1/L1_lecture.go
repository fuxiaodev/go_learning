package L1

import (
	"fmt"
	"math"
)

/*
interfaces allow you to focus on what a type does rather than how its build
they can help you write more flexible and reusable code by defining behaviors that different types can share
interfaces are just collections of method signatures
a type implements an interface if it has methods that match the interface's method signatures
*/

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64{
	return 2 * r.width + 2 * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return math.Pi * 2 * c.radius
}

// when a type implements an interface, it can then be used as that interface type
func printShapeData(s shape) {
	fmt.Printf("Area: %v - Perimeter: %v\n", s.area(), s.perimeter())
}

// because the empty interface doesn't require a type to have any methods at all
// every type automatically implements the empty interface
type empty interface{}
