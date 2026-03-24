package L6

import (
	"math"
	"fmt"
)

/*
when working with interfaces, sometimes you will need access to the underlying
type of an interface value
you can cast an interface to its underlying type using a type assertion
*/
type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

/*
if ok is true, c is indeed a circle, otherwise, it is not
*/
func printShapeInfo(s shape) {
	c, ok := s.(circle)
	if ok {
		radius := c.radius
		fmt.Printf("s is a circle, radius: %v\n", radius)
		return
	}
}
