package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}
func main() {
	s := square{
		side: 5,
	}

	c := circle{
		radius: 5,
	}

	fmt.Println("area of square is ", info(s))
	fmt.Println("area of circle is ", info(c))

}
